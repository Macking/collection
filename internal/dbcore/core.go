package dbcore

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	globalDB     *gorm.DB
	globalConfig *DBConfig
	injectors    []func(db *gorm.DB)
)

type ctxTransactionKey struct{}

func Connect(cfg *DBConfig) {
	cfg = defaultDbConfig(cfg)
	globalConfig = cfg

	dsn := fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DSN,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Mysql connect open error ", err)
	}
	globalDB = db
}

func GetDB(ctx context.Context) *gorm.DB {
	iFace := ctx.Value(ctxTransactionKey{})
	if iFace != nil {
		tx, ok := iFace.(*gorm.DB)
		if !ok {
			return nil
		}
		return tx
	}
	return globalDB.WithContext(ctx)
}

func GetDBConfig() DBConfig {
	return *globalConfig
}
