data_source "prometheus" {
	type = "prometheus"
	access = "proxy",
	url = "http://mydatasource.com"
    database = "ohmytsdb"
    user = "johndoe"
	password = "cookieisstrong"
    is_default = true

    basic_auth {
        user = "johndoe",
        password = "cookieisstrong"
    }

    json_data = <<EOF
{
    "authType": "keys",
    "defaultRegion": "us-west-1"
}
EOF
}

