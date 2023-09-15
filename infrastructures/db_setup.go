package infrastructures

import (
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	dbOnce sync.Once
)

func InitDBConnection(dbType DBType, dsn string) {
	dbOnce.Do(func() {
		// only postgres is implemented here, rest all will be in same way
		if dbType == POSTGRES {
			db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				log.Printf("Failed to connect to database: %v", err)
				panic(err)
			}
			DB = db
		}
	})
}
