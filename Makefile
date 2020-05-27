all: build clean

ARCH?=amd64
VERSION?=1.0.0

build: clean
		GOARCH=$(ARCH) CGO_ENABLED=0 go build  cmd/main.go -o k8s-version-reminder   github.com/AliyunContainerService/k8s-version-reminder
clean:
		rm -f k8s-version-reminder