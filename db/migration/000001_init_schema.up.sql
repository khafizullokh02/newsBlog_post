CREATE TYPE "post_type" AS ENUM (
  'blog',
  'news',
  ' '
);

CREATE TABLE "category" (
  "id" bigserial PRIMARY KEY,
  "title" int NOT NULL,
  "post_count" int NOT NULL,
  "article_count" int NOT NULL
);

CREATE TABLE "comment" (
  "id" bigserial PRIMARY KEY,
  "comment" varchar NOT NULL,
  "parent_id" int DEFAULT NULL,
  "user_id" integer NOT NULL,
  "like_count" integer NOT NULL,
  "post_id" int NOT NULL
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "full_name" varchar NOT NULL,
  "user_avatar" int NOT NULL,
  "count_posts" int NOT NULL,
  "email" varchar NOT NULL,
  "pasword" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "post" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" int NOT NULL,
  "title" varchar NOT NULL,
  "content" varchar NOT NULL,
  "category_id" int NOT NULL,
  "post_type" post_type NOT NULL DEFAULT ' ',
  "like_count" int NOT NULL,
  "comment_count" int NOT NULL,
  "view_count" int NOT NULL,
  "published_at" timestamp,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

ALTER TABLE "post" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comment" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "post" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");
