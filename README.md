Packer ovftool post-processor
=============================

Exposes VMware's [OVF Tool](https://www.vmware.com/support/developer/ovf/) as
a [Packer](http://www.packer.io) post-processor, enabling VMware builds to
produce OVA/OVF artifacts.

For the sake of simplicity, only a few of ovftool's options are currently
exposed, but adding more would be a simple task. Contributions are welcome.

Usage
-----
Add the post-processor to your packer template:

    {
        "post-processors": [{
            "type": "ovftool",
            "only": ["vmware"],
            "format": "ova"
        }]
    }

Available configuration options:

* `target`: The path where the artifact should be created, without the file 
  extension. This is a [configuration template](http://www.packer.io/docs/templates/configuration-templates.html). 
  Defaults to `packer_{{.BuildName}}_{{.Provider}}`.
* `format`:      The target type to create, either "ova" or "ovf". Defaults 
  to "ovf" if not specified.
* `compression`: The compression level to use when creating the artifact. A 
  number from 1-9. Default value is 0, meaning no compression.


Installation
------------
Run:

    $ go get github.com/rullmann/packer-post-processor-ovftool
    $ go install github.com/rullmann/packer-post-processor-ovftool

Copy the binary to `~/.packer.d/plugins/`:

    $ cp ${GOPATH}/bin/packer-post-processor-ovftool ~/.packer.d/plugins/
