<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->

<a href="https://terraform.io">
    <img src="https://upload.wikimedia.org/wikipedia/commons/0/04/Terraform_Logo.svg" alt="Terraform logo" title="Terraform" height="30" />
</a>
&nbsp;
<a href="https://opennebula.io/">
    <img src="https://opennebula.io/wp-content/uploads/2013/12/opennebula_cloud_logo_white_bg.png" alt="OpenNebula logo" title="OpenNebula" height="30" />
</a>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

# Terraform Provider for OpenNebula

[![CI](https://github.com/OpenNebula/terraform-provider-opennebula/actions/workflows/ci.yaml/badge.svg)](https://github.com/OpenNebula/terraform-provider-opennebula/actions/workflows/ci.yaml)

## Quick start

* [Install Terraform](https://learn.hashicorp.com/terraform/getting-started/install)
* [Use the Provider](https://registry.terraform.io/providers/OpenNebula/opennebula/latest/docs)

## Example usage

```hcl
terraform {
  required_providers {
    opennebula = {
      source = "OpenNebula/opennebula"
      version = "~> 1.4"
    }
  }
}

provider "opennebula" {
  endpoint      = "https://example.com:2633/RPC2"
}

resource "opennebula_group" "group" {
  name = "OpenNebula"
}
```

More details [here](./website/docs/index.html.markdown).

## OpenNebula versions support

* `~> 6.10` (current)
* `~> 6.4` (LTS)

See OpenNebula's [Release Policy](https://github.com/OpenNebula/one/wiki/Release-Policy) for more details.

## Contributing

Please refer to [CONTRIBUTING.md](./CONTRIBUTING.md).

## License

Please refer to [LICENSE](./LICENSE).

## References

Other Projects about Terraform provider exist. This project has been inspired by [Runtastic](https://github.com/runtastic/terraform-provider-opennebula) and [BlackBerry](https://github.com/blackberry/terraform-provider-opennebula) projects.

## Acknowledgements

Some of the software features included in this repository have been made possible through the funding of the following innovation projects:

* [ONEnextgen](http://onenextgen.eu/) (Grant Agreement UNICO IPCEI-2023-003), supported by the Spanish Ministry for Digital Transformation and Civil Service through the UNICO IPCEI Program, co-funded by the European Union – NextGenerationEU through the Recovery and Resilience Facility (RRF).
