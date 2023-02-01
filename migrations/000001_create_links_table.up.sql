CREATE TABLE IF NOT EXISTS links(
    id VARCHAR(10),
    hashed VARCHAR(2048) NOT NULL,
    original VARCHAR(2048) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY(id)
);

CREATE INDEX links_created_at_index ON links (created_at DESC);
