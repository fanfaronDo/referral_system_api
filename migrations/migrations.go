package migrations

import (
	"gorm.io/gorm"
)

type Migrator struct {
	db *gorm.DB
}

func NewMigrator(db *gorm.DB) *Migrator {
	return &Migrator{db: db}
}

func (m *Migrator) MigrateUp(items ...interface{}) error {
	for _, item := range items {
		if !m.db.Migrator().HasTable(item) {
			if err := m.db.Migrator().CreateTable(item); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *Migrator) MigrateDown(items ...interface{}) error {
	for _, item := range items {
		if m.db.Migrator().HasTable(item) {
			if err := m.db.Migrator().DropTable(item); err != nil {
				return err
			}
		}
	}

	return nil
}
