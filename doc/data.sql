# ************************************************************
# Sequel Pro SQL dump
# Version 5446
#
# https://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.22)
# Database: mons
# Generation Time: 2019-09-16 03:35:07 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table roles
# ------------------------------------------------------------

DROP TABLE IF EXISTS `roles`;

CREATE TABLE `roles` (
  `id` varchar(36) NOT NULL DEFAULT '' COMMENT 'ID',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '角色名称',
  `description` varchar(1000) NOT NULL DEFAULT '' COMMENT '角色描述',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;

INSERT INTO `roles` (`id`, `name`, `description`, `created_at`, `updated_at`)
VALUES
	('05e3b81e-a2b2-11e9-9b83-165f98df9650','editor','编辑','2019-07-10 09:28:29.814','2019-07-10 09:28:29.814'),
	('1782725e-a2b2-11e9-9b83-165f98df9650','admin','管理员','2019-07-10 09:28:59.381','2019-07-10 09:28:59.381');

/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user_roles
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user_roles`;

CREATE TABLE `user_roles` (
  `id` varchar(36) NOT NULL DEFAULT '' COMMENT 'ID',
  `user_id` varchar(36) NOT NULL DEFAULT '' COMMENT '用户ID',
  `role_id` varchar(36) NOT NULL DEFAULT '' COMMENT '角色ID',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `user_roles` WRITE;
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;

INSERT INTO `user_roles` (`id`, `user_id`, `role_id`, `created_at`)
VALUES
	('d36077aa-a2b2-11e9-9b83-165f98df9650','23f9db62-a145-11e9-ae8f-132379f358df','1782725e-a2b2-11e9-9b83-165f98df9650','2019-07-10 09:34:14.570'),
	('e00b86f2-a2b2-11e9-9b83-165f98df9650','23f9db62-a145-11e9-ae8f-132379f358df','05e3b81e-a2b2-11e9-9b83-165f98df9650','2019-07-10 09:34:35.824');

/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` varchar(36) NOT NULL DEFAULT '' COMMENT 'ID',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `nickname` varchar(100) NOT NULL DEFAULT '' COMMENT '昵称',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `intro` varchar(1024) NOT NULL DEFAULT '' COMMENT '简介',
  `is_approved` tinyint(1) NOT NULL COMMENT '是否审核',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `username`, `nickname`, `password`, `intro`, `is_approved`, `created_at`, `updated_at`)
VALUES
	('23f9db62-a145-11e9-ae8f-132379f358df','venjiang-1568604263775630','vj','','ddd miscro service',0,'2019-07-25 17:45:29.000','2019-08-30 16:57:34.007'),
	('58027af0-b378-11e9-af26-d43a6506c0c8','abc','a','','',0,'2019-07-31 17:48:26.552','2019-07-31 17:48:26.552'),
	('58030b3c-b378-11e9-af26-d43a6506c0c8','DomainService-user 2','ds','','',0,'2019-07-31 17:48:26.552','2019-07-31 17:48:26.552'),
	('acec32ae-aec3-11e9-9f85-784f43946143','abc','','','',0,'2019-07-25 18:05:05.453','2019-07-25 18:05:05.453'),
	('acec497e-aec3-11e9-9f85-784f43946143','DomainService-user 2','','','',0,'2019-07-25 18:05:05.453','2019-07-25 18:05:05.453'),
	('b2d0aab0-aec3-11e9-9f85-784f43946143','abc','','','',0,'2019-07-25 18:05:15.339','2019-07-25 18:05:15.339'),
	('b2d0cdb0-aec3-11e9-9f85-784f43946143','DomainService-user 2','','','',0,'2019-07-25 18:05:15.339','2019-07-25 18:05:15.339');

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
