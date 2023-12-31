# get the repo root and output path
REPO_ROOT:=$(shell pwd)
export REPO_ROOT
OUT_DIR=$(REPO_ROOT)/bin
# record the source commit in the binary, overridable
COMMIT?=$(shell git describe --tags --always --dirty 2>/dev/null)
INSTALL?=install
GOPATH := $(shell go env GOPATH)
# make install will place binaries here
INSTALL_DIR?=$(GOPATH)/bin
# the output binary name, overridden when cross compiling
BINARY_NAME?=kubetest2-eks
BINARY_PATH?=.
# the container cli to use e.g. docker,podman
DOCKER?=$(shell which docker || which podman || echo "docker")
export DOCKER
# ========================= Setup Go With Gimme ================================
# go version to use for build etc.
# setup correct go version with gimme
# PATH:=$(shell . hack/build/setup-go.sh && echo "$${PATH}")
# go1.9+ can autodetect GOROOT, but if some other tool sets it ...
GOROOT:=
# enable modules
GO111MODULE=on
# disable CGO by default for static binaries
CGO_ENABLED=0
export PATH GOROOT GO111MODULE CGO_ENABLED
# work around broken PATH export
SPACE:=$(subst ,, )
SHELL:=env PATH=$(subst $(SPACE),\$(SPACE),$(PATH)) $(SHELL)
BUILD_FLAGS ?= -trimpath -ldflags="-buildid= -X=github.com/aws/aws-k8s-tester/kubetest2/internal.Version=$(COMMIT)"

.PHONY: deployers
deployers: deployer-eksctl deployer-eksapi

deployer-eksctl:
	CGO_ENABLED=1 go build -v $(BUILD_FLAGS) -o $(OUT_DIR)/kubetest2-eksctl kubetest2-eksctl/main.go

deployer-eksapi:
	CGO_ENABLED=1 go build -v $(BUILD_FLAGS) -o $(OUT_DIR)/kubetest2-eksapi kubetest2-eksapi/main.go

.PHONY: testers
testers: tester-multi tester-ginkgo-v1

tester-multi:
	GGO_ENABLED=1 go build -v $(BUILD_FLAGS) -o $(OUT_DIR)/kubetest2-tester-multi kubetest2-tester-multi/main.go

tester-ginkgo-v1:
	GGO_ENABLED=1 go build -v $(BUILD_FLAGS) -o $(OUT_DIR)/kubetest2-tester-ginkgo-v1 kubetest2-tester-ginkgo-v1/main.go

.PHONY: install
install: deployers testers
	$(INSTALL) -d $(INSTALL_DIR)
	$(INSTALL) $(OUT_DIR)/* $(INSTALL_DIR)/

.PHONY: clean
clean:
	rm -rf $(OUT_DIR)
