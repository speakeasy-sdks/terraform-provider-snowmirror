terraform {
  required_providers {
    snowmirror = {
      source  = "mikevdberge/snowmirror"
      version = "0.1.0"
    }
  }
}

provider "snowmirror" {
  # Configuration options
}