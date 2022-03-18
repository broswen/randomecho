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
        percent = 50
      }

      target "corsair" {
        percent = 50
      }
    }

    network {
      port "http" {}
    }

    task "container" {
      driver = "docker"

      config {
        image = "broswen/randomecho:1.5.0"
        ports = ["http"]
      }

      resources {
        cpu    = 256
        memory = 128
      }
    }
  }
}
