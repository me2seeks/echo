SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS `user_relation`;
-- ----------------------------
-- user_relation
-- ----------------------------
CREATE TABLE `user_relation` (
    `id` bigint NOT NULL,
    `follower_id` bigint NOT NULL COMMENT '关注者的用户ID',
    `followee_id` bigint NOT NULL COMMENT '被关注者的用户ID',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `delete_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `del_state` tinyint NOT NULL DEFAULT '0',
    `version` bigint UNSIGNED NOT NULL DEFAULT '0' COMMENT '版本号',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_follower_followee` (`follower_id`, `followee_id`),
    KEY `idx_follower_id` (`follower_id`),
    KEY `idx_followee_id` (`followee_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户关注关系表';
SET FOREIGN_KEY_CHECKS = 1;