/*
 Navicat Premium Data Transfer

 Source Server         : promotion
 Source Server Type    : MySQL
 Source Server Version : 50720
 Source Host           : 192.168.230.144:3306
 Source Schema         : alpha

 Target Server Type    : MySQL
 Target Server Version : 50720
 File Encoding         : 65001

 Date: 05/08/2019 17:49:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `icon` varchar(255) NOT NULL DEFAULT '' COMMENT 'Icon\n',
  `label` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '菜单名',
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父id',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '菜单的url',
  `level` tinyint(4) NOT NULL DEFAULT '0' COMMENT '层级',
  `cond` varchar(2000) NOT NULL DEFAULT '' COMMENT '条件',
  `component` varchar(500) NOT NULL DEFAULT '' COMMENT '前端引入的组件',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='权限表';

-- ----------------------------
-- Records of permission
-- ----------------------------
BEGIN;
INSERT INTO `permission` VALUES (2, '', '仪表盘', 0, '/dashboard/workplace', 1, 'alpha/handler/admin/role.List,alpha/handler/admin/role.Get,alpha/handler/admin/permission.List', 'RouteView', 0, '2019-07-11 16:58:21', '2019-08-05 09:40:24');
INSERT INTO `permission` VALUES (3, '', '设置', 0, '/form', 1, '    ', 'PageView', 500, '2019-07-11 18:03:04', '2019-08-05 09:40:40');
INSERT INTO `permission` VALUES (4, '', '分析页', 2, '/dashboard/analysis', 1, 'alpha/handler/admin/user.Information', 'analysis', 500, '2019-07-11 18:03:20', '2019-08-05 09:40:57');
INSERT INTO `permission` VALUES (5, '', '权限', 2, 'https://www.baidu.com/', 1, '', '', 501, '2019-07-11 18:03:24', '2019-08-05 09:41:43');
INSERT INTO `permission` VALUES (6, '', '角色', 2, '/dashboard/workplace', 1, '', 'workplace', 500, '2019-07-11 18:03:30', '2019-08-05 09:41:41');
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '名称',
  `description` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '简介',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` VALUES (6, '测试新增', '1111sdadasd', '2019-07-12 16:05:10', '2019-07-12 16:05:25');
COMMIT;

-- ----------------------------
-- Table structure for role_permission
-- ----------------------------
DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE `role_permission` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色id',
  `permission_id` int(11) NOT NULL DEFAULT '0' COMMENT '权限id',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `role_id` (`role_id`,`permission_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of role_permission
-- ----------------------------
BEGIN;
INSERT INTO `role_permission` VALUES (24, 2, 2, '2019-07-12 08:04:54', '2019-07-22 08:52:55');
INSERT INTO `role_permission` VALUES (25, 2, 1, '2019-07-12 08:04:54', '2019-07-31 07:22:19');
INSERT INTO `role_permission` VALUES (30, 6, 6, '2019-07-12 08:05:25', '2019-07-12 08:05:25');
INSERT INTO `role_permission` VALUES (31, 6, 7, '2019-07-12 08:05:25', '2019-07-12 08:05:25');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '姓名',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '用户手机号',
  `password` varchar(100) NOT NULL DEFAULT '' COMMENT '用户密码',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像',
  `last_time` datetime NOT NULL COMMENT '上次登录时间',
  `last_ip` varchar(20) CHARACTER SET utf8 NOT NULL COMMENT '上次登录ip',
  `is_root` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否为超级管理员1是2否',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '用户状态1正常2冻结',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='后台用户表';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, '17628293814', '测试角色', '17628293814', '$2a$10$oBccugDIZzPfGVWGhuVMf.sopoV5m5/Wqr15YbNfHMVx5yQqvidGK', 'dsadasd.sd', '2019-07-22 16:53:03', '192.168.230.1', 0, 1, '2019-07-15 15:53:22', '2019-07-22 16:53:03');
INSERT INTO `user` VALUES (2, '176282938141', '测试角色123', '17628293815', '$2a$10$8Sol2bURmVetCM3cvfubPu4nH0gBNb4wfa9GDPI0.gvpVkNk6FIfi', 'dsadasd.sd11', '2019-08-05 17:43:43', '192.168.230.1', 1, 1, '2019-07-16 15:38:10', '2019-08-05 17:43:43');
COMMIT;

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '管理员id',
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色id',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `member_id` (`user_id`,`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user_role
-- ----------------------------
BEGIN;
INSERT INTO `user_role` VALUES (1, 1, 1, '2019-07-15 07:53:22', '2019-07-15 07:53:22');
INSERT INTO `user_role` VALUES (2, 1, 2, '2019-07-15 07:53:22', '2019-07-15 07:53:22');
INSERT INTO `user_role` VALUES (3, 1, 3, '2019-07-15 07:53:22', '2019-07-15 07:53:22');
INSERT INTO `user_role` VALUES (7, 2, 1, '2019-07-29 08:15:36', '2019-07-29 08:15:36');
INSERT INTO `user_role` VALUES (8, 2, 2, '2019-07-29 08:15:36', '2019-07-29 08:15:36');
INSERT INTO `user_role` VALUES (9, 2, 3, '2019-07-29 08:15:36', '2019-07-29 08:15:36');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
