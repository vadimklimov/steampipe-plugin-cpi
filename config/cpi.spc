connection "cpi" {
  plugin = "vadimklimov/cpi"

  # Base URL.
  # In the service key, the `url` attribute in the `oauth` section.
  # base_url = "https://xxxxxxxxxx.it-cpi000.cfapps.xx00-000.hana.ondemand.com"

  # Token URL.
  # In the service key: the `tokenurl` attribute in the `oauth` section.
  # token_url = "https://xxxxxxxxxx.authentication.xx00.hana.ondemand.com/oauth/token"

  # Client ID.
  # In the service key: the `clientid` attribute in the `oauth` section.
  # client_id = "sb-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx!x000000|it!x00000"

  # Client secret.
  # In the service key: the `clientsecret` attribute in the `oauth` section.
  # client_secret = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx$xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

  # Maximum concurrency of requests to APIs of an SAP Cloud Integration tenant.
  # If not set, the default value is the number of logical CPUs available on the system.
  # max_concurrency = 8

  # Timeout for requests to APIs of an SAP Cloud Integration tenant.
  # Valid time units: ns, us/µs, ms, s, m, h.
  # If not set, the default value is "30s" (30 seconds).
  # timeout = "30s"
}
