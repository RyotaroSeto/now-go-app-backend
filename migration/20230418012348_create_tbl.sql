-- +goose NO TRANSACTION

-- +goose Up

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    status INT NOT NULL DEFAULT 1,
    email VARCHAR(255) NOT NULL,
    created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users_details (
    id SERIAL PRIMARY KEY,
    user_id INT,
    date_of_birth INT NOT NULL,
    gender VARCHAR(255) NOT NULL,
    residence VARCHAR(255) NOT NULL,
    occupation VARCHAR(255) NOT NULL,
    height INT,
    weight INT,
    created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX ON users_details (user_id);

CREATE TABLE IF NOT EXISTS likes (
    user_id INT,
    liked_user_id INT,
    liked_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status INT NOT NULL DEFAULT 1,
    message_body VARCHAR(255),
    PRIMARY KEY (user_id, liked_user_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (liked_user_id) REFERENCES users(id)
);

CREATE INDEX ON likes (user_id);
CREATE INDEX ON likes (liked_user_id);


CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    sender_user_id INT,
    receiver_user_id INT,
    message_body VARCHAR(255),
    status INT NOT NULL DEFAULT 1,
    sent_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sender_user_id) REFERENCES users(id),
    FOREIGN KEY (receiver_user_id) REFERENCES users(id)
);

CREATE INDEX ON messages (sender_user_id);
CREATE INDEX ON messages (receiver_user_id);

CREATE TABLE IF NOT EXISTS boards (
    id SERIAL PRIMARY KEY,
    user_id INT,
    -- 目的
    -- 場所
    body VARCHAR(255),
    created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX ON boards (user_id);

-- CREATE TABLE IF NOT EXISTS match (
--     user_id INT,
--     matched_user_id INT,
--     matched_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,,
--     PRIMARY KEY (user_id, matched_user_id),
--     FOREIGN KEY (user_id) REFERENCES users(id),
--     FOREIGN KEY (matched_user_id) REFERENCES users(id)
-- );

-- CREATE INDEX ON match (user_id);
-- CREATE INDEX ON match (matched_user_id);
-- +goose Down

DROP TABLE users;

DROP TABLE users_details;

-- DROP TABLE match;

DROP TABLE messages;

DROP TABLE go_post;

DROP TABLE likes;
