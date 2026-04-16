package mycasbin

import (
	"fmt"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"gorm.io/gorm"
)

const (
	SuperAdmin = "superadmin"
)

// Initialize the model from a string.
var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && (p.obj == "*" || keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`

var (
	enforcer *casbin.SyncedEnforcer
	once     sync.Once
)

func Setup(db *gorm.DB, _ string) *casbin.SyncedEnforcer {
	once.Do(func() {
		Apter, err := gormAdapter.NewAdapterByDBUseTableName(db, "", "sys_casbin_rule")
		if err != nil && err.Error() != "invalid DDL" {
			panic(err)
		}

		m, err := model.NewModelFromString(text)
		if err != nil {
			panic(err)
		}
		enforcer, err = casbin.NewSyncedEnforcer(m, Apter)
		if err != nil {
			panic(err)
		}

		enforcer.EnableLog(true)

		err = enforcer.LoadPolicy()
		if err != nil {
			panic(err)
		}

		ok, err := enforcer.HasPolicy(SuperAdmin, "*", "*")
		if err != nil {
			panic(err) // 立即终止程序
		}
		if !ok {
			_, err := enforcer.AddPolicy(SuperAdmin, "*", "*")
			if err != nil {
				panic(err)
			}
		}

	})

	return enforcer
}

// SyncUserRoleGroupingPolicies 根据 sys_user_role 和 sys_role 表，
// 确保每个用户在 Casbin 中都有正确的 grouping policy (g, user_X, roleKey)。
// 这样即使 sys_casbin_rule 表数据丢失，启动时也能自动恢复。
// 必须在数据库迁移（种子数据）完成之后调用。
func SyncUserRoleGroupingPolicies(db *gorm.DB) {
	l := logger.NewHelper(sdk.Runtime.GetLogger())

	if enforcer == nil {
		l.Warn("casbin syncUserRoleGroupingPolicies skipped: enforcer not initialized")
		return
	}

	// 先重新加载策略，确保内存与数据库一致。
	if err := enforcer.LoadPolicy(); err != nil {
		l.Errorf("casbin syncUserRoleGroupingPolicies LoadPolicy error: %v", err)
		return
	}

	type userRoleRow struct {
		UserID  int    `gorm:"column:user_id"`
		RoleKey string `gorm:"column:role_key"`
	}

	var rows []userRoleRow
	err := db.Table("sys_user_role").
		Select("sys_user_role.user_id, sys_role.role_key").
		Joins("LEFT JOIN sys_role ON sys_role.role_id = sys_user_role.role_id").
		Where("sys_role.role_key IS NOT NULL AND sys_role.role_key != ''").
		Find(&rows).Error
	if err != nil {
		l.Errorf("casbin syncUserRoleGroupingPolicies query error: %v", err)
		return
	}

	added := 0
	for _, row := range rows {
		sub := fmt.Sprintf("user_%d", row.UserID)
		has, err := enforcer.HasGroupingPolicy(sub, row.RoleKey)
		if err != nil {
			l.Errorf("casbin HasGroupingPolicy error: %v", err)
			continue
		}
		if !has {
			if _, err := enforcer.AddGroupingPolicy(sub, row.RoleKey); err != nil {
				l.Errorf("casbin AddGroupingPolicy(%s, %s) error: %v", sub, row.RoleKey, err)
			} else {
				added++
			}
		}
	}
	if added > 0 {
		l.Infof("casbin syncUserRoleGroupingPolicies: added %d grouping policies", added)
	}
}

func UpdateCallback(msg string) {
	l := logger.NewHelper(sdk.Runtime.GetLogger())
	l.Infof("casbin updateCallback msg: %v", msg)
	err := enforcer.LoadPolicy()
	if err != nil {
		l.Errorf("casbin LoadPolicy err: %v", err)
	}
}
