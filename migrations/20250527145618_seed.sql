-- +goose Up
-- +goose StatementBegin
INSERT INTO tasks (name, points)
VALUES
    ('join_telegram', 50),
    ('follow_twitter', 70),
    ('invite_friend', 100),
    ('fill_profile', 10);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM tasks WHERE name IN ('join_telegram', 'follow_twitter', 'invite_friend', fill_profile);
-- +goose StatementEnd
