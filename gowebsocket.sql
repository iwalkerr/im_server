/*
 Navicat Premium Data Transfer

 Source Server         : 47.108.146.201
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : 47.108.146.201:3306
 Source Schema         : gowebsocket

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : 65001

 Date: 01/03/2021 08:48:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for im_order
-- ----------------------------
DROP TABLE IF EXISTS `im_order`;
CREATE TABLE `im_order` (
  `order_id` bigint(11) NOT NULL AUTO_INCREMENT,
  `is_team` tinyint(4) DEFAULT '0' COMMENT '是否是群，0不是，1是',
  `team_id` int(11) DEFAULT '0' COMMENT '群ID',
  `team_ower` int(11) DEFAULT NULL COMMENT '群主ID',
  `sponsor` int(11) DEFAULT NULL COMMENT '发起人ID',
  `money` int(11) DEFAULT '0' COMMENT '总金额 分',
  `is_over` tinyint(4) DEFAULT '0' COMMENT '0未结算，1已结算',
  `team_fee` int(11) DEFAULT NULL COMMENT '群主服务费 %',
  `server_fee` int(11) DEFAULT NULL COMMENT '平台维护费',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`order_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of im_order
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for im_order_detail
-- ----------------------------
DROP TABLE IF EXISTS `im_order_detail`;
CREATE TABLE `im_order_detail` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_id` int(11) NOT NULL COMMENT '订单ID(群收款订单号)',
  `pay_user` int(11) DEFAULT NULL COMMENT '支付用户ID',
  `money` int(11) DEFAULT NULL COMMENT '金额',
  `is_pay` tinyint(4) DEFAULT NULL COMMENT '是否已经支付: 0未支付 1已支付',
  `pay_time` datetime DEFAULT NULL COMMENT '支付时间',
  `pay_orderno` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '支付订单号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of im_order_detail
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for im_payinfo
-- ----------------------------
DROP TABLE IF EXISTS `im_payinfo`;
CREATE TABLE `im_payinfo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pay_name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '支付名称',
  `pay_code` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '支付编码',
  `hand_fee` int(11) DEFAULT NULL COMMENT '手续费 %',
  `status` tinyint(4) DEFAULT NULL COMMENT '0正常，1禁用',
  `icon` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '图标',
  `remark` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `param` json DEFAULT NULL COMMENT '支付json参数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of im_payinfo
-- ----------------------------
BEGIN;
INSERT INTO `im_payinfo` VALUES (1, '支付宝', 'alipay', 60, 0, '', '支付宝支付', NULL);
INSERT INTO `im_payinfo` VALUES (2, '微信支付', 'wxpay', 60, 0, '', '微信支付', NULL);
INSERT INTO `im_payinfo` VALUES (3, '余额支付', 'fdpay', 0, 0, '', '凤蝶余额支付', NULL);
COMMIT;

-- ----------------------------
-- Table structure for im_team_chat
-- ----------------------------
DROP TABLE IF EXISTS `im_team_chat`;
CREATE TABLE `im_team_chat` (
  `team_chat_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `deleted` tinyint(1) DEFAULT NULL COMMENT '是否删除 1删除 0正常',
  `team_id` int(11) DEFAULT '0' COMMENT '群ID',
  `from_user_id` int(11) DEFAULT NULL COMMENT '消息发送者',
  `content` varchar(255) DEFAULT NULL COMMENT '消息内容',
  `created_time` datetime DEFAULT NULL COMMENT '发送时间',
  `msg_type` tinyint(4) DEFAULT '0' COMMENT '消息类型',
  PRIMARY KEY (`team_chat_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='群聊天表';

-- ----------------------------
-- Records of im_team_chat
-- ----------------------------
BEGIN;
INSERT INTO `im_team_chat` VALUES (1, 0, 0, 1, '我发送了消息', '2021-02-24 11:20:10', 0);
COMMIT;

-- ----------------------------
-- Table structure for im_team_member
-- ----------------------------
DROP TABLE IF EXISTS `im_team_member`;
CREATE TABLE `im_team_member` (
  `team_member_id` int(11) NOT NULL AUTO_INCREMENT,
  `deleted` tinyint(1) DEFAULT NULL COMMENT '是否删除1是 0正常',
  `team_id` int(11) NOT NULL DEFAULT '0' COMMENT '群ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `last_chat_id` bigint(20) DEFAULT '0' COMMENT '进群消息位置ID',
  `status` tinyint(4) DEFAULT '0' COMMENT '0 待审核 1 通过 2 拒绝',
  `invitation` int(11) DEFAULT '0' COMMENT '邀请人ID',
  `created_time` datetime DEFAULT NULL COMMENT '进群时间',
  PRIMARY KEY (`team_member_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='群成员表';

-- ----------------------------
-- Records of im_team_member
-- ----------------------------
BEGIN;
INSERT INTO `im_team_member` VALUES (1, 0, 1, 1, 1, 0, 0, '2021-02-24 11:19:26');
COMMIT;

-- ----------------------------
-- Table structure for im_teams
-- ----------------------------
DROP TABLE IF EXISTS `im_teams`;
CREATE TABLE `im_teams` (
  `team_id` int(11) NOT NULL AUTO_INCREMENT,
  `team_identifier` varchar(64) DEFAULT NULL COMMENT '群ID',
  `head_picture` varchar(255) DEFAULT NULL COMMENT '群头像',
  `team_qrcode` varchar(100) DEFAULT '' COMMENT '群二维码',
  `name` varchar(64) DEFAULT NULL COMMENT '群名称',
  `member` int(11) DEFAULT '1000' COMMENT '最大成员数量，默认为1000',
  `announcement` varchar(255) DEFAULT NULL COMMENT '群公告',
  `creator_id` int(11) DEFAULT NULL COMMENT '群主',
  `status` int(11) DEFAULT '0' COMMENT '群状态：0待审核,1通过,2拒绝,3解散',
  `join_check` tinyint(4) DEFAULT '0' COMMENT '加入群是否需要审核,1审核，0不审核',
  `allow_user` tinyint(4) DEFAULT '0' COMMENT '互加好友 1允许 0禁止',
  `service_fee` int(11) DEFAULT '0' COMMENT '平台维护费 %',
  `team_fee` int(11) DEFAULT '0' COMMENT '团长服务费 %',
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`team_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='群';

-- ----------------------------
-- Records of im_teams
-- ----------------------------
BEGIN;
INSERT INTO `im_teams` VALUES (1, '10001', 'http:/www.baidu.com', 'http:/www.baidu.com', '我是群名称', 1000, '群公告，支付宝支付', 1, 0, 0, 0, 1, 10, '2021-02-28 22:33:04', '2021-02-24 11:18:52');
COMMIT;

-- ----------------------------
-- Table structure for im_user_drawal
-- ----------------------------
DROP TABLE IF EXISTS `im_user_drawal`;
CREATE TABLE `im_user_drawal` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tx_no` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '体现订单编号',
  `user_id` int(11) DEFAULT NULL,
  `type` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '体现通道编号',
  `account_info` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '体现账号信息',
  `money` int(11) DEFAULT NULL COMMENT '体现金额 %',
  `pay_fee` int(11) DEFAULT NULL COMMENT '提现费用 %',
  `create_time` datetime DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL COMMENT '1处理中，2取消，3完成',
  `done_time` datetime DEFAULT NULL COMMENT '完成时间',
  `out_no` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '第三方单号',
  `pay_order_no` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '第三方资金号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of im_user_drawal
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for im_user_friend
-- ----------------------------
DROP TABLE IF EXISTS `im_user_friend`;
CREATE TABLE `im_user_friend` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL COMMENT '用户ID',
  `friend_id` int(11) DEFAULT NULL COMMENT '好友ID',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '好友请求备注',
  `status` tinyint(4) DEFAULT NULL COMMENT '0 待通过 1 通过 2 拒绝',
  `add_time` datetime DEFAULT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of im_user_friend
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for im_user_record
-- ----------------------------
DROP TABLE IF EXISTS `im_user_record`;
CREATE TABLE `im_user_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL COMMENT '用户ID',
  `type` tinyint(4) DEFAULT NULL COMMENT '1资金，2积分',
  `order_type` tinyint(4) DEFAULT NULL COMMENT '1充值,2红包，3活动账单，4提现到支付宝余额，5活动账单佣金',
  `orderno` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '订单号',
  `change_type` tinyint(4) DEFAULT NULL COMMENT '1增加，2减少',
  `change_fee` int(11) DEFAULT NULL COMMENT '自动变动金额 %',
  `balance` int(11) DEFAULT NULL COMMENT '账户金额 %',
  `check_code` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'md5加密校验字符串',
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of im_user_record
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for im_users
-- ----------------------------
DROP TABLE IF EXISTS `im_users`;
CREATE TABLE `im_users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `deleted` tinyint(4) DEFAULT '0' COMMENT '是否注销账号，默认0，1注销',
  `username` varchar(10) NOT NULL COMMENT '用户唯一名称',
  `name` varchar(128) DEFAULT '' COMMENT '用户显示名',
  `head_picture` varchar(255) DEFAULT '' COMMENT '用户头像',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '加密的密码',
  `signature` varchar(256) DEFAULT '' COMMENT '个性签名',
  `phone` varchar(20) DEFAULT '' COMMENT '手机号码',
  `sex` tinyint(1) DEFAULT '0' COMMENT '1男，2女，0保密',
  `user_addr_id` int(11) DEFAULT '0' COMMENT '用户地址ID',
  `openid` varchar(100) DEFAULT NULL COMMENT '第三方登陆OpendID',
  `face_img` varchar(100) DEFAULT NULL COMMENT '二维码地址',
  `status` tinyint(4) DEFAULT NULL COMMENT '0 正常 1 锁定',
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='APP用户表';

-- ----------------------------
-- Records of im_users
-- ----------------------------
BEGIN;
INSERT INTO `im_users` VALUES (1, 0, 'ewe3442', '张三👏', '47.108.146.201:8080/static/public/images/default/18.jpeg', 'f22b0a53ac96db2b1a535899eafc1eee', '2werr', '13881887710', 0, 0, NULL, NULL, NULL, '2021-02-28 17:03:46', '2021-02-24 03:23:09');
INSERT INTO `im_users` VALUES (6, 0, 'JMa24914', '张三', '47.108.146.201:8080/static/public/images/default/28.jpeg', 'cb499c9e3f98c79ce1d1956a826137da', 'Iwi17y', '13881887711', 1, 0, NULL, NULL, NULL, '2021-02-28 16:31:48', '2021-02-28 16:31:48');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
