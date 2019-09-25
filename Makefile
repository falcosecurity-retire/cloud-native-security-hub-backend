.PHONY: test build push

test:
	go test -v ./...

dev:
	RESOURCES_PATH=test/fixtures/resources VENDOR_PATH=test/fixtures/vendors go run cmd/server/main.go

watch:
	ag -l | entr -c go test -v ./...

build:
	docker build -t gcr.io/mateo-burillo-ns/securityhub-backend .

push: build
	docker push gcr.io/mateo-burillo-ns/securityhub-backend

deploy: deploy-backend deploy-frontend

deploy-backend:
	kubectl -n securityhub patch deployment securityhub-backend -p "{\"spec\": {\"template\": {\"metadata\": { \"labels\": {  \"redeploy\": \"$(shell date +%s)\"}}}}}"

deploy-frontend:
	kubectl -n securityhub patch deployment securityhub-frontend -p "{\"spec\": {\"template\": {\"metadata\": { \"labels\": {  \"redeploy\": \"$(shell date +%s)\"}}}}}"
