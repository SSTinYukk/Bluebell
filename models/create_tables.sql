DROP TABLE IF EXIST `user`
CREATE TABLE `user`(
    `id` bigint(20) NOT NULL,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,             
    `email` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `gender` tinyint(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP UPDATE ON CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    UNIQUE KEY  `idx_username` (`username`) USING IN BTREE,
    UNIQUE KEY  `idx_user_id` (`user_id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;