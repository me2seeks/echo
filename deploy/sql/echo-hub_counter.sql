SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- ----------------------------
-- Table structure for feed_count
-- ----------------------------
DROP TABLE IF EXISTS `feed_count`;
CREATE TABLE `feed_count` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `feed_id` bigint NULL DEFAULT NULL,
    `comment_count` int NULL DEFAULT NULL,
    `view_count` int NULL DEFAULT NULL,
    `like_count` int NULL DEFAULT NULL,
    `repost_count` int NULL DEFAULT NULL,
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `del_state` tinyint NOT NULL DEFAULT 0,
    `version` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本号',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;
-- ----------------------------
-- Table structure for comment_count
-- ----------------------------
DROP TABLE IF EXISTS `comment_count`;
CREATE TABLE `comment_count` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `comment_id` bigint NULL DEFAULT NULL,
    `comment_count` int NULL DEFAULT NULL,
    `view_count` int NULL DEFAULT NULL,
    `like_count` int NULL DEFAULT NULL,
    `repost_count` int NULL DEFAULT NULL,
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `del_state` tinyint NOT NULL DEFAULT 0,
    `version` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本号',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;
SET FOREIGN_KEY_CHECKS = 1;