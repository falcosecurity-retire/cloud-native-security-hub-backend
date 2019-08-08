.PHONY: test build

test:
	go test -v ./...

watch:
	ag -l | entr -c go test -v ./...

build:
	docker build -t sysdiglabs/cloud-native-visibility-hub-server .
