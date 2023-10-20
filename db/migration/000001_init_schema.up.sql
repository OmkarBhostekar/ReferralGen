CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "templates" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "name" varchar,
  "template" text,
  "params" varchar[],
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "templates" ("name");

ALTER TABLE "templates" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");