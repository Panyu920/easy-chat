CREATE TABLE `users` (
  `id` varchar(24) NOT NULL DEFAULT '' COMMENT '用户ID',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像',
  `nickname` varchar(25) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(191) NOT NULL DEFAULT '' COMMENT '密码',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `last_login_at` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
  `phone` varchar(128) NOT NULL DEFAULT '' COMMENT '手机号',
  `phone_verified` tinyint(1) NOT NULL DEFAULT 0 COMMENT '手机号是否验证',
  `status` tinyint NOT NULL DEFAULT 0 COMMENT '用户状态(0: 正常,1:禁用, 2: 删除)',
  `sex` tinyint NOT NULL DEFAULT 0 COMMENT '性别(0: 男,1: 女)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_phone` (`phone`),
  KEY `idx_nickname` (`nickname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';