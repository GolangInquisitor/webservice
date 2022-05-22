/*
 Navicat Premium Data Transfer

 Source Server         : mainmase
 Source Server Type    : PostgreSQL
 Source Server Version : 140002
 Source Host           : localhost:5432
 Source Catalog        : postgres
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140002
 File Encoding         : 65001

 Date: 20/05/2022 17:08:48
*/


-- ----------------------------
-- Sequence structure for price_inc
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."price_inc";
CREATE SEQUENCE "public"."price_inc" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS "public"."orders";
CREATE TABLE "public"."orders" (
  "uuid" uuid NOT NULL DEFAULT gen_random_uuid(),
  "id" int8,
  "product" uuid,
  "user_id" uuid,
  "price" numeric,
  "description" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Table structure for price
-- ----------------------------
DROP TABLE IF EXISTS "public"."price";
CREATE TABLE "public"."price" (
  "id" int8 NOT NULL DEFAULT nextval('price_inc'::regclass),
  "product" uuid NOT NULL,
  "currency" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "amount" numeric NOT NULL
)
;

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS "public"."product";
CREATE TABLE "public"."product" (
  "uuid" uuid NOT NULL DEFAULT gen_random_uuid(),
  "description" text COLLATE "pg_catalog"."default",
  "left_in_stock" int8,
  "delete" bool DEFAULT false
)
;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "uuid" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" varchar(255) COLLATE "pg_catalog"."default",
  "surname" varchar(255) COLLATE "pg_catalog"."default",
  "midlename" varchar(255) COLLATE "pg_catalog"."default",
  "fio" varchar(255) COLLATE "pg_catalog"."default",
  "gender" varchar(255) COLLATE "pg_catalog"."default",
  "age" int4
)
;

-- ----------------------------
-- Function structure for create_product
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."create_product"("decription" varchar, "cur" varchar, "amo" numeric, "left_in_stock" int4);
CREATE OR REPLACE FUNCTION "public"."create_product"("decription" varchar, "cur" varchar, "amo" numeric, "left_in_stock" int4)
  RETURNS "pg_catalog"."int4" AS $BODY$
	  DECLARE
  res uuid;
	res2 int;
	BEGIN 
	-- Routine body goes here...
 

  insert into product (description, left_in_stock) values (decription, left_in_stock) returning product.uuid into res;
  insert into price  (currency, amount, product) VALUES (cur,amo,res) returning price.id into res2;

	
	RETURN res2;
END$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Function structure for fio_create
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."fio_create"();
CREATE OR REPLACE FUNCTION "public"."fio_create"()
  RETURNS "pg_catalog"."trigger" AS $BODY$
	BEGIN
	-- Ro
	NEW.fio=concat_ws(' ',NEW.SURNAME,NEW.NAME,NEW.MIDLENAME);
 
	RETURN new;
END$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
SELECT setval('"public"."price_inc"', 19, true);

-- ----------------------------
-- Primary Key structure for table orders
-- ----------------------------
ALTER TABLE "public"."orders" ADD CONSTRAINT "order_pkey" PRIMARY KEY ("uuid");

-- ----------------------------
-- Primary Key structure for table price
-- ----------------------------
ALTER TABLE "public"."price" ADD CONSTRAINT "price_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table product
-- ----------------------------
ALTER TABLE "public"."product" ADD CONSTRAINT "product_pkey" PRIMARY KEY ("uuid");

-- ----------------------------
-- Triggers structure for table users
-- ----------------------------
CREATE TRIGGER "fio_before_ins" BEFORE INSERT OR UPDATE ON "public"."users"
FOR EACH ROW
EXECUTE PROCEDURE "public"."fio_create"();

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("uuid");

-- ----------------------------
-- Foreign Keys structure for table orders
-- ----------------------------
ALTER TABLE "public"."orders" ADD CONSTRAINT "product_fk" FOREIGN KEY ("product") REFERENCES "public"."product" ("uuid") ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE "public"."orders" ADD CONSTRAINT "user_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("uuid") ON DELETE CASCADE ON UPDATE RESTRICT;

-- ----------------------------
-- Foreign Keys structure for table price
-- ----------------------------
ALTER TABLE "public"."price" ADD CONSTRAINT "product_fk" FOREIGN KEY ("product") REFERENCES "public"."product" ("uuid") ON DELETE CASCADE ON UPDATE RESTRICT;
