SHELL_PATH = /bin/bash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# Define dependencies

GOLANG          := golang:1.20
ALPINE          := alpine:3.18
KIND            := kindest/node:v1.27.3

KIND_CLUSTER    := time-tracker-cluster
NAMESPACE       := default
APP             := time-tracker
BASE_IMAGE_NAME := github.com/mahdi-mk
SERVICE_NAME    := time-tracker-api
VERSION         := 0.0.1
SERVICE_IMAGE   := $(BASE_IMAGE_NAME)/$(SERVICE_NAME):$(VERSION)


# ==============================================================================
# Building containers

all: build

build:
	docker build \
		-f zarf/docker/dockerfile \
		-t $(SERVICE_IMAGE) \
		.

# ==============================================================================
# Running from within k8s/kind

kind-up:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/kind-config.yaml

kind-down:
	kind delete cluster --name $(KIND_CLUSTER)

# ------------------------------------------------------------------------------

kind-load:
	kind load docker-image $(SERVICE_IMAGE) --name $(KIND_CLUSTER)

kind-apply:
	kubectl apply -f zarf/k8s/database/namespace.yaml
	kubectl apply -f zarf/k8s/database/deployment.yaml
	kubectl apply -f zarf/k8s/database/service.yaml

	kubectl wait --namespace=database-system --timeout=120s --for=condition=Available deployment/database

	kubectl apply -f zarf/k8s/time-tracker/deployment.yaml
	kubectl apply -f zarf/k8s/time-tracker/service.yaml

kind-restart:
	kubectl rollout restart deployment $(APP) --namespace=$(NAMESPACE)

kind-update: all kind-load kind-restart

kind-update-apply: all kind-load kind-apply kind-restart

# ------------------------------------------------------------------------------

kind-logs:
	kubectl logs -l app=$(APP) --namespace=$(NAMESPACE) --all-containers=true -f --tail=100

kind-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide --all-namespaces
	kubectl get pods -o wide --watch --all-namespaces

kind-describe:
	kubectl describe nodes
	kubectl describe svc

kind-describe-deployment:
	kubectl describe deployment --namespace=$(NAMESPACE) $(APP)

kind-describe-sales:
	kubectl describe pod --namespace=$(NAMESPACE) -l app=$(APP)

# ------------------------------------------------------------------------------
