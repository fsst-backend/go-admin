-- 开始初始化数据;

-- sys_config 表结构: id, config_name, config_key, config_value, config_type, is_frontend, remark, create_by, update_by, created_at, updated_at, deleted_at
INSERT INTO sys_config VALUES (1, '皮肤样式', 'sys_index_skinName', 'skin-green', 'Y', '0', '主框架页-默认皮肤样式名称:蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow', 1, 1, '2021-05-13 19:56:37.913', '2021-06-05 13:50:13.123', NULL);
INSERT INTO sys_config VALUES (3, '侧栏主题', 'sys_index_sideTheme', 'theme-dark', 'Y', '0', '主框架页-侧边栏主题:深色主题theme-dark，浅色主题theme-light', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO sys_config VALUES (4, '系统名称', 'sys_app_name', 'go-admin管理系统', 'Y', '1', '', 1, 0, '2021-03-17 08:52:06.067', '2021-05-28 10:08:25.248', NULL);
--INSERT INTO sys_config VALUES (5, '系统logo', 'sys_app_logo', 'https://doc-image.zhangwj.com/img/go-admin.png', 'Y', '1', '', 1, 0, '2021-03-17 08:53:19.462', '2021-03-17 08:53:19.462', NULL);

-- sys_dept 表结构: dept_id, parent_id, dept_path, dept_name, dept_catalog, sort, leader, phone, email, status, create_by, update_by, created_at, updated_at, deleted_at
-- dept_catalog 类型说明: finance(财务), hr(人力资源), customer_service(客服), rnd(研发), sales(销售), market(市场), operation(运营), ops(运维)
INSERT INTO sys_dept VALUES (1, 0, '/0/1/', 'futren', '', 0, 'futren', '17777777789', 'futren@gamil.com', 1, 1, 1, '2021-05-13 19:56:37.913', '2021-06-05 17:06:44.960', NULL);
INSERT INTO sys_dept VALUES (7, 1, '/0/1/7/', '研发部', 'rnd', 1, 'futren', '17777777789', 'futren@gamil.com', 1, 1, 1, '2021-05-13 19:56:37.913', '2021-06-16 21:35:00.109', NULL);
INSERT INTO sys_dept VALUES (8, 1, '/0/1/8/', '运维部', 'ops', 0, 'futren', '17777777789', 'futren@gamil.com', 1, 1, 1, '2021-05-13 19:56:37.913', '2021-06-16 21:41:39.747', NULL);
INSERT INTO sys_dept VALUES (9, 1, '/0/1/9/', '客服部', 'customer_service', 0, 'futren', '17777777789', 'futren@gamil.com', 1, 1, 1, '2021-05-13 19:56:37.913', '2021-06-05 17:07:05.993', NULL);
INSERT INTO sys_dept VALUES (10, 1, '/0/1/10/', '人力资源', 'hr', 3, 'futren', '17777777789', 'futren@gamil.com', 1, 1, 1, '2021-05-13 19:56:37.913', '2021-06-05 17:07:08.503', NULL);

-- sys_dict_data 表结构: dict_code, dict_sort, dict_label, dict_value, dict_type, css_class, list_class, is_default, status, default_value, remark, create_by, update_by, created_at, updated_at, deleted_at
INSERT INTO sys_dict_data VALUES (1, 0, '正常', '2', 'sys_normal_disable', '', '', '', 1, '', '系统正常', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.168', NULL);
INSERT INTO sys_dict_data VALUES (2, 0, '停用', '1', 'sys_normal_disable', '', '', '', 1, '', '系统停用', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (3, 0, '男', '0', 'sys_user_sex', '', '', '', 1, '', '性别男', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (4, 0, '女', '1', 'sys_user_sex', '', '', '', 1, '', '性别女', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (5, 0, '未知', '2', 'sys_user_sex', '', '', '', 1, '', '性别未知', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (6, 0, '显示', '0', 'sys_show_hide', '', '', '', 1, '', '显示菜单', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (7, 0, '隐藏', '1', 'sys_show_hide', '', '', '', 1, '', '隐藏菜单', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (8, 0, '是', 'Y', 'sys_yes_no', '', '', '', 1, '', '系统默认是', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (9, 0, '否', 'N', 'sys_yes_no', '', '', '', 1, '', '系统默认否', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (10, 0, '正常', '2', 'sys_job_status', '', '', '', 1, '', '正常状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (11, 0, '停用', '1', 'sys_job_status', '', '', '', 1, '', '停用状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (12, 0, '默认', 'DEFAULT', 'sys_job_group', '', '', '', 1, '', '默认分组', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (13, 0, '系统', 'SYSTEM', 'sys_job_group', '', '', '', 1, '', '系统分组', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (14, 0, '通知', '1', 'sys_notice_type', '', '', '', 1, '', '通知', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (15, 0, '公告', '2', 'sys_notice_type', '', '', '', 1, '', '公告', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (16, 0, '正常', '2', 'sys_common_status', '', '', '', 1, '', '正常状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (17, 0, '关闭', '1', 'sys_common_status', '', '', '', 1, '', '关闭状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (18, 0, '新增', '1', 'sys_oper_type', '', '', '', 1, '', '新增操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (19, 0, '修改', '2', 'sys_oper_type', '', '', '', 1, '', '修改操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (20, 0, '删除', '3', 'sys_oper_type', '', '', '', 1, '', '删除操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (21, 0, '授权', '4', 'sys_oper_type', '', '', '', 1, '', '授权操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (22, 0, '导出', '5', 'sys_oper_type', '', '', '', 1, '', '导出操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (23, 0, '导入', '6', 'sys_oper_type', '', '', '', 1, '', '导入操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (24, 0, '强退', '7', 'sys_oper_type', '', '', '', 1, '', '强退操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (25, 0, '生成代码', '8', 'sys_oper_type', '', '', '', 1, '', '生成操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (26, 0, '清空数据', '9', 'sys_oper_type', '', '', '', 1, '', '清空操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (27, 0, '成功', '0', 'sys_notice_status', '', '', '', 1, '', '成功状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (28, 0, '失败', '1', 'sys_notice_status', '', '', '', 1, '', '失败状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (29, 0, '登录', '10', 'sys_oper_type', '', '', '', 1, '', '登录操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (30, 0, '退出', '11', 'sys_oper_type', '', '', '', 1, '', '', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (31, 0, '获取验证码', '12', 'sys_oper_type', '', '', '', 1, '', '获取验证码', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_data VALUES (32, 0, '正常', '1', 'sys_content_status', '', '', '', 1, '', '', 1, 1, '2021-05-13 19:56:40.845', '2021-05-13 19:56:40.845', NULL);
INSERT INTO sys_dict_data VALUES (33, 1, '禁用', '2', 'sys_content_status', '', '', '', 1, '', '', 1, 1, '2021-05-13 19:56:40.845', '2021-05-13 19:56:40.845', NULL);

-- sys_dict_type 表结构: dict_id, dict_name, dict_type, status, remark, create_by, update_by, created_at, updated_at, deleted_at
INSERT INTO sys_dict_type VALUES (1, '系统开关', 'sys_normal_disable', 1, '系统开关列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (2, '用户性别', 'sys_user_sex', 1, '用户性别列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (3, '菜单状态', 'sys_show_hide', 1, '菜单状态列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (4, '系统是否', 'sys_yes_no', 1, '系统是否列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (5, '任务状态', 'sys_job_status', 1, '任务状态列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (6, '任务分组', 'sys_job_group', 1, '任务分组列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (7, '通知类型', 'sys_notice_type', 1, '通知类型列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (8, '系统状态', 'sys_common_status', 1, '登录状态列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (9, '操作类型', 'sys_oper_type', 1, '操作类型列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (10, '通知状态', 'sys_notice_status', 1, '通知状态列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO sys_dict_type VALUES (11, '内容状态', 'sys_content_status', 1, '', 1, 1, '2021-05-13 19:56:40.813', '2021-05-13 19:56:40.813', NULL);

-- sys_job 表结构: job_id, job_name, job_group, job_type, cron_expression, invoke_target, args, misfire_policy, concurrent, status, entry_id, create_by, update_by, created_at, updated_at, deleted_at
INSERT INTO sys_job (job_id, job_name, job_group, job_type, cron_expression, invoke_target, args, misfire_policy, concurrent, status, entry_id, create_by, update_by, created_at, updated_at, deleted_at) VALUES (1, '接口测试', 'DEFAULT', 1, '0/5 * * * * ', 'http://localhost:13348', '', 1, 1, 1, 0, 1, 1, '2021-05-13 19:56:37.914', '2021-06-14 20:59:55.417', NULL);
INSERT INTO sys_job (job_id, job_name, job_group, job_type, cron_expression, invoke_target, args, misfire_policy, concurrent, status, entry_id, create_by, update_by, created_at, updated_at, deleted_at) VALUES (2, '函数测试', 'DEFAULT', 2, '0/5 * * * * ', 'ExamplesOne', '参数', 1, 1, 1, 0, 1, 1, '2021-05-13 19:56:37.914', '2021-05-31 23:55:37.221', NULL);

-- sys_post 表结构: post_id, post_name, post_code, sort, status, remark, create_by, update_by, created_at, updated_at, deleted_at
INSERT INTO sys_post VALUES (1, '首席执行官', 'CEO', 0, 1, '首席执行官', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO sys_post VALUES (2, '首席技术执行官', 'CTO', 2, 1, '首席技术执行官', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO sys_post VALUES (3, '首席运营官', 'COO', 3, 1, '测试工程师', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);

-- sys_role 表结构: role_id, role_name, status, role_key, role_sort, flag, remark, admin, data_scope, create_by, update_by, created_at, updated_at, deleted_at
INSERT INTO sys_role VALUES (1, '系统管理员', '2', 'superadmin', 1, '', '', true, '', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);

-- sys_user 表结构: user_id, uuid, username, password, nick_name, phone, salt, avatar, sex, email, dept_id, post_id, remark, status, create_by, update_by, created_at, updated_at, deleted_at
INSERT INTO sys_user (user_id, uuid, username, password, nick_name, phone, salt, avatar, sex, email, dept_id, post_id, remark, status, create_by, update_by, created_at, updated_at, deleted_at) VALUES (1, '8ab590bb-f729-493e-ba2f-c1f863ecd6f5', 'admin', '$2a$10$TzZ1dVFwOOq/7L7Y.LaTle./qbDg3R3fqZ.6zAA4F3Vv7cp1QO3uO', 'admin', '13818888888', '', '', '1', 'admin@email.com', 1, 1, '', 2, 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);

-- sys_user_role 表结构: id, user_id, role_id, create_by, update_by, created_at, updated_at, deleted_at
INSERT INTO sys_user_role VALUES (1, 1, 1, 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);

-- sys_casbin_rule 表结构: id, ptype, v0, v1, v2, v3, v4, v5
-- ptype='p' 表示策略(policy): v0=角色, v1=资源路径, v2=操作方法
-- ptype='g' 表示角色继承/用户角色关系(grouping): v0=用户, v1=角色
INSERT INTO sys_casbin_rule (id, ptype, v0, v1, v2, v3, v4, v5) VALUES (1, 'p', 'superadmin', '*', '*', '', '', '');
INSERT INTO sys_casbin_rule (id, ptype, v0, v1, v2, v3, v4, v5) VALUES (2, 'g', 'user_1', 'superadmin', '', '', '', '');

-- 数据完成 ;
