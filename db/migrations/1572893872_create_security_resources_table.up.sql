CREATE TABLE IF NOT EXISTS security_resources(
   id serial PRIMARY KEY,
   raw jsonb
);

CREATE INDEX IF NOT EXISTS security_resources_indexginp ON security_resources USING GIN (raw jsonb_path_ops);
