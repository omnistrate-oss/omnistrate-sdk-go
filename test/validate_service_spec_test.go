// Package test holds hand-written tests for the generated Omnistrate SDK.
//
// It deliberately sits outside v1/ and fleet/, which `make clean` deletes and the generator
// rewrites; a test that only exists until the next generation cannot protect anything.
package test

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	v1 "github.com/omnistrate-oss/omnistrate-sdk-go/v1"
)

// ValidateServiceSpec is a read-only endpoint whose entire value depends on the server seeing the
// exact bytes the caller meant to validate. These tests drive the generated client against an
// httptest server and assert on the raw wire bytes rather than on the client's own round trip,
// because a symmetric encode/decode bug is invisible to a round-trip assertion and would still
// mean the server validated something the user never wrote.

const validateSpecPath = "/2022-09-01-00/service/spec/validate"

// capturedRequest is what the fake server saw.
type capturedRequest struct {
	method string
	path   string
	header http.Header
	body   []byte
}

// newValidationServer serves respond() at the validation route and records the request. It fails
// the test if the client sends anything to any other path, which is how "no fallback to the legacy
// build endpoint" is observable from here.
func newValidationServer(t *testing.T, respond func(w http.ResponseWriter)) (*v1.APIClient, *capturedRequest) {
	t.Helper()

	captured := &capturedRequest{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)

		captured.method = r.Method
		captured.path = r.URL.Path
		captured.header = r.Header.Clone()
		captured.body = body

		if r.URL.Path != validateSpecPath {
			t.Errorf("client sent %s %s; validation must never reach another endpoint", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		respond(w)
	}))
	t.Cleanup(server.Close)

	cfg := v1.NewConfiguration()
	cfg.Servers = v1.ServerConfigurations{{URL: server.URL}}
	return v1.NewAPIClient(cfg), captured
}

// respondWithResult writes a 200 result body.
func respondWithResult(t *testing.T, body string) func(w http.ResponseWriter) {
	t.Helper()
	return func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := io.WriteString(w, body)
		require.NoError(t, err)
	}
}

// minimalResult is the smallest response satisfying the contract: a status, a version, a digest and
// four arrays that are present even when empty.
const minimalResult = `{
  "status": "VALID",
  "validationVersion": "1",
  "inputDigest": "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
  "checks": [],
  "diagnostics": [],
  "requiredArtifacts": [],
  "validatedArtifacts": []
}`

// buildArchive returns a small gzipped tar and the digest of its exact compressed bytes, which is
// what the wire contract says sha256 covers.
func buildArchive(t *testing.T, name, content string) (raw []byte, digest string) {
	t.Helper()

	var compressed bytes.Buffer
	gz := gzip.NewWriter(&compressed)
	tw := tar.NewWriter(gz)

	require.NoError(t, tw.WriteHeader(&tar.Header{
		Name: name,
		Mode: 0o600,
		Size: int64(len(content)),
	}))
	_, err := tw.Write([]byte(content))
	require.NoError(t, err)
	require.NoError(t, tw.Close())
	require.NoError(t, gz.Close())

	raw = compressed.Bytes()
	sum := sha256.Sum256(raw)
	return raw, hex.EncodeToString(sum[:])
}

// The field names on the wire are the contract. A Go-side rename is free; a JSON-side rename is a
// silently rejected or, worse, silently ignored request.
func TestValidateServiceSpec_RequestWireFieldNames(t *testing.T) {
	t.Parallel()

	client, captured := newValidationServer(t, respondWithResult(t, minimalResult))

	specBytes := []byte("name: mysql\n")
	fileContent := base64.StdEncoding.EncodeToString(specBytes)

	req := v1.NewValidateServiceSpecRequest2(fileContent, "MySQL multi-writer service", "service-plan")
	req.SetEnvironment("dev")
	req.SetEnvironmentType("DEV")
	req.SetDescription("A MySQL SaaS")
	req.SetServiceLogoURL("https://example.com/logo.png")
	req.SetReleaseVersionName("v1")
	req.SetConfigs(map[string]string{"app-config": base64.StdEncoding.EncodeToString([]byte("k: v"))})
	req.SetSecrets(map[string]string{"app-secret": base64.StdEncoding.EncodeToString([]byte("s3cr3t"))})

	_, httpRes, err := client.ServiceApiAPI.
		ServiceApiValidateServiceSpec(context.Background()).
		ValidateServiceSpecRequest2(*req).
		Execute()
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, httpRes.StatusCode)

	assert.Equal(t, http.MethodPost, captured.method)
	assert.Equal(t, validateSpecPath, captured.path)
	assert.Equal(t, "application/json", captured.header.Get("Content-Type"))

	var body map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(captured.body, &body))

	for _, field := range []string{
		"name", "specType", "fileContent", "environment", "environmentType",
		"description", "serviceLogoURL", "releaseVersionName", "configs", "secrets",
	} {
		assert.Contains(t, body, field, "%s must be sent under exactly this name", field)
	}

	// The token travels in the Authorization header. A body copy would be a credential in every
	// request log that records bodies.
	assert.NotContains(t, body, "token")
	// There is no mode switch on this endpoint, so there is nothing here that could turn it into a
	// build, and no tenant identity for a caller to forge.
	for _, forbidden := range []string{"dryrun", "dryRun", "apply", "orgId", "userId", "serviceId"} {
		assert.NotContains(t, body, forbidden)
	}
}

// An explicitly false flag is not the same statement as an omitted flag, and `omitempty` on a bare
// bool would erase the difference. The pointer types are what keep "the user said release=false"
// distinguishable from "the user said nothing", so the false must actually reach the wire.
func TestValidateServiceSpec_ExplicitlyFalseFlagsAreSerialized(t *testing.T) {
	t.Parallel()

	client, captured := newValidationServer(t, respondWithResult(t, minimalResult))

	req := v1.NewValidateServiceSpecRequest2("Zm9v", "svc", "compose")
	req.SetRelease(false)
	req.SetReleaseAsPreferred(false)
	req.SetForceCreateNewServicePlanVersion(false)

	_, _, err := client.ServiceApiAPI.
		ServiceApiValidateServiceSpec(context.Background()).
		ValidateServiceSpecRequest2(*req).
		Execute()
	require.NoError(t, err)

	var body map[string]any
	require.NoError(t, json.Unmarshal(captured.body, &body))

	for _, flag := range []string{"release", "releaseAsPreferred", "forceCreateNewServicePlanVersion"} {
		value, present := body[flag]
		require.True(t, present, "%s was set to false and must appear on the wire, not be dropped", flag)
		assert.Equal(t, false, value, "%s must serialize as the boolean false", flag)
	}
}

// The other half of the same rule: a flag nobody set must not appear at all, so that the server can
// apply its own default instead of being told "false" by the client's zero value.
func TestValidateServiceSpec_OmittedFlagsAreAbsent(t *testing.T) {
	t.Parallel()

	client, captured := newValidationServer(t, respondWithResult(t, minimalResult))

	req := v1.NewValidateServiceSpecRequest2("Zm9v", "svc", "compose")

	_, _, err := client.ServiceApiAPI.
		ServiceApiValidateServiceSpec(context.Background()).
		ValidateServiceSpecRequest2(*req).
		Execute()
	require.NoError(t, err)

	var body map[string]any
	require.NoError(t, json.Unmarshal(captured.body, &body))

	assert.Equal(t, []string{"fileContent", "name", "specType"}, sortedKeys(body),
		"a discovery request carries only what the caller actually stated")
}

func sortedKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	for i := 1; i < len(keys); i++ {
		for j := i; j > 0 && keys[j] < keys[j-1]; j-- {
			keys[j], keys[j-1] = keys[j-1], keys[j]
		}
	}
	return keys
}

// The credential goes in the Authorization header with the Bearer scheme, and nowhere else.
func TestValidateServiceSpec_AuthorizationHeader(t *testing.T) {
	t.Parallel()

	client, captured := newValidationServer(t, respondWithResult(t, minimalResult))

	const token = "test-jwt-token-value"
	ctx := context.WithValue(context.Background(), v1.ContextAccessToken, token)

	req := v1.NewValidateServiceSpecRequest2("Zm9v", "svc", "compose")
	_, _, err := client.ServiceApiAPI.
		ServiceApiValidateServiceSpec(ctx).
		ValidateServiceSpecRequest2(*req).
		Execute()
	require.NoError(t, err)

	assert.Equal(t, "Bearer "+token, captured.header.Get("Authorization"))
	assert.NotContains(t, string(captured.body), token,
		"the token must not also be copied into the request body")
}

// Validation is performed on the bytes the client sends. If base64 transport mutates them by even
// one byte, every downstream diagnostic describes content the user does not have.
func TestValidateServiceSpec_ArtifactContentBytesSurviveTheWire(t *testing.T) {
	t.Parallel()

	client, captured := newValidationServer(t, respondWithResult(t, minimalResult))

	archive, digest := buildArchive(t, "main.tf", "resource \"null_resource\" \"a\" {}\n")

	artifact := v1.NewValidationArtifactInput(
		base64.StdEncoding.EncodeToString(archive),
		int64(len(archive)),
		"tar+gzip+base64",
		"terraform/network",
		digest,
	)

	req := v1.NewValidateServiceSpecRequest2("Zm9v", "svc", "service-plan")
	req.SetArtifacts([]v1.ValidationArtifactInput{*artifact})

	_, _, err := client.ServiceApiAPI.
		ServiceApiValidateServiceSpec(context.Background()).
		ValidateServiceSpecRequest2(*req).
		Execute()
	require.NoError(t, err)

	var body struct {
		Artifacts []struct {
			LogicalPath         string `json:"logicalPath"`
			Encoding            string `json:"encoding"`
			ArchiveContent      string `json:"archiveContent"`
			Sha256              string `json:"sha256"`
			CompressedSizeBytes int64  `json:"compressedSizeBytes"`
		} `json:"artifacts"`
	}
	require.NoError(t, json.Unmarshal(captured.body, &body))
	require.Len(t, body.Artifacts, 1)

	got := body.Artifacts[0]
	assert.Equal(t, "terraform/network", got.LogicalPath)
	assert.Equal(t, "tar+gzip+base64", got.Encoding)
	assert.Equal(t, digest, got.Sha256)
	assert.Equal(t, int64(len(archive)), got.CompressedSizeBytes)

	decoded, err := base64.StdEncoding.DecodeString(got.ArchiveContent)
	require.NoError(t, err, "archiveContent must be plain base64 the server can decode")
	assert.Equal(t, archive, decoded, "the server must receive the exact compressed bytes")

	recomputed := sha256.Sum256(decoded)
	assert.Equal(t, digest, hex.EncodeToString(recomputed[:]),
		"the digest the server recomputes must match the one the client sent")

	// Reading the archive back proves it is a real tar.gz and not, say, base64 of base64.
	gz, err := gzip.NewReader(bytes.NewReader(decoded))
	require.NoError(t, err)
	entry, err := tar.NewReader(gz).Next()
	require.NoError(t, err)
	assert.Equal(t, "main.tf", entry.Name)
}

// The discovery response from specification 03. A client that cannot read this cannot package the
// content the server asked for.
func TestValidateServiceSpec_DecodesIncompleteDiscoveryResult(t *testing.T) {
	t.Parallel()

	const discovery = `{
  "status": "INCOMPLETE",
  "validationVersion": "1",
  "inputDigest": "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
  "observedAt": "2026-09-05T12:00:00Z",
  "existingTarget": {
    "serviceID": "s-existing",
    "serviceEnvironmentID": "se-existing",
    "productTierID": "pt-existing"
  },
  "checks": [
    {"name": "syntax", "status": "PASSED"},
    {"name": "artifact-content", "status": "INCOMPLETE"}
  ],
  "diagnostics": [
    {
      "code": "ARTIFACT_CONTENT_REQUIRED",
      "severity": "incomplete",
      "path": "/services/0/terraformConfigurations/configurationPerCloudProvider/aws",
      "resourceKey": "network",
      "message": "Local artifact content is required for validation."
    }
  ],
  "requiredArtifacts": [
    {
      "logicalPath": "terraform/network",
      "uses": [
        {"resourceKey": "network", "kind": "terraform", "provider": "aws", "path": "/services/0/terraformConfigurations/configurationPerCloudProvider/aws"}
      ]
    }
  ],
  "validatedArtifacts": []
}`

	client, _ := newValidationServer(t, respondWithResult(t, discovery))

	req := v1.NewValidateServiceSpecRequest2("Zm9v", "svc", "service-plan")
	result, httpRes, err := client.ServiceApiAPI.
		ServiceApiValidateServiceSpec(context.Background()).
		ValidateServiceSpecRequest2(*req).
		Execute()

	// INCOMPLETE is a successfully processed request, not a transport failure.
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, httpRes.StatusCode)
	require.NotNil(t, result)

	assert.Equal(t, "INCOMPLETE", result.GetStatus())
	assert.Equal(t, "1", result.GetValidationVersion())
	assert.Equal(t, "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", result.GetInputDigest())
	assert.Equal(t, time.Date(2026, 9, 5, 12, 0, 0, 0, time.UTC), result.GetObservedAt())

	target := result.GetExistingTarget()
	assert.Equal(t, "s-existing", target.GetServiceID())
	assert.Equal(t, "se-existing", target.GetServiceEnvironmentID())
	assert.Equal(t, "pt-existing", target.GetProductTierID())

	require.Len(t, result.GetChecks(), 2)
	assert.Equal(t, "syntax", result.GetChecks()[0].GetName())
	assert.Equal(t, "PASSED", result.GetChecks()[0].GetStatus())
	assert.Equal(t, "artifact-content", result.GetChecks()[1].GetName())
	assert.Equal(t, "INCOMPLETE", result.GetChecks()[1].GetStatus())

	require.Len(t, result.GetDiagnostics(), 1)
	diag := result.GetDiagnostics()[0]
	assert.Equal(t, "ARTIFACT_CONTENT_REQUIRED", diag.GetCode())
	assert.Equal(t, "incomplete", diag.GetSeverity())
	assert.Equal(t, "network", diag.GetResourceKey())
	assert.Equal(t, "/services/0/terraformConfigurations/configurationPerCloudProvider/aws", diag.GetPath())

	require.Len(t, result.GetRequiredArtifacts(), 1)
	requirement := result.GetRequiredArtifacts()[0]
	assert.Equal(t, "terraform/network", requirement.GetLogicalPath())
	require.Len(t, requirement.GetUses(), 1, "a requirement is never reported with an empty uses list")
	use := requirement.GetUses()[0]
	assert.Equal(t, "network", use.GetResourceKey())
	assert.Equal(t, "terraform", use.GetKind())
	assert.Equal(t, "aws", use.GetProvider())
	_, hasOnPrem := use.GetOnPremPlatformOk()
	assert.False(t, hasOnPrem, "a cloud-provider use carries no on-prem platform")

	assert.NotNil(t, result.GetValidatedArtifacts())
	assert.Empty(t, result.GetValidatedArtifacts(),
		"nothing was supplied, so nothing was validated; the array is present and empty")
}

// The CLI prints these arrays. An explicit [] and an omitted field are different for it, so the
// empty case has to decode into something it can range over rather than something it must nil-check.
func TestValidateServiceSpec_EmptyArraysDecodeAsEmptyNotAbsent(t *testing.T) {
	t.Parallel()

	client, _ := newValidationServer(t, respondWithResult(t, minimalResult))

	req := v1.NewValidateServiceSpecRequest2("Zm9v", "svc", "compose")
	result, _, err := client.ServiceApiAPI.
		ServiceApiValidateServiceSpec(context.Background()).
		ValidateServiceSpecRequest2(*req).
		Execute()
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.NotNil(t, result.Checks)
	assert.NotNil(t, result.Diagnostics)
	assert.NotNil(t, result.RequiredArtifacts)
	assert.NotNil(t, result.ValidatedArtifacts)

	_, hasObservedAt := result.GetObservedAtOk()
	assert.False(t, hasObservedAt, "no snapshot was read, so no observation time is reported")
	_, hasTarget := result.GetExistingTargetOk()
	assert.False(t, hasTarget, "a candidate with no existing target reports none")
	_, hasLimits := result.GetLimitsOk()
	assert.False(t, hasLimits, "limits are optional")
}

// A validated artifact is the acknowledgement the client checks before it believes its bytes were
// used: same canonical path, same digest, and the uses whose validators actually consumed them.
func TestValidateServiceSpec_DecodesValidatedArtifacts(t *testing.T) {
	t.Parallel()

	const validated = `{
  "status": "VALID",
  "validationVersion": "1",
  "inputDigest": "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
  "checks": [{"name": "artifact-content", "status": "PASSED"}],
  "diagnostics": [],
  "requiredArtifacts": [],
  "validatedArtifacts": [
    {
      "logicalPath": "terraform/network",
      "sha256": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
      "compressedSizeBytes": 1234,
      "uses": [
        {"resourceKey": "network", "kind": "terraform", "provider": "aws", "path": "/p/aws"},
        {"resourceKey": "network", "kind": "terraform", "onPremPlatform": "Generic", "path": "/p/onprem"}
      ]
    }
  ]
}`

	client, _ := newValidationServer(t, respondWithResult(t, validated))

	req := v1.NewValidateServiceSpecRequest2("Zm9v", "svc", "service-plan")
	result, _, err := client.ServiceApiAPI.
		ServiceApiValidateServiceSpec(context.Background()).
		ValidateServiceSpecRequest2(*req).
		Execute()
	require.NoError(t, err)

	require.Len(t, result.GetValidatedArtifacts(), 1)
	artifact := result.GetValidatedArtifacts()[0]
	assert.Equal(t, "terraform/network", artifact.GetLogicalPath())
	assert.Equal(t, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", artifact.GetSha256())
	assert.Equal(t, int64(1234), artifact.GetCompressedSizeBytes())

	require.Len(t, artifact.GetUses(), 2, "one archive can serve several uses and reports all of them")
	assert.Equal(t, "aws", artifact.GetUses()[0].GetProvider())
	assert.Equal(t, "Generic", artifact.GetUses()[1].GetOnPremPlatform())
	_, hasProvider := artifact.GetUses()[1].GetProviderOk()
	assert.False(t, hasProvider)
}

// A client that wants to fail early instead of sending 24 MiB it knows will be refused needs every
// bound the server enforces, not a subset.
func TestValidateServiceSpec_DecodesAdvertisedLimits(t *testing.T) {
	t.Parallel()

	const withLimits = `{
  "status": "INCOMPLETE",
  "validationVersion": "1",
  "inputDigest": "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
  "checks": [],
  "diagnostics": [],
  "requiredArtifacts": [],
  "validatedArtifacts": [],
  "limits": {
    "maxRequestBodyBytes": 25165824,
    "maxTotalCompressedArtifactBytes": 16777216,
    "maxTotalExtractedBytes": 67108864,
    "maxTotalSpecBytes": 1048576,
    "maxArtifacts": 64,
    "maxArchiveEntries": 4096,
    "maxArchiveMemberPathBytes": 1024,
    "requestDeadlineSeconds": 120,
    "maxConcurrentContentValidations": 4
  }
}`

	client, _ := newValidationServer(t, respondWithResult(t, withLimits))

	req := v1.NewValidateServiceSpecRequest2("Zm9v", "svc", "service-plan")
	result, _, err := client.ServiceApiAPI.
		ServiceApiValidateServiceSpec(context.Background()).
		ValidateServiceSpecRequest2(*req).
		Execute()
	require.NoError(t, err)

	limits, hasLimits := result.GetLimitsOk()
	require.True(t, hasLimits)
	assert.Equal(t, int64(24*1024*1024), limits.GetMaxRequestBodyBytes())
	assert.Equal(t, int64(16*1024*1024), limits.GetMaxTotalCompressedArtifactBytes())
	assert.Equal(t, int64(64*1024*1024), limits.GetMaxTotalExtractedBytes())
	assert.Equal(t, int64(1024*1024), limits.GetMaxTotalSpecBytes())
	assert.Equal(t, int64(64), limits.GetMaxArtifacts())
	assert.Equal(t, int64(4096), limits.GetMaxArchiveEntries())
	assert.Equal(t, int64(1024), limits.GetMaxArchiveMemberPathBytes())
	assert.Equal(t, int64(120), limits.GetRequestDeadlineSeconds())
	assert.Equal(t, int64(4), limits.GetMaxConcurrentContentValidations())
}

// An invalid specification is a finding, not a transport failure. If the client turned this into an
// error it would have no diagnostics to print.
func TestValidateServiceSpec_InvalidSpecIsNotAnHTTPError(t *testing.T) {
	t.Parallel()

	const invalid = `{
  "status": "INVALID",
  "validationVersion": "1",
  "inputDigest": "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
  "checks": [{"name": "syntax", "status": "FAILED"}],
  "diagnostics": [
    {"code": "SPEC_SYNTAX_INVALID", "severity": "error", "message": "Unexpected token.", "path": "/services/0"}
  ],
  "requiredArtifacts": [],
  "validatedArtifacts": []
}`

	client, _ := newValidationServer(t, respondWithResult(t, invalid))

	req := v1.NewValidateServiceSpecRequest2("Zm9v", "svc", "compose")
	result, httpRes, err := client.ServiceApiAPI.
		ServiceApiValidateServiceSpec(context.Background()).
		ValidateServiceSpecRequest2(*req).
		Execute()

	require.NoError(t, err, "an INVALID specification is still a 200")
	assert.Equal(t, http.StatusOK, httpRes.StatusCode)
	assert.Equal(t, "INVALID", result.GetStatus())
	require.Len(t, result.GetDiagnostics(), 1)
	assert.Equal(t, "error", result.GetDiagnostics()[0].GetSeverity())
}

// Envelope failures are HTTP statuses, and the client has to be able to tell them apart: 413 means
// send less, 429 means come back later, 504 means the deadline passed, and only 500 is a defect.
func TestValidateServiceSpec_EnvelopeFailuresSurfaceTheirStatus(t *testing.T) {
	t.Parallel()

	cases := []struct {
		status int
		name   string
	}{
		{http.StatusBadRequest, "bad_request"},
		{http.StatusUnauthorized, "auth_failure"},
		{http.StatusForbidden, "forbidden"},
		{http.StatusNotFound, "not_found"},
		{http.StatusConflict, "invalid_state"},
		{http.StatusRequestEntityTooLarge, "request_too_large"},
		{http.StatusTooManyRequests, "too_many_requests"},
		{http.StatusInternalServerError, "failed_request"},
		{http.StatusGatewayTimeout, "timeout"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d_%s", tc.status, tc.name), func(t *testing.T) {
			t.Parallel()

			payload := fmt.Sprintf(
				`{"name":%q,"id":"req-1","message":"boom","temporary":false,"timeout":%t,"fault":%t}`,
				tc.name, tc.status == http.StatusGatewayTimeout, tc.status == http.StatusInternalServerError)

			client, _ := newValidationServer(t, func(w http.ResponseWriter) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tc.status)
				_, err := io.WriteString(w, payload)
				require.NoError(t, err)
			})

			req := v1.NewValidateServiceSpecRequest2("Zm9v", "svc", "compose")
			result, httpRes, err := client.ServiceApiAPI.
				ServiceApiValidateServiceSpec(context.Background()).
				ValidateServiceSpecRequest2(*req).
				Execute()

			require.Error(t, err, "HTTP %d must not be reported as success", tc.status)
			require.NotNil(t, httpRes)
			assert.Equal(t, tc.status, httpRes.StatusCode,
				"the caller must be able to distinguish this status from the others")
			assert.Nil(t, result)

			apiErr := &v1.GenericOpenAPIError{}
			require.ErrorAs(t, err, &apiErr)
			decoded, ok := apiErr.Model().(v1.Error)
			require.True(t, ok, "HTTP %d must decode into the typed error envelope", tc.status)
			assert.Equal(t, tc.name, decoded.GetName())
		})
	}
}

// The three required fields are required on decode too, so a truncated or hand-built body is
// rejected where it is cheap to notice rather than deep inside a handler.
func TestValidateServiceSpecRequest_RequiredFieldsAreEnforcedOnDecode(t *testing.T) {
	t.Parallel()

	full := `{"name":"svc","specType":"compose","fileContent":"Zm9v"}`

	var ok v1.ValidateServiceSpecRequest2
	require.NoError(t, json.Unmarshal([]byte(full), &ok))
	assert.Equal(t, "compose", ok.GetSpecType())

	for _, missing := range []string{
		`{"specType":"compose","fileContent":"Zm9v"}`,
		`{"name":"svc","fileContent":"Zm9v"}`,
		`{"name":"svc","specType":"compose"}`,
	} {
		var into v1.ValidateServiceSpecRequest2
		assert.Error(t, json.Unmarshal([]byte(missing), &into),
			"a request missing a required field must not decode: %s", missing)
	}
}

// Same rule on the response: a result without its arrays is not a result the client can act on.
func TestValidateServiceSpecResult_RequiredFieldsAreEnforcedOnDecode(t *testing.T) {
	t.Parallel()

	var ok v1.ValidateServiceSpecResult
	require.NoError(t, json.Unmarshal([]byte(minimalResult), &ok))
	assert.Equal(t, "VALID", ok.GetStatus())

	for _, missing := range []string{
		`{"validationVersion":"1","inputDigest":"d","checks":[],"diagnostics":[],"requiredArtifacts":[],"validatedArtifacts":[]}`,
		`{"status":"VALID","validationVersion":"1","inputDigest":"d","diagnostics":[],"requiredArtifacts":[],"validatedArtifacts":[]}`,
		`{"status":"VALID","validationVersion":"1","inputDigest":"d","checks":[],"diagnostics":[],"requiredArtifacts":[]}`,
	} {
		var into v1.ValidateServiceSpecResult
		assert.Error(t, json.Unmarshal([]byte(missing), &into),
			"a result missing a required field must not decode: %s", missing)
	}
}

// All five artifact fields travel together. Accepting four of them would mean accepting content the
// server has no independent way to check.
func TestValidationArtifactInput_RequiredFieldsAreEnforcedOnDecode(t *testing.T) {
	t.Parallel()

	full := `{"logicalPath":"terraform/network","encoding":"tar+gzip+base64","archiveContent":"H4sI",` +
		`"sha256":"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855","compressedSizeBytes":1234}`

	var ok v1.ValidationArtifactInput
	require.NoError(t, json.Unmarshal([]byte(full), &ok))
	assert.Equal(t, "tar+gzip+base64", ok.GetEncoding())
	assert.Equal(t, int64(1234), ok.GetCompressedSizeBytes())

	for _, field := range []string{"logicalPath", "encoding", "archiveContent", "sha256", "compressedSizeBytes"} {
		var generic map[string]any
		require.NoError(t, json.Unmarshal([]byte(full), &generic))
		delete(generic, field)

		partial, err := json.Marshal(generic)
		require.NoError(t, err)

		var into v1.ValidationArtifactInput
		assert.Error(t, json.Unmarshal(partial, &into),
			"an artifact without %s must not decode", field)
	}
}

// The legacy build request is untouched by this change: same required fields, same dryrun flag, so
// existing callers keep working while the adapter changes what the server does with them.
func TestLegacyBuildRequest_ContractUnchanged(t *testing.T) {
	t.Parallel()

	var into v1.BuildServiceFromServicePlanSpecRequest2
	require.NoError(t, json.Unmarshal([]byte(`{"name":"svc","fileContent":"Zm9v"}`), &into))
	assert.Equal(t, "svc", into.GetName())
	_, hasDryrun := into.GetDryrunOk()
	assert.False(t, hasDryrun, "dryrun stays optional")

	into.SetDryrun(false)
	body, err := json.Marshal(into)
	require.NoError(t, err)
	assert.Contains(t, string(body), `"dryrun":false`,
		"an explicit dryrun=false must reach the server rather than being dropped")
}
