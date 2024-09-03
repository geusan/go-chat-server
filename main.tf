terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0.1"
    }
  }
}

provider "docker" {}

resource "docker_image" "mysql" {
  name         = "mysql:8.0"
  keep_locally = true
}

resource "docker_image" "redis" {
  name         = "redis:7.4.0"
  keep_locally = true
}

# TODO: make docker container
# chat-server
# resource "docker_image" "chat-server" {
#     name = "chat-server"
#     keep_locally = false
#     build {
#         context = "${terraform.workspace}/chat-server"
#         tag = ["chat-server:develop"]
#     }
# }

# # api-server
# resource "docker_image" "api-server" {
#     name = "api-server"
#     keep_locally = false
#     build {
#         context = "${terraform.workspace}/server"
#         tag = ["api-server:develop"]
#     }
# }

resource "docker_container" "mysql" {
  image = docker_image.mysql.image_id
  name  = "mysql"

  ports {
    internal = 3306
    external = 3306
  }

  env = [
    # password for local => root:localhost
    "MYSQL_ROOT_PASSWORD=localhost",
    "MYSQL_DATABASE=chat",
  ]

  volumes {
    host_path = "${abspath(terraform.workspace)}/db/data"
    container_path = "/var/lib/mysql"
  }

  volumes {
    host_path = "${abspath(terraform.workspace)}/db/my.cnf"
    container_path = "/etc/mysql/my.cnf"
  }
}

resource "docker_container" "redis" {
  image = docker_image.redis.image_id
  name  = "redis"

  ports {
    internal = 6397
    external = 6397
  }
}
