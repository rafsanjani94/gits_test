/*
 Navicat Premium Data Transfer

 Source Server         : pgsql-docker
 Source Server Type    : PostgreSQL
 Source Server Version : 140003
 Source Host           : localhost:5432
 Source Catalog        : postgres
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140003
 File Encoding         : 65001

 Date: 16/02/2023 22:08:08
*/


-- ----------------------------
-- Sequence structure for author_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."author_id_seq";
CREATE SEQUENCE "public"."author_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for book_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."book_id_seq";
CREATE SEQUENCE "public"."book_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for publisher_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."publisher_id_seq";
CREATE SEQUENCE "public"."publisher_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for users_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."users_id_seq";
CREATE SEQUENCE "public"."users_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Table structure for author
-- ----------------------------
DROP TABLE IF EXISTS "public"."author";
CREATE TABLE "public"."author" (
  "id" int8 NOT NULL DEFAULT nextval('author_id_seq'::regclass),
  "name" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of author
-- ----------------------------

-- ----------------------------
-- Table structure for book
-- ----------------------------
DROP TABLE IF EXISTS "public"."book";
CREATE TABLE "public"."book" (
  "id" int8 NOT NULL DEFAULT nextval('book_id_seq'::regclass),
  "name" text COLLATE "pg_catalog"."default",
  "author_id" int8,
  "publisher_id" int8
)
;

-- ----------------------------
-- Records of book
-- ----------------------------
INSERT INTO "public"."book" VALUES (1, 'menikah', 1, 1);
INSERT INTO "public"."book" VALUES (4, 'dolanan', 2, 1);
INSERT INTO "public"."book" VALUES (6, 'kuliner cuy', 1, 1);
INSERT INTO "public"."book" VALUES (11, 'agama', 1, 1);

-- ----------------------------
-- Table structure for publisher
-- ----------------------------
DROP TABLE IF EXISTS "public"."publisher";
CREATE TABLE "public"."publisher" (
  "id" int8 NOT NULL DEFAULT nextval('publisher_id_seq'::regclass),
  "name" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of publisher
-- ----------------------------

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" int8 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
  "author_id" int8,
  "username" text COLLATE "pg_catalog"."default",
  "password" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO "public"."users" VALUES (1, 1, 'aliraf', '146e5ea66b4b6036c6bbf395cc90c30e');
INSERT INTO "public"."users" VALUES (2, 2, 'sanjani', '146e5ea66b4b6036c6bbf395cc90c30e');

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."author_id_seq"
OWNED BY "public"."author"."id";
SELECT setval('"public"."author_id_seq"', 2, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."book_id_seq"
OWNED BY "public"."book"."id";
SELECT setval('"public"."book_id_seq"', 12, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."publisher_id_seq"
OWNED BY "public"."publisher"."id";
SELECT setval('"public"."publisher_id_seq"', 2, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."users_id_seq"
OWNED BY "public"."users"."id";
SELECT setval('"public"."users_id_seq"', 3, true);

-- ----------------------------
-- Primary Key structure for table author
-- ----------------------------
ALTER TABLE "public"."author" ADD CONSTRAINT "author_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table book
-- ----------------------------
ALTER TABLE "public"."book" ADD CONSTRAINT "book_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table publisher
-- ----------------------------
ALTER TABLE "public"."publisher" ADD CONSTRAINT "publisher_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");
