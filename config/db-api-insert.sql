-- ============================================
-- sys_api 表数据插入脚本
-- 从 Swagger 文档自动生成
-- ============================================

-- 表结构:
-- id: 主键ID
-- handle: 处理器名称(暂时为空)
-- title: API标题/描述
-- path: API路径
-- action: HTTP方法(GET/POST/PUT/DELETE)
-- type: 接口类型(SYS=系统/BUS=业务)
-- create_by: 创建者ID
-- update_by: 更新者ID
-- created_at: 创建时间
-- updated_at: 更新时间
-- deleted_at: 删除时间(软删除)

INSERT INTO sys_api VALUES (1, '', '获取系统前台配置信息，主要注意这里不在验证权限', '/lotus/api/v1/app-config', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (2, '', '获取验证码', '/lotus/api/v1/captcha', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (3, '', '分页列表数据 / page list data', '/lotus/api/v1/db/columns/page', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (4, '', '分页列表数据 / page list data', '/lotus/api/v1/db/tables/page', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (5, '', '删除部门', '/lotus/api/v1/dept', 'DELETE', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (6, '', '分页部门列表数据', '/lotus/api/v1/dept', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (7, '', '添加部门', '/lotus/api/v1/dept', 'POST', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (8, '', '修改部门', '/lotus/api/v1/dept', 'PUT', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (9, '', '获取部门数据', '/lotus/api/v1/dept/get', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (10, '', '数据字典根据key获取', '/lotus/api/v1/dict-data/option-select', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (11, '', '删除字典数据', '/lotus/api/v1/dict/data', 'DELETE', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (12, '', '字典数据列表', '/lotus/api/v1/dict/data', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (13, '', '添加字典数据', '/lotus/api/v1/dict/data', 'POST', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (14, '', '修改字典数据', '/lotus/api/v1/dict/data', 'PUT', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (15, '', '通过编码获取字典数据', '/lotus/api/v1/dict/data/get', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (16, '', '删除字典类型', '/lotus/api/v1/dict/type', 'DELETE', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (17, '', '字典类型列表数据', '/lotus/api/v1/dict/type', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (18, '', '添加字典类型', '/lotus/api/v1/dict/type', 'POST', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (19, '', '修改字典类型', '/lotus/api/v1/dict/type', 'PUT', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (20, '', '字典类型全部数据 代码生成使用接口', '/lotus/api/v1/dict/type-option-select', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (21, '', '字典类型通过字典id获取', '/lotus/api/v1/dict/type/get', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (22, '', '获取个人信息', '/lotus/api/v1/getinfo', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (23, '', '删除定时任务', '/lotus/api/v1/job/remove', 'POST', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (24, '', '启动定时任务', '/lotus/api/v1/job/start', 'POST', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (25, '', '登陆', '/lotus/api/v1/login', 'POST', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (26, '', '删除菜单', '/lotus/api/v1/menu', 'DELETE', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (27, '', 'Menu列表数据', '/lotus/api/v1/menu', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (28, '', '创建菜单', '/lotus/api/v1/menu', 'POST', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (29, '', '修改菜单', '/lotus/api/v1/menu', 'PUT', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (30, '', 'Menu详情数据', '/lotus/api/v1/menu/get', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (31, '', '角色修改使用的菜单列表', '/lotus/api/v1/menuTreeselect', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (32, '', '根据登录角色名称获取菜单列表数据（左菜单使用）', '/lotus/api/v1/menurole', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (33, '', '删除岗位', '/lotus/api/v1/post', 'DELETE', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (34, '', '岗位列表数据', '/lotus/api/v1/post', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (35, '', '添加岗位', '/lotus/api/v1/post', 'POST', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (36, '', '修改岗位', '/lotus/api/v1/post', 'PUT', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (37, '', '获取岗位信息', '/lotus/api/v1/post/get', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (38, '', '上传图片', '/lotus/api/v1/public/uploadFile', 'POST', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (39, '', '删除用户角色', '/lotus/api/v1/role', 'DELETE', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (40, '', '获取SysRole列表', '/lotus/api/v1/role', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (41, '', '创建SysRole', '/lotus/api/v1/role', 'POST', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (42, '', '修改用户角色', '/lotus/api/v1/role', 'PUT', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (43, '', '修改用户角色', '/lotus/api/v1/role-status', 'PUT', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (44, '', '获取Role数据', '/lotus/api/v1/role/get', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (45, '', '更新角色数据权限', '/lotus/api/v1/roledatascope', 'PUT', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (46, '', '获取系统信息', '/lotus/api/v1/server/monitor', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (47, '', '获取配置', '/lotus/api/v1/set-config', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (48, '', '设置配置', '/lotus/api/v1/set-config', 'PUT', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (49, '', '删除接口管理', '/lotus/api/v1/sys-api', 'DELETE', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (50, '', '获取接口管理列表', '/lotus/api/v1/sys-api', 'GET', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (51, '', '修改接口管理', '/lotus/api/v1/sys-api', 'PUT', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (52, '', '获取接口管理', '/lotus/api/v1/sys-api/get', 'GET', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (53, '', '删除配置管理', '/lotus/api/v1/sys-config', 'DELETE', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (54, '', '获取配置管理列表', '/lotus/api/v1/sys-config', 'GET', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (55, '', '创建配置管理', '/lotus/api/v1/sys-config', 'POST', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (56, '', '修改配置管理', '/lotus/api/v1/sys-config', 'PUT', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (57, '', '根据Key获取SysConfig的Service', '/lotus/api/v1/sys-config/get', 'GET', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (58, '', '登录日志删除', '/lotus/api/v1/sys-login-log', 'DELETE', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (59, '', '登录日志列表', '/lotus/api/v1/sys-login-log', 'GET', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (60, '', '登录日志通过id获取', '/lotus/api/v1/sys-login-log/get', 'GET', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (61, '', '删除操作日志', '/lotus/api/v1/sys-opera-log', 'DELETE', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (62, '', '操作日志列表', '/lotus/api/v1/sys-opera-log', 'GET', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (63, '', '操作日志通过id获取', '/lotus/api/v1/sys-opera-log/get', 'GET', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (64, '', '删除权限定义', '/lotus/api/v1/sys-permission', 'DELETE', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (65, '', '获取权限定义列表', '/lotus/api/v1/sys-permission', 'GET', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (66, '', '创建权限定义', '/lotus/api/v1/sys-permission', 'POST', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (67, '', '修改权限定义', '/lotus/api/v1/sys-permission', 'PUT', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (68, '', '获取单个权限定义', '/lotus/api/v1/sys-permission/get', 'GET', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (69, '', '删除用户数据', '/lotus/api/v1/sys-user', 'DELETE', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (70, '', '列表用户信息数据', '/lotus/api/v1/sys-user', 'GET', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (71, '', '创建用户', '/lotus/api/v1/sys-user', 'POST', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (72, '', '修改用户数据', '/lotus/api/v1/sys-user', 'PUT', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (73, '', '获取用户', '/lotus/api/v1/sys-user/get', 'GET', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (74, '', '设置用户角色', '/lotus/api/v1/sys-user/role', 'PUT', 'SYS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (75, '', '添加表结构', '/lotus/api/v1/sys/tables/info', 'POST', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (76, '', '修改表结构', '/lotus/api/v1/sys/tables/info', 'PUT', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (77, '', '删除表结构', '/lotus/api/v1/sys/tables/info/{tableId}', 'DELETE', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (78, '', '获取配置', '/lotus/api/v1/sys/tables/info/{tableId}', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (79, '', '分页列表数据', '/lotus/api/v1/sys/tables/page', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (80, '', '修改头像', '/lotus/api/v1/user/avatar', 'POST', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (81, '', '获取个人中心用户', '/lotus/api/v1/user/profile', 'GET', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (82, '', '重置用户密码', '/lotus/api/v1/user/pwd/reset', 'PUT', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (83, '', '修改密码', '/lotus/api/v1/user/pwd/set', 'PUT', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (84, '', '修改用户状态', '/lotus/api/v1/user/status', 'PUT', 'BUS', 1, 1, NOW(), NOW(), NULL);
INSERT INTO sys_api VALUES (85, '', '退出登录', '/lotus/logout', 'POST', 'BUS', 1, 1, NOW(), NOW(), NULL);

-- 总计: 85 条 API 记录

-- 按标签分组统计:
-- 个人中心: 3 个接口
-- 公共接口: 1 个接口
-- 其他: 1 个接口
-- 字典数据: 6 个接口
-- 字典类型: 6 个接口
-- 定时任务: 2 个接口
-- 岗位: 5 个接口
-- 工具 / 生成工具: 7 个接口
-- 接口管理: 4 个接口
-- 操作日志: 3 个接口
-- 权限定义: 5 个接口
-- 用户: 9 个接口
-- 登录日志: 3 个接口
-- 登陆: 2 个接口
-- 系统监控: 1 个接口
-- 菜单: 7 个接口
-- 角色管理: 7 个接口
-- 部门: 5 个接口
-- 配置管理: 8 个接口