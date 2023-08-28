CREATE TABLE IF NOT EXISTS outbox
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    exchange VARCHAR(255) NOT NULL,
    route VARCHAR(255) NOT NULL,
    payload JSONB NOT NULL,
    aggregate_id VARCHAR(255) NOT NULL,
    event_status INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);