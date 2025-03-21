DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password` varchar(64) NOT NULL,              
    `email` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `gender` tinyint(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP  ON UPDATE  CURRENT_TIMESTAMP,
    UNIQUE KEY  `idx_username` (`username`) USING BTREE,
    UNIQUE KEY  `idx_user_id` (`user_id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `community`;
CREATE TABLE `community` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `community_id` int(10) unsigned NOT NULL,
  `community_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `introduction` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_community_id` (`community_id`),
  UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `post_id` bigint(20) NOT NULL COMMENT '帖子id',
  `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `author_id` bigint(20) NOT NULL COMMENT '作者的用户id',
  `community_id` bigint(20) NOT NULL COMMENT '所属社区',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '帖子状态',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_post_id` (`post_id`),
  KEY `idx_author_id` (`author_id`),
  KEY `idx_community_id` (`community_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 创建 resources 表
DROP TABLE IF EXISTS `resources`;
CREATE TABLE `resources` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `filename` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件名',
    `path` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件存储路径',
    `author_id` bigint(20) NOT NULL COMMENT '上传者的用户 ID',
    `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '资源状态，0: 待审核，1: 已通过，2: 已拒绝',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    KEY `idx_author_id` (`author_id`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 在 post 表中添加 resource_id 字段
ALTER TABLE `post`
ADD COLUMN `resource_id` bigint(20) DEFAULT NULL COMMENT '关联的资源 ID';

-- 为 post 表的 resource_id 字段添加索引
ALTER TABLE `post`
ADD INDEX `idx_resource_id` (`resource_id`);

-- 在 community 表中添加 resource_id 字段
ALTER TABLE `community`
ADD COLUMN `resource_id` bigint(20) DEFAULT NULL COMMENT '关联的资源 ID';

-- 为 community 表的 resource_id 字段添加索引
ALTER TABLE `community`
ADD INDEX `idx_resource_id` (`resource_id`);

--测试数据
INSERT INTO `community` (`community_id`, `community_name`, `introduction`, `create_time`, `update_time`) 
VALUES (1, 'STM32', '学习STM32', '2025-01-01 08:10:10', '2025-11-01 08:10:10');

INSERT INTO `community` (`community_id`, `community_name`, `introduction`, `create_time`, `update_time`) 
VALUES (2, 'C51', '学习C51', '2025-01-01 08:00:00', '2025-01-01 08:00:00');

INSERT INTO `community` (`community_id`, `community_name`, `introduction`, `create_time`, `update_time`) 
VALUES (3, '8086', '学习8086', '2025-03-07 08:30:00', '2025-08-07 08:30:00');

INSERT INTO `community` (`community_id`, `community_name`, `introduction`, `create_time`, `update_time`) 
VALUES (4, '机器学习', '学习机器学习', '2025-01-01 08:00:00', '2025-01-01 08:00:00');