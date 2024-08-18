SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` bigint UNSIGNED NOT NULL,
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `delete_at` datetime NULL DEFAULT NULL,
    `version` bigint UNSIGNED NOT NULL DEFAULT '0' COMMENT '版本号',
    `email` char(254) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `salt` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `sex` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别 0:男 1:女',
    `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `bio` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_mobile` (`email`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表';
-- ----------------------------
-- user_relation
-- ----------------------------
CREATE TABLE `user_relation` (
    `id` bigint UNSIGNED NOT NULL,
    `follower_id` bigint UNSIGNED NOT NULL COMMENT '关注者的用户ID',
    `followee_id` bigint UNSIGNED NOT NULL COMMENT '被关注者的用户ID',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `delete_at` datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_follower_followee` (`follower_id`, `followee_id`),
    KEY `idx_follower_id` (`follower_id`),
    KEY `idx_followee_id` (`followee_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户关注关系表';
SET FOREIGN_KEY_CHECKS = 1;