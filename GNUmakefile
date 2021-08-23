default: testacc

bin = terraform-provider-kamaz

# Run acceptance tests
.PHONY: testacc local build fmt fmtcheck

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -w -s ./internal/provider

fmtcheck:
	@echo "==> Checking source code against gofmt..."
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

testacc: fmtcheck
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

build: fmtcheck
	go build -o $(bin)

generate: build
	go generate  ./...

local: build
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/hashicorp/kamaz/0.1.0/darwin_amd64
	mv $(bin) ~/.terraform.d/plugins/registry.terraform.io/hashicorp/kamaz/0.1.0/darwin_amd64

