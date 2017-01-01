/*
Navicat PGSQL Data Transfer

Source Server         : pgsql
Source Server Version : 90502
Source Host           : localhost:5432
Source Database       : pg_test
Source Schema         : public

Target Server Type    : PGSQL
Target Server Version : 90502
File Encoding         : 65001

Date: 2016-12-30 18:32:48
*/


-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS "public"."categories";
CREATE TABLE "public"."categories" (
"id" serial,
"name" varchar COLLATE "default",
"created_at" timestamp(6),
"updated_at" timestamp(6)
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Records of categories
-- ----------------------------
INSERT INTO "public"."categories" VALUES ('0', '休闲益智', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('1', '竞速游戏', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('2', '大作', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('3', '射击游戏', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('4', '角色扮演', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('5', '动作游戏', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('6', '卡牌游戏', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('7', 'bt网游', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('8', '格斗游戏', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('9', '体育运动', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('10', '养成游戏', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('11', '飞行游戏', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('12', '策略塔防', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('13', '音乐游戏', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('14', '实用软件', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('15', '装机必备', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('16', '生活实用', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('17', '影音娱乐', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('18', '社交聊天', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');
INSERT INTO "public"."categories" VALUES ('19', '系统工具', '2016-12-30 16:15:24.765703', '2016-12-30 16:15:24.765703');

-- ----------------------------
-- Alter Sequences Owned By 
-- ----------------------------

-- ----------------------------
-- Primary Key structure for table categories
-- ----------------------------
ALTER TABLE "public"."categories" ADD PRIMARY KEY ("id");
