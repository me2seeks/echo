## 交互池（Interaction Service）
- `comment-service`：处理评论和回复
- `message-service`：处理私信和通知
- `follow-service`：处理用户之间的关注关系

```sql
CREATE TABLE `comments` (
    `id` BIGINT NOT NULL  COMMENT '主键ID',
    `post_id` BIGINT NOT NULL COMMENT '关联的帖子或文章ID',
    `user_id` BIGINT NOT NULL COMMENT '发表评论的用户ID',
    `parent_id` BIGINT DEFAULT NULL COMMENT '父评论ID，表示回复',
    `content` TEXT NOT NULL COMMENT '评论内容',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态，1表示正常，0表示删除',
    PRIMARY KEY (`id`),
    INDEX `idx_post_id` (`post_id`),
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户的评论和回复';

CREATE TABLE `messages` (
    `id` BIGINT NOT NULL  COMMENT '主键ID',
    `sender_id` BIGINT NOT NULL COMMENT '发送者用户ID',
    `receiver_id` BIGINT NOT NULL COMMENT '接收者用户ID',
    `content` TEXT NOT NULL COMMENT '私信内容',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发送时间',
    `read_at` DATETIME DEFAULT NULL COMMENT '阅读时间',
    `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态，1表示正常，0表示删除',
    PRIMARY KEY (`id`),
    INDEX `idx_sender_id` (`sender_id`),
    INDEX `idx_receiver_id` (`receiver_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户之间的私信';
```