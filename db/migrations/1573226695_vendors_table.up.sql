CREATE TABLE IF NOT EXISTS vendors(
   id serial PRIMARY KEY,
   raw jsonb
);

CREATE INDEX IF NOT EXISTS vendors_indexginp ON vendors USING GIN (raw jsonb_path_ops);
CREATE UNIQUE INDEX IF NOT EXISTS vendors_name_unique ON vendors(
  (raw ->> 'id'));
