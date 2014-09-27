.PHONY: ring test doc

ring: test doc

test:
	go test -v

doc:
	godocdown > README.md
