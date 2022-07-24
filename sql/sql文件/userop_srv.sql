/*
 Navicat Premium Data Transfer

 Source Server         : centos7
 Source Server Type    : MySQL
 Source Server Version : 50731
 Source Host           : 192.168.0.104:3306
 Source Schema         : mxshop_userop_srv

 Target Server Type    : MySQL
 Target Server Version : 50731
 File Encoding         : 65001

 Date: 03/01/2021 16:38:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(0) NOT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  `update_time` datetime(0) NOT NULL,
  `user` int(11) NOT NULL,
  `province` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `city` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `district` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `signer_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `signer_mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of address
-- ----------------------------
INSERT INTO `address` VALUES (2, '2020-11-19 20:55:34', 0, '2020-11-19 20:55:34', 1, '陕西省', '宜昌市', '华中', 'Helen Lewis', 'George Allen', '18686868686');
INSERT INTO `address` VALUES (3, '2020-11-19 23:48:50', 0, '2020-11-19 23:48:50', 1, '天津市', '天津市', '和平区', '北京市', 'bobby', '18782222220');
INSERT INTO `address` VALUES (4, '2020-12-11 23:44:06', 0, '2020-12-11 23:44:06', 1, '北京市', '北京市', '西城区', '北京市', 'bobby', '18787878787');
INSERT INTO `address` VALUES (5, '2020-12-15 18:11:27', 0, '2020-12-15 18:11:27', 1, '河北省', '唐山市', '路南区', 'bobby', 'bobby', '18789898987');

-- ----------------------------
-- Table structure for leavingmessages
-- ----------------------------
DROP TABLE IF EXISTS `leavingmessages`;
CREATE TABLE `leavingmessages`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(0) NOT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  `update_time` datetime(0) NOT NULL,
  `user` int(11) NOT NULL,
  `message_type` int(11) NOT NULL,
  `subject` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `message` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `file` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of leavingmessages
-- ----------------------------
INSERT INTO `leavingmessages` VALUES (1, '2020-11-19 20:56:58', 0, '2020-11-19 20:56:58', 1, 1, 'occaecat aute voluptate dolor', 'ad', 'rlogin://nsmyiuyc.tt/dxrieny');
INSERT INTO `leavingmessages` VALUES (4, '2020-12-15 18:17:34', 0, '2020-12-15 18:17:34', 1, 5, '口罩 ', '继续一批kn95口罩', 'http://mxshop-files.oss-cn-hangzhou.aliyuncs.com/mxshop-images/csp2019真题.zip');

-- ----------------------------
-- Table structure for userfav
-- ----------------------------
DROP TABLE IF EXISTS `userfav`;
CREATE TABLE `userfav`  (
  `add_time` datetime(0) NOT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  `update_time` datetime(0) NOT NULL,
  `user` int(11) NOT NULL,
  `goods` int(11) NOT NULL,
  PRIMARY KEY (`user`, `goods`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of userfav
-- ----------------------------
INSERT INTO `userfav` VALUES ('2020-12-14 22:03:32', 0, '2020-12-14 22:03:32', 1, 423);
INSERT INTO `userfav` VALUES ('2020-12-14 23:12:46', 0, '2020-12-14 23:12:46', 1, 430);

SET FOREIGN_KEY_CHECKS = 1;
