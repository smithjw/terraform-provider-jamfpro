---
page_title: "jamfpro_smtp_server"
description: |-
  
---

# jamfpro_smtp_server (Resource)


## Example Usage
```terraform
# Basic Authentication Example
resource "jamfpro_smtp_server" "basic_auth" {
  enabled             = true
  authentication_type = "BASIC"

  connection_settings {
    host               = "smtp.sendgrid.net"
    port               = 587
    encryption_type    = "TLS_1_2"
    connection_timeout = 5
  }

  sender_settings {
    display_name  = "Jamf Pro Server"
    email_address = "user@company.com"
  }

  basic_auth_credentials {
    username = "sample-username"
    password = "password"
  }
}

# Graph API Authentication Example
resource "jamfpro_smtp_server" "graph_api" {
  enabled             = true
  authentication_type = "GRAPH_API"

  sender_settings {
    email_address = "noreply@yourdomain.onmicrosoft.com"
  }

  graph_api_credentials {
    tenant_id     = "c84b7b82-c277-411b-975d-7431b4ce40ac"
    client_id     = "5294f9d1-f723-419c-93db-ff040bf7c947"
    client_secret = "password"
  }
}

# Google Mail Authentication Example
resource "jamfpro_smtp_server" "google_mail" {
  enabled             = true
  authentication_type = "GOOGLE_MAIL"

  sender_settings {
    email_address = "exampleEmail@gmail.com"
  }

  google_mail_credentials {
    client_id     = "012345678901-abcdefghijklmnopqrstuvwxyz123456.apps.googleusercontent.com"
    client_secret = "password"
  }

  # Note: authentications block is computed and will be populated by the API
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `authentication_type` (String) Type of SMTP authentication type to use
- `enabled` (Boolean) Whether the SMTP server is enabled to allow Jamf Pro to send emails and invitations
- `sender_settings` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--sender_settings))

### Optional

- `basic_auth_credentials` (Block List, Max: 1) (see [below for nested schema](#nestedblock--basic_auth_credentials))
- `connection_settings` (Block List, Max: 1) (see [below for nested schema](#nestedblock--connection_settings))
- `google_mail_credentials` (Block List, Max: 1) (see [below for nested schema](#nestedblock--google_mail_credentials))
- `graph_api_credentials` (Block List, Max: 1) (see [below for nested schema](#nestedblock--graph_api_credentials))
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `authentications` (Set of Object) (see [below for nested schema](#nestedatt--authentications))
- `id` (String) The ID of this resource.

<a id="nestedblock--sender_settings"></a>
### Nested Schema for `sender_settings`

Required:

- `email_address` (String) Email address of the sender

Optional:

- `display_name` (String) Display name for the sender


<a id="nestedblock--basic_auth_credentials"></a>
### Nested Schema for `basic_auth_credentials`

Required:

- `password` (String, Sensitive) Password for basic authentication
- `username` (String) Username for basic authentication


<a id="nestedblock--connection_settings"></a>
### Nested Schema for `connection_settings`

Required:

- `encryption_type` (String) Type of encryption to use
- `host` (String) SMTP server hostname
- `port` (Number) SMTP server port

Optional:

- `connection_timeout` (Number) Connection timeout in seconds


<a id="nestedblock--google_mail_credentials"></a>
### Nested Schema for `google_mail_credentials`

Required:

- `client_id` (String) Google client ID for Gmail
- `client_secret` (String, Sensitive) Google client secret for Gmail


<a id="nestedblock--graph_api_credentials"></a>
### Nested Schema for `graph_api_credentials`

Required:

- `client_id` (String) Microsoft client ID for Graph API. Must be a valid GUID/UUID.
- `client_secret` (String, Sensitive) Microsoft client secret for Graph API
- `tenant_id` (String) Microsoft tenant ID for Graph API. Must be a valid GUID/UUID.


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)


<a id="nestedatt--authentications"></a>
### Nested Schema for `authentications`

Read-Only:

- `email_address` (String)
- `status` (String)