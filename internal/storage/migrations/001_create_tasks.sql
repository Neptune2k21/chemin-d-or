-- migrate:up
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS tasks (
    id          UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    title       TEXT        NOT NULL,
    description TEXT,
    due_date    TIMESTAMP,
    created_at  TIMESTAMP   NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP   NOT NULL DEFAULT NOW()
);

-- migrate:down
DROP TABLE IF EXISTS tasks;
