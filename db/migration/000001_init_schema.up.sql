CREATE TABLE "articles" (
  "id" bigserial PRIMARY KEY,
  "poster_id" bigint NOT NULL,
  "content" varchar NOT NULL,
  "title" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "article_id" bigint NOT NULL,
  "content" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "articles" ("id");

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "comments" ("id");

CREATE INDEX ON "comments" ("article_id");

ALTER TABLE "articles" ADD FOREIGN KEY ("poster_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("article_id") REFERENCES "articles" ("id");
