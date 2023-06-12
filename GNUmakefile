default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

.PHONY: swagger
swagger:
	@swagger generate client -f swagger/public.swagger.json -t internal

.PHONY: tools
tools:
	@go install github.com/go-swagger/go-swagger/cmd/swagger@latest