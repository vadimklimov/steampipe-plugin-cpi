---
title: "Steampipe Table: cpi_integration_package - Query SAP Cloud Integration integration packages using SQL"
description: "Allows users to query integration packages in an SAP Cloud Integration tenant, providing insights into package configurations, versions, and metadata including vendor information and update status."
folder: "Integration Package"
---

# Table: cpi_integration_package - Query SAP Cloud Integration integration packages using SQL

SAP Cloud Integration integration packages are containers that organize and group related integration artifacts such as integration flows, value mappings, and script collections. These packages serve as the primary organizational unit for integration content, enabling structured development, deployment, and maintenance of integration scenarios. Integration packages can be created by different vendors, including SAP and custom developments, and support versioning and updates to manage the lifecycle of integration content.

## Table Usage Guide

The `cpi_integration_package` table provides insights into integration packages within your SAP Cloud Integration tenant. As an integration developer or administrator, you can use this table to explore package configurations, track versions, monitor available updates, and understand the organization of integration content. This table is particularly valuable for managing package lifecycles, identifying update opportunities, and maintaining an organized integration landscape.

## Examples

### List all integration packages in a tenant
Explore all integration packages to get a comprehensive view of your integration content. This helps in understanding the overall organization of integration artifacts and their current status.

```sql+postgres
select
  id,
  version,
  name,
  vendor,
  mode,
  modified_at
from
  cpi_integration_package;
```

```sql+sqlite
select
  id,
  version,
  name,
  vendor,
  mode,
  modified_at
from
  cpi_integration_package;
```

### List all integration packages provided by SAP with 'adapter' in the package ID (case-insensitive) and that have available updates
Identify SAP-provided adapter packages that need updates. This query is useful for maintaining up-to-date integration content and ensuring you're using the latest adapter versions.

```sql+postgres
select
  id,
  version,
  name,
  mode,
  modified_at
from
  cpi_integration_package
where
  vendor = 'SAP'
  and id ilike '%adapter%'
  and update_available;
```

```sql+sqlite
select
  id,
  version,
  name,
  mode,
  modified_at
from
  cpi_integration_package
where
  vendor = 'SAP'
  and id like '%adapter%'
  and update_available;
```

### Get a specific integration package
Retrieve detailed information about a particular integration package. This helps in reviewing the configuration and status of a specific package when performing maintenance or updates.

```sql+postgres
select
  id,
  version,
  name,
  vendor,
  mode,
  modified_at
from
  cpi_integration_package
where
  id = 'Examples';
```

```sql+sqlite
select
  id,
  version,
  name,
  vendor,
  mode,
  modified_at
from
  cpi_integration_package
where
  id = 'Examples';
```
