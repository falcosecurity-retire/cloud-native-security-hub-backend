terraform {
  backend "gcs" {
    bucket  = "demo-environments-state"
    prefix  = "terraform/cloud-native-visibility-hub"
  }
}
