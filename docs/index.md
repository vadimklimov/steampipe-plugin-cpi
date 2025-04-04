---
organization: vadimklimov
category: ["paas"]
icon_url: "/images/plugins/vadimklimov/cpi.svg"
brand_color: "#0070F2"
display_name: "SAP Cloud Integration"
short_name: "cpi"
description: "Steampipe plugin for querying artifacts from SAP Cloud Integration."
og_description: "Query SAP Cloud Integration with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/vadimklimov/cpi-social-graphic.png"
---

# SAP Cloud Integration + Steampipe

SAP Cloud Integration is part of [SAP Integration Suite](https://www.sap.com/products/technology-platform/integration-suite.html), an integration platform-as-a-service (iPaaS). It enables the development and execution of integration flows across hybrid landscapes, supporting application-to-application (A2A), business-to-business (B2B), and business-to-government (B2G) scenarios.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

For example:

```sql
select
  name,
  version,
  vendor,
  mode
from
  cpi_integration_package;
```

```text
+------------------------------------------------------------------------+---------+--------------+--------------+
| name                                                                   | version | vendor       | mode         |
+------------------------------------------------------------------------+---------+--------------+--------------+
| Advanced Event Mesh Adapter for SAP Integration Suite                  | 1.3.4   | SAP          | READ_ONLY    |
| Examples                                                               | 1.0.0   | Vadim Klimov | EDIT_ALLOWED |
| Process Integration Pipeline - Generic Integration Flows and Templates | 1.0.10  | SAP          | EDIT_ALLOWED |
+------------------------------------------------------------------------+---------+--------------+--------------+
```

## Documentation

- **[Table definitions & examples →](tables)**

## Get started

### Install

Download and install the latest CPI plugin:

```shell
steampipe plugin install vadimklimov/cpi
```

### Credentials

The CPI plugin makes use of open (public) APIs of SAP Cloud Integration to retrieve the required information about queried artifacts. APIs are OAuth-protected and support the client credentials flow. This authentication mechanism is employed by the CPI plugin to get API calls to an SAP Cloud Integration tenant authenticated and authorized.

To enable the CPI plugin to access the necessary APIs of an SAP Cloud Integration tenant, it is necessary to create an OAuth client for it in the SAP Business Technology Platform subaccount where the corresponding subscription for SAP Integration Suite has been created. In a Cloud Foundry environment, a service instance represents an OAuth client - hence, a service instance and a service key for it must be created.

1. Create a service instance for the `Process Integration Runtime` service and the `api` service plan. Ensurue that the `Client Credentials` grant type is selected when configuring service instance parameters. Assign the required roles. Refer to [Permissions](#permissions) for the list of roles required to retrieve data from the corresponding tables.

2. Create a service key of the `ClientId/Secret` type for the service instance created in the previous step.

| Item        | Description                                                                                            |
| ----------- | ------------------------------------------------------------------------------------------------------ |
| Credentials | The CPI plugin uses a client ID and client secret to authenticate calls to SAP Cloud Integration APIs. |
| Permissions | Assign the required roles (refer to [Permissions](#permissions)) to the service instance.              |
| Radius      | Each connection represents a single SAP Cloud Integration tenant.                                      |
| Resolution  | Credentials explicitly set in a Steampipe configuration file (`~/.steampipe/config/cpi.spc`).          |

### Permissions

| Table                     | Required role           |
| ------------------------- | ----------------------- |
| `cpi_integration_flow`    | `WorkspacePackagesRead` |
| `cpi_integration_package` | `WorkspacePackagesRead` |
| `cpi_message_mapping`     | `WorkspacePackagesRead` |
| `cpi_script_collection`   | `WorkspacePackagesRead` |
| `cpi_value_mapping`       | `WorkspacePackagesRead` |

### Configuration

Installing the latest CPI plugin will create a configuration file `~/.steampipe/config/cpi.spc` with a single connection named `cpi`:

```hcl
connection "cpi" {
  plugin = "vadimklimov/cpi"

  # Base URL.
  # In the service key, the `url` attribute in the `oauth` section.
  # base_url = "https://xxxxxxxxxx.it-cpi000.cfapps.xx00-000.hana.ondemand.com"

  # Token URL.
  # In the service key: the `tokenurl` attribute in the `oauth` section.
  # token_url = "https://xxxxxxxxxx.authentication.xx00.hana.ondemand.com/oauth/token"

  # Client ID.
  # In the service key: the `clientid` attribute in the `oauth` section.
  # client_id = "sb-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx!x000000|it!x00000"

  # Client secret.
  # In the service key: the `clientsecret` attribute in the `oauth` section.
  # client_secret = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx$xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

  # Maximum concurrency of requests to APIs of an SAP Cloud Integration tenant.
  # If not set, the default value is the number of logical CPUs available on the system.
  # max_concurrency = 8

  # Timeout for requests to APIs of an SAP Cloud Integration tenant.
  # Valid time units: ns, us/µs, ms, s, m, h.
  # If not set, the default value is "30s" (30 seconds).
  # timeout = "30s"
}
```

## Get involved

- Open source: https://github.com/vadimklimov/steampipe-plugin-cpi
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
