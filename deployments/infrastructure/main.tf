provider "google" {
  credentials = "${file("~/.config/sysdig/demo-environments/service-account.json")}"
  project     = "mateo-burillo-ns"
  region      = "us-central1"
  zone        = "us-central1-a"
}

resource "google_container_cluster" "cluster" {
  name               = "cloud-native-visibility-hub"
  initial_node_count = 3

  node_config {
    machine_type = "n1-standard-2"
  }

  network_policy {
    enabled = true
  }

  monitoring_service = "none"
  logging_service = "none"
}

resource "google_storage_bucket" "resources" {
  name          = "cloud-native-visibility-hub-resources"

  versioning {
    enabled = true
  }
}
