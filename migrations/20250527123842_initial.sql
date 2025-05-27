-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
        user_id         SERIAL          NOT NULL,
        login           TEXT            NOT NULL UNIQUE,
        password_hash   TEXT            NOT NULL,
        name            TEXT,
        points          INTEGER         DEFAULT 0,
        referrer_id     INTEGER,
        created_at      TIMESTAMPTZ     DEFAULT CURRENT_TIMESTAMP,
        updated_at      TIMESTAMPTZ     DEFAULT CURRENT_TIMESTAMP,

        PRIMARY KEY (user_id),
        FOREIGN KEY (referrer_id) REFERENCES users (user_id) ON DELETE SET NULL
);

CREATE TYPE task_name_enum AS ENUM (
    'join_telegram',
    'follow_twitter',
    'invite_friend',
    'fill_profile'
);

CREATE TABLE IF NOT EXISTS tasks (
    task_id         SERIAL              NOT NULL,
    name            task_name_enum      NOT NULL UNIQUE,
    description     TEXT,
    points          INTEGER             NOT NULL,
    created_at      TIMESTAMPTZ         DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (task_id)
);

CREATE TABLE IF NOT EXISTS user_tasks (
    user_task_id    SERIAL      NOT NULL,
    user_id         INTEGER     NOT NULL,
    task_id         INTEGER     NOT NULL,
    completed_at    TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (user_task_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (task_id) REFERENCES tasks(task_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_tasks;
DROP TABLE IF EXISTS tasks;
DROP TYPE IF EXISTS task_name_enum;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
