package migrations

import (
	"fmt"

	"gorm.io/gorm"
)

type AlterFieldTypes struct{}

func (a AlterFieldTypes) Name() string {
	return "alter_field_types"
}

func (a AlterFieldTypes) Up(tx *gorm.DB) error {
	// 修改 sys_api 表字段类型
	if err := alterSysApiFields(tx); err != nil {
		return err
	}

	// 修改 sys_config 表字段类型
	if err := alterSysConfigFields(tx); err != nil {
		return err
	}

	// 修改 sys_dept 表字段类型
	if err := alterSysDeptFields(tx); err != nil {
		return err
	}

	// 修改 sys_dict_data 表字段类型
	if err := alterSysDictDataFields(tx); err != nil {
		return err
	}

	// 修改 sys_dict_type 表字段类型
	if err := alterSysDictTypeFields(tx); err != nil {
		return err
	}

	// 修改 sys_login_log 表字段类型
	if err := alterSysLoginLogFields(tx); err != nil {
		return err
	}

	// 修改 sys_menu 表字段类型
	if err := alterSysMenuFields(tx); err != nil {
		return err
	}

	// 修改 sys_opera_log 表字段类型
	if err := alterSysOperaLogFields(tx); err != nil {
		return err
	}

	// 修改 sys_permission 表字段类型
	if err := alterSysPermissionFields(tx); err != nil {
		return err
	}

	// 修改 sys_permission_api 表字段类型
	if err := alterSysPermissionApiFields(tx); err != nil {
		return err
	}

	// 修改 sys_post 表字段类型
	if err := alterSysPostFields(tx); err != nil {
		return err
	}

	// 修改 sys_role 表字段类型
	if err := alterSysRoleFields(tx); err != nil {
		return err
	}

	// 修改 sys_role_dept 表字段类型
	if err := alterSysRoleDeptFields(tx); err != nil {
		return err
	}

	// 修改 sys_role_menu 表字段类型
	if err := alterSysRoleMenuFields(tx); err != nil {
		return err
	}

	// 修改 sys_role_permission 表字段类型
	if err := alterSysRolePermissionFields(tx); err != nil {
		return err
	}

	// 修改 sys_user 表字段类型
	if err := alterSysUserFields(tx); err != nil {
		return err
	}

	// 修改 sys_user_role 表字段类型
	if err := alterSysUserRoleFields(tx); err != nil {
		return err
	}

	return nil
}

func (a AlterFieldTypes) Down(tx *gorm.DB) error {
	// 由于字段类型修改通常不可逆，这里只提供提示信息
	fmt.Println("Down migration for field type changes is not implemented as it may cause data loss.")
	return nil
}

func alterSysApiFields(tx *gorm.DB) error {
	var sqls []string

	// 检查并修改字段类型
	if dbType := tx.Dialector.Name(); dbType == "mysql" {
		sqls = []string{
			"ALTER TABLE sys_api MODIFY COLUMN handle VARCHAR(128) COMMENT 'handle';",
			"ALTER TABLE sys_api MODIFY COLUMN title VARCHAR(128) COMMENT '标题';",
			"ALTER TABLE sys_api MODIFY COLUMN path VARCHAR(128) COMMENT '地址';",
			"ALTER TABLE sys_api MODIFY COLUMN action VARCHAR(16) COMMENT '请求类型';",
			"ALTER TABLE sys_api MODIFY COLUMN type VARCHAR(16) COMMENT '接口类型';",
			"ALTER TABLE sys_api MODIFY COLUMN tag VARCHAR(128) COMMENT '标签';",
		}
	} else if dbType == "postgres" {
		sqls = []string{
			"ALTER TABLE sys_api ALTER COLUMN handle TYPE VARCHAR(128);",
			"ALTER TABLE sys_api ALTER COLUMN title TYPE VARCHAR(128);",
			"ALTER TABLE sys_api ALTER COLUMN path TYPE VARCHAR(128);",
			"ALTER TABLE sys_api ALTER COLUMN action TYPE VARCHAR(16);",
			"ALTER TABLE sys_api ALTER COLUMN type TYPE VARCHAR(16);",
			"ALTER TABLE sys_api ALTER COLUMN tag TYPE VARCHAR(128);",
		}
	} else if dbType == "sqlite" {
		// SQLite 不支持直接修改列类型，需要重建表
		return fmt.Errorf("SQLite does not support direct column type modification. Please handle this manually.")
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			return fmt.Errorf("alter sys_api fields error: %v", err)
		}
	}
	return nil
}

func alterSysConfigFields(tx *gorm.DB) error {
	var sqls []string

	if dbType := tx.Dialector.Name(); dbType == "mysql" {
		sqls = []string{
			"ALTER TABLE sys_config MODIFY COLUMN config_name VARCHAR(128) COMMENT 'ConfigName';",
			"ALTER TABLE sys_config MODIFY COLUMN config_key VARCHAR(128) COMMENT 'ConfigKey';",
			"ALTER TABLE sys_config MODIFY COLUMN config_value VARCHAR(255) COMMENT 'ConfigValue';",
			"ALTER TABLE sys_config MODIFY COLUMN config_type VARCHAR(64) COMMENT 'ConfigType';",
			"ALTER TABLE sys_config MODIFY COLUMN is_frontend VARCHAR(64) COMMENT '是否前台';",
			"ALTER TABLE sys_config MODIFY COLUMN remark VARCHAR(128) COMMENT 'Remark';",
		}
	} else if dbType == "postgres" {
		sqls = []string{
			"ALTER TABLE sys_config ALTER COLUMN config_name TYPE VARCHAR(128);",
			"ALTER TABLE sys_config ALTER COLUMN config_key TYPE VARCHAR(128);",
			"ALTER TABLE sys_config ALTER COLUMN config_value TYPE VARCHAR(255);",
			"ALTER TABLE sys_config ALTER COLUMN config_type TYPE VARCHAR(64);",
			"ALTER TABLE sys_config ALTER COLUMN is_frontend TYPE VARCHAR(64);",
			"ALTER TABLE sys_config ALTER COLUMN remark TYPE VARCHAR(128);",
		}
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			return fmt.Errorf("alter sys_config fields error: %v", err)
		}
	}
	return nil
}

func alterSysDeptFields(tx *gorm.DB) error {
	var sqls []string

	if dbType := tx.Dialector.Name(); dbType == "mysql" {
		sqls = []string{
			"ALTER TABLE sys_dept MODIFY COLUMN dept_path VARCHAR(255);",
			"ALTER TABLE sys_dept MODIFY COLUMN dept_name VARCHAR(128) COMMENT '部门名称';",
			"ALTER TABLE sys_dept MODIFY COLUMN leader VARCHAR(128) COMMENT '负责人';",
			"ALTER TABLE sys_dept MODIFY COLUMN phone VARCHAR(11) COMMENT '手机';",
			"ALTER TABLE sys_dept MODIFY COLUMN email VARCHAR(64) COMMENT '邮箱';",
		}
	} else if dbType == "postgres" {
		sqls = []string{
			"ALTER TABLE sys_dept ALTER COLUMN dept_path TYPE VARCHAR(255);",
			"ALTER TABLE sys_dept ALTER COLUMN dept_name TYPE VARCHAR(128);",
			"ALTER TABLE sys_dept ALTER COLUMN leader TYPE VARCHAR(128);",
			"ALTER TABLE sys_dept ALTER COLUMN phone TYPE VARCHAR(11);",
			"ALTER TABLE sys_dept ALTER COLUMN email TYPE VARCHAR(64);",
		}
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			return fmt.Errorf("alter sys_dept fields error: %v", err)
		}
	}
	return nil
}

func alterSysDictDataFields(tx *gorm.DB) error {
	var sqls []string

	if dbType := tx.Dialector.Name(); dbType == "mysql" {
		sqls = []string{
			"ALTER TABLE sys_dict_data MODIFY COLUMN dict_label VARCHAR(128) COMMENT 'DictLabel';",
			"ALTER TABLE sys_dict_data MODIFY COLUMN dict_value VARCHAR(255) COMMENT 'DictValue';",
			"ALTER TABLE sys_dict_data MODIFY COLUMN dict_type VARCHAR(64) COMMENT 'DictType';",
			"ALTER TABLE sys_dict_data MODIFY COLUMN css_class VARCHAR(128) COMMENT 'CssClass';",
			"ALTER TABLE sys_dict_data MODIFY COLUMN list_class VARCHAR(128) COMMENT 'ListClass';",
			"ALTER TABLE sys_dict_data MODIFY COLUMN is_default VARCHAR(8) COMMENT 'IsDefault';",
			"ALTER TABLE sys_dict_data MODIFY COLUMN default_value VARCHAR(8) COMMENT 'Default';",
			"ALTER TABLE sys_dict_data MODIFY COLUMN remark VARCHAR(255) COMMENT 'Remark';",
		}
	} else if dbType == "postgres" {
		sqls = []string{
			"ALTER TABLE sys_dict_data ALTER COLUMN dict_label TYPE VARCHAR(128);",
			"ALTER TABLE sys_dict_data ALTER COLUMN dict_value TYPE VARCHAR(255);",
			"ALTER TABLE sys_dict_data ALTER COLUMN dict_type TYPE VARCHAR(64);",
			"ALTER TABLE sys_dict_data ALTER COLUMN css_class TYPE VARCHAR(128);",
			"ALTER TABLE sys_dict_data ALTER COLUMN list_class TYPE VARCHAR(128);",
			"ALTER TABLE sys_dict_data ALTER COLUMN is_default TYPE VARCHAR(8);",
			"ALTER TABLE sys_dict_data ALTER COLUMN default_value TYPE VARCHAR(8);",
			"ALTER TABLE sys_dict_data ALTER COLUMN remark TYPE VARCHAR(255);",
		}
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			return fmt.Errorf("alter sys_dict_data fields error: %v", err)
		}
	}
	return nil
}

func alterSysDictTypeFields(tx *gorm.DB) error {
	var sqls []string

	if dbType := tx.Dialector.Name(); dbType == "mysql" {
		sqls = []string{
			"ALTER TABLE sys_dict_type MODIFY COLUMN dict_name VARCHAR(128) COMMENT 'DictName';",
			"ALTER TABLE sys_dict_type MODIFY COLUMN dict_type VARCHAR(128) COMMENT 'DictType';",
			"ALTER TABLE sys_dict_type MODIFY COLUMN remark VARCHAR(255) COMMENT 'Remark';",
		}
	} else if dbType == "postgres" {
		sqls = []string{
			"ALTER TABLE sys_dict_type ALTER COLUMN dict_name TYPE VARCHAR(128);",
			"ALTER TABLE sys_dict_type ALTER COLUMN dict_type TYPE VARCHAR(128);",
			"ALTER TABLE sys_dict_type ALTER COLUMN remark TYPE VARCHAR(255);",
		}
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			return fmt.Errorf("alter sys_dict_type fields error: %v", err)
		}
	}
	return nil
}

func alterSysLoginLogFields(tx *gorm.DB) error {
	var sqls []string

	if dbType := tx.Dialector.Name(); dbType == "mysql" {
		sqls = []string{
			"ALTER TABLE sys_login_log MODIFY COLUMN username VARCHAR(128) COMMENT '用户名';",
			"ALTER TABLE sys_login_log MODIFY COLUMN status TINYINT COMMENT '状态';",
			"ALTER TABLE sys_login_log MODIFY COLUMN ipaddr VARCHAR(255) COMMENT 'ip地址';",
			"ALTER TABLE sys_login_log MODIFY COLUMN login_location VARCHAR(255) COMMENT '归属地';",
			"ALTER TABLE sys_login_log MODIFY COLUMN browser VARCHAR(255) COMMENT '浏览器';",
			"ALTER TABLE sys_login_log MODIFY COLUMN os VARCHAR(255) COMMENT '系统';",
			"ALTER TABLE sys_login_log MODIFY COLUMN platform VARCHAR(255) COMMENT '固件';",
			"ALTER TABLE sys_login_log MODIFY COLUMN remark VARCHAR(255) COMMENT '备注';",
			"ALTER TABLE sys_login_log MODIFY COLUMN msg VARCHAR(255) COMMENT '信息';",
		}
	} else if dbType == "postgres" {
		sqls = []string{
			"ALTER TABLE sys_login_log ALTER COLUMN username TYPE VARCHAR(128);",
			"ALTER TABLE sys_login_log ALTER COLUMN status TYPE SMALLINT;",
			"ALTER TABLE sys_login_log ALTER COLUMN ipaddr TYPE VARCHAR(255);",
			"ALTER TABLE sys_login_log ALTER COLUMN login_location TYPE VARCHAR(255);",
			"ALTER TABLE sys_login_log ALTER COLUMN browser TYPE VARCHAR(255);",
			"ALTER TABLE sys_login_log ALTER COLUMN os TYPE VARCHAR(255);",
			"ALTER TABLE sys_login_log ALTER COLUMN platform TYPE VARCHAR(255);",
			"ALTER TABLE sys_login_log ALTER COLUMN remark TYPE VARCHAR(255);",
			"ALTER TABLE sys_login_log ALTER COLUMN msg TYPE VARCHAR(255);",
		}
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			return fmt.Errorf("alter sys_login_log fields error: %v", err)
		}
	}
	return nil
}

func alterSysMenuFields(tx *gorm.DB) error {
	var sqls []string

	if dbType := tx.Dialector.Name(); dbType == "mysql" {
		sqls = []string{
			"ALTER TABLE sys_menu MODIFY COLUMN menu_name VARCHAR(128);",
			"ALTER TABLE sys_menu MODIFY COLUMN title VARCHAR(128);",
			"ALTER TABLE sys_menu MODIFY COLUMN menu_type VARCHAR(10) COMMENT 'M: Menu/Directory   C: Component/Page  F: Function/Button';",
			"ALTER TABLE sys_menu MODIFY COLUMN menu_path VARCHAR(128);",
			"ALTER TABLE sys_menu MODIFY COLUMN path VARCHAR(128);",
			"ALTER TABLE sys_menu MODIFY COLUMN perm VARCHAR(255);",
			"ALTER TABLE sys_menu MODIFY COLUMN component VARCHAR(255);",
			"ALTER TABLE sys_menu MODIFY COLUMN icon VARCHAR(128);",
			"ALTER TABLE sys_menu MODIFY COLUMN external_link VARCHAR(255);",
			"ALTER TABLE sys_menu MODIFY COLUMN text_badge VARCHAR(10);",
			"ALTER TABLE sys_menu MODIFY COLUMN active_path VARCHAR(100);",
			"ALTER TABLE sys_menu MODIFY COLUMN status VARCHAR(1);",
			"ALTER TABLE sys_menu MODIFY COLUMN permission_code VARCHAR(128) COMMENT '关联权限表code';",
		}
	} else if dbType == "postgres" {
		sqls = []string{
			"ALTER TABLE sys_menu ALTER COLUMN menu_name TYPE VARCHAR(128);",
			"ALTER TABLE sys_menu ALTER COLUMN title TYPE VARCHAR(128);",
			"ALTER TABLE sys_menu ALTER COLUMN menu_type TYPE VARCHAR(10);",
			"ALTER TABLE sys_menu ALTER COLUMN menu_path TYPE VARCHAR(128);",
			"ALTER TABLE sys_menu ALTER COLUMN path TYPE VARCHAR(128);",
			"ALTER TABLE sys_menu ALTER COLUMN perm TYPE VARCHAR(255);",
			"ALTER TABLE sys_menu ALTER COLUMN component TYPE VARCHAR(255);",
			"ALTER TABLE sys_menu ALTER COLUMN icon TYPE VARCHAR(128);",
			"ALTER TABLE sys_menu ALTER COLUMN external_link TYPE VARCHAR(255);",
			"ALTER TABLE sys_menu ALTER COLUMN text_badge TYPE VARCHAR(10);",
			"ALTER TABLE sys_menu ALTER COLUMN active_path TYPE VARCHAR(100);",
			"ALTER TABLE sys_menu ALTER COLUMN status TYPE VARCHAR(1);",
			"ALTER TABLE sys_menu ALTER COLUMN permission_code TYPE VARCHAR(128);",
		}
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			return fmt.Errorf("alter sys_menu fields error: %v", err)
		}
	}
	return nil
}

func alterSysOperaLogFields(tx *gorm.DB) error {
	var sqls []string

	if dbType := tx.Dialector.Name(); dbType == "mysql" {
		sqls = []string{
			"ALTER TABLE sys_opera_log MODIFY COLUMN title VARCHAR(255) COMMENT '操作模块';",
			"ALTER TABLE sys_opera_log MODIFY COLUMN business_type VARCHAR(128) COMMENT '操作类型';",
			"ALTER TABLE sys_opera_log MODIFY COLUMN business_types VARCHAR(128) COMMENT 'BusinessTypes';",
			"ALTER TABLE sys_opera_log MODIFY COLUMN method VARCHAR(128) COMMENT '函数';",
			"ALTER TABLE sys_opera_log MODIFY COLUMN request_method VARCHAR(128) COMMENT '请求方式 GET POST PUT DELETE';",
			"ALTER TABLE sys_opera_log MODIFY COLUMN operator_type VARCHAR(128) COMMENT '操作类型';",
			"ALTER TABLE sys_opera_log MODIFY COLUMN oper_name VARCHAR(128) COMMENT '操作者';",
			"ALTER TABLE sys_opera_log MODIFY COLUMN dept_name VARCHAR(128) COMMENT '部门名称';",
			"ALTER TABLE sys_opera_log MODIFY COLUMN oper_url VARCHAR(255) COMMENT '访问地址';",
			"ALTER TABLE sys_opera_log MODIFY COLUMN oper_ip VARCHAR(128) COMMENT '客户端ip';",
			"ALTER TABLE sys_opera_log MODIFY COLUMN oper_location VARCHAR(128) COMMENT '访问位置';",
			"ALTER TABLE sys_opera_log MODIFY COLUMN status TINYINT COMMENT '操作状态 1:正常 2:关闭';",
			"ALTER TABLE sys_opera_log MODIFY COLUMN remark VARCHAR(255) COMMENT '备注';",
			"ALTER TABLE sys_opera_log MODIFY COLUMN latency_time VARCHAR(128) COMMENT '耗时';",
			"ALTER TABLE sys_opera_log MODIFY COLUMN user_agent VARCHAR(255) COMMENT 'ua';",
		}
	} else if dbType == "postgres" {
		sqls = []string{
			"ALTER TABLE sys_opera_log ALTER COLUMN title TYPE VARCHAR(255);",
			"ALTER TABLE sys_opera_log ALTER COLUMN business_type TYPE VARCHAR(128);",
			"ALTER TABLE sys_opera_log ALTER COLUMN business_types TYPE VARCHAR(128);",
			"ALTER TABLE sys_opera_log ALTER COLUMN method TYPE VARCHAR(128);",
			"ALTER TABLE sys_opera_log ALTER COLUMN request_method TYPE VARCHAR(128);",
			"ALTER TABLE sys_opera_log ALTER COLUMN operator_type TYPE VARCHAR(128);",
			"ALTER TABLE sys_opera_log ALTER COLUMN oper_name TYPE VARCHAR(128);",
			"ALTER TABLE sys_opera_log ALTER COLUMN dept_name TYPE VARCHAR(128);",
			"ALTER TABLE sys_opera_log ALTER COLUMN oper_url TYPE VARCHAR(255);",
			"ALTER TABLE sys_opera_log ALTER COLUMN oper_ip TYPE VARCHAR(128);",
			"ALTER TABLE sys_opera_log ALTER COLUMN oper_location TYPE VARCHAR(128);",
			"ALTER TABLE sys_opera_log ALTER COLUMN status TYPE SMALLINT;",
			"ALTER TABLE sys_opera_log ALTER COLUMN remark TYPE VARCHAR(255);",
			"ALTER TABLE sys_opera_log ALTER COLUMN latency_time TYPE VARCHAR(128);",
			"ALTER TABLE sys_opera_log ALTER COLUMN user_agent TYPE VARCHAR(255);",
		}
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			return fmt.Errorf("alter sys_opera_log fields error: %v", err)
		}
	}
	return nil
}

func alterSysPermissionFields(tx *gorm.DB) error {
	var sqls []string

	if dbType := tx.Dialector.Name(); dbType == "mysql" {
		sqls = []string{
			"ALTER TABLE sys_permission MODIFY COLUMN code VARCHAR(128) COMMENT '权限唯一编码';",
			"ALTER TABLE sys_permission MODIFY COLUMN name VARCHAR(128) COMMENT '权限名称';",
			"ALTER TABLE sys_permission MODIFY COLUMN type VARCHAR(16) COMMENT '权限类型(menu/button/api/page)';",
			"ALTER TABLE sys_permission MODIFY COLUMN remark VARCHAR(255) COMMENT '备注说明';",
		}
	} else if dbType == "postgres" {
		sqls = []string{
			"ALTER TABLE sys_permission ALTER COLUMN code TYPE VARCHAR(128);",
			"ALTER TABLE sys_permission ALTER COLUMN name TYPE VARCHAR(128);",
			"ALTER TABLE sys_permission ALTER COLUMN type TYPE VARCHAR(16);",
			"ALTER TABLE sys_permission ALTER COLUMN remark TYPE VARCHAR(255);",
		}
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			return fmt.Errorf("alter sys_permission fields error: %v", err)
		}
	}
	return nil
}

func alterSysPermissionApiFields(tx *gorm.DB) error {
	var sqls []string

	if dbType := tx.Dialector.Name(); dbType == "mysql" {
		sqls = []string{
			"ALTER TABLE sys_permission_api MODIFY COLUMN permission_id INT COMMENT '权限ID';",
			"ALTER TABLE sys_permission_api MODIFY COLUMN api_id INT COMMENT 'API ID';",
		}
	} else if dbType == "postgres" {
		sqls = []string{
			"ALTER TABLE sys_permission_api ALTER COLUMN permission_id TYPE INTEGER;",
			"ALTER TABLE sys_permission_api ALTER COLUMN api_id TYPE INTEGER;",
		}
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			return fmt.Errorf("alter sys_permission_api fields error: %v", err)
		}
	}
	return nil
}

func alterSysPostFields(tx *gorm.DB) error {
	var sqls []string

	if dbType := tx.Dialector.Name(); dbType == "mysql" {
		sqls = []string{
			"ALTER TABLE sys_post MODIFY COLUMN post_name VARCHAR(128) COMMENT '岗位名称';",
			"ALTER TABLE sys_post MODIFY COLUMN post_code VARCHAR(128) COMMENT '岗位代码';",
			"ALTER TABLE sys_post MODIFY COLUMN remark VARCHAR(255) COMMENT '描述';",
		}
	} else if dbType == "postgres" {
		sqls = []string{
			"ALTER TABLE sys_post ALTER COLUMN post_name TYPE VARCHAR(128);",
			"ALTER TABLE sys_post ALTER COLUMN post_code TYPE VARCHAR(128);",
			"ALTER TABLE sys_post ALTER COLUMN remark TYPE VARCHAR(255);",
		}
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			return fmt.Errorf("alter sys_post fields error: %v", err)
		}
	}
	return nil
}

func alterSysRoleFields(tx *gorm.DB) error {
	var sqls []string

	if dbType := tx.Dialector.Name(); dbType == "mysql" {
		sqls = []string{
			"ALTER TABLE sys_role MODIFY COLUMN role_name VARCHAR(128) COMMENT '角色名称';",
			"ALTER TABLE sys_role MODIFY COLUMN status TINYINT COMMENT '状态 1禁用 2正常';",
			"ALTER TABLE sys_role MODIFY COLUMN role_key VARCHAR(128) COMMENT '角色代码';",
			"ALTER TABLE sys_role MODIFY COLUMN flag VARCHAR(128);",
			"ALTER TABLE sys_role MODIFY COLUMN remark VARCHAR(255) COMMENT '备注';",
			"ALTER TABLE sys_role MODIFY COLUMN data_scope VARCHAR(128) COMMENT '数据范围';",
		}
	} else if dbType == "postgres" {
		sqls = []string{
			"ALTER TABLE sys_role ALTER COLUMN role_name TYPE VARCHAR(128);",
			"ALTER TABLE sys_role ALTER COLUMN status TYPE SMALLINT;",
			"ALTER TABLE sys_role ALTER COLUMN role_key TYPE VARCHAR(128);",
			"ALTER TABLE sys_role ALTER COLUMN flag TYPE VARCHAR(128);",
			"ALTER TABLE sys_role ALTER COLUMN remark TYPE VARCHAR(255);",
			"ALTER TABLE sys_role ALTER COLUMN data_scope TYPE VARCHAR(128);",
		}
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			return fmt.Errorf("alter sys_role fields error: %v", err)
		}
	}
	return nil
}

func alterSysRoleDeptFields(tx *gorm.DB) error {
	var sqls []string

	if dbType := tx.Dialector.Name(); dbType == "mysql" {
		sqls = []string{
			"ALTER TABLE sys_role_dept MODIFY COLUMN role_id INT COMMENT '角色编码';",
			"ALTER TABLE sys_role_dept MODIFY COLUMN dept_id INT COMMENT '部门编码';",
		}
	} else if dbType == "postgres" {
		sqls = []string{
			"ALTER TABLE sys_role_dept ALTER COLUMN role_id TYPE INTEGER;",
			"ALTER TABLE sys_role_dept ALTER COLUMN dept_id TYPE INTEGER;",
		}
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			return fmt.Errorf("alter sys_role_dept fields error: %v", err)
		}
	}
	return nil
}

func alterSysRoleMenuFields(tx *gorm.DB) error {
	var sqls []string

	if dbType := tx.Dialector.Name(); dbType == "mysql" {
		sqls = []string{
			"ALTER TABLE sys_role_menu MODIFY COLUMN role_id INT COMMENT '角色编码';",
			"ALTER TABLE sys_role_menu MODIFY COLUMN menu_id INT COMMENT '菜单编码';",
		}
	} else if dbType == "postgres" {
		sqls = []string{
			"ALTER TABLE sys_role_menu ALTER COLUMN role_id TYPE INTEGER;",
			"ALTER TABLE sys_role_menu ALTER COLUMN menu_id TYPE INTEGER;",
		}
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			return fmt.Errorf("alter sys_role_menu fields error: %v", err)
		}
	}
	return nil
}

func alterSysRolePermissionFields(tx *gorm.DB) error {
	var sqls []string

	if dbType := tx.Dialector.Name(); dbType == "mysql" {
		sqls = []string{
			"ALTER TABLE sys_role_permission MODIFY COLUMN role_id INT COMMENT '角色ID';",
			"ALTER TABLE sys_role_permission MODIFY COLUMN permission_id INT COMMENT '权限ID';",
		}
	} else if dbType == "postgres" {
		sqls = []string{
			"ALTER TABLE sys_role_permission ALTER COLUMN role_id TYPE INTEGER;",
			"ALTER TABLE sys_role_permission ALTER COLUMN permission_id TYPE INTEGER;",
		}
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			return fmt.Errorf("alter sys_role_permission fields error: %v", err)
		}
	}
	return nil
}

func alterSysUserFields(tx *gorm.DB) error {
	var sqls []string

	if dbType := tx.Dialector.Name(); dbType == "mysql" {
		sqls = []string{
			"ALTER TABLE sys_user MODIFY COLUMN uuid VARCHAR(255) COMMENT 'UUID';",
			"ALTER TABLE sys_user MODIFY COLUMN username VARCHAR(64) COMMENT '用户名';",
			"ALTER TABLE sys_user MODIFY COLUMN password VARCHAR(128) COMMENT '密码';",
			"ALTER TABLE sys_user MODIFY COLUMN nick_name VARCHAR(128) COMMENT '昵称';",
			"ALTER TABLE sys_user MODIFY COLUMN phone VARCHAR(11) COMMENT '手机号';",
			"ALTER TABLE sys_user MODIFY COLUMN salt VARCHAR(255) COMMENT '加盐';",
			"ALTER TABLE sys_user MODIFY COLUMN avatar VARCHAR(255) COMMENT '头像';",
			"ALTER TABLE sys_user MODIFY COLUMN sex VARCHAR(255) COMMENT '性别';",
			"ALTER TABLE sys_user MODIFY COLUMN email VARCHAR(128) COMMENT '邮箱';",
			"ALTER TABLE sys_user MODIFY COLUMN remark VARCHAR(255) COMMENT '备注';",
			"ALTER TABLE sys_user MODIFY COLUMN status TINYINT COMMENT '状态';",
		}
	} else if dbType == "postgres" {
		sqls = []string{
			"ALTER TABLE sys_user ALTER COLUMN uuid TYPE VARCHAR(255);",
			"ALTER TABLE sys_user ALTER COLUMN username TYPE VARCHAR(64);",
			"ALTER TABLE sys_user ALTER COLUMN password TYPE VARCHAR(128);",
			"ALTER TABLE sys_user ALTER COLUMN nick_name TYPE VARCHAR(128);",
			"ALTER TABLE sys_user ALTER COLUMN phone TYPE VARCHAR(11);",
			"ALTER TABLE sys_user ALTER COLUMN salt TYPE VARCHAR(255);",
			"ALTER TABLE sys_user ALTER COLUMN avatar TYPE VARCHAR(255);",
			"ALTER TABLE sys_user ALTER COLUMN sex TYPE VARCHAR(255);",
			"ALTER TABLE sys_user ALTER COLUMN email TYPE VARCHAR(128);",
			"ALTER TABLE sys_user ALTER COLUMN remark TYPE VARCHAR(255);",
			"ALTER TABLE sys_user ALTER COLUMN status TYPE SMALLINT;",
		}
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			return fmt.Errorf("alter sys_user fields error: %v", err)
		}
	}
	return nil
}

func alterSysUserRoleFields(tx *gorm.DB) error {
	var sqls []string

	if dbType := tx.Dialector.Name(); dbType == "mysql" {
		sqls = []string{
			"ALTER TABLE sys_user_role MODIFY COLUMN user_id INT COMMENT '用户ID';",
			"ALTER TABLE sys_user_role MODIFY COLUMN role_id INT COMMENT '角色ID';",
		}
	} else if dbType == "postgres" {
		sqls = []string{
			"ALTER TABLE sys_user_role ALTER COLUMN user_id TYPE INTEGER;",
			"ALTER TABLE sys_user_role ALTER COLUMN role_id TYPE INTEGER;",
		}
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			return fmt.Errorf("alter sys_user_role fields error: %v", err)
		}
	}
	return nil
}
