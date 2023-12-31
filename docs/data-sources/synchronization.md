---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "snowmirror_synchronization Data Source - terraform-provider-snowmirror"
subcategory: ""
description: |-
  Synchronization DataSource
---

# snowmirror_synchronization (Data Source)

Synchronization DataSource

## Example Usage

```terraform
data "snowmirror_synchronization" "my_synchronization" {
  columns = [
    {
      name = "sys_id"
    },
  ]
  columns_to_exclude = [
    {
      name = "sys_created_by"
    },
  ]
  id              = 9
  run_immediately = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `columns` (Attributes List) (see [below for nested schema](#nestedatt--columns))
- `columns_to_exclude` (Attributes List) (see [below for nested schema](#nestedatt--columns_to_exclude))
- `id` (Number) Id of the synchronization.
- `run_immediately` (Boolean) Determines whether initial synchronization should be done

### Read-Only

- `active` (Boolean) true - synchronization is active and can be scheduled to synchronize data from ServiceNow
false - synchronization is deactivated and cannot be scheduled to synchronize data from ServiceNowNow
- `allow_inherited_columns` (Boolean) SnowMirror checks if columns exist in ServiceNow. If this flag is set to true,
- `attachment_directory` (String)
- `auto_schema_update` (String) Configures how to check for schema changes in ServiceNow.

Enabled (true) - whenever a synchronization is executed, SnowMirror checks for schema changes in ServiceNow. Automatically adds, updates (data type, max. length of a column) and removes columns. If a new column is added SnowMirror clears the mirror table and downloads all records from scratch.
Enabled (no truncation) (ENABLED_WITHOUT_TRUNCATION) - the same as Enabled option. It handles new columns differently, though. If a new column is added SnowMirror does not clear the mirror table. Instead, it creates the column and populates it with a default value (which is defined in ServiceNow).
- `delete_strategy` (String) must be one of ["AUDIT", "TRUNCATE", "DIFF", "NONE"]
- `encoded_query` (String)
- `format` (String) must be one of ["CSV", "XML"]
How to store backups. "CSV" - comma separated file. "XML" - XML files.
- `full_load_scheduler` (Attributes) (see [below for nested schema](#nestedatt--full_load_scheduler))
- `master_table` (String)
- `mirror_table` (String) Name of the table in mirror database where the data will be migrated.
- `name` (String) Display name of the synchronization.
- `reference_field_type` (String) Defines how to synchronize reference field types.
- `retention_period` (Number) How many days to keep backups
- `scheduler` (Attributes) (see [below for nested schema](#nestedatt--scheduler))
- `scheduler_priority` (String) must be one of ["HIGHEST", "HIGH", "NORMAL", "LOW", "LOWEST"]
- `synchronization_type` (String)
- `table` (String) Name of the table in ServiceNow.
- `update_before_synchronization_run` (String)
- `view` (String) Name of the view in ServiceNow.

<a id="nestedatt--columns"></a>
### Nested Schema for `columns`

Optional:

- `name` (String)


<a id="nestedatt--columns_to_exclude"></a>
### Nested Schema for `columns_to_exclude`

Optional:

- `name` (String)


<a id="nestedatt--full_load_scheduler"></a>
### Nested Schema for `full_load_scheduler`

Read-Only:

- `begin_date` (String)
- `execution_type` (String) must be one of ["CLEAN_AND_SYNCHRONIZE", "DIFFERENTIAL."]
- `time` (String)
- `type` (String) must be one of ["MANUALLY", "DAILY", "WEEKLY", "PERIODICALLY", "CRON"]
- `visible` (Boolean)


<a id="nestedatt--scheduler"></a>
### Nested Schema for `scheduler`

Read-Only:

- `begin_date` (String)
- `days` (List of String)
- `inc_load_execution_type` (String)
- `time` (String)
- `type` (String) must be one of ["MANUALLY", "DAILY", "WEEKLY", "PERIODICALLY", "CRON"]
Specifies when the incremental load synchronization will run
- `visible` (Boolean)


