# Table: cpi_integration_package

Retrieve information about integration packages within an SAP Cloud Integration tenant's workspace.

## Examples

### List all integration packages in a tenant

```sql
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

```sql
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

### Get a specific integration package

```sql
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
