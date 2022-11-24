// @Title DB manager
// @Description DB manager proxy
// @Author omega
package manager

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

// Manager dn instance
type Manager struct {
	dbs map[string]map[string][]*sql.DB
}

// NewManager init manager
func NewManager() *Manager {
	m := new(Manager)
	m.dbs = make(map[string]map[string][]*sql.DB)
	return m
}

// addDB add db
func (m *Manager) addDB(gn string, im bool, db *sql.DB) {
	dr := "master"
	if !im {
		dr = "slave"
	}
	group, ok := m.dbs[gn]
	if !ok {
		group = make(map[string][]*sql.DB)
	}
	_, ok = group[dr]
	if ok {
		group[dr] = append(group[dr], db)
	} else {
		group[dr] = []*sql.DB{db}
	}
	m.dbs[gn] = group
}

// getDB get db
func (m *Manager) getDB(gn, dr string, names ...string) (*sql.DB, error) {
	if gn == "" || dr == "" {
		return nil, fmt.Errorf("gn or dr should not be empty")
	}
	dbs, ok := m.dbs[gn][dr]
	if ok {
		max := len(dbs)
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(max)
		return dbs[i], nil
	}
	return nil, fmt.Errorf("DataBase %s:%s not found", gn, dr)
}

//  getReadDB get slave db
func (m *Manager) getReadDB(gn string) (*sql.DB, error) {
	return m.getDB(gn, "slave")
}
