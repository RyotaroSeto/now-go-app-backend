-- +goose NO TRANSACTION

-- +goose Up

CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL PRIMARY KEY,
    "user_name" VARCHAR(255) UNIQUE NOT NULL,
    "hashed_password" VARCHAR(255) NOT NULL,
    "status" INT NOT NULL DEFAULT 1,
    "email" VARCHAR(255) NOT NULL,
    "created_date" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_date" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "sessions" (
    "id" UUID PRIMARY KEY,
    "user_name" VARCHAR(255) UNIQUE NOT NULL,
    "refresh_token" VARCHAR NOT NULL,
    "user_agent" VARCHAR NOT NULL,
    "client_ip" VARCHAR NOT NULL,
    "is_blocked" BOOLEAN NOT NULL DEFAULT false,
    "expires_date" TIMESTAMP NOT NULL,
    "created_date" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("user_name") REFERENCES "users" ("user_name")
);

CREATE TABLE IF NOT EXISTS "users_details" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INT,
    "name" VARCHAR(255) NOT NULL,
	"age" INT NOT NULL,
	"gender" VARCHAR(255) NOT NULL,
	"height" INT,
	"location" VARCHAR(255),
	"work" VARCHAR(255),
	"graduation" VARCHAR(255),
	"hobby" VARCHAR(255),
	"passion" VARCHAR(255),
	"tweet" VARCHAR,
	"introduction" VARCHAR,
    "created_date" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_date" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

CREATE INDEX ON "users_details" ("user_id");

CREATE TABLE IF NOT EXISTS "likes" (
    "user_id" INT,
    "liked_user_id" INT,
    "liked_date" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "status" INT NOT NULL DEFAULT 1,
    "message_body" VARCHAR(255),
    PRIMARY KEY ("user_id", "liked_user_id"),
    FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
    FOREIGN KEY ("liked_user_id") REFERENCES "users" ("id")
);

CREATE INDEX ON "likes" ("user_id");
CREATE INDEX ON "likes" ("liked_user_id");


CREATE TABLE IF NOT EXISTS "messages" (
    "id" SERIAL PRIMARY KEY,
    "sender_user_id" INT,
    "receiver_user_id" INT,
    "message_body" VARCHAR(255),
    "status" INT NOT NULL DEFAULT 1,
    "sent_date" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("sender_user_id") REFERENCES "users" ("id"),
    FOREIGN KEY ("receiver_user_id") REFERENCES "users" ("id")
);

CREATE INDEX ON "messages" ("sender_user_id");
CREATE INDEX ON "messages" ("receiver_user_id");

CREATE TABLE IF NOT EXISTS "boards" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INT,
    "body" VARCHAR(255),
    "created_date" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_date" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

CREATE INDEX ON "boards" ("user_id");


-- +goose Down

DROP TABLE "users";

DROP TABLE "users_details";

DROP TABLE "messages";

DROP TABLE "go_post";

DROP TABLE "likes";
