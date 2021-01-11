export GO111MODULE=on
version = 99.0.0

fmt:
	go fmt ./...

build: fmt
	go build -o terraform-provider-twilio

install_macos: build
	mkdir -p ~/.terraform.d/plugins/local/mpon/twilio/$(version)/darwin_amd64
	cp terraform-provider-twilio ~/.terraform.d/plugins/local/mpon/twilio/$(version)/darwin_amd64/terraform-provider-twilio_v$(version)

test: fmt
	go test -v `go list ./... | tail -n +2`

plan: install_macos
	terraform init
	terraform fmt
	terraform plan

apply: plan
	terraform apply

.PHONY: fmt build test plan apply
