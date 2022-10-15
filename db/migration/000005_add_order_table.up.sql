CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "status" varchar NOT NULL,
  "delivery_fee" numeric NOT NULL,
  "subtotal" numeric NOT NULL,
  "total" numeric NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);
ALTER TABLE "orders" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

CREATE INDEX ON "orders" ("owner");
