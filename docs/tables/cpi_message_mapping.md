---
title: "Steampipe Table: cpi_message_mapping - Query SAP Cloud Integration message mappings using SQL"
description: "Allows users to query message mappings in an SAP Cloud Integration tenant, providing insights into message transformation and structure conversion rules."
folder: "Message Mapping"
---

# Table: cpi_message_mapping - Query SAP Cloud Integration message mappings using SQL

SAP Cloud Integration message mappings are crucial components that define how message structures should be transformed between different formats and standards. These mappings enable complex message transformations across integration scenarios by providing a graphical way to map source structures to target structures. Message mappings are organized within integration packages and can be versioned and reused across multiple integration flows, ensuring consistency in message transformations.

## Table Usage Guide

The `cpi_message_mapping` table provides insights into message mappings within your SAP Cloud Integration tenant. As an integration developer or administrator, you can use this table to explore and manage message mapping definitions, track their versions, and understand their relationships with integration packages. This table is particularly valuable for maintaining message transformation consistency, version control, and understanding message mapping usage across your integration landscape.

## Examples

### List all message mappings in a tenant, sorted by package ID
Explore all message mappings organized by their package IDs to get a comprehensive view of your message transformation rules. This helps in understanding the distribution of mappings across different packages and their current versions.

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
Analyze the message mappings within a particular package to understand its message transformation components. This is useful when focusing on mapping management within a specific integration scenario or project.

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
Retrieve details about the currently active version of a particular message mapping. This helps in verifying the current mapping configuration and ensuring the correct version is being used in integration flows.

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
Examine a particular version of a message mapping to track changes or review historical configurations. This is particularly useful for version control and audit purposes.

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
