## 内容池（Content Service）
- `post-service`：管理用户发布的帖子和文章
- `media-service`：管理用户上传的图片和视频

```sql
CREATE TABLE `posts` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id` BIGINT NOT NULL COMMENT '发布用户的ID',
    `content` TEXT NOT NULL COMMENT '帖子或文章的内容',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态，1表示正常，0表示删除',
    PRIMARY KEY (`id`),
    INDEX `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户发布的帖子和文章';
CREATE TABLE `media` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id` BIGINT NOT NULL COMMENT '上传用户的ID',
    `post_id` BIGINT DEFAULT NULL COMMENT '关联的帖子或文章ID',
    `type` ENUM('image', 'video') NOT NULL COMMENT '媒体类型，image表示图片，video表示视频',
    `url` VARCHAR(255) NOT NULL COMMENT '媒体文件的URL',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
    `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态，1表示正常，0表示删除',
    PRIMARY KEY (`id`),
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_post_id` (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户上传的图片和视频';
```

