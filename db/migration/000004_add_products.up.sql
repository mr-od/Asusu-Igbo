CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "owner" varchar NOT NULL,
  "price" numeric NOT NULL,
  "description" varchar NOT NULL,
  "imgs_url" text[],
  "imgs_name" text[],
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);



ALTER TABLE "products" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

CREATE INDEX ON "products" ("owner");

ALTER TABLE "products" ADD COLUMN "tsv" tsvector;

UPDATE "products" SET "tsv" =
    setweight(to_tsvector(name), 'A') ||
    setweight(to_tsvector(description), 'B');


CREATE INDEX ix_products_tsv ON "products" USING GIN(tsv);