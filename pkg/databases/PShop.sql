CREATE TABLE "users" (
  "id" bigint PRIMARY KEY DEFAULT CONCAT('U', LPAD(NEXTVAL('users_id_seq')::TEXT, 6, '0')),
  "username" VARCHAR UNIQUE NOT NULL,
  "password" VARCHAR NOT NULL,
  "email" VARCHAR UNIQUE,
  "role_id" INT NOT NULL ,
  "is_deleted" BOOLEAN NULL,
  "created_by" BIGINT NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_by" BIGINT NOT NULL,
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "oauth" (
  "id" bigint PRIMARY KEY,
  "user_id" bigint,
  "access_token" VARCHAR,
  "refresh_token" VARCHAR,
  "is_deleted" BOOLEAN NULL,
  "created_by" BIGINT NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_by" BIGINT NOT NULL,
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "roles" (
  "id" bigint PRIMARY KEY,
  "title" VARCHAR,
  "is_deleted" bool,
  "created_by" bigint,
  "created_at" timestamp,
  "updated_by" bigint,
  "updated_at" timestamp
);

CREATE TABLE "images" (
  "id" bigint PRIMARY KEY,
  "filename" VARCHAR,
  "url" VARCHAR,
  "product_id" VARCHAR,
  "is_deleted" bool,
  "created_by" bigint,
  "created_at" timestamp,
  "updated_by" bigint,
  "updated_at" timestamp
);

CREATE TABLE "products" (
  "id" bigint PRIMARY KEY,
  "title" VARCHAR,
  "description" VARCHAR,
  "price" float,
  "is_deleted" bool,
  "created_by" bigint,
  "created_at" timestamp,
  "updated_by" bigint,
  "updated_at" timestamp
);

CREATE TABLE "products_categories" (
  "product_id" bitint,
  "category_id" bitint,
  "is_deleted" bool,
  "created_by" bigint,
  "created_at" timestamp,
  "updated_by" bigint,
  "updated_at" timestamp
);

CREATE TABLE "categories" (
  "id" bigint PRIMARY KEY,
  "title" VARCHAR UNIQUE,
  "is_deleted" bool,
  "created_by" bigint,
  "created_at" timestamp,
  "updated_by" bigint,
  "updated_at" timestamp
);

CREATE TABLE "orders" (
  "id" bigint PRIMARY KEY,
  "user_id" bigint,
  "contact" VARCHAR,
  "address" VARCHAR,
  "transfer_slip" jsonb,
  "status" VARCHAR,
  "is_deleted" bool,
  "created_by" bigint,
  "created_at" timestamp,
  "updated_by" bigint,
  "updated_at" timestamp
);

CREATE TABLE "products_orders" (
  "id" bigint PRIMARY KEY,
  "order_id" VARCHAR,
  "qty" INT,
  "product" jsonb,
  "is_deleted" bool,
  "created_by" bigint,
  "created_at" timestamp,
  "updated_by" bigint,
  "updated_at" timestamp
);

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");
ALTER TABLE "oauth" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "images" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
ALTER TABLE "products_categories" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
ALTER TABLE "products_categories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "products_orders" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

COMMIT;
