-- 开始初始化数据;
INSERT INTO sys_config VALUES (1, '皮肤样式', 'sys_index_skinName', 'skin-green', 'Y', '0', '主框架页-默认皮肤样式名称:蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow', 1, 1, '2021-05-13 19:56:37.913', '2021-06-05 13:50:13.123', NULL);
INSERT INTO sys_config VALUES (2, '初始密码', 'sys_user_initPassword', '123456', 'Y', '0', '用户管理-账号初始密码:123456', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO sys_config VALUES (3, '侧栏主题', 'sys_index_sideTheme', 'theme-dark', 'Y', '0', '主框架页-侧边栏主题:深色主题theme-dark，浅色主题theme-light', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO sys_config VALUES (4, '系统名称', 'sys_app_name', 'go-admin管理系统', 'Y', '1', '', 1, 0, '2021-03-17 08:52:06.067', '2021-05-28 10:08:25.248', NULL);
INSERT INTO sys_config VALUES (5, '系统logo', 'sys_app_logo', 'https://doc-image.zhangwj.com/img/go-admin.png', 'Y', '1', '', 1, 0, '2021-03-17 08:53:19.462', '2021-03-17 08:53:19.462', NULL);

INSERT INTO sys_dept VALUES (1, 0, '/0/1/', 'futren', 0, 'futren', '17777777789', 'futren@gamil.com', '2', 1, 1, '2021-05-13 19:56:37.913', '2021-06-05 17:06:44.960', NULL);
INSERT INTO sys_dept VALUES (7, 1, '/0/1/7/', '研发部', 1, 'futren', '17777777789', 'futren@gamil.com', '2', 1, 1, '2021-05-13 19:56:37.913', '2021-06-16 21:35:00.109', NULL);
INSERT INTO sys_dept VALUES (8, 1, '/0/1/8/', '运维部', 0, 'futren', '17777777789', 'futren@gamil.com', '2', 1, 1, '2021-05-13 19:56:37.913', '2021-06-16 21:41:39.747', NULL);
INSERT INTO sys_dept VALUES (9, 1, '/0/1/9/', '客服部', 0, 'futren', '17777777789', 'futren@gamil.com', '2', 1, 1, '2021-05-13 19:56:37.913', '2021-06-05 17:07:05.993', NULL);
INSERT INTO sys_dept VALUES (10, 1, '/0/1/10/', '人力资源', 3, 'futren', '17777777789', 'futren@gamil.com', '1', 1, 1, '2021-05-13 19:56:37.913', '2021-06-05 17:07:08.503', NULL);

INSERT INTO sys_dict_data VALUES (1, 0, '正常', '2', 'sys_normal_disable', '', '', '', '2', '', '系统正常', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.168', NULL);
INSERT INTO sys_dict_data VALUES (2, 0, '停用', '1', 'sys_normal_disable', '', '', '', '2', '', '系统停用', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (3, 0, '男', '0', 'sys_user_sex', '', '', '', '2', '', '性别男', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (4, 0, '女', '1', 'sys_user_sex', '', '', '', '2', '', '性别女', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (5, 0, '未知', '2', 'sys_user_sex', '', '', '', '2', '', '性别未知', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (6, 0, '显示', '0', 'sys_show_hide', '', '', '', '2', '', '显示菜单', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (7, 0, '隐藏', '1', 'sys_show_hide', '', '', '', '2', '', '隐藏菜单', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (8, 0, '是', 'Y', 'sys_yes_no', '', '', '', '2', '', '系统默认是', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (9, 0, '否', 'N', 'sys_yes_no', '', '', '', '2', '', '系统默认否', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (10, 0, '正常', '2', 'sys_job_status', '', '', '', '2', '', '正常状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (11, 0, '停用', '1', 'sys_job_status', '', '', '', '2', '', '停用状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (12, 0, '默认', 'DEFAULT', 'sys_job_group', '', '', '', '2', '', '默认分组', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (13, 0, '系统', 'SYSTEM', 'sys_job_group', '', '', '', '2', '', '系统分组', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (14, 0, '通知', '1', 'sys_notice_type', '', '', '', '2', '', '通知', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (15, 0, '公告', '2', 'sys_notice_type', '', '', '', '2', '', '公告', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (16, 0, '正常', '2', 'sys_common_status', '', '', '', '2', '', '正常状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (17, 0, '关闭', '1', 'sys_common_status', '', '', '', '2', '', '关闭状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (18, 0, '新增', '1', 'sys_oper_type', '', '', '', '2', '', '新增操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (19, 0, '修改', '2', 'sys_oper_type', '', '', '', '2', '', '修改操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (20, 0, '删除', '3', 'sys_oper_type', '', '', '', '2', '', '删除操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (21, 0, '授权', '4', 'sys_oper_type', '', '', '', '2', '', '授权操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (22, 0, '导出', '5', 'sys_oper_type', '', '', '', '2', '', '导出操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (23, 0, '导入', '6', 'sys_oper_type', '', '', '', '2', '', '导入操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (24, 0, '强退', '7', 'sys_oper_type', '', '', '', '2', '', '强退操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (25, 0, '生成代码', '8', 'sys_oper_type', '', '', '', '2', '', '生成操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (26, 0, '清空数据', '9', 'sys_oper_type', '', '', '', '2', '', '清空操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (27, 0, '成功', '0', 'sys_notice_status', '', '', '', '2', '', '成功状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (28, 0, '失败', '1', 'sys_notice_status', '', '', '', '2', '', '失败状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (29, 0, '登录', '10', 'sys_oper_type', '', '', '', '2', '', '登录操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (30, 0, '退出', '11', 'sys_oper_type', '', '', '', '2', '', '', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (31, 0, '获取验证码', '12', 'sys_oper_type', '', '', '', '2', '', '获取验证码', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (32, 0, '正常', '1', 'sys_content_status', '', '', '', '1', '', '', 1, 1, '2021-05-13 19:56:40.845', '2021-05-13 19:56:40.845', NULL);
INSERT INTO sys_dict_data VALUES (33, 1, '禁用', '2', 'sys_content_status', '', '', '', '1', '', '', 1, 1, '2021-05-13 19:56:40.845', '2021-05-13 19:56:40.845', NULL);

INSERT INTO sys_dict_type VALUES (1, '系统开关', 'sys_normal_disable', '2', '系统开关列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (2, '用户性别', 'sys_user_sex', '2', '用户性别列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (3, '菜单状态', 'sys_show_hide', '2', '菜单状态列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (4, '系统是否', 'sys_yes_no', '2', '系统是否列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (5, '任务状态', 'sys_job_status', '2', '任务状态列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (6, '任务分组', 'sys_job_group', '2', '任务分组列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (7, '通知类型', 'sys_notice_type', '2', '通知类型列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (8, '系统状态', 'sys_common_status', '2', '登录状态列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (9, '操作类型', 'sys_oper_type', '2', '操作类型列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (10, '通知状态', 'sys_notice_status', '2', '通知状态列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (11, '内容状态', 'sys_content_status', '2', '', 1, 1, '2021-05-13 19:56:40.813', '2021-05-13 19:56:40.813', NULL);

INSERT INTO sys_job VALUES (1, '接口测试', 'DEFAULT', 1, '0/5 * * * * ', 'http://localhost:8000', '', 1, 1, 1, 0, '2021-05-13 19:56:37.914', '2021-06-14 20:59:55.417', NULL, 1, 1);
INSERT INTO sys_job VALUES (2, '函数测试', 'DEFAULT', 2, '0/5 * * * * ', 'ExamplesOne', '参数', 1, 1, 1, 0, '2021-05-13 19:56:37.914', '2021-05-31 23:55:37.221', NULL, 1, 1);


INSERT INTO sys_post VALUES (1, '首席执行官', 'CEO', 0, '2','首席执行官', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO sys_post VALUES (2, '首席技术执行官', 'CTO', 2, '2','首席技术执行官', 1, 1,'2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO sys_post VALUES (3, '首席运营官', 'COO', 3, '2','测试工程师', 1, 1,'2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO sys_role VALUES (1, '系统管理员', '2', 'superadmin', 1, '', '', true, '', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
-- sys_user 表结构: user_id, uuid, username, password, nick_name, phone, salt, avatar, sex, email, dept_id, post_id, remark, status, create_by, update_by, created_at, updated_at, deleted_at
INSERT INTO sys_user VALUES (1, 'b270eb86-6b4e-4bc1-9bb1-aa4a807a3ab5', 'admin@761242.com', '$2a$10$679frSflZc5DhKHb7k9uduoU8V9IcwQ1H4JahWd44WjM2BzGhHxE2', 'futren', '17777777789', '', '', '1', '1@qq.com', 1, 1, '', '2', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);

-- sys_user_role 表结构: id, user_id, role_id, create_by, update_by, created_at, updated_at, deleted_at
INSERT INTO sys_user_role VALUES (1, 1, 1, 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);

-- sys_casbin_rule 表结构: id, ptype, v0, v1, v2, v3, v4, v5
-- ptype='p' 表示策略(policy): v0=角色, v1=资源路径, v2=操作方法
-- ptype='g' 表示角色继承/用户角色关系(grouping): v0=用户, v1=角色
INSERT INTO sys_casbin_rule VALUES (1, 'p', 'superadmin', '*', '*', '', '', '');
INSERT INTO sys_casbin_rule VALUES (2, 'g', 'user_1', 'superadmin', '', '', '', '');



INSERT INTO sys_api VALUES (1, '', '获取系统前台配置信息，主要注意这里不在验证权限', '/lotus/api/v1/app-config', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (2, '', '获取验证码', '/lotus/api/v1/captcha', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (3, '', '分页列表数据 / page list data', '/lotus/api/v1/db/columns/page', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (4, '', '分页列表数据 / page list data', '/lotus/api/v1/db/tables/page', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (5, '', '删除部门', '/lotus/api/v1/dept', 'DELETE', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (6, '', '分页部门列表数据', '/lotus/api/v1/dept', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (7, '', '添加部门', '/lotus/api/v1/dept', 'POST', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (8, '', '修改部门', '/lotus/api/v1/dept', 'PUT', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (9, '', '获取部门数据', '/lotus/api/v1/dept/get', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (10, '', '数据字典根据key获取', '/lotus/api/v1/dict-data/option-select', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (11, '', '删除字典数据', '/lotus/api/v1/dict/data', 'DELETE', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (12, '', '字典数据列表', '/lotus/api/v1/dict/data', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (13, '', '添加字典数据', '/lotus/api/v1/dict/data', 'POST', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (14, '', '修改字典数据', '/lotus/api/v1/dict/data', 'PUT', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (15, '', '通过编码获取字典数据', '/lotus/api/v1/dict/data/get', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (16, '', '删除字典类型', '/lotus/api/v1/dict/type', 'DELETE', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (17, '', '字典类型列表数据', '/lotus/api/v1/dict/type', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (18, '', '添加字典类型', '/lotus/api/v1/dict/type', 'POST', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (19, '', '修改字典类型', '/lotus/api/v1/dict/type', 'PUT', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (20, '', '字典类型全部数据 代码生成使用接口', '/lotus/api/v1/dict/type-option-select', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (21, '', '字典类型通过字典id获取', '/lotus/api/v1/dict/type/get', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (22, '', '获取个人信息', '/lotus/api/v1/getinfo', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (23, '', '删除定时任务', '/lotus/api/v1/job/remove', 'POST', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (24, '', '启动定时任务', '/lotus/api/v1/job/start', 'POST', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (25, '', '登陆', '/lotus/api/v1/login', 'POST', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (26, '', '删除菜单', '/lotus/api/v1/menu', 'DELETE', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (27, '', 'Menu列表数据', '/lotus/api/v1/menu', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (28, '', '创建菜单', '/lotus/api/v1/menu', 'POST', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (29, '', '修改菜单', '/lotus/api/v1/menu', 'PUT', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (30, '', 'Menu详情数据', '/lotus/api/v1/menu/get', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (31, '', '角色修改使用的菜单列表', '/lotus/api/v1/menuTreeselect', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (32, '', '根据登录角色名称获取菜单列表数据（左菜单使用）', '/lotus/api/v1/menurole', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (33, '', '删除岗位', '/lotus/api/v1/post', 'DELETE', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (34, '', '岗位列表数据', '/lotus/api/v1/post', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (35, '', '添加岗位', '/lotus/api/v1/post', 'POST', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (36, '', '修改岗位', '/lotus/api/v1/post', 'PUT', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (37, '', '获取岗位信息', '/lotus/api/v1/post/get', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (38, '', '上传图片', '/lotus/api/v1/public/uploadFile', 'POST', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (39, '', '删除用户角色', '/lotus/api/v1/role', 'DELETE', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (40, '', '获取SysRole列表', '/lotus/api/v1/role', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (41, '', '创建SysRole', '/lotus/api/v1/role', 'POST', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (42, '', '修改用户角色', '/lotus/api/v1/role', 'PUT', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (43, '', '修改用户角色', '/lotus/api/v1/role-status', 'PUT', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (44, '', '获取Role数据', '/lotus/api/v1/role/get', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (45, '', '更新角色数据权限', '/lotus/api/v1/roledatascope', 'PUT', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (46, '', '获取系统信息', '/lotus/api/v1/server/monitor', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (47, '', '获取配置', '/lotus/api/v1/set-config', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (48, '', '设置配置', '/lotus/api/v1/set-config', 'PUT', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (49, '', '删除接口管理', '/lotus/api/v1/sys-api', 'DELETE', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (50, '', '获取接口管理列表', '/lotus/api/v1/sys-api', 'GET', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (51, '', '修改接口管理', '/lotus/api/v1/sys-api', 'PUT', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (52, '', '获取接口管理', '/lotus/api/v1/sys-api/get', 'GET', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (53, '', '删除配置管理', '/lotus/api/v1/sys-config', 'DELETE', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (54, '', '获取配置管理列表', '/lotus/api/v1/sys-config', 'GET', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (55, '', '创建配置管理', '/lotus/api/v1/sys-config', 'POST', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (56, '', '修改配置管理', '/lotus/api/v1/sys-config', 'PUT', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (57, '', '根据Key获取SysConfig的Service', '/lotus/api/v1/sys-config/get', 'GET', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (58, '', '登录日志删除', '/lotus/api/v1/sys-login-log', 'DELETE', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (59, '', '登录日志列表', '/lotus/api/v1/sys-login-log', 'GET', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (60, '', '登录日志通过id获取', '/lotus/api/v1/sys-login-log/get', 'GET', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (61, '', '删除操作日志', '/lotus/api/v1/sys-opera-log', 'DELETE', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (62, '', '操作日志列表', '/lotus/api/v1/sys-opera-log', 'GET', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (63, '', '操作日志通过id获取', '/lotus/api/v1/sys-opera-log/get', 'GET', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (64, '', '删除权限定义', '/lotus/api/v1/sys-permission', 'DELETE', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (65, '', '获取权限定义列表', '/lotus/api/v1/sys-permission', 'GET', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (66, '', '创建权限定义', '/lotus/api/v1/sys-permission', 'POST', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (67, '', '修改权限定义', '/lotus/api/v1/sys-permission', 'PUT', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (68, '', '获取单个权限定义', '/lotus/api/v1/sys-permission/get', 'GET', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (69, '', '删除用户数据', '/lotus/api/v1/sys-user', 'DELETE', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (70, '', '列表用户信息数据', '/lotus/api/v1/sys-user', 'GET', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (71, '', '创建用户', '/lotus/api/v1/sys-user', 'POST', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (72, '', '修改用户数据', '/lotus/api/v1/sys-user', 'PUT', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (73, '', '获取用户', '/lotus/api/v1/sys-user/get', 'GET', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (74, '', '设置用户角色', '/lotus/api/v1/sys-user/role', 'PUT', 'SYS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (75, '', '添加表结构', '/lotus/api/v1/sys/tables/info', 'POST', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (76, '', '修改表结构', '/lotus/api/v1/sys/tables/info', 'PUT', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (77, '', '删除表结构', '/lotus/api/v1/sys/tables/info/{tableId}', 'DELETE', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (78, '', '获取配置', '/lotus/api/v1/sys/tables/info/{tableId}', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (79, '', '分页列表数据', '/lotus/api/v1/sys/tables/page', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (80, '', '修改头像', '/lotus/api/v1/user/avatar', 'POST', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (81, '', '获取个人中心用户', '/lotus/api/v1/user/profile', 'GET', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (82, '', '重置用户密码', '/lotus/api/v1/user/pwd/reset', 'PUT', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (83, '', '修改密码', '/lotus/api/v1/user/pwd/set', 'PUT', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (84, '', '修改用户状态', '/lotus/api/v1/user/status', 'PUT', 'BUS', NOW(), NOW(), NULL, 1, 1);
INSERT INTO sys_api VALUES (85, '', '退出登录', '/lotus/logout', 'POST', 'BUS', NOW(), NOW(), NULL, 1, 1);


-- 数据完成 ;