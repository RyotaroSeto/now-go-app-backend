-- +goose NO TRANSACTION

-- +goose Up

CREATE TABLE IF NOT EXISTS nowgo.users (
    id INT PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS nowgo.users_detail (
    id INT PRIMARY KEY,
    user_id INT,
    date_of_birth DATE NOT NULL,
    gender VARCHAR(255) NOT NULL,
    residence VARCHAR(255) NOT NULL,
    occupation VARCHAR(255) NOT NULL,
    height INT,
    weight INT,
    FOREIGN KEY (user_id) REFERENCES nowgo.users(id)
);

CREATE INDEX ON nowgo.users_detail (user_id);

CREATE TABLE IF NOT EXISTS nowgo.likes (
    user_id INT,
    liked_user_id INT,
    liked_date DATE,
    PRIMARY KEY (user_id, liked_user_id),
    FOREIGN KEY (user_id) REFERENCES nowgo.users(id),
    FOREIGN KEY (liked_user_id) REFERENCES nowgo.users(id)
);

CREATE INDEX ON nowgo.likes (user_id);
CREATE INDEX ON nowgo.likes (liked_user_id);

CREATE TABLE IF NOT EXISTS nowgo.match (
    user_id INT,
    matched_user_id INT,
    matched_date DATE,
    PRIMARY KEY (user_id, matched_user_id),
    FOREIGN KEY (user_id) REFERENCES nowgo.users(id),
    FOREIGN KEY (matched_user_id) REFERENCES nowgo.users(id)
);

CREATE INDEX ON nowgo.match (user_id);
CREATE INDEX ON nowgo.match (matched_user_id);

CREATE TABLE IF NOT EXISTS nowgo.message (
    id INT PRIMARY KEY,
    sender_user_id INT,
    receiver_user_id INT,
    message_body VARCHAR(255),
    sent_date DATE,
    FOREIGN KEY (sender_user_id) REFERENCES nowgo.users(id),
    FOREIGN KEY (receiver_user_id) REFERENCES nowgo.users(id)
);

CREATE INDEX ON nowgo.message (sender_user_id);
CREATE INDEX ON nowgo.message (receiver_user_id);

-- +goose Down

DROP TABLE nowgo.users;

DROP TABLE nowgo.users_detail;

DROP TABLE nowgo.likes;

DROP TABLE nowgo.match;

DROP TABLE nowgo.message;
