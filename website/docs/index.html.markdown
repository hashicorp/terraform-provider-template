---
layout: "template"
page_title: "Provider: Template"
sidebar_current: "docs-template-index"
description: |-
  The Template provider is used to template strings for other Terraform resources.
---

# Template Provider

~> **This provider is deprecated.** We now recommend that you use one of the approaches in the Deprecation section below.

The template provider exposes data sources to use templates to generate
strings for other Terraform resources or outputs.

Use the navigation to the left to read about the available data sources.

## Example Usage

```hcl
# Template for initial configuration bash script
data "template_file" "init" {
  template = "${file("init.tpl")}"
  vars = {
    consul_address = "${aws_instance.consul.private_ip}"
  }
}

# Create a web server
resource "aws_instance" "web" {
  # ...

  user_data = "${data.template_file.init.rendered}"
}
```

## Deprecation

The template provider is deprecated and the provider has been archived in accordance with HashiCorp's [provider archiving process](https://terraform.io/docs/internals/archiving.html). While released versions of the provider will remain available, we recommend that you replace usages of this provider as follows.

### `template_file`

For Terraform 0.12 and later, the `template_file` data source has been superseded by [the `templatefile` function](/docs/configuration/functions/templatefile.html), which can be used directly in expressions without creating a separate data resource.

### `template_dir`

The [`hashicorp/dir/template`](https://registry.terraform.io/modules/hashicorp/dir/template) module offers an improved version of the functionality available in `template_dir`.

### `template_cloudinit_config`

This resource has been moved to a new provider, [`terraform-provider-cloudinit`](https://github.com/hashicorp/terraform-provider-cloudinit) as `cloudinit_config`, with no change in functionality.
