---
title: "Steampipe Table: cpi_integration_package - Query SAP Cloud Integration integration packages using SQL"
description: "Allows users to query integration packages in an SAP Cloud Integration tenant. This table provides information about integration packages, including integration package ID, version, name, description, vendor, mode, and more."
folder: "Integration Package"
---

# Table: cpi_integration_package

Retrieve information about integration packages within an SAP Cloud Integration tenant's workspace.

## Examples

### List all integration packages in a tenant

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
