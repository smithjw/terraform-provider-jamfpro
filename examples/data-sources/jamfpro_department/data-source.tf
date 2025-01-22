# Create department resource
resource "jamfpro_department" "test" {
  name = "Engineering"
}

# Test data source by ID
data "jamfpro_department" "test_by_id" {
  id = jamfpro_department.test.id
}

# Test data source by name
data "jamfpro_department" "test_by_name" {
  name = jamfpro_department.test.name
}

# Outputs to verify
output "department_by_id" {
  value = {
    id   = data.jamfpro_department.test_by_id.id
    name = data.jamfpro_department.test_by_id.name
  }
}

output "department_by_name" {
  value = {
    id   = data.jamfpro_department.test_by_name.id
    name = data.jamfpro_department.test_by_name.name
  }
}

# Validation
output "matching_ids" {
  value = data.jamfpro_department.test_by_id.id == data.jamfpro_department.test_by_name.id
}