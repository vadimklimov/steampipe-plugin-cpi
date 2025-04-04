# Table: cpi_integration_flow

Retrieve information about integration flows within an SAP Cloud Integration tenant's workspace.

## Examples

### List all integration flows in a tenant, sorted by package ID

```sql
select
  id,
  version,
  name,
  package_id,
  modified_at
from
  cpi_integration_flow
order by
  package_id;
```

### List all integration flows in a specific integration package

```sql
select
  id,
  version,
  name,
  modified_at
from
  cpi_integration_flow
where
  package_id = 'Examples';
```

### Get an active version of a specific integration flow

```sql
select
  id,
  version,
  name,
  modified_at
from
  cpi_integration_flow
where
  id = 'ExampleFlowOne';
```

### Get a specific version of a specific integration flow

```sql
select
  id,
  version,
  name,
  modified_at
from
  cpi_integration_flow
where
  id = 'ExampleFlowOne'
  and version = '1.2.3';
```
