PACKAGE_NAME=omnistrategosdk
PACKAGE_VERSION=0.0.1
ORG_NAME=omnistrate
REPO_NAME=omnistrate-sdk-go
OPEN_API_SPEC?=https://api.omnistrate.cloud/2022-09-01-00/openapi.yaml
FLEET_OPEN_API_SPEC?=https://api.omnistrate.cloud/2022-09-01-00/fleet/openapi.yaml

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
gen-go-sdk:
	echo "Generate Go SDK"
	openapi-generator generate \
  -i ${OPEN_API_SPEC} \
  -g go \
  -o v1 \
  -t templates/go \
  --additional-properties packageName=v1,packageVersion=${PACKAGE_VERSION},withGoMod=false,isGoSubmodule=true,generateInterfaces=true \
  --git-user-id ${ORG_NAME} \
  --git-repo-id ${REPO_NAME}

.PHONY: gen-fleet-go-sdk
gen-fleet-go-sdk:
	echo "Generate Fleet Go SDK"
	openapi-generator generate \
  -i ${FLEET_OPEN_API_SPEC} \
  -g go \
  -o fleet \
  --additional-properties packageName=fleet,packageVersion=${PACKAGE_VERSION},withGoMod=false,isGoSubmodule=true,generateInterfaces=true \
  --git-user-id ${ORG_NAME} \
  --git-repo-id ${REPO_NAME}

.PHONY: validate
validate:
	echo "Verify Go OpenAPI specs"
	openapi-generator validate -i ${OPEN_API_SPEC}
	openapi-generator validate -i ${FLEET_OPEN_API_SPEC}

.PHONY: tidy
tidy:
	echo "Tidy dependency modules"
	go mod tidy

.PHONY: download
download:
	echo "Download dependency modules"
	go mod download

.PHONY: go-template
go-template:
	echo "Generate meta template"
	openapi-generator meta -l go