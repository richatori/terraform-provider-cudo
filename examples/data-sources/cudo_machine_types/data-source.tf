data "cudo_machine_types" "machine_types" {
  search_params = {
    memory_gib = 4
  }
}