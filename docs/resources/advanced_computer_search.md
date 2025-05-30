---
page_title: "jamfpro_advanced_computer_search"
description: |-
  
---

# jamfpro_advanced_computer_search (Resource)


## Example Usage
```terraform
resource "jamfpro_advanced_computer_search" "advanced_computer_search_001" {
  name    = "tf - Example computer device Search"
  view_as = "Standard Web Page"

  sort1 = "Serial Number"
  sort2 = "Username"
  sort3 = "Department"

  criteria {
    name          = "Building"
    priority      = 0
    and_or        = "and"
    search_type   = "is"
    value         = "square"
    opening_paren = true
    closing_paren = false
  }

  criteria {
    name          = "Model"
    priority      = 1
    and_or        = "and"
    search_type   = "is"
    value         = "macbook air"
    opening_paren = false
    closing_paren = true
  }

  criteria {
    name          = "Computer Name"
    priority      = 2
    and_or        = "or"
    search_type   = "matches regex"
    value         = "thing"
    opening_paren = true
    closing_paren = false
  }

  criteria {
    name          = "Licensed Software"
    priority      = 3
    and_or        = "and"
    search_type   = "has"
    value         = "office"
    opening_paren = false
    closing_paren = true
  }

  site_id = "1"

  display_fields = [
    "Activation Lock Manageable",
    "Apple Silicon",
    "Architecture Type",
    "Available RAM Slots"
  ]

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The unique name of the advanced computer search

### Optional

- `criteria` (Block List) (see [below for nested schema](#nestedblock--criteria))
- `display_fields` (List of String) List of displayfields
- `site_id` (Number) Jamf Pro Site-related settings of the policy.
- `sort1` (String) First sorting criteria for the advanced computer search
- `sort2` (String) Second sorting criteria for the advanced computer search
- `sort3` (String) Third sorting criteria for the advanced computer search
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `view_as` (String) View type of the advanced computer search

### Read-Only

- `id` (String) The unique identifier of the advanced computer search

<a id="nestedblock--criteria"></a>
### Nested Schema for `criteria`

Required:

- `name` (String)
- `search_type` (String)
- `value` (String)

Optional:

- `and_or` (String)
- `closing_paren` (Boolean)
- `opening_paren` (Boolean)
- `priority` (Number)


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)