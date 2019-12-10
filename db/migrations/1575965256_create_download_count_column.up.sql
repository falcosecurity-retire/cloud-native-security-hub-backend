ALTER TABLE security_resources ADD COLUMN download_count INTEGER DEFAULT 0;
ALTER TABLE latest_security_resources ADD COLUMN download_count INTEGER DEFAULT 0;