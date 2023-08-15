resource "snowmirror_synchronization" "my_synchronization" {
    sync = {
        active = true
        allow_inherited_columns = false
        auto_schema_update = true
        columns = [
            {
                name = "Stuart Stiedemann"
            },
        ]
        columns_to_exclude = [
            {
                name = "Sabrina Oberbrunner"
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
        name = "Raquel Bednar"
        reference_field_type = "...my_reference_field_type..."
        run_immediately = false
        scheduler = {
            begin_date = "...my_begin_date..."
            type = "...my_type..."
        }
        scheduler_priority = "...my_scheduler_priority..."
        table = "...my_table..."
        view = "...my_view..."
    }
        }