/*
 Navicat Premium Data Transfer

 Source Server         : centos7
 Source Server Type    : MySQL
 Source Server Version : 50731
 Source Host           : 192.168.0.104:3306
 Source Schema         : mxshop_user_srv

 Target Server Type    : MySQL
 Target Server Version : 50731
 File Encoding         : 65001

 Date: 03/01/2021 16:38:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `nick_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `head_url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `birthday` date NULL DEFAULT NULL,
  `address` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `gender` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `role` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_mobile`(`mobile`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '18782222220', '$pbkdf2-sha256$29000$O8cYw1hLCUHIOYcwhnCu1Q$ehwTTeBznh2t/ZI8md3NfCO6.C.brO9KsF7bb3nBUM0', 'bobby3', NULL, '2020-12-02', NULL, NULL, 'male', 2);
INSERT INTO `user` VALUES (2, '18782222221', '$pbkdf2-sha256$29000$UgrBmJNSCmEspVTK2dsbAw$YIMk8ngYxIQ/ZhmpKFl3mLQrDzoVECfEPntNuuVCAFE', 'bobby1', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (3, '18782222222', '$pbkdf2-sha256$29000$SUlJyZlTag2h9N7bmzMmZA$nWqeo6fTuQHNfFKB/3dNnvRJmYCA8h5iSIBefSaGF8g', 'bobby2', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (4, '18782222223', '$pbkdf2-sha256$29000$KGUMAcC4V6qVkrLWOicEYA$sV7zszUAdUQiGClq5VIihJ.5nhAkyTa6Bf/uxMYt8ow', 'bobby3', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (5, '18782222224', '$pbkdf2-sha256$29000$szYGYCzlXIvROidkTMmZcw$482iZfrktoPzCOujrSCHDQfW1vAjArbzqPIYB6r3Y5s', 'bobby4', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (6, '18782222225', '$pbkdf2-sha256$29000$cU5JaQ2B8F5rLcWYM.b8nw$yFAA74UdkyvZUedVQ7rPWF3RtaxX.KP8bhg9LRkm98Q', 'bobby5', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (7, '18782222226', '$pbkdf2-sha256$29000$ZizlvJeyVopRSilFSKmVsg$aAQAXK2X/sa8UKRZLkgUDM7nvmv0rQUGNmnNwnT9780', 'bobby6', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (8, '18782222227', '$pbkdf2-sha256$29000$zzknRGiNUQrB.D8H4HwPQQ$FMhuapUWWSVi6ncvQ9cBm8QNVxqh3LLfABjS7rkv.GY', 'bobby7', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (9, '18782222228', '$pbkdf2-sha256$29000$TokRAmAsxdi7d25NibF2rg$A1ubNTbbLSqazTvOJqt9Rwh1J.BFhVtJ2EONopZT970', 'bobby8', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (10, '18782222229', '$pbkdf2-sha256$29000$9f5/b.0959y7957TupfSeg$bshZP2GW.Of8hJN2iB09dOauO77jpTKGN3VgwrG4P1Y', 'bobby9', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (11, '18787878787', '$pbkdf2-sha256$29000$0zpn7P0/5/xfCyEEQChFSA$f6bRPkJbIhH4a44PubspvunEB/ZmdQzg8OEn8F81zqA', 'bobby', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (12, '18785685656', '$pbkdf2-sha256$29000$.f./t/aeU0qJUQrBeO8dww$1xesYRK2USctuwAKH8lkpb1bR6iLJHK.nCNk9vGNWcI', '18785685656', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (13, '18782902568', '$pbkdf2-sha256$29000$MsZYa21tLYUwhhAiBGCMMQ$CAS0EZjn61dm/vcC8x5FmdMww8.hWsGo0uEULP2.R9U', '18782902568', NULL, NULL, NULL, NULL, NULL, 1);

SET FOREIGN_KEY_CHECKS = 1;
