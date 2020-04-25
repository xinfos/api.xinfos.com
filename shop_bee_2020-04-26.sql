# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.21)
# Database: shop_bee
# Generation Time: 2020-04-25 23:32:20 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table t_attachment
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_attachment`;

CREATE TABLE `t_attachment` (
  `attach_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '附件ID',
  `g_id` bigint(20) unsigned NOT NULL COMMENT '附件组ID',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '附件名称',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '附件地址',
  `attach_type` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '附件类型 [1:图片 | 2:视频]',
  `attach_size` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '附件大小',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`attach_id`),
  KEY `idx_g_id_is_delete` (`g_id`,`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table t_brand
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_brand`;

CREATE TABLE `t_brand` (
  `brand_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '品牌ID',
  `brand_name` varchar(60) NOT NULL DEFAULT '' COMMENT '品牌名称',
  `brand_logo` varchar(255) NOT NULL DEFAULT '' COMMENT '品牌图片',
  `brand_desc` varchar(255) NOT NULL DEFAULT '' COMMENT '品牌描述',
  `cat_id` bigint(20) unsigned NOT NULL COMMENT '品牌分类ID',
  `sort_order` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '品牌在前台页面的显示顺序,数字越大越靠前',
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '品牌是否显示 [1:是 | 2:否 ]',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '品牌创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '品牌修改时间',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否被删除: 1:删除时间戳 2: 正常',
  PRIMARY KEY (`brand_id`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_show_order` (`is_show`,`sort_order`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='产品品牌表';

LOCK TABLES `t_brand` WRITE;
/*!40000 ALTER TABLE `t_brand` DISABLE KEYS */;

INSERT INTO `t_brand` (`brand_id`, `brand_name`, `brand_logo`, `brand_desc`, `cat_id`, `sort_order`, `is_show`, `created_at`, `updated_at`, `is_delete`)
VALUES
	(10000,'阿迪达斯','http://www.baidu.com','全球最大的运动品牌',10010,1,1,'2020-02-16 14:20:03','2020-04-25 15:59:30',1),
	(10001,'中国耐克','http://www.nike.com','美国最大的运动品牌',10010,9999,1,'2020-02-16 14:20:45','2020-04-25 16:13:48',2),
	(10008,'耐1克','http://www.baidu.com','全球最大的运动品牌',10010,1,1,'2020-04-25 15:47:09','2020-04-25 15:47:09',2),
	(10009,'Nike','http://www.nike.com','全球最大的运动品牌',10010,999,1,'2020-04-25 15:53:18','2020-04-25 16:05:50',1);

/*!40000 ALTER TABLE `t_brand` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table t_category
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_category`;

CREATE TABLE `t_category` (
  `cat_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类ID',
  `parent_cat_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '父级分类ID',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '分类名称',
  `alias` varchar(30) NOT NULL COMMENT '分类别名',
  `desc` varchar(255) NOT NULL COMMENT '分类描述',
  `show_in_nav` tinyint(2) unsigned NOT NULL DEFAULT '2' COMMENT '导航栏是否展示  [1:是 | 2: 否]',
  `is_show` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '前端是否显示 [1:是 | 2: 否]',
  `is_parent` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '是否为父级分类 [1:是 | 2: 否]',
  `state` smallint(6) unsigned NOT NULL DEFAULT '1' COMMENT '分类状态[1:正常]',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_delete` tinyint(2) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`cat_id`),
  UNIQUE KEY `uk_name_is_delete` (`name`,`is_delete`),
  KEY `idx_cat_id_pid_is_delete` (`cat_id`,`parent_cat_id`,`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `t_category` WRITE;
/*!40000 ALTER TABLE `t_category` DISABLE KEYS */;

INSERT INTO `t_category` (`cat_id`, `parent_cat_id`, `name`, `alias`, `desc`, `show_in_nav`, `is_show`, `is_parent`, `state`, `created_at`, `updated_at`, `is_delete`)
VALUES
	(10001,10000,'游戏话费','游戏话费','游戏话费',1,0,1,1,'2020-02-14 23:19:14','2020-02-14 23:35:06',2),
	(10002,10000,'服装鞋包','服装鞋包','服装鞋包',1,0,1,1,'2020-02-14 23:20:22','2020-02-14 23:35:06',2),
	(10003,10000,'手机数码','手机数码','手机数码',1,0,1,1,'2020-02-14 23:20:38','2020-02-14 23:35:06',2),
	(10004,10000,'家用电器','家用电器','家用电器',1,0,1,1,'2020-02-14 23:20:55','2020-02-14 23:35:06',2),
	(10005,10000,'美妆饰品','美妆饰品','美妆饰品',1,0,1,1,'2020-02-14 23:21:15','2020-02-14 23:35:06',2),
	(10006,10000,'母婴用品','母婴用品','母婴用品',1,0,1,1,'2020-02-14 23:21:30','2020-02-14 23:35:06',2),
	(10007,10000,'家具建材','家具建材','家具建材',1,0,1,1,'2020-02-14 23:21:48','2020-02-15 13:18:11',2),
	(10008,10000,'百货食品','百货食品','百货食品',1,0,1,1,'2020-02-14 23:22:05','2020-02-14 23:35:06',2),
	(10009,10000,'运动户外','运动户外','运动户外',1,0,1,1,'2020-02-14 23:22:18','2020-02-14 23:35:06',2),
	(10010,10002,'流行男鞋','流行男鞋','流行男鞋',1,0,1,1,'2020-02-14 23:24:47','2020-02-14 23:35:06',2),
	(10011,10002,'男鞋','男鞋','男鞋',1,1,1,1,'2020-02-14 23:28:03','2020-02-15 14:11:53',2),
	(10012,10011,'T恤','T恤','T恤',1,0,0,1,'2020-02-14 23:28:46','2020-02-14 23:35:06',2),
	(10013,10011,'风衣','风衣','风衣',1,0,0,1,'2020-02-14 23:33:08','2020-02-14 23:33:08',2),
	(10014,10011,'夹克','夹克','夹克',1,0,0,1,'2020-02-14 23:34:31','2020-02-14 23:34:31',2),
	(10016,10011,'衬衫','衬衫','衬衫',1,1,1,1,'2020-02-14 23:51:16','2020-02-19 14:40:16',2),
	(10017,10000,'asdfasdf','asdfasdf','asdfasdfasdfasdf',0,0,0,1,'2020-02-19 11:44:32','2020-02-19 14:40:00',2),
	(10018,10000,'asdfasfsafasdf','asdfasdf','sdfasfasdf',0,0,0,1,'2020-02-19 11:44:45','2020-02-19 14:40:00',2),
	(10019,10000,'aaaaaaaaaa','撒打算放','asdfasfasdfadsf',0,0,0,1,'2020-02-19 11:45:21','2020-02-19 14:40:00',2),
	(10020,10000,'阿萨德飞洒','大事发生','撒打发斯蒂芬',0,0,0,1,'2020-02-19 14:53:30','2020-02-19 14:53:30',2),
	(10021,10018,'我的撒发生的','阿斯顿发发阿法士大夫','按时发送',0,0,0,1,'2020-02-19 14:55:27','2020-02-19 14:55:27',2),
	(10022,0,'','','',2,1,1,1,'2020-04-21 11:23:07','2020-04-21 11:23:07',2);

/*!40000 ALTER TABLE `t_category` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table t_item
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_item`;

CREATE TABLE `t_item` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `item_id` bigint(20) unsigned NOT NULL COMMENT 'SKU',
  `product_id` int(11) DEFAULT NULL,
  `name` int(11) DEFAULT NULL,
  `img` int(11) DEFAULT NULL,
  `sales_price` int(11) DEFAULT NULL,
  `market_price` int(11) DEFAULT NULL,
  `stock` int(11) DEFAULT NULL,
  `warning_stock` int(11) DEFAULT NULL,
  `code` int(11) DEFAULT NULL,
  `barcode` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `is_delete` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table t_item_spec
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_item_spec`;

CREATE TABLE `t_item_spec` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table t_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_menu`;

CREATE TABLE `t_menu` (
  `menu_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `menu_pid` int(11) unsigned NOT NULL COMMENT '菜单父级ID',
  `title` varchar(20) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `addr` varchar(100) NOT NULL DEFAULT '' COMMENT '菜单地址',
  `is_menu` tinyint(11) unsigned NOT NULL DEFAULT '1' COMMENT '是否为菜单',
  `group_title` varchar(20) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_delete` int(11) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table t_product_img
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_product_img`;

CREATE TABLE `t_product_img` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `product_id` bigint(20) unsigned NOT NULL COMMENT '商品ID',
  `attach_url` varchar(255) NOT NULL DEFAULT '' COMMENT '图片URL',
  `displayorder` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '显示顺序',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_delete` tinyint(1) unsigned NOT NULL COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table t_proudct
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_proudct`;

CREATE TABLE `t_proudct` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `product_id` bigint(20) unsigned NOT NULL COMMENT '商品ID',
  `cat_id` bigint(20) unsigned NOT NULL COMMENT '分类ID',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '商品名称',
  `store_id` bigint(11) unsigned NOT NULL COMMENT '商家店铺ID',
  `type_id` int(11) unsigned NOT NULL COMMENT '商品类型',
  `brand_id` bigint(20) unsigned NOT NULL COMMENT '品牌ID',
  `sketch` varchar(120) NOT NULL DEFAULT '' COMMENT '商品简述',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '商品描述',
  `keywords` varchar(50) NOT NULL DEFAULT '' COMMENT '关键字',
  `state` smallint(6) unsigned NOT NULL DEFAULT '0' COMMENT '商品状态',
  `is_package` tinyint(3) unsigned NOT NULL,
  `is_integral` tinyint(3) unsigned NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table t_seller
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_seller`;

CREATE TABLE `t_seller` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '用户手机号',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '用户姓名',
  `email` varchar(30) NOT NULL DEFAULT '' COMMENT '邮箱',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `position` varchar(20) NOT NULL DEFAULT '' COMMENT '职位',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '登录密码',
  `state` tinyint(3) unsigned NOT NULL DEFAULT '2' COMMENT '状态',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_mobile_is_delete` (`mobile`,`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `t_seller` WRITE;
/*!40000 ALTER TABLE `t_seller` DISABLE KEYS */;

INSERT INTO `t_seller` (`id`, `mobile`, `name`, `email`, `avatar`, `position`, `password`, `state`, `created_at`, `updated_at`, `is_delete`)
VALUES
	(1,'13520928996','杨骏','psj474@163.com','','','$2y$10$YoE.6W11gp6R.YjUFXtpr.7/hwJe9yY1eXzBje7RFywMNzAl1J9eO',2,'2020-02-07 13:35:03','2020-02-07 13:35:03',2),
	(2,'17610258996','17610258996','','','','$2y$10$o8Pk20J1.VR2F2wx0jS1G.uHFx1W/0oMXwpLlh00Kf8DKAf7x9uI.',2,'2020-02-19 19:36:18','2020-02-19 19:36:18',2);

/*!40000 ALTER TABLE `t_seller` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table t_shop
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_shop`;

CREATE TABLE `t_shop` (
  `s_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '店铺ID',
  `seller_id` bigint(20) unsigned NOT NULL COMMENT '卖家ID',
  `cat_id` bigint(20) unsigned NOT NULL COMMENT '店铺主营分类',
  `version` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '店铺版本 [1:单店铺 | 2:多店铺]',
  `type` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '店铺类型 [1:普通商城 | 2: 新零售]',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '店铺名称',
  `desc` varchar(100) NOT NULL DEFAULT '' COMMENT '店铺描述',
  `logo` varchar(255) NOT NULL DEFAULT '' COMMENT '店铺LOGO',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '店铺地址',
  `cert_type` tinyint(4) unsigned NOT NULL COMMENT '店铺认证类型',
  `state` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '店铺状态[1:营业中 | 2: 打样]',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`s_id`),
  UNIQUE KEY `uk_name_is_delete` (`name`,`is_delete`),
  KEY `idx_seller_id_is_delete` (`seller_id`,`is_delete`),
  KEY `idx_sid_is_delete` (`s_id`,`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='店铺表';

LOCK TABLES `t_shop` WRITE;
/*!40000 ALTER TABLE `t_shop` DISABLE KEYS */;

INSERT INTO `t_shop` (`s_id`, `seller_id`, `cat_id`, `version`, `type`, `name`, `desc`, `logo`, `url`, `cert_type`, `state`, `created_at`, `updated_at`, `is_delete`)
VALUES
	(1,1,1,1,1,'舌尖上的中国','舌尖上的中国','','',0,1,'2020-02-09 22:46:46','2020-02-10 23:56:25',2);

/*!40000 ALTER TABLE `t_shop` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table t_shop_address
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_shop_address`;

CREATE TABLE `t_shop_address` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table t_shop_attach_group
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_shop_attach_group`;

CREATE TABLE `t_shop_attach_group` (
  `g_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '附件组ID',
  `s_id` bigint(20) unsigned NOT NULL COMMENT '店铺ID',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '附件组名称',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`g_id`),
  UNIQUE KEY `uk_s_id_name_is_delete` (`s_id`,`name`,`is_delete`),
  KEY `idx_g_id_is_delete` (`g_id`,`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table t_shop_permission
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_shop_permission`;

CREATE TABLE `t_shop_permission` (
  `permission_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '权限ID',
  `permission_name` varchar(20) NOT NULL DEFAULT '' COMMENT '权限名称',
  `permission_desc` varchar(30) NOT NULL DEFAULT '' COMMENT '权限描述',
  `permission_type` tinyint(3) unsigned NOT NULL COMMENT '权限类型 [1:menu | 2: button]',
  `route` varchar(30) NOT NULL DEFAULT '' COMMENT '权限路由',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`permission_id`),
  UNIQUE KEY `uk_name_route_type_is_delete` (`permission_name`,`route`,`permission_type`,`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table t_shop_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_shop_role`;

CREATE TABLE `t_shop_role` (
  `role_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(20) NOT NULL DEFAULT '' COMMENT '角色名称',
  `role_type` smallint(6) unsigned NOT NULL COMMENT '角色类型',
  `role_desc` varchar(30) NOT NULL DEFAULT '' COMMENT '角色描述',
  `is_admin` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否为超管 [1:是 | 2: 否]',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`role_id`),
  UNIQUE KEY `uk_name_type_is_delete` (`role_name`,`role_type`,`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table t_shop_role_permission
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_shop_role_permission`;

CREATE TABLE `t_shop_role_permission` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `role_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
  `permission_id` bigint(20) unsigned NOT NULL COMMENT '权限ID',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `is_delete` tinyint(1) unsigned NOT NULL COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table t_shop_seller
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_shop_seller`;

CREATE TABLE `t_shop_seller` (
  `s_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `seller_id` bigint(20) unsigned NOT NULL,
  `role_id` bigint(20) unsigned NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`s_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table t_shop_staff
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_shop_staff`;

CREATE TABLE `t_shop_staff` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `staff_no` varchar(20) NOT NULL DEFAULT '' COMMENT '员工编号',
  `s_id` bigint(20) unsigned NOT NULL COMMENT '员工所属店铺ID',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '员工姓名',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '员工联系方式',
  `state` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '员工状态[1:启用 | 2: 停用]',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_statf_no_sid_is_delete` (`staff_no`,`s_id`,`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='店铺员工表';



# Dump of table t_sku_images
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_sku_images`;

CREATE TABLE `t_sku_images` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table t_system_sku_spec
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_system_sku_spec`;

CREATE TABLE `t_system_sku_spec` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `s_spec_id` bigint(20) unsigned NOT NULL COMMENT '系统规格ID',
  `name` varchar(15) NOT NULL DEFAULT '' COMMENT '规格名称',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `t_system_sku_spec` WRITE;
/*!40000 ALTER TABLE `t_system_sku_spec` DISABLE KEYS */;

INSERT INTO `t_system_sku_spec` (`id`, `s_spec_id`, `name`, `created_at`, `updated_at`, `is_delete`)
VALUES
	(1,1000301,'机身颜色','2020-04-22 17:21:31','2020-04-22 17:24:06',2),
	(2,1000302,'存储容量','2020-04-22 17:21:36','2020-04-22 17:24:10',2),
	(3,1000303,'版本类型','2020-04-22 17:22:35','2020-04-22 17:24:14',2);

/*!40000 ALTER TABLE `t_system_sku_spec` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table t_system_sku_spec_value
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_system_sku_spec_value`;

CREATE TABLE `t_system_sku_spec_value` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `s_spec_id` bigint(20) unsigned NOT NULL COMMENT '系统规格ID',
  `s_spec_val_id` bigint(20) NOT NULL COMMENT '规格值ID',
  `value` varchar(30) NOT NULL DEFAULT '' COMMENT '规格值',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_delete` int(11) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `t_system_sku_spec_value` WRITE;
/*!40000 ALTER TABLE `t_system_sku_spec_value` DISABLE KEYS */;

INSERT INTO `t_system_sku_spec_value` (`id`, `s_spec_id`, `s_spec_val_id`, `value`, `created_at`, `updated_at`, `is_delete`)
VALUES
	(1,1000301,10003011,'黑色','2020-04-22 17:24:24','2020-04-22 17:26:37',2),
	(2,1000301,10003011,'土灰色','2020-04-22 17:25:51','2020-04-22 17:26:51',2);

/*!40000 ALTER TABLE `t_system_sku_spec_value` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table t_system_spu_attr
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_system_spu_attr`;

CREATE TABLE `t_system_spu_attr` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `s_attr_id` bigint(20) unsigned NOT NULL COMMENT '系统商品属性ID',
  `s_group_id` bigint(20) unsigned NOT NULL COMMENT '系统商品属性组ID',
  `cat_id` bigint(20) unsigned NOT NULL COMMENT '系统商品分类ID',
  `name` varchar(15) NOT NULL DEFAULT '' COMMENT '系统商品属性名称',
  `fill_type` tinyint(3) unsigned NOT NULL DEFAULT '2' COMMENT '参数填充类型 [1: 选项框 | 2: 输入框]',
  `is_numeric` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否是数字类型参数 [1:是 | 2: 否]',
  `unit` varchar(10) NOT NULL DEFAULT '' COMMENT '数字类型参数的单位，非数字类型可以为空',
  `is_generic` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否是SPU通用属性 [1:是 | 2: 否]',
  `is_searching` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否用于搜索过滤 [1:是 | 2: 否]',
  `is_required` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否为必填属性 [1:是 | 2: 否]',
  `segments` varchar(200) NOT NULL DEFAULT '' COMMENT '数值类型参数，如果需要搜索，则添加分段间隔值，如CPU频率间隔：0.5-1.0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `t_system_spu_attr` WRITE;
/*!40000 ALTER TABLE `t_system_spu_attr` DISABLE KEYS */;

INSERT INTO `t_system_spu_attr` (`id`, `s_attr_id`, `s_group_id`, `cat_id`, `name`, `fill_type`, `is_numeric`, `unit`, `is_generic`, `is_searching`, `is_required`, `segments`, `created_at`, `updated_at`, `is_delete`)
VALUES
	(1,1000311,100031,10003,'入网型号',2,2,' ',2,2,2,'','2020-04-21 11:27:22','2020-04-22 17:06:31',2),
	(2,1000312,100031,10003,'产品名称',2,2,' ',2,2,2,'','2020-04-21 11:27:32','2020-04-22 17:06:31',2),
	(3,1000313,100031,10003,'上市年份',2,2,' ',2,2,2,'','2020-04-21 11:27:41','2020-04-22 17:06:31',2),
	(4,1000314,100031,10003,'上市月份',2,2,' ',2,2,2,'','2020-04-21 11:28:21','2020-04-22 17:06:31',2),
	(5,1000315,100032,10003,'机身长度（mm）',2,2,' ',2,2,2,'','2020-04-21 11:28:39','2020-04-22 17:06:31',2),
	(6,1000316,100032,10003,'机身重量（g）',2,2,' ',2,2,2,'','2020-04-21 11:28:45','2020-04-22 17:06:31',2),
	(7,1000317,100032,10003,'机身材质工艺',2,2,' ',2,2,2,'','2020-04-21 11:28:47','2020-04-22 17:06:31',2),
	(8,1000318,100032,10003,'机身宽度（mm）',2,2,' ',2,2,2,'','2020-04-21 11:28:48','2020-04-22 17:06:31',2),
	(9,1000319,100032,10003,'机身材质分类',2,2,' ',2,2,2,'','2020-04-21 11:28:49','2020-04-22 17:06:31',2),
	(10,1000320,100032,10003,'机身厚度（mm）',2,2,' ',2,2,2,'','2020-04-21 11:28:49','2020-04-22 17:06:31',2),
	(11,1000321,100032,10003,'运营商标志或内容',2,2,' ',2,2,2,'','2020-04-21 11:28:50','2020-04-22 17:06:31',2),
	(12,1000322,100032,10003,'CPU品牌',2,2,' ',2,2,2,'','2020-04-21 11:28:51','2020-04-22 17:06:31',2);

/*!40000 ALTER TABLE `t_system_spu_attr` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table t_system_spu_attr_group
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_system_spu_attr_group`;

CREATE TABLE `t_system_spu_attr_group` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `s_group_id` bigint(20) unsigned NOT NULL COMMENT '系统商品属性组ID',
  `cat_id` bigint(20) unsigned NOT NULL COMMENT '系统商品分类ID',
  `name` varchar(15) NOT NULL DEFAULT '' COMMENT '系统商品属性组名称',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `t_system_spu_attr_group` WRITE;
/*!40000 ALTER TABLE `t_system_spu_attr_group` DISABLE KEYS */;

INSERT INTO `t_system_spu_attr_group` (`id`, `s_group_id`, `cat_id`, `name`, `created_at`, `updated_at`, `is_delete`)
VALUES
	(1,100031,10003,'主体','2020-04-21 11:23:55','2020-04-21 14:37:42',2),
	(2,100032,10003,'基本信息','2020-04-21 11:24:33','2020-04-21 14:37:39',2),
	(3,100033,10003,'主芯片','2020-04-21 11:24:41','2020-04-21 14:37:36',2);

/*!40000 ALTER TABLE `t_system_spu_attr_group` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table t_system_spu_attr_group_map
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_system_spu_attr_group_map`;

CREATE TABLE `t_system_spu_attr_group_map` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `cat_id` bigint(20) unsigned NOT NULL COMMENT '分类ID',
  `s_group_id` bigint(20) unsigned NOT NULL COMMENT '系统属性组ID',
  `s_attr_id` bigint(20) unsigned NOT NULL COMMENT '系统属性ID',
  `displayorder` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '显示顺序',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_cat_group_attr_isdelete` (`cat_id`,`s_group_id`,`s_attr_id`,`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `t_system_spu_attr_group_map` WRITE;
/*!40000 ALTER TABLE `t_system_spu_attr_group_map` DISABLE KEYS */;

INSERT INTO `t_system_spu_attr_group_map` (`id`, `cat_id`, `s_group_id`, `s_attr_id`, `displayorder`, `created_at`, `updated_at`, `is_delete`)
VALUES
	(1,10003,100031,0,1,'2020-04-21 11:16:40','2020-04-21 11:25:17',2),
	(2,10003,100032,0,1,'2020-04-21 11:25:27','2020-04-21 11:25:48',2),
	(3,10003,100033,0,1,'2020-04-21 11:25:42','2020-04-21 11:25:51',2);

/*!40000 ALTER TABLE `t_system_spu_attr_group_map` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table t_system_spu_attr_value
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_system_spu_attr_value`;

CREATE TABLE `t_system_spu_attr_value` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `s_attr_val_id` bigint(20) unsigned NOT NULL COMMENT '系统SPU属性值ID',
  `s_attr_id` bigint(20) unsigned NOT NULL COMMENT '系统SPU属性ID',
  `cat_id` bigint(20) unsigned DEFAULT NULL COMMENT '系统商品分类ID',
  `value` varchar(30) NOT NULL DEFAULT '' COMMENT '属性值',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除 [1:是 | 2: 否]',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `t_system_spu_attr_value` WRITE;
/*!40000 ALTER TABLE `t_system_spu_attr_value` DISABLE KEYS */;

INSERT INTO `t_system_spu_attr_value` (`id`, `s_attr_val_id`, `s_attr_id`, `cat_id`, `value`, `created_at`, `updated_at`, `is_delete`)
VALUES
	(1,13001,1000311,10003,'IN2020','2020-04-21 14:10:29','2020-04-21 16:25:21',2),
	(2,13002,1000311,10003,'IN2021','2020-04-21 14:10:44','2020-04-21 16:25:23',2),
	(3,13003,1000311,10003,'IN2022','2020-04-21 14:10:50','2020-04-21 16:25:27',2),
	(4,13004,1000311,10003,'IN2023','2020-04-21 14:10:53','2020-04-21 16:25:30',2),
	(5,13005,1000311,10003,'IN2024','2020-04-21 14:11:11','2020-04-21 16:25:33',2),
	(6,13100,1000313,10003,'2020年 春','2020-04-21 14:11:54','2020-04-21 16:25:36',2),
	(7,13101,1000313,10003,'2019年 冬','2020-04-21 14:13:59','2020-04-21 16:25:40',2),
	(8,13102,1000313,10003,'2019年 秋','2020-04-21 14:14:00','2020-04-21 16:25:44',2),
	(9,13103,1000313,10003,'2019年 夏','2020-04-21 14:14:01','2020-04-21 16:25:48',2),
	(10,13200,1000322,10003,'高通(Qualcomm)','2020-04-21 14:16:39','2020-04-21 16:25:51',2),
	(11,13201,1000322,10003,'联发科','2020-04-21 14:16:40','2020-04-21 16:25:54',2);

/*!40000 ALTER TABLE `t_system_spu_attr_value` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table t_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_user`;

CREATE TABLE `t_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `nick_name` varchar(255) NOT NULL DEFAULT '',
  `phone` varchar(20) NOT NULL DEFAULT '',
  `id_card` varchar(18) NOT NULL DEFAULT '',
  `birthday` date NOT NULL,
  `gender` tinyint(3) NOT NULL DEFAULT '0',
  `age` tinyint(11) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_delete` tinyint(1) NOT NULL DEFAULT '2',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_idcard` (`id_card`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table t_user_adderss
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_user_adderss`;

CREATE TABLE `t_user_adderss` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `consignee_name` varchar(50) NOT NULL DEFAULT '' COMMENT '收货人姓名',
  `consignee_phone` varchar(30) NOT NULL DEFAULT '' COMMENT '收货人手机号',
  `country` int(11) unsigned NOT NULL COMMENT '国家编码',
  `province` int(11) unsigned NOT NULL COMMENT '省份编码',
  `city` int(11) unsigned NOT NULL COMMENT '城市编码',
  `area` int(11) unsigned NOT NULL COMMENT '区域编码',
  `street` int(11) unsigned NOT NULL COMMENT '街道编码',
  `address` varchar(120) NOT NULL DEFAULT '' COMMENT '详细地址',
  `post_code` int(11) unsigned NOT NULL COMMENT '邮政编码',
  `is_default` tinyint(1) NOT NULL DEFAULT '2' COMMENT '是否设置为默认收货地址 [1:是 | 2:否]',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更改时间',
  `is_delete` tinyint(1) NOT NULL DEFAULT '2' COMMENT '是否删除 [1:是 | 2:否]',
  PRIMARY KEY (`id`),
  KEY `idx_user_id_is_deleted` (`user_id`,`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table t_user_profile
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_user_profile`;

CREATE TABLE `t_user_profile` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `avatar` int(11) DEFAULT NULL,
  `realname` int(11) DEFAULT NULL,
  `gender` int(11) DEFAULT NULL,
  `birthday` int(11) DEFAULT NULL,
  `user_level` int(11) DEFAULT NULL,
  `experience` int(11) DEFAULT NULL,
  `personal_sign` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
