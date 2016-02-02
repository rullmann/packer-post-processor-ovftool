package main

import (
	"github.com/rozha/packer-post-processor-ovftool/ovftool"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(ovftool.OVFPostProcessor))
	server.Serve()
}
