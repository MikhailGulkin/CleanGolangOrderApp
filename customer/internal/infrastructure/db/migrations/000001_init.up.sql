CREATE TABLE IF NOT EXISTS events
(
    event_id    UUID PRIMARY KEY,
    event_type  VARCHAR(255) NOT NULL,
    event_data  JSONB        NOT NULL,
    entity_type VARCHAR(255) NOT NULL,
    entity_id   UUID         NOT NULL
);
CREATE TABLE IF NOT EXISTS entities
(
    entity_type    VARCHAR(1000),
    entity_id      VARCHAR(1000),
    entity_version VARCHAR(1000) NOT NULL,
    PRIMARY KEY (entity_type, entity_id)
);