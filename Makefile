.PHONY: all docs speakeasy check-speakeasy
all: speakeasy docs

docs:
	go generate ./...

openapi.yaml:
	curl https://raw.githubusercontent.com/mikevdberge/snowmirror-api/main/reference/snowmirror-api.yaml > openapi.yaml

speakeasy: check-speakeasy openapi.yaml
	speakeasy generate sdk --lang terraform -o . -s openapi.yaml

check-speakeasy:
	@command -v speakeasy >/dev/null 2>&1 || { echo >&2 "speakeasy CLI is not installed. Please install before continuing."; exit 1; }