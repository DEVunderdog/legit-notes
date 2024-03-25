CREATE TABLE "users" (
  "id" SERIAL,
  "username" VARCHAR UNIQUE NOT NULL PRIMARY KEY,
  "email" VARCHAR UNIQUE NOT NULL,
  "hashed_password" VARCHAR NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "notes" (
  "id" SERIAL,
  "user_id" VARCHAR UNIQUE NOT NULL PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp
);

ALTER TABLE "notes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("username");
