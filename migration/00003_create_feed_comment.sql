CREATE DATABASE eureka;

CREATE TABLE IF NOT EXISTS feed (
    feed_id TEXT PRIMARY KEY,
    account_id TEXT NOT NULL,
    message TEXT NOT NULL,
    image_url TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS comment (
    comment_id TEXT PRIMARY KEY,
    feed_id TEXT NOT NULL,
    account_id TEXT NOT NULL,
    message TEXT NOT NULL,
    image_url TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    FOREIGN KEY (feed_id) REFERENCES feed(feed_id)
);