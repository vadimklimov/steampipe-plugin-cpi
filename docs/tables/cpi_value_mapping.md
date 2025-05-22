---
title: "Steampipe Table: cpi_value_mapping - Query SAP Cloud Integration value mappings using SQL"
description: "Allows users to query value mappings in an SAP Cloud Integration tenant. This table provides information about value mappings, including value mapping ID, version, name and description."
folder: "Value Mapping"
---

# Table: cpi_value_mapping

Retrieve information about value mappings within an SAP Cloud Integration tenant's workspace.

## Examples

### List all value mappings in a tenant, sorted by package ID

```sql+postgres
select
  id,
  version,
  name,
  package_id
from
  cpi_value_mapping
order by
  package_id;
```

```sql+sqlite
select
  id,
  version,
  name,
  package_id
from
  cpi_value_mapping
order by
  package_id;
```

### List all value mappings in a specific integration package

```sql+postgres
select
  id,
  version,
  name
from
  cpi_value_mapping
where
  package_id = 'Examples';
```

```sql+sqlite
select
  id,
  version,
  name
from
  cpi_value_mapping
where
  package_id = 'Examples';
```

### Get an active version of a specific value mapping

```sql+postgres
select
  id,
  version,
  name
from
  cpi_value_mapping
where
  id = 'ExampleValueMappingOne';
```

```sql+sqlite
select
  id,
  version,
  name
from
  cpi_value_mapping
where
  id = 'ExampleValueMappingOne';
```

### Get a specific version of a specific value mapping

```sql+postgres
select
  id,
  version,
  name
from
  cpi_value_mapping
where
  id = 'ExampleValueMappingOne'
  and version = '1.2.3';
```

```sql+sqlite
select
  id,
  version,
  name
from
  cpi_value_mapping
where
  id = 'ExampleValueMappingOne'
  and version = '1.2.3';
```
