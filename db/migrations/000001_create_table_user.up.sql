CREATE TABLE users (
    id char(26) PRIMARY KEY,
    nip VARCHAR(20) NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL,
    password VARCHAR(100),
    role VARCHAR(10) NOT NULL,
    identity_card_scan_img TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_user_name ON users (name);
CREATE INDEX IF NOT EXISTS idx_user_nip ON users (nip);
CREATE INDEX IF NOT EXISTS idx_user_role ON users (role);
CREATE INDEX IF NOT EXISTS idx_user_created_at ON users (created_at);
