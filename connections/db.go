package connections

import (
	"log"

	"github.com/steelthedev/supra-paints/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(url string) *gorm.DB {
	DB, err := gorm.Open(postgres.Open(url), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              false,
	})

	if err != nil {
		log.Fatalf("Error: %s", err)
		panic(err)
	}

	DB.AutoMigrate(data.User{})

	return DB
}
