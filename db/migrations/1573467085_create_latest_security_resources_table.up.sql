CREATE TABLE IF NOT EXISTS latest_security_resources(
   id serial PRIMARY KEY,
   raw jsonb
);

CREATE INDEX IF NOT EXISTS latest_security_resources_indexginp ON latest_security_resources USING GIN (raw jsonb_path_ops);
CREATE UNIQUE INDEX IF NOT EXISTS latest_security_resources_type_name_unique ON latest_security_resources(
  (raw ->> 'id'),
  (raw ->> 'kind'));
