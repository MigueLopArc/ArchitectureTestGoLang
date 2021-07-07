CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS  pgcrypto;

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "email" VARCHAR(320) NOT NULL,
  "first_name" VARCHAR(100),
  "last_name" VARCHAR(100),
  "password" VARCHAR(1000) NOT NULL,
  "creation_date" timestamp DEFAULT (NOW())
);

CREATE TABLE "notes" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "user_id" uuid NOT NULL,
  "title" VARCHAR(200) NOT NULL,
  "content" TEXT,
  "creation_date" timestamp DEFAULT (NOW()),
  "modification_date" timestamp DEFAULT (NOW())
);



ALTER TABLE "notes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");