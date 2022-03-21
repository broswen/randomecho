job "randomecho" {
  datacenters = ["dc1"]

  type = "service"

  spread {
    attribute = "${node.datacenter}"
  }

  group "api" {
    count = 5

    spread {
      attribute = "${node.unique.name}"
      target "silverstone" {
        percent = 50
      }

      target "corsair" {
        percent = 50
      }
    }

    network {
      port "http" {}
    }

    task "randomecho" {
      driver = "docker"

      config {
        image = "broswen/randomecho:2.0.0"
        ports = ["http"]
      }

      service {
        name = "randomecho"
        port = "http"
        check {
          name = "alive"
          type = "http"
          protocol = "http"
          method = "GET"
          path = "/time"
          interval = "10s"
          timeout = "2s"
        }
      }

      resources {
        cpu    = 256
        memory = 128
      }
    }
  }

  group "cache" {
    count = 1

    network {
      port "db" {
        to = 6379
      }
    }

    restart {
      attempts = 10
      interval = "5m"
      delay = "30s"
      mode = "delay"
    }


    task "redis" {
      driver = "docker"

      config {
        image = "redis:latest"
        ports = ["db"]
      }

      service {
        name = "redis"
        port = "db"
        check {
          name = "alive"
          type = "tcp"
          interval = "10s"
          timeout = "2s"
        }
      }

      resources {
        cpu = 500
        memory = 256
      }
    }
  }
}