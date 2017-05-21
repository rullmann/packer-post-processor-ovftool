package main

import (
	"github.com/rullmann/packer-post-processor-ovftool/ovftool"
	"github.com/hashicorp/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(ovftool.OVFPostProcessor))
	server.Serve()
}
