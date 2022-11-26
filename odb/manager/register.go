package manager

import (
	"database/sql"
	"fmt"
)

func (m *Manager) RegisterDB(conf ODBConf) error {
	if m == nil {
		return fmt.Errorf("please first init manager")
	}
	for _, conf := range conf.DBConfs {
		db, err := sql.Open(conf.Driver, conf.Dsn)
		if err != nil {
			return err
		}
		if conf.MaxLifetime > 0 {
			db.SetConnMaxLifetime(conf.MaxLifetime)
		}
		if conf.MaxIdleConns > 0 {
			db.SetMaxIdleConns(conf.MaxIdleConns)
		}
		if conf.MaxOpenConns > 0 {
			db.SetMaxOpenConns(conf.MaxOpenConns)
		}
		m.addDB(conf.DBName, conf.IsMaster, db)
	}
	return nil
}
