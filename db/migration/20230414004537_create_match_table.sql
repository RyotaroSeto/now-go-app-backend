-- user table
CREATE TABLE user (
    id INT PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
);

-- user_detail table
CREATE TABLE user_detail (
    id INT PRIMARY KEY,
    user_id INT,
    date_of_birth DATE NOT NULL,
    gender VARCHAR(255) NOT NULL,
    residence VARCHAR(255) NOT NULL,
    occupation VARCHAR(255) NOT NULL,
    education VARCHAR(255) NOT NULL,
    height INT,
    weight INT,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

-- like table
CREATE TABLE LIKE (
    user_id INT,
    liked_user_id INT,
    liked_date DATE,
    PRIMARY KEY (user_id, liked_user_id),
    FOREIGN KEY (user_id) REFERENCES USER(id),
    FOREIGN KEY (liked_user_id) REFERENCES USER(id)
);


-- match table
CREATE TABLE match (
    user_id INT,
    matched_user_id INT,
    matched_date DATE,
    PRIMARY KEY (user_id, matched_user_id),
    FOREIGN KEY (user_id) REFERENCES USER(id),
    FOREIGN KEY (matched_user_id) REFERENCES USER(id)
);


-- message table
CREATE TABLE message (
    id INT PRIMARY KEY,
    sender_user_id INT,
    receiver_user_id INT,
    message_body VARCHAR(255),
    sent_date DATE,
    FOREIGN KEY (sender_user_id) REFERENCES user(id),
    FOREIGN KEY (receiver_user_id) REFERENCES user(id)
);
