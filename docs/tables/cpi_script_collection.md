# Table: cpi_script_collection

Retrieve information about script collections within an SAP Cloud Integration tenant's workspace.

## Examples

### List all script collections in a tenant, sorted by package ID

```sql
select
  id,
  version,
  name,
  package_id
from
  cpi_script_collection
order by
  package_id;
```

### List all script collections in a specific integration package

```sql
select
  id,
  version,
  name
from
  cpi_script_collection
where
  package_id = 'Examples';
```

### Get an active version of a specific script collection

```sql
select
  id,
  version,
  name
from
  cpi_script_collection
where
  id = 'ExampleScriptCollectionOne';
```

### Get a specific version of a specific script collection

```sql
select
  id,
  version,
  name
from
  cpi_script_collection
where
  id = 'ExampleScriptCollectionOne'
  and version = '1.2.3';
```
