CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE habits (
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL,
    title VARCHAR(100) NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT TRUE,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW()
);