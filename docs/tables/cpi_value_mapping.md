# Table: cpi_value_mapping

Retrieve information about value mappings within an SAP Cloud Integration tenant's workspace.

## Examples

### List all value mappings in a tenant, sorted by package ID

```sql
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

```sql
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

```sql
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

```sql
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
