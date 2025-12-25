-- 创建用户与角色关系表
-- sys_user_role 表用于存储用户与角色的多对多关系

-- MySQL 语法
CREATE TABLE IF NOT EXISTS `sys_user_role` (
  `user_id` int(11) NOT NULL COMMENT '用户ID',
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  `create_by` int(11) DEFAULT NULL COMMENT '创建者',
  `update_by` int(11) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`user_id`,`role_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户与角色关联表';

-- 示例：为用户分配角色关系
-- INSERT INTO sys_user_role (user_id, role_id, create_by, created_at, updated_at) 
-- VALUES (1, 1, 1, NOW(), NOW());

-- PostgreSQL 语法
-- CREATE TABLE IF NOT EXISTS sys_user_role (
--   user_id int NOT NULL,
--   role_id int NOT NULL,
--   create_by int DEFAULT NULL,
--   update_by int DEFAULT NULL,
--   created_at timestamp DEFAULT NULL,
--   updated_at timestamp DEFAULT NULL,
--   deleted_at timestamp DEFAULT NULL,
--   PRIMARY KEY (user_id, role_id)
-- );
-- CREATE INDEX idx_user_role_user_id ON sys_user_role(user_id);
-- CREATE INDEX idx_user_role_role_id ON sys_user_role(role_id);

-- SQLite 语法
-- CREATE TABLE IF NOT EXISTS sys_user_role (
--   user_id INTEGER NOT NULL,
--   role_id INTEGER NOT NULL,
--   create_by INTEGER DEFAULT NULL,
--   update_by INTEGER DEFAULT NULL,
--   created_at TEXT DEFAULT NULL,
--   updated_at TEXT DEFAULT NULL,
--   deleted_at TEXT DEFAULT NULL,
--   PRIMARY KEY (user_id, role_id)
-- );
-- CREATE INDEX idx_user_role_user_id ON sys_user_role(user_id);
-- CREATE INDEX idx_user_role_role_id ON sys_user_role(role_id);
