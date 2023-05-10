BEGIN;


DROP TRIGGER IF EXISTS set_updated_at_timestamp_users_table ON "users";
DROP TRIGGER IF EXISTS set_updated_at_timestamp_oauth_table ON "oauth";
DROP TRIGGER IF EXISTS set_updated_at_timestamp_products_table ON "products";
DROP TRIGGER IF EXISTS set_updated_at_timestamp_images_table ON "images";
DROP TRIGGER IF EXISTS set_updated_at_timestamp_orders_table ON "orders";
DROP TRIGGER IF EXISTS set_updated_at_timestamp_products_orders_table ON "products_orders";
DROP TRIGGER IF EXISTS set_updated_at_timestamp_categories_table ON "categories";
DROP TRIGGER IF EXISTS set_updated_at_timestamp_products_categories_table ON "products_categories";
DROP TRIGGER IF EXISTS set_updated_at_timestamp_roles_table ON "roles";

DROP FUNCTION IF EXISTS set_updated_at_column();

DROP TABLE IF EXISTS "users" CASCADE;
DROP TABLE IF EXISTS "oauth" CASCADE;
DROP TABLE IF EXISTS "roles" CASCADE;
DROP TABLE IF EXISTS "products" CASCADE;
DROP TABLE IF EXISTS "categories" CASCADE;
DROP TABLE IF EXISTS "products_categories" CASCADE;
DROP TABLE IF EXISTS "images" CASCADE;
DROP TABLE IF EXISTS "orders" CASCADE;
DROP TABLE IF EXISTS "products_orders" CASCADE;

DROP SEQUENCE IF EXISTS users_id_seq;
DROP SEQUENCE IF EXISTS products_id_seq;
DROP SEQUENCE IF EXISTS orders_id_seq;

DROP TYPE IF EXISTS "order_status";

COMMIT;