SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS `feeds`;
CREATE TABLE `feeds` (
    `id` bigint NOT NULL,
    `user_id` bigint NOT NULL,
    `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `media0` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `media1` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `media2` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `media3` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `del_state` tinyint NOT NULL DEFAULT 0,
    `version` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本号',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `user_id`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
    `id` BIGINT NOT NULL COMMENT '主键ID',
    `feed_id` BIGINT NOT NULL COMMENT '关联的帖子或文章ID',
    `user_id` BIGINT NOT NULL COMMENT '发表评论的用户ID',
    `parent_id` BIGINT DEFAULT NULL COMMENT '父评论ID,表示回复',
    `content` TEXT NOT NULL COMMENT '评论内容',
    `media0` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `media1` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `media2` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `media3` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `del_state` tinyint NOT NULL DEFAULT 0,
    `version` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本号',
    PRIMARY KEY (`id`),
    INDEX `idx_feed_id` (`feed_id`),
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_parent_id` (`parent_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '用户的评论和回复';
DROP TABLE IF EXISTS `user_last_request`;
CREATE TABLE `user_last_request` (
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `last_request_time` DATETIME NOT NULL COMMENT '用户最后一次请求时间',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `del_state` TINYINT NOT NULL DEFAULT 0 COMMENT '删除状态',
    `version` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本号',
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '用户最后一次请求时间';
SET FOREIGN_KEY_CHECKS = 1;