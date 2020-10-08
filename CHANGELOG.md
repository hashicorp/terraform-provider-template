## 2.2.0 (October 08, 2020)

* Ship arm64 binary

## 2.1.2 (May 01, 2019)

* This release includes another upgrade of the Terraform SDK, primarily for the purposes of aligning versions with other providers prior to the v0.12.0 release. The changes in the SDK should not affect the behavior of this particular provider.

## 2.1.1 (April 11, 2019)

OTHER:

* This release includes an upgrade of the Terraform SDK to the latest version, primarily for the purposes of aligning versions with other providers prior to the v0.12.0 release. The changes in the SDK should not affect the behavior of this particular provider.

## 2.1.0 (February 28, 2019)

IMPROVEMENTS:

* The previous release contains an SDK incompatible with TF 0.12. Fortunately 0.12 was not released yet so upgrading the vendored sdk makes this release compatible with 0.12.

## 2.0.0 (January 14, 2019)

UPGRADE COMPATIBILITY NOTES:

* Template provider v2.0.0 switches template implementations from HIL to HCL 2.0 templates and uses the set of functions from Terraform v0.12. The new version is broadly compatible for most straightforward template usage, but there are some particular differences to watch out for:
  * HCL 2.0 template syntax includes a new `%{ ... }` construct for control structures, like `%{ if a == b } ... %{ endif }`. If your existing templates contain any `%{` sequences you will need to now escape them as `%%{` to ensure correct parsing.
  * The `format` function from Terraform is now a new implementation using the HCL type system rather than based on the Go `fmt` package. The basic features are still compatible, but there are some differences in the handling of some more complex verbs, like `%#v` which now uses JSON syntax instead of Go-like syntax.
  * HIL apparently incorrectly treated `$$` (without a following `{` as an escape for `$` even though such escaping was not actually required, while the new parser correctly only deals with `$${` as an escape for `${`. This issue was unfortunately not known at the time of the 2.0.0 release. If you have a template that was using a sequence like `$$FOO` as an escaped form of `$FOO`, remove the extra dollar sign to get the expected interpretation. Dollar sign escaping is only needed as part of escaping the full `${` sequence, not dollar signs in isolation.
* The `filename` argument is no longer supported in `template_file`. This argument was deprecated long before this provider was split into its own codebase. Replace uses of `filename = "foo"` with `template = file("foo")` to get a similar result.
* In `template_cloudinit_config`, we now require gzip output to be base64 encoded, because Terraform Core expects all strings to be valid UTF-8 and would previously silently corrupt raw gzip output in some cases. Base64 encoding is the standard way to represent small binary payloads in Terraform configuration, and so other providers that expect binary data should have mechanisms to accept Base64 data. For example, in `aws_instance` you can use `user_data_base64` instead of `user_data`.
* Please note that Terraform v0.12 now has a built-in function `templatefile` which behaves in much the same way as the `template_file` data source. You may wish to switch to using the function rather than the data source after upgrading to Terraform v0.12.

IMPROVEMENTS:

* The provider is now compatible with Terraform v0.12, while retaining compatibility with prior versions.
* The template language now supports simple control structures. The `for` construct allows repetition and the `if` construct allows conditional output. For example:
  ```
  %{ for addr in ip_addrs ~}
  backend ${addr}
  %{ endfor ~}
  ```
  This is the same syntax used for string templates within the Terraform language itself, from Terraform v0.12.
* Expressions within interpolation sequences `${ ... }` have the same new functionality added in Terraform v0.12, including the expanded set of available functions.

## 1.0.0 (September 26, 2017)

* No changes from 0.1.1; just adjusting to [the new version numbering scheme](https://www.hashicorp.com/blog/hashicorp-terraform-provider-versioning/).

## 0.1.1 (June 21, 2017)

NOTES:

Bumping the provider version to get around provider caching issues - still same functionality

## 0.1.0 (June 21, 2017)

NOTES:

* Same functionality as that of Terraform 0.9.8. Repacked as part of [Provider Splitout](https://www.hashicorp.com/blog/upcoming-provider-changes-in-terraform-0-10/)
