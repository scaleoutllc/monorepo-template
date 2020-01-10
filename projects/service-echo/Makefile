NDEF = $(if $(value $(1)),,$(error $(1) not set))
PROJECT_FOLDER ?= $(notdir $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST))))))
PROJECT_META ?= $(shell spruce json manifest.yml)
PROJECT_FILES ?= $(shell echo '$(PROJECT_META)' | jq -r '.project.triggers.build[]' | tr '\n' ' ')
PROJECT_IMAGE ?= $(shell echo '$(PROJECT_META)' | jq -r '.project.repository.host + "/" + .project.name')
DEPLOYMENTS_PATH ?= "infrastructure/$(if $(SAVE_DEPLOYMENT_NAME),$(SAVE_DEPLOYMENT_NAME),$(DEPLOYMENT_NAME))"
VERSION_FILE_PATH ?= "$(DEPLOYMENTS_PATH)/version"
VERSION_FILE_CONTENT ?= $(shell cat $(VERSION_FILE_PATH) 2> /dev/null)
VERSION_TAG ?= $(if $(TAG),$(TAG),$(if $(VERSION_FILE_CONTENT),$(VERSION_FILE_CONTENT),$(CURRENT_VERSION)))
CURRENT_VERSION ?= $(shell ../../bin/versions $(PROJECT_FILES) | head -n1)
VERSION_PACKAGE = $(shell go list -m)/shared/components/go/version

export CGO_ENABLED = 0

all: lint fmt vet test build

fmt:
	go fmt ./...

lint:
	golint ./...

vet:
	go vet ./...

test:
	go test -v -short ./...

build:
	go generate ./...
	go build -o dist/main -ldflags "-X $(VERSION_PACKAGE).NUMBER=$(CURRENT_VERSION) -X $(VERSION_PACKAGE).COMMIT=$(CURRENT_VERSION)" ./src

run:
	go run ./...

clean:
	rm -rf dist .cache

# Get the latest commit that changed files we consider to be part of the version
# of this project. If SAVE_DEPLOYMENT_NAME is set, save the version retreived
# to a file.
version:
	@>&2 printf $(if $(SAVE_DEPLOYMENT_NAME),"Saving version to $(VERSION_FILE_PATH).\n","")
	@echo $(CURRENT_VERSION) | tee $(if $(SAVE_DEPLOYMENT_NAME),$(VERSION_FILE_PATH),"/dev/null")

# Render a deployable manifest. Tag container based on the following criteria:
# 1. If TAG is explictly passed, use it.
# 2. If no TAG is provided, try to use version file in the deployment directory.
# 3. If no TAG is provided and no version file is found, calculate latest SHA.
manifest:
	$(call NDEF,DEPLOYMENT_NAME)
	@>&2 echo $(if $(TAG),"Using explicitly provided TAG ($(TAG)).",$(if $(VERSION_FILE_CONTENT),"Using version file for TAG ($(VERSION_FILE_CONTENT)).","Using current valid commit sha for TAG ($(CURRENT_VERSION))."))
	@IMAGE=$(PROJECT_IMAGE) TAG=$(VERSION_TAG) render-manifest $(DEPLOYMENTS_PATH)

# List which environments this should run in.
environments:
	@find infrastructure -mindepth 1 -maxdepth 1 -type d -exec basename {} \;

# Show changes since the last (chosen) deployment.
changelog:
	$(call NDEF,DEPLOYMENT_NAME)
ifeq ($(strip $(VERSION_FILE_CONTENT)),)
	@echo "No deployment version file found for $(DEPLOYMENT_NAME)."
	@exit 1
endif
	$(call NDEF,VERSION_FILE_CONTENT)
	@../../bin/changes $(VERSION_FILE_CONTENT) $(PROJECT_FILES)

.PHONY: all fmt lint vet test build run clean version manifest environments changelog
