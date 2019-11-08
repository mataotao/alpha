server = "alpha"
.IGNORE:love vet
all: clean gotool
	@GO111MODULE=on GOPROXY=https://goproxy.io go build -o ${server} -v .
clean:
	@rm -f ${server}
#	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}
gotool:
	@gofmt -w .
	@GO111MODULE=on go vet .
love:
	@GO111MODULE=on GOPROXY=https://goproxy.io go run -race main.go
push: vet
	@gofmt -w .
	@git commit -m'$(m)'
	@git push origin $(b)
execute:
	@chmod +x ${server}
vet:
	@GO111MODULE=on golangci-lint run
.PHONY: clean gotool
