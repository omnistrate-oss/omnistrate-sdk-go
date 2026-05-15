ORG_NAME=omnistrate-oss
REPO_NAME=omnistrate-sdk-go
OPEN_API_SPEC?=https://api.omnistrate.cloud/2022-09-01-00/openapi.yaml
FLEET_OPEN_API_SPEC?=https://api.omnistrate.cloud/2022-09-01-00/fleet/openapi.yaml
OPENAPI_GENERATOR_VERSION=7.22.0
OPENAPI_GENERATOR_DIR=$(CURDIR)/.openapi-generator
OPENAPI_GENERATOR_JAR=$(OPENAPI_GENERATOR_DIR)/openapi-generator-cli-$(OPENAPI_GENERATOR_VERSION).jar
OPENAPI_GENERATOR_DOWNLOAD_URL=https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$(OPENAPI_GENERATOR_VERSION)/openapi-generator-cli-$(OPENAPI_GENERATOR_VERSION).jar
OPENAPI_GENERATOR=java -jar $(OPENAPI_GENERATOR_JAR)

.DEFAULT_GOAL := all

$(OPENAPI_GENERATOR_JAR):
	mkdir -p $(OPENAPI_GENERATOR_DIR)
	curl -fsSL --retry 3 -o $@.tmp $(OPENAPI_GENERATOR_DOWNLOAD_URL)
	mv $@.tmp $@

.PHONY: openapi-generator-version
openapi-generator-version: $(OPENAPI_GENERATOR_JAR)
	$(OPENAPI_GENERATOR) version

.PHONY: all
all: clean gen-go-sdk gen-fleet-go-sdk tidy build

.PHONY: build
build:
	echo "Build Go SDK"
	go build -v ./...

.PHONY: clean
clean:
	echo "Clean Go SDK"
	rm -rf fleet
	rm -rf v1

.PHONY: gen-go-sdk
gen-go-sdk: $(OPENAPI_GENERATOR_JAR)
	echo "Generate Go SDK"
	_JAVA_OPTIONS=-DmaxYamlCodePoints=99999999 $(OPENAPI_GENERATOR) generate \
	  -i ${OPEN_API_SPEC} \
	  -g go \
	  -o v1 \
  -t templates/go \
  --additional-properties packageName=v1,withGoMod=false,isGoSubmodule=true,generateInterfaces=true,disallowAdditionalPropertiesIfNotPresent=false,structPrefix=false \
  --git-user-id ${ORG_NAME} \
  --git-repo-id ${REPO_NAME}

.PHONY: gen-fleet-go-sdk
gen-fleet-go-sdk: $(OPENAPI_GENERATOR_JAR)
	echo "Generate Fleet Go SDK"
	_JAVA_OPTIONS=-DmaxYamlCodePoints=99999999 $(OPENAPI_GENERATOR) generate \
	  -i ${FLEET_OPEN_API_SPEC} \
	  -g go \
	  -o fleet \
  --additional-properties packageName=fleet,withGoMod=false,isGoSubmodule=true,generateInterfaces=true,disallowAdditionalPropertiesIfNotPresent=false,structPrefix=false \
  --git-user-id ${ORG_NAME} \
  --git-repo-id ${REPO_NAME}

.PHONY: validate
validate: $(OPENAPI_GENERATOR_JAR)
	echo "Verify Go OpenAPI specs"
	$(OPENAPI_GENERATOR) validate -i ${OPEN_API_SPEC}
	$(OPENAPI_GENERATOR) validate -i ${FLEET_OPEN_API_SPEC}

.PHONY: tidy
tidy:
	echo "Tidy dependency modules"
	go mod tidy

.PHONY: download
download:
	echo "Download dependency modules"
	go mod download

.PHONY: go-template
go-template: $(OPENAPI_GENERATOR_JAR)
	echo "Generate meta template"
	$(OPENAPI_GENERATOR) meta -l go
