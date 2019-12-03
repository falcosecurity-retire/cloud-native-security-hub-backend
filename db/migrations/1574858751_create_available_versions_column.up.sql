ALTER TABLE IF EXISTS security_resources
  ADD COLUMN IF NOT EXISTS available_versions text[];

ALTER TABLE IF EXISTS latest_security_resources
  ADD COLUMN IF NOT EXISTS available_versions text[];
