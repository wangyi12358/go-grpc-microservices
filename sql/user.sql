/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : PostgreSQL
 Source Server Version : 140007
 Source Host           : localhost:5432
 Source Catalog        : user
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140007
 File Encoding         : 65001

 Date: 09/06/2024 22:19:39
*/


-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "public"."user";
CREATE TABLE "public"."user" (
  "id" text COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "username" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "status" int4 NOT NULL DEFAULT 1,
  "created_id" text COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::text,
  "updated_id" text COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::text,
  "created_at" timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(3) NOT NULL
)
;
ALTER TABLE "public"."user" OWNER TO "postgres";

-- ----------------------------
-- Indexes structure for table user
-- ----------------------------
CREATE UNIQUE INDEX "User_username_key" ON "public"."user" USING btree (
  "username" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table user
-- ----------------------------
ALTER TABLE "public"."user" ADD CONSTRAINT "User_pkey" PRIMARY KEY ("id");
