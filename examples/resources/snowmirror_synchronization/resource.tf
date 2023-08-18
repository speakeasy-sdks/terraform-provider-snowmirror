resource "snowmirror_synchronization" "my_synchronization" {
  active                  = true
  allow_inherited_columns = false
  auto_schema_update      = "true"
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
  delete_strategy = "AUDIT"
  encoded_query   = ""
  full_load_scheduler = {
    begin_date     = "2014-08-07"
    execution_type = "CLEAN_AND_SYNCHRONIZE"
    time           = "...my_time..."
    type           = "MANUALLY"
    visible        = false
  }
  id                   = 3
  mirror_table         = "incident"
  name                 = "incident"
  reference_field_type = "...my_reference_field_type..."
  run_immediately      = false
  scheduler = {
    begin_date = "2023-08-11"
    days = [
      "SATURDAY",
    ]
    inc_load_execution_type = "...my_inc_load_execution_type..."
    time                    = "02:00"
    type                    = "MANUALLY"
    visible                 = false
  }
  scheduler_priority = "HIGH"
  table              = "incident"
  view               = "...my_view..."
}