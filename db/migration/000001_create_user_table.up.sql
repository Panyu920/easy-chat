CREATE TABLE `users` (
  `id` varchar(24) PRIMARY KEY NOT NULL COMMENT '用户ID',
  `avatar` varchar(255) COMMENT '用户头像',
  `nickname` varchar(25) NOT NULL COMMENT '用户名',
  `password` varchar(191) NOT NULL COMMENT '密码',
  `create_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP) COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP) COMMENT '更新时间',
  `last_login_at` timestamp COMMENT '最后登录时间',
  `phone` varchar(128) UNIQUE NOT NULL COMMENT '手机号',
  `phone_verified` bool NOT NULL DEFAULT 0 COMMENT '手机号是否验证',
  `status` tinyint NOT NULL DEFAULT 0 COMMENT '用户状态(0: 正常,1:禁用, 2: 删除)',
  `sex` tinyint NOT NULL DEFAULT 0 COMMENT '性别(0: 男,1: 女)'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

CREATE INDEX `idx_nickname` ON `users` (`nickname`);

CREATE UNIQUE INDEX `idx_phone` ON `users` (`phone`);