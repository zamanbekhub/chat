-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS chat.chat(
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,
    chat_id SERIAL PRIMARY KEY,
    name VARCHAR,
    description VARCHAR,
    type_code VARCHAR NOT NULL
);
CREATE INDEX idx_chat_chat_id ON chat.chat(chat_id);
CREATE INDEX idx_chat_type_code ON chat.chat(type_code);

CREATE TABLE IF NOT EXISTS chat.user_chat_role(
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,
    chat_id INT REFERENCES chat.chat(chat_id) NOT NULL,
    user_id INT NOT NULL,
    role_id INT NOT NULL
);

CREATE INDEX idx_user_chat_role_chat_id ON chat.user_chat_role(chat_id);
CREATE INDEX idx_user_chat_role_user_id ON chat.user_chat_role(user_id);
CREATE INDEX idx_user_chat_role_role_id ON chat.user_chat_role(role_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX chat.idx_user_chat_role_role_id;
DROP INDEX chat.idx_user_chat_role_user_id;
DROP INDEX chat.idx_user_chat_role_chat_id;
DROP TABLE chat.chat_user_role;

DROP INDEX chat.idx_chat_type_code;
DROP INDEX chat.idx_chat_chat_id;
DROP TABLE chat.chat;
-- +goose StatementEnd
