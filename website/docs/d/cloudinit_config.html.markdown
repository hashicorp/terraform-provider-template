---
layout: "template"
page_title: "Template: cloudinit_multipart"
sidebar_current: "docs-template-datasource-cloudinit-config"
description: |-
  Renders a multi-part cloud-init config from source files.
---

# template_cloudinit_config

Renders a [multipart MIME configuration](https://cloudinit.readthedocs.io/en/latest/topics/format.html#mime-multi-part-archive)
for use with [Cloud-init](https://cloudinit.readthedocs.io/).

Cloud-init is a commonly-used startup configuration utility for cloud compute
instances. It accepts configuration via provider-specific user data mechanisms,
such as `user_data` for Amazon EC2 instances. Multipart MIME is one of the
data formats it accepts. For more information, see
[User-Data Formats](https://cloudinit.readthedocs.io/en/latest/topics/format.html)
in the Cloud-init manual.

This is not a generalized utility for producing multipart MIME messages. Its
featureset is specialized for the features of cloud-init.

## Example Usage

```hcl
# Render a part using a `template_file`
data "template_file" "script" {
  template = "${file("${path.module}/init.tpl")}"

  vars = {
    consul_address = "${aws_instance.consul.private_ip}"
  }
}

# Render a multi-part cloud-init config making use of the part
# above, and other source files
data "template_cloudinit_config" "config" {
  gzip          = true
  base64_encode = true

  # Main cloud-config configuration file.
  part {
    filename     = "init.cfg"
    content_type = "text/cloud-config"
    content      = "${data.template_file.script.rendered}"
  }

  part {
    content_type = "text/x-shellscript"
    content      = "baz"
  }

  part {
    content_type = "text/x-shellscript"
    content      = "ffbaz"
  }
}

# Start an AWS instance with the cloud-init config as user data
resource "aws_instance" "web" {
  ami              = "ami-d05e75b8"
  instance_type    = "t2.micro"
  user_data_base64 = "${data.template_cloudinit_config.config.rendered}"
}
```

## Argument Reference

The following arguments are supported:

* `gzip` - (Optional) Specify whether or not to gzip the rendered output. Defaults to `true`.

* `base64_encode` - (Optional) Base64 encoding of the rendered output. Defaults to `true`,
  and cannot be disabled if `gzip` is `true`.

* `part` - (Required) A nested block type which adds a file to the generated
  cloud-init configuration. Use multiple `part` blocks to specify multiple
  files, which will be included in order of declaration in the final MIME
  document.

Each `part` block expects the following arguments:

* `content` - (Required) Body content for the part.

* `filename` - (Optional) A filename to report in the header for the part.

* `content_type` - (Optional) A MIME-style content type to report in the header for the part.

* `merge_type` - (Optional) A value for the `X-Merge-Type` header of the part,
  to control [cloud-init merging behavior](https://cloudinit.readthedocs.io/en/latest/topics/merging.html).

## Attributes Reference

The following attributes are exported:

* `rendered` - The final rendered multi-part cloud-init config.
