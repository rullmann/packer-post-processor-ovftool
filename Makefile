.PHONY: .depend \
        all \
        clean \
        help \
        release 

PROJECT = packer-post-processor-ovftool
VERSION = 0.0.1

all: $(PROJECT)

$(PROJECT): .depend clean
	go build -v .

help:
	@echo "clean     remove testing artifacts"
	@echo "dist      cross-compile binaries for distribution"
	@echo "help      show this page"

.depend:
	go get -d github.com/hashicorp/packer

clean:
	go clean -x
	rm -rf dist

dist:
	gox --osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/$(PROJECT)-$(VERSION)-{{ .OS }}_{{ .Arch }}"
