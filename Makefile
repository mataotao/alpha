server = "alpha"
.IGNORE:love
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
push:
	@gofmt -w .
	@git commit -m'$(m)'
	@git push origin $(b)
execute:
	@chmod +x ${server}
.PHONY: clean gotool
