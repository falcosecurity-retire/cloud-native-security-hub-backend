.PHONY: test build push

test: generate
	#go test -v ./...
	CGO_ENABLED=1 ginkgo -r -randomizeSuites -race

generate:
	go generate -x ./...

dev: generate
	RESOURCES_PATH=test/fixtures/resources VENDOR_PATH=test/fixtures/vendors go run cmd/server/main.go

watch: generate
	ag -l | entr -c ginkgo -r -randomizeSuites

build: generate
	docker build -f Dockerfile.server -t gcr.io/mateo-burillo-ns/securityhub-backend .
	docker build -f Dockerfile.dbimport -t gcr.io/mateo-burillo-ns/securityhub-dbimport .

push: build
	docker push gcr.io/mateo-burillo-ns/securityhub-backend
	docker push gcr.io/mateo-burillo-ns/securityhub-dbimport

deploy: deploy-backend deploy-frontend

deploy-backend:
	kubectl -n securityhub patch deployment backend -p "{\"spec\": {\"template\": {\"metadata\": { \"labels\": {  \"redeploy\": \"$(shell date +%s)\"}}}}}"

deploy-frontend:
	kubectl -n securityhub patch deployment frontend -p "{\"spec\": {\"template\": {\"metadata\": { \"labels\": {  \"redeploy\": \"$(shell date +%s)\"}}}}}"
