package main

import (
	"github.com/mitchellh/packer/packer/plugin"
	"github.com/june07/packer-builder-xenserver/builder/xenserver/template"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(template.Builder))
	server.Serve()
}
