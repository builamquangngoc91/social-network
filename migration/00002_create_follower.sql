CREATE TABLE IF NOT EXISTS follower (
    id TEXT PRIMARY KEY,
    account_id TEXT NOT NULL,
    follower_id TEXT NOT NULL,
    followed_date TIMESTAMPTZ DEFAULT NOW(),
    unfollowed_date TIMESTAMPTZ DEFAULT NOW(),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (account_id) REFERENCES account(account_id),
    FOREIGN KEY (follower_id) REFERENCES account(account_id)
);