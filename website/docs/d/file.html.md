---
layout: "template"
page_title: "Template: template_file"
sidebar_current: "docs-template-datasource-file"
description: |-
  Renders a template from a file.
---

# template_file

The `template_file` data source renders a template from a template string,
which is usually loaded from an external file.

~> **Note** In Terraform 0.12 and later,
[the `templatefile` function](/docs/configuration/functions/templatefile.html)
offers a built-in mechanism for rendering a template from a file. Use that
function instead, unless you are using Terraform 0.11 or earlier.

## Example Usage

```hcl
data "template_file" "init" {
  template = "${file("${path.module}/init.tpl")}"
  vars = {
    consul_address = "${aws_instance.consul.private_ip}"
  }
}
```

Inside `init.tpl` you can include the value of `consul_address`. For example:

```bash
#!/bin/bash

echo "CONSUL_ADDRESS = ${consul_address}" > /tmp/iplist
```

Although in principle `template_file` can be used with an inline template
string, we don't recommend this approach because it requires awkward escaping.
Instead, just use [template syntax](/docs/configuration/expressions.html#string-templates)
directly in the configuration. For example:

```hcl
  user_data = <<-EOT
    echo "CONSUL_ADDRESS = ${aws_instance.consul.private_ip}" > /tmp/iplist
  EOT
```

## Argument Reference

The following arguments are supported:

* `template` - (Required) The contents of the template, as a string using
  [Terraform template syntax](/docs/configuration/expressions.html#string-templates).
  Use [the `file` function](/docs/configuration/functions/file.html) to load
  the template source from a separate file on disk.

* `vars` - (Optional) Variables for interpolation within the template. Note
  that variables must all be primitives. Direct references to lists or maps
  will cause a validation error.

Earlier versions of `template_file` accepted another argument `filename` as
an alternative to `template`. This has now been removed. Use the `template`
argument with the `file` function to get the same effect.

## Template Syntax

The `template` argument is processed as
[Terraform template syntax](/docs/configuration/expressions.html#string-templates).

However, this provider has its own copy of the template engine embedded in it,
separate from Terraform itself, and so which features are available are decided
based on what Terraform version the provider was compiled against, and not
on which Terraform version you are running.

For more consistent results, Terraform 0.12 has a built in function
[`templatefile`](/docs/configuration/functions/templatefile.html) which serves
the same purpose as this data source. Use that function instead if you are
using Terraform 0.12 or later. Its template and expression capabilities will
always match the version of Terraform you are using.

## Attributes Reference

The following attributes are exported:

* `template` - See Argument Reference above.
* `vars` - See Argument Reference above.
* `rendered` - The final rendered template.
