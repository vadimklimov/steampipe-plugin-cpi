---
title: "Steampipe Table: cpi_script_collection - Query SAP Cloud Integration script collections using SQL"
description: "Allows users to query script collections in an SAP Cloud Integration tenant, providing insights into reusable script components and their configurations."
folder: "Script Collection"
---

# Table: cpi_script_collection - Query SAP Cloud Integration script collections using SQL

SAP Cloud Integration script collections are reusable components that contain custom scripts and functions used across integration flows. These collections enable developers to maintain consistent scripting logic, reduce code duplication, and manage script versions effectively. Script collections are organized within integration packages and can be referenced by multiple integration flows, making them a crucial part of integration development and maintenance.

## Table Usage Guide

The `cpi_script_collection` table provides insights into script collections within your SAP Cloud Integration tenant. As an integration developer or administrator, you can use this table to explore and manage script collections, track their versions, and understand their relationships with integration packages. This table is particularly valuable for maintaining script consistency, version control, and understanding script usage across your integration landscape.

## Examples

### List all script collections in a tenant, sorted by package ID
Explore all script collections organized by their package IDs to get a comprehensive view of your scripting components. This helps in understanding the distribution of scripts across different packages and their current versions.

```sql+postgres
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

```sql+sqlite
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
Analyze the script collections within a particular package to understand its reusable components. This is useful when focusing on script management within a specific integration scenario or project.

```sql+postgres
select
  id,
  version,
  name
from
  cpi_script_collection
where
  package_id = 'Examples';
```

```sql+sqlite
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
Retrieve details about the currently active version of a particular script collection. This helps in verifying the current script configuration and ensuring the correct version is being used in integration flows.

```sql+postgres
select
  id,
  version,
  name
from
  cpi_script_collection
where
  id = 'ExampleScriptCollectionOne';
```

```sql+sqlite
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
Examine a particular version of a script collection to track changes or review historical configurations. This is particularly useful for version control and audit purposes.

```sql+postgres
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

```sql+sqlite
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
