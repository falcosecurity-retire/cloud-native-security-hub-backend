ALTER TABLE IF EXISTS security_resources
  DROP COLUMN IF EXISTS available_versions;

ALTER TABLE IF EXISTS latest_security_resources
  DROP COLUMN IF EXISTS available_versions;
