SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- ----------------------------
-- Table structure for feed_count
-- ----------------------------
DROP TABLE IF EXISTS `feed_count`;
CREATE TABLE `feed_count` (
    `feed_id` bigint NULL DEFAULT NULL,
    `comment_count` int NOT NULL DEFAULT 0,
    `view_count` int NOT NULL DEFAULT 0,
    `like_count` int NOT NULL DEFAULT 0,
    `repost_count` int NOT NULL DEFAULT 0,
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `del_state` tinyint NOT NULL DEFAULT 0,
    `version` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本号',
    PRIMARY KEY (`feed_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;
-- ----------------------------
-- Table structure for comment_count
-- ----------------------------
DROP TABLE IF EXISTS `comment_count`;
CREATE TABLE `comment_count` (
    `comment_id` bigint NOT NULL,
    `comment_count` int NOT NULL DEFAULT 0,
    `view_count` int NOT NULL DEFAULT 0,
    `like_count` int NOT NULL DEFAULT 0,
    `repost_count` int NOT NULL DEFAULT 0,
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `del_state` tinyint NOT NULL DEFAULT 0,
    `version` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本号',
    PRIMARY KEY (`comment_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;
DROP TABLE IF EXISTS `user_state`;
CREATE TABLE `user_state` (
    `user_id` bigint NOT NULL,
    `following_count` int NOT NULL DEFAULT NULL DEFAULT 0 COMMENT '关注数',
    `follower_count` int NOT NULL DEFAULT NULL DEFAULT 0 COMMENT '粉丝数',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `del_state` tinyint NOT NULL DEFAULT 0,
    `version` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本号',
    PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;
SET FOREIGN_KEY_CHECKS = 1;