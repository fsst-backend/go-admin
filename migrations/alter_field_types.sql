-- 数据库字段类型修改脚本
-- 适用于 MySQL 数据库

-- 修改 sys_api 表字段类型
ALTER TABLE sys_api MODIFY COLUMN handle VARCHAR(128) COMMENT 'handle';
ALTER TABLE sys_api MODIFY COLUMN title VARCHAR(128) COMMENT '标题';
ALTER TABLE sys_api MODIFY COLUMN path VARCHAR(128) COMMENT '地址';
ALTER TABLE sys_api MODIFY COLUMN action VARCHAR(16) COMMENT '请求类型';
ALTER TABLE sys_api MODIFY COLUMN type VARCHAR(16) COMMENT '接口类型';
ALTER TABLE sys_api MODIFY COLUMN tag VARCHAR(128) COMMENT '标签';

-- 修改 sys_config 表字段类型
ALTER TABLE sys_config MODIFY COLUMN config_name VARCHAR(128) COMMENT 'ConfigName';
ALTER TABLE sys_config MODIFY COLUMN config_key VARCHAR(128) COMMENT 'ConfigKey';
ALTER TABLE sys_config MODIFY COLUMN config_value VARCHAR(255) COMMENT 'ConfigValue';
ALTER TABLE sys_config MODIFY COLUMN config_type VARCHAR(64) COMMENT 'ConfigType';
ALTER TABLE sys_config MODIFY COLUMN is_frontend VARCHAR(64) COMMENT '是否前台';
ALTER TABLE sys_config MODIFY COLUMN remark VARCHAR(128) COMMENT 'Remark';

-- 修改 sys_dept 表字段类型
ALTER TABLE sys_dept MODIFY COLUMN dept_path VARCHAR(255);
ALTER TABLE sys_dept MODIFY COLUMN dept_name VARCHAR(128) COMMENT '部门名称';
ALTER TABLE sys_dept MODIFY COLUMN leader VARCHAR(128) COMMENT '负责人';
ALTER TABLE sys_dept MODIFY COLUMN phone VARCHAR(11) COMMENT '手机';
ALTER TABLE sys_dept MODIFY COLUMN email VARCHAR(64) COMMENT '邮箱';

-- 修改 sys_dict_data 表字段类型
ALTER TABLE sys_dict_data MODIFY COLUMN dict_label VARCHAR(128) COMMENT 'DictLabel';
ALTER TABLE sys_dict_data MODIFY COLUMN dict_value VARCHAR(255) COMMENT 'DictValue';
ALTER TABLE sys_dict_data MODIFY COLUMN dict_type VARCHAR(64) COMMENT 'DictType';
ALTER TABLE sys_dict_data MODIFY COLUMN css_class VARCHAR(128) COMMENT 'CssClass';
ALTER TABLE sys_dict_data MODIFY COLUMN list_class VARCHAR(128) COMMENT 'ListClass';
ALTER TABLE sys_dict_data MODIFY COLUMN is_default VARCHAR(8) COMMENT 'IsDefault';
ALTER TABLE sys_dict_data MODIFY COLUMN default_value VARCHAR(8) COMMENT 'Default';
ALTER TABLE sys_dict_data MODIFY COLUMN remark VARCHAR(255) COMMENT 'Remark';

-- 修改 sys_dict_type 表字段类型
ALTER TABLE sys_dict_type MODIFY COLUMN dict_name VARCHAR(128) COMMENT 'DictName';
ALTER TABLE sys_dict_type MODIFY COLUMN dict_type VARCHAR(128) COMMENT 'DictType';
ALTER TABLE sys_dict_type MODIFY COLUMN remark VARCHAR(255) COMMENT 'Remark';

-- 修改 sys_login_log 表字段类型
ALTER TABLE sys_login_log MODIFY COLUMN username VARCHAR(128) COMMENT '用户名';
ALTER TABLE sys_login_log MODIFY COLUMN status TINYINT COMMENT '状态';
ALTER TABLE sys_login_log MODIFY COLUMN ipaddr VARCHAR(255) COMMENT 'ip地址';
ALTER TABLE sys_login_log MODIFY COLUMN login_location VARCHAR(255) COMMENT '归属地';
ALTER TABLE sys_login_log MODIFY COLUMN browser VARCHAR(255) COMMENT '浏览器';
ALTER TABLE sys_login_log MODIFY COLUMN os VARCHAR(255) COMMENT '系统';
ALTER TABLE sys_login_log MODIFY COLUMN platform VARCHAR(255) COMMENT '固件';
ALTER TABLE sys_login_log MODIFY COLUMN remark VARCHAR(255) COMMENT '备注';
ALTER TABLE sys_login_log MODIFY COLUMN msg VARCHAR(255) COMMENT '信息';

-- 修改 sys_menu 表字段类型
ALTER TABLE sys_menu MODIFY COLUMN menu_name VARCHAR(128);
ALTER TABLE sys_menu MODIFY COLUMN title VARCHAR(128);
ALTER TABLE sys_menu MODIFY COLUMN menu_type VARCHAR(10) COMMENT 'M: Menu/Directory   C: Component/Page  F: Function/Button';
ALTER TABLE sys_menu MODIFY COLUMN menu_path VARCHAR(128);
ALTER TABLE sys_menu MODIFY COLUMN path VARCHAR(128);
ALTER TABLE sys_menu MODIFY COLUMN perm VARCHAR(255);
ALTER TABLE sys_menu MODIFY COLUMN component VARCHAR(255);
ALTER TABLE sys_menu MODIFY COLUMN icon VARCHAR(128);
ALTER TABLE sys_menu MODIFY COLUMN external_link VARCHAR(255);
ALTER TABLE sys_menu MODIFY COLUMN text_badge VARCHAR(10);
ALTER TABLE sys_menu MODIFY COLUMN active_path VARCHAR(100);
ALTER TABLE sys_menu MODIFY COLUMN status VARCHAR(1);
ALTER TABLE sys_menu MODIFY COLUMN permission_code VARCHAR(128) COMMENT '关联权限表code';

-- 修改 sys_opera_log 表字段类型
ALTER TABLE sys_opera_log MODIFY COLUMN title VARCHAR(255) COMMENT '操作模块';
ALTER TABLE sys_opera_log MODIFY COLUMN business_type VARCHAR(128) COMMENT '操作类型';
ALTER TABLE sys_opera_log MODIFY COLUMN business_types VARCHAR(128) COMMENT 'BusinessTypes';
ALTER TABLE sys_opera_log MODIFY COLUMN method VARCHAR(128) COMMENT '函数';
ALTER TABLE sys_opera_log MODIFY COLUMN request_method VARCHAR(128) COMMENT '请求方式 GET POST PUT DELETE';
ALTER TABLE sys_opera_log MODIFY COLUMN operator_type VARCHAR(128) COMMENT '操作类型';
ALTER TABLE sys_opera_log MODIFY COLUMN oper_name VARCHAR(128) COMMENT '操作者';
ALTER TABLE sys_opera_log MODIFY COLUMN dept_name VARCHAR(128) COMMENT '部门名称';
ALTER TABLE sys_opera_log MODIFY COLUMN oper_url VARCHAR(255) COMMENT '访问地址';
ALTER TABLE sys_opera_log MODIFY COLUMN oper_ip VARCHAR(128) COMMENT '客户端ip';
ALTER TABLE sys_opera_log MODIFY COLUMN oper_location VARCHAR(128) COMMENT '访问位置';
ALTER TABLE sys_opera_log MODIFY COLUMN status TINYINT COMMENT '操作状态 1:正常 2:关闭';
ALTER TABLE sys_opera_log MODIFY COLUMN remark VARCHAR(255) COMMENT '备注';
ALTER TABLE sys_opera_log MODIFY COLUMN latency_time VARCHAR(128) COMMENT '耗时';
ALTER TABLE sys_opera_log MODIFY COLUMN user_agent VARCHAR(255) COMMENT 'ua';

-- 修改 sys_permission 表字段类型
ALTER TABLE sys_permission MODIFY COLUMN code VARCHAR(128) COMMENT '权限唯一编码';
ALTER TABLE sys_permission MODIFY COLUMN name VARCHAR(128) COMMENT '权限名称';
ALTER TABLE sys_permission MODIFY COLUMN type VARCHAR(16) COMMENT '权限类型(menu/button/api/page)';
ALTER TABLE sys_permission MODIFY COLUMN remark VARCHAR(255) COMMENT '备注说明';

-- 修改 sys_permission_api 表字段类型
ALTER TABLE sys_permission_api MODIFY COLUMN permission_id INT COMMENT '权限ID';
ALTER TABLE sys_permission_api MODIFY COLUMN api_id INT COMMENT 'API ID';

-- 修改 sys_post 表字段类型
ALTER TABLE sys_post MODIFY COLUMN post_name VARCHAR(128) COMMENT '岗位名称';
ALTER TABLE sys_post MODIFY COLUMN post_code VARCHAR(128) COMMENT '岗位代码';
ALTER TABLE sys_post MODIFY COLUMN remark VARCHAR(255) COMMENT '描述';

-- 修改 sys_role 表字段类型
ALTER TABLE sys_role MODIFY COLUMN role_name VARCHAR(128) COMMENT '角色名称';
ALTER TABLE sys_role MODIFY COLUMN status TINYINT COMMENT '状态 1禁用 2正常';
ALTER TABLE sys_role MODIFY COLUMN role_key VARCHAR(128) COMMENT '角色代码';
ALTER TABLE sys_role MODIFY COLUMN flag VARCHAR(128);
ALTER TABLE sys_role MODIFY COLUMN remark VARCHAR(255) COMMENT '备注';
ALTER TABLE sys_role MODIFY COLUMN data_scope VARCHAR(128) COMMENT '数据范围';

-- 修改 sys_role_dept 表字段类型
ALTER TABLE sys_role_dept MODIFY COLUMN role_id INT COMMENT '角色编码';
ALTER TABLE sys_role_dept MODIFY COLUMN dept_id INT COMMENT '部门编码';

-- 修改 sys_role_menu 表字段类型
ALTER TABLE sys_role_menu MODIFY COLUMN role_id INT COMMENT '角色编码';
ALTER TABLE sys_role_menu MODIFY COLUMN menu_id INT COMMENT '菜单编码';

-- 修改 sys_role_permission 表字段类型
ALTER TABLE sys_role_permission MODIFY COLUMN role_id INT COMMENT '角色ID';
ALTER TABLE sys_role_permission MODIFY COLUMN permission_id INT COMMENT '权限ID';

-- 修改 sys_user 表字段类型
-- 单设备登录字段（由 migration 1739000000000 自动添加，或手动执行）:
-- ALTER TABLE sys_user ADD COLUMN token_version INT DEFAULT 0 COMMENT '登录令牌版本';
ALTER TABLE sys_user MODIFY COLUMN uuid VARCHAR(255) COMMENT 'UUID';
ALTER TABLE sys_user MODIFY COLUMN username VARCHAR(64) COMMENT '用户名';
ALTER TABLE sys_user MODIFY COLUMN password VARCHAR(128) COMMENT '密码';
ALTER TABLE sys_user MODIFY COLUMN nick_name VARCHAR(128) COMMENT '昵称';
ALTER TABLE sys_user MODIFY COLUMN phone VARCHAR(11) COMMENT '手机号';
ALTER TABLE sys_user MODIFY COLUMN salt VARCHAR(255) COMMENT '加盐';
ALTER TABLE sys_user MODIFY COLUMN avatar VARCHAR(255) COMMENT '头像';
ALTER TABLE sys_user MODIFY COLUMN sex VARCHAR(255) COMMENT '性别';
ALTER TABLE sys_user MODIFY COLUMN email VARCHAR(128) COMMENT '邮箱';
ALTER TABLE sys_user MODIFY COLUMN remark VARCHAR(255) COMMENT '备注';
ALTER TABLE sys_user MODIFY COLUMN status TINYINT COMMENT '状态';

-- 修改 sys_user_role 表字段类型
ALTER TABLE sys_user_role MODIFY COLUMN user_id INT COMMENT '用户ID';
ALTER TABLE sys_user_role MODIFY COLUMN role_id INT COMMENT '角色ID';