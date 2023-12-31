# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

variable "resource_group_name" {
  description = "The name of the Resource Group in which this Virtual Machine should be created."
}

variable "prefix" {
  description = "The prefix used for all resources used by this Virtual Machine"
}
