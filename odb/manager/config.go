// @Title DB config
// @Description DB connection configuration message
// @Author omega
package manager

import "time"

// DBConf DB configuration
type DBConf struct {
	DBName       string
	IsMaster     bool   // main library
	Driver       string // driver
	Dsn          string
	MaxLifetime  time.Duration
	MaxIdleConns int
	MaxOpenConns int
}

// ODBConf multiple configuration
type ODBConf struct {
	DBConfs []DBConf
}
