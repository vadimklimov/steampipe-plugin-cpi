---
title: "Steampipe Table: cpi_value_mapping - Query SAP Cloud Integration value mappings using SQL"
description: "Allows users to query value mappings in an SAP Cloud Integration tenant, providing insights into data transformation and standardization rules."
folder: "Value Mapping"
---

# Table: cpi_value_mapping - Query SAP Cloud Integration value mappings using SQL

SAP Cloud Integration value mappings are essential components that define how values should be transformed between different systems and standards. These mappings enable consistent data transformation across integration scenarios by providing a centralized way to manage value conversions. Value mappings are organized within integration packages and can be versioned and reused across multiple integration flows, ensuring consistency in data transformations.

## Table Usage Guide

The `cpi_value_mapping` table provides insights into value mappings within your SAP Cloud Integration tenant. As an integration developer or administrator, you can use this table to explore and manage value mapping definitions, track their versions, and understand their relationships with integration packages. This table is particularly valuable for maintaining data transformation consistency, version control, and understanding value mapping usage across your integration landscape.

## Examples

### List all value mappings in a tenant, sorted by package ID
Explore all value mappings organized by their package IDs to get a comprehensive view of your data transformation rules. This helps in understanding the distribution of mappings across different packages and their current versions.

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
Analyze the value mappings within a particular package to understand its data transformation components. This is useful when focusing on mapping management within a specific integration scenario or project.

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
Retrieve details about the currently active version of a particular value mapping. This helps in verifying the current mapping configuration and ensuring the correct version is being used in integration flows.

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
Examine a particular version of a value mapping to track changes or review historical configurations. This is particularly useful for version control and audit purposes.

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
