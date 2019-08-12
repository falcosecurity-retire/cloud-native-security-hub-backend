.PHONY: test build push

test:
	go test -v ./...

watch:
	ag -l | entr -c go test -v ./...

build:
	docker build -t gcr.io/mateo-burillo-ns/cnvh-backend .

push: build
	docker push gcr.io/mateo-burillo-ns/cnvh-backend
