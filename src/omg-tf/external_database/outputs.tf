output "sql_db_port" {
  value = "3306"
}

output "opsman_sql_db_name" {
  value = "${google_sql_database.opsman.0.name}"
}

output "opsman_sql_username" {
  value = "${random_id.opsman_db_username.0.b64}"
}

output "opsman_sql_password" {
  sensitive = true
  value = "${random_id.opsman_db_password.0.b64}"
}
