---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "snowmirror Provider"
subcategory: ""
description: |-
  SnowMirror API: Example Hashicups through Speakeasy
---

# snowmirror Provider

SnowMirror API: Example Hashicups through Speakeasy

## Example Usage

```terraform
terraform {
  required_providers {
    snowmirror = {
      source  = "mikevdberge/snowmirror"
      version = "0.1.1"
    }
  }
}

provider "snowmirror" {
  # Configuration options
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `password` (String, Sensitive)
- `server_url` (String) Server URL (defaults to http://localhost:9090)
- `username` (String, Sensitive)