data "snowmirror_synchronization" "my_synchronization" {
    id = 5
            sync = {
        active = false
        allow_inherited_columns = false
        auto_schema_update = false
        columns = [
            {
                name = "Miriam Huel"
                sys_id = "...my_sys_id..."
            },
        ]
        columns_to_exclude = [
            {
                name = "Erica Bogisich III"
                sys_id = "...my_sys_id..."
            },
        ]
        delete_strategy = "...my_delete_strategy..."
        encoded_query = "...my_encoded_query..."
        full_load_scheduler = {
            begin_date = "...my_begin_date..."
            execution_type = "...my_execution_type..."
            type = "...my_type..."
        }
        mirror_table = "...my_mirror_table..."
        name = "Timmy Satterfield"
        reference_field_type = "...my_reference_field_type..."
        run_immediately = true
        scheduler = {
            begin_date = "...my_begin_date..."
            type = "...my_type..."
        }
        scheduler_priority = "...my_scheduler_priority..."
        table = "...my_table..."
        view = "...my_view..."
    }
        }