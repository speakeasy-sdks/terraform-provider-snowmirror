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