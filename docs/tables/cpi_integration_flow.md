---
title: "Steampipe Table: cpi_integration_flow - Query SAP Cloud Integration integration flows using SQL"
description: "Allows users to query integration flows in an SAP Cloud Integration tenant, providing insights into flow configurations, versions, and metadata."
folder: "Integration Flow"
---

# Table: cpi_integration_flow - Query SAP Cloud Integration integration flows using SQL

SAP Cloud Integration integration flows are the core components that define the integration scenarios in SAP Cloud Integration. Each integration flow represents a specific integration process that can include message transformations, routing logic, and connectivity configurations. Integration flows are organized within integration packages and can be versioned, deployed, and monitored to ensure successful integration between different systems and applications.

## Table Usage Guide

The `cpi_integration_flow` table provides insights into integration flows within your SAP Cloud Integration tenant. As an integration developer or administrator, you can use this table to explore detailed information about integration flows, including their versions, configurations, and relationships with integration packages. This table is particularly valuable for managing integration scenarios, tracking changes across versions, and maintaining an overview of your integration landscape.

## Examples

### List all integration flows in a tenant, sorted by package ID
Explore all integration flows organized by their package IDs to get a comprehensive view of your integration landscape. This helps in understanding the distribution of integration flows across different packages and their current versions.

```sql+postgres
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

```sql+sqlite
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
Analyze the integration flows within a particular package to understand its components and recent modifications. This is useful when focusing on a specific integration scenario or project.

```sql+postgres
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

```sql+sqlite
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
Retrieve details about the currently active version of a particular integration flow. This helps in verifying the current configuration and last modification time of a specific integration scenario.

```sql+postgres
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

```sql+sqlite
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
Examine a particular version of an integration flow to track changes or review historical configurations. This is particularly useful for version control and audit purposes.

```sql+postgres
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

```sql+sqlite
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
