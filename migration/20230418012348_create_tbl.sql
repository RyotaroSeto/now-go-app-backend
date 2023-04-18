-- +goose NO TRANSACTION

-- +goose Up

CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users_detail (
    id INT PRIMARY KEY,
    user_id INT,
    date_of_birth TIMESTAMP NOT NULL,
    gender VARCHAR(255) NOT NULL,
    residence VARCHAR(255) NOT NULL,
    occupation VARCHAR(255) NOT NULL,
    height INT,
    weight INT,
    created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX ON users_detail (user_id);

CREATE TABLE IF NOT EXISTS likes (
    user_id INT,
    liked_user_id INT,
    liked_date DATE,
    PRIMARY KEY (user_id, liked_user_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (liked_user_id) REFERENCES users(id)
);

CREATE INDEX ON likes (user_id);
CREATE INDEX ON likes (liked_user_id);

CREATE TABLE IF NOT EXISTS match (
    user_id INT,
    matched_user_id INT,
    matched_date DATE,
    PRIMARY KEY (user_id, matched_user_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (matched_user_id) REFERENCES users(id)
);

CREATE INDEX ON match (user_id);
CREATE INDEX ON match (matched_user_id);

CREATE TABLE IF NOT EXISTS message (
    id INT PRIMARY KEY,
    sender_user_id INT,
    receiver_user_id INT,
    message_body VARCHAR(255),
    sent_date DATE,
    FOREIGN KEY (sender_user_id) REFERENCES users(id),
    FOREIGN KEY (receiver_user_id) REFERENCES users(id)
);

CREATE INDEX ON message (sender_user_id);
CREATE INDEX ON message (receiver_user_id);

-- +goose Down

DROP TABLE users;

DROP TABLE users_detail;

DROP TABLE likes;

DROP TABLE match;

DROP TABLE message;
