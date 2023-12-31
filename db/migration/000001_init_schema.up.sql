CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "fullname" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "templates" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "name" varchar NOT NULL,
  "template" text NOT NULL,
  "params" varchar[],
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "generations" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "created_date" varchar NOT NULL,
  "count" int DEFAULT 0
);

CREATE INDEX ON "templates" ("name");

CREATE INDEX ON "generations" ("created_date", "user_id");

ALTER TABLE "templates" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "generations" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
