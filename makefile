.PHONY: ring test doc

ring: test doc

test:
	go fmt
	go test -v

doc:
	PATH=$$PATH:$$GOPATH/bin godocdown > README.md
