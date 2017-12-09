grafana_data_source "prometheus" {
    type = "prometheus"
}

grafana_dashboard "traefik_detailed_stats" {
    title = "Traefik Detailed Stats"

    templating = [
        "${grafana_dashboard_template_var.service_name.id}"
    ]
}

grafana_dashboard_template_var "service_name" {
    label = "Service"
    name = "service_name"
    type = "query"
    hide = "none"

    query_options {
        data_source = "${var.data_source}"
        refresh = "on_dashboard_load"
        query = "label_values(traefik_request_duration_seconds_sum, exported_service)"
        regex = "/^(?!.*http|https$).*$/"
        sort = "asc"
    }

    selection_options {
        multi_value = false
        include_all_option = false
    }
}

