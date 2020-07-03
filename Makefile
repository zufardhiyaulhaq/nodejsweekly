REPOSITORY?=
TAG?=

build:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o nodeweekly cmd/nodeweekly/*.go 
	docker build -t ${REPOSITORY}:${TAG} .
	rm nodeweekly

