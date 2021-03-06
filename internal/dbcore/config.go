package dbcore

type DBConfig struct {
	DSN                string
	MaxIdleConnections int
	MaxOpenConnections int
	AutoMigrate        bool
}

func defaultDbConfig(cfg *DBConfig) *DBConfig {
	newCfg := *cfg

	if newCfg.MaxIdleConnections == 0 {
		newCfg.MaxIdleConnections = 10
	}

	if newCfg.MaxOpenConnections == 0 {
		newCfg.MaxOpenConnections = 20
	}

	newCfg.DSN = "root:root@tcp(192.168.0.1:3306)/dbname"

	return &newCfg
}
