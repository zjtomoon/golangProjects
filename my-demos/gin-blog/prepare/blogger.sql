create database if not exists ginblog;
use ginblog;

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `category_id` int(10) unsigned NOT NULL,
  `summary` varchar(255) NOT NULL DEFAULT '' COMMENT '摘要',
  `content` text NOT NULL,
  `title` varchar(100) NOT NULL DEFAULT '',
  `view_count` int(11) DEFAULT '0',
  `comment_count` int(11) unsigned DEFAULT NULL COMMENT '0',
  `username` varchar(80) DEFAULT '',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES ('1', '1', 'sum1', '测试1测试1测试1', '测试1', '0', '0', 'haima', '1', '2019-11-10 09:48:40', '2019-11-10 09:48:40');
INSERT INTO `article` VALUES ('2', '1', 'sum2', '测试2测试2测试2测试2', '测试2', '0', '0', 'haima', '1', '2019-11-10 10:38:53', '2019-11-10 10:38:53');

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` bigint(20) unsigned NOT NULL,
  `category_name` varchar(255) NOT NULL DEFAULT '',
  `category_no` int(10) unsigned DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES ('1', '分类一', '10001', '2019-11-10 09:12:16', '2019-11-10 09:12:16');
INSERT INTO `category` VALUES ('2', '分类二', '10002', '2019-11-10 09:24:52', '2019-11-10 09:24:52');

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL,
  `content` varchar(255) NOT NULL,
  `username` varchar(60) NOT NULL DEFAULT '',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1',
  `article_id` bigint(20) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of comment
-- ----------------------------

-- ----------------------------
-- Table structure for leave
-- ----------------------------
DROP TABLE IF EXISTS `leave`;
CREATE TABLE `leave` (
  `id` bigint(20) NOT NULL,
  `username` varchar(60) NOT NULL,
  `email` varchar(160) NOT NULL DEFAULT '',
  `content` varchar(255) DEFAULT '',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of leave
-- ----------------------------
