mock_resource "opslevel_component_type" {
  defaults = {
    # id intentionally omitted - will be assigned a random string
    name        = "API"
    alias       = "api"
    category    = "default"
    description = "An API component type"
    icon = {
      color = "#F59E0B"
      name  = "PhCloud"
    }
    properties = {
      "api_version" = {
        name                    = "API Version"
        description             = "The version of the API"
        allowed_in_config_files = true
        display_status          = "visible"
        locked_status           = "unlocked"
        schema                  = "{\"type\":\"string\"}"
      }
    }
    relationships = {}
  }
}

override_resource {
  target = opslevel_component_type.redis_cloud
  values = {
    name        = "Redis Cloud Database"
    alias       = "redis_cloud_database"
    category    = "infrastructure"
    description = "A Redis Enterprise Cloud managed database instance"
    icon = {
      color = "#d63031"
      name  = "PhDatabase"
    }
    properties = {
      "engine" = {
        name                    = "Engine"
        description             = "The database engine"
        allowed_in_config_files = false
        display_status          = "visible"
        locked_status           = "ui_locked"
        schema                  = "{\"enum\":[\"memcached\",\"redis\"],\"type\":\"string\"}"
      }
    }
    relationships = {}
  }
}

