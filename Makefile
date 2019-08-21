.PHONY: test build push

test:
	go test -v ./...

dev:
	RESOURCES_PATH=test/fixtures/resources VENDOR_PATH=test/fixtures/vendors go run cmd/server/main.go

watch:
	ag -l | entr -c go test -v ./...

build:
	docker build -t gcr.io/mateo-burillo-ns/cnvh-backend .

push: build
	docker push gcr.io/mateo-burillo-ns/cnvh-backend
