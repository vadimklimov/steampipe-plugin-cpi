---
title: "Steampipe Table: cpi_message_mapping - Query SAP Cloud Integration message mappings using SQL"
description: "Allows users to query message mappings in an SAP Cloud Integration tenant. This table provides information about message mappings, including message mapping ID, version, name and description."
folder: "Message Mapping"
---

# Table: cpi_message_mapping

Retrieve information about message mappings within an SAP Cloud Integration tenant's workspace.

## Examples

### List all message mappings in a tenant, sorted by package ID

```sql+postgres
select
  id,
  version,
  name,
  package_id
from
  cpi_message_mapping
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
  cpi_message_mapping
order by
  package_id;
```

### List all message mappings in a specific integration package

```sql+postgres
select
  id,
  version,
  name
from
  cpi_message_mapping
where
  package_id = 'Examples';
```

```sql+sqlite
select
  id,
  version,
  name
from
  cpi_message_mapping
where
  package_id = 'Examples';
```

### Get an active version of a specific message mapping

```sql+postgres
select
  id,
  version,
  name
from
  cpi_message_mapping
where
  id = 'ExampleMessageMappingOne';
```

```sql+sqlite
select
  id,
  version,
  name
from
  cpi_message_mapping
where
  id = 'ExampleMessageMappingOne';
```

### Get a specific version of a specific message mapping

```sql+postgres
select
  id,
  version,
  name
from
  cpi_message_mapping
where
  id = 'ExampleMessageMappingOne'
  and version = '1.2.3';
```

```sql+sqlite
select
  id,
  version,
  name
from
  cpi_message_mapping
where
  id = 'ExampleMessageMappingOne'
  and version = '1.2.3';
```
