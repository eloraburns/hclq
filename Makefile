default:
	docker run --rm -v $(shell pwd):/usr/src/hclq -w /usr/src/hclq golang:1.10 make go

go:
	./build.sh
