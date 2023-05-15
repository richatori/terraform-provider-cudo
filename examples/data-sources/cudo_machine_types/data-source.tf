data "cudo_machine_types" "types" {
  search_params = {
    memory_gib = 4
  }
}