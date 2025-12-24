package database

import (
	"log"

	"github.com/Alippp1/tes-golang/internal/models"
)

func AutoMigrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Supplier{},
		&models.Item{},
		&models.Purchasing{},
		&models.PurchasingDetail{},
	)

	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	log.Println("✅ All tables migrated")
}
