CREATE TABLE IF NOT EXISTS follower (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    follower_id TEXT NOT NULL,
    followed_date TIMESTAMPTZ DEFAULT NOW(),
    unfollowed_date TIMESTAMPTZ DEFAULT NOW(),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);