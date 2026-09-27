-- SQL dump generated using DBML (dbml.dbdiagram.io)
-- Database: MySQL
-- Generated at: 2026-09-27T06:27:43.934Z

CREATE TABLE `friends` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '关系ID',
  `user_id` varchar(24) NOT NULL COMMENT '用户ID',
  `friend_id` varchar(24) NOT NULL COMMENT '好友ID',
  `remark` varchar(255) COMMENT '备注',
  `add_source` tinyint NOT NULL DEFAULT 0 COMMENT '添加好友来源(0: 好友请求,1: 好友添加)',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_friend_id` (`friend_id`),
  UNIQUE KEY `idx_user_id_friend_id` (`user_id`, `friend_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='好友关系' COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `friends_requests` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '好友请求ID',
  `user_id` varchar(24) NOT NULL COMMENT '用户ID',
  `friend_id` varchar(24) NOT NULL COMMENT '好友ID',
  `req_msg` varchar(255) COMMENT '请求消息',
  `req_time` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '请求时间',
  `req_status` tinyint NOT NULL DEFAULT 0 COMMENT '请求状态(0: 待处理,1: 已同意,2: 已拒绝)',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_friend_id` (`friend_id`),
  UNIQUE KEY `idx_user_id_friend_id` (`user_id`, `friend_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='好友请求' COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `groups` (
  `id` varchar(24) NOT NULL COMMENT '群组ID',
  `name` varchar(255) NOT NULL COMMENT '群组名称',
  `desc` varchar(255) COMMENT '群组描述',
  `avatar` varchar(255) COMMENT '群头像',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_user_id` varchar(24) NOT NULL COMMENT '创建用户ID',
  `type` tinyint NOT NULL DEFAULT 0 COMMENT '群组类型(0: 普通群组,1: 私有群组)',
  `is_verify` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否验证',
  `status` tinyint NOT NULL DEFAULT 0 COMMENT '群组状态(0: 正常,1:禁用, 2: 删除)',
  `notification` text COMMENT '通知',
  `notification_user_id` varchar(24) COMMENT '通知用户ID',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='群组表' COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `group_members` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '群组成员ID',
  `group_id` varchar(24) NOT NULL COMMENT '群组ID',
  `user_id` varchar(24) NOT NULL COMMENT '用户ID',
  `join_time` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '加入时间',
  `join_source` tinyint NOT NULL DEFAULT 0 COMMENT '加入群组来源(0: 群组邀请,1: 群组添加)',
  `role_level` tinyint NOT NULL DEFAULT 0 COMMENT '角色等级(0: 普通成员,1: 管理员,2: 群主)',
  `inviter_user_id` varchar(24) COMMENT '邀请用户ID',
  `handler_user_id` varchar(24) COMMENT '处理用户ID',
  PRIMARY KEY (`id`),
  KEY `idx_group_id` (`group_id`),
  KEY `idx_user_id` (`user_id`),
  UNIQUE KEY `idx_group_id_user_id` (`group_id`, `user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='群组成员' COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `group_requests` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '群组请求ID',
  `group_id` varchar(24) NOT NULL COMMENT '群组ID',
  `user_id` varchar(24) NOT NULL COMMENT '用户ID',
  `req_msg` varchar(255) COMMENT '请求消息',
  `req_time` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '请求时间',
  `req_status` tinyint NOT NULL DEFAULT 0 COMMENT '请求状态(0: 待处理,1: 已同意,2: 已拒绝)',
  `join_source` tinyint NOT NULL DEFAULT 0 COMMENT '加入群组来源(0: 群组邀请,1: 群组添加)',
  `inviter_user_id` varchar(24) COMMENT '邀请用户ID',
  `handler_user_id` varchar(24) COMMENT '处理用户ID',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_group_id` (`group_id`),
  KEY `idx_user_id` (`user_id`),
  UNIQUE KEY `idx_group_id_user_id` (`group_id`, `user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='群组请求' COLLATE=utf8mb4_unicode_ci;
