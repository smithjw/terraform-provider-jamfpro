---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "jamfpro_computer_extension_attribute Resource - terraform-provider-jamfpro"
subcategory: ""
description: |-
  
---

# jamfpro_computer_extension_attribute (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled` (Boolean) Indicates if the computer extension attribute is enabled.
- `name` (String) The unique name of the Jamf Pro computer extension attribute.

### Optional

- `data_type` (String) Data type of the computer extension attribute. Can be String / Integer / Date (YYYY-MM-DD hh:mm:ss). Value defaults to `String`.
- `description` (String) Description of the computer extension attribute.
- `input_type` (Block List, Max: 1) Input type details of the computer extension attribute. (see [below for nested schema](#nestedblock--input_type))
- `inventory_display` (String) Display details for inventory for the computer extension attribute. Value defaults to `General`.
- `recon_display` (String) Display details for recon for the computer extension attribute.
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The unique identifier of the computer extension attribute.

<a id="nestedblock--input_type"></a>
### Nested Schema for `input_type`

Required:

- `type` (String)

Optional:

- `choices` (List of String)
- `platform` (String) Platform type for the computer extension attribute.
- `script` (String)


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)
