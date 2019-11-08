CREATE UNIQUE INDEX IF NOT EXISTS security_resources_type_name_vendor_unique ON security_resources(
  (raw ->> 'id'),
  (raw ->> 'kind'),
  (raw ->> 'version'));
