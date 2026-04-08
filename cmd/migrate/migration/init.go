package migration

import (
	"log"
	"path/filepath"
	"sort"
	"sync"

	commonmodels "go-admin/common/models"

	"gorm.io/gorm"
)

var Migrate = &Migration{
	version: make(map[string]func(db *gorm.DB, version string) error),
}

type Migration struct {
	db      *gorm.DB
	version map[string]func(db *gorm.DB, version string) error
	mutex   sync.Mutex
}

func (e *Migration) GetDb() *gorm.DB {
	return e.db
}

func (e *Migration) SetDb(db *gorm.DB) {
	e.db = db
}

func (e *Migration) SetVersion(k string, f func(db *gorm.DB, version string) error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.version[k] = f
}

func (e *Migration) Migrate() {
	if err := e.repairLedgerForLegacyDB(); err != nil {
		log.Fatalln(err)
	}

	versions := make([]string, 0)
	for k := range e.version {
		versions = append(versions, k)
	}
	if !sort.StringsAreSorted(versions) {
		sort.Strings(versions)
	}
	var err error
	var count int64
	for _, v := range versions {
		err = e.db.Table("sys_migration").Where("version = ?", v).Count(&count).Error
		if err != nil {
			log.Fatalln(err)
		}
		if count > 0 {
			log.Println(count)
			count = 0
			continue
		}
		err = (e.version[v])(e.db.Debug(), v)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

const version159919 = "1599190683659"

// repairLedgerForLegacyDB 兼容：早期只有 AutoMigrate、未写 sys_migration，或首次建表后才有的版本表。
// 若库中已有 sys_user 等业务表，但缺少 159919 记录，则补登该版本，避免再次执行 InitDb / 全量建表迁移。
func (e *Migration) repairLedgerForLegacyDB() error {
	db := e.db
	if db == nil {
		return nil
	}
	var n int64
	if err := db.Table("sys_migration").Where("version = ?", version159919).Count(&n).Error; err != nil {
		return err
	}
	if n > 0 {
		return nil
	}
	if !db.Migrator().HasTable("sys_user") {
		return nil
	}
	rec := commonmodels.Migration{Version: version159919}
	if err := db.Create(&rec).Error; err != nil {
		if isDuplicateKeyError(err) {
			return nil
		}
		return err
	}
	log.Printf("migration repair: 已补登 sys_migration.version=%s（检测到已有业务表、无该版本记录）", version159919)
	return nil
}

func GetFilename(s string) string {
	s = filepath.Base(s)
	return s[:13]
}
