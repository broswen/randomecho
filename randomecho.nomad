job "randomecho" {
  datacenters = ["dc1"]

  type = "service"

  spread {
    attribute = "${node.datacenter}"
  }

  group "echo" {
    count = 5

    spread {
      attribute = "${node.unique.name}"
      target "silverstone" {
        percent = 66
      }

      target "corsair" {
        percent = 34
      }
    }

    network {
      port "http" {}
    }

    task "container" {
      driver = "docker"

      service {
        tags = ["echo", "go", "api", "${NOMAD_ALLOC_ID}"]
        port = "http"
        check {
          type = "http"
          protocol = "http"
          method = "GET"
          path = "/time"
          interval = "10s"
          timeout = "2s"

          check_restart {
            limit = 3
            grace = "3s"
          }
        }
      }

      env = {
        TEST = true
      }

      config {
        image = "broswen/randomecho:2.0.0"
        ports = ["http"]
      }

      resources {
        cpu    = 256
        memory = 128
      }
    }
  }
}
