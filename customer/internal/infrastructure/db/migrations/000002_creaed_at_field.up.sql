ALTER TABLE IF EXISTS events
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE IF EXISTS entities
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;