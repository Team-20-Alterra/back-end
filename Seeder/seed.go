package Seeder

import (
	"geinterra/models"
	"log"

	"gorm.io/gorm"
)

var bank = []models.Bank{
	models.Bank{
		Name:    "Bank BRI",
		Code: "002",
		Logo:    "https://i2.wp.com/febi.uinsaid.ac.id/wp-content/uploads/2020/11/Logo-BRI-Bank-Rakyat-Indonesia-PNG-Terbaru.png?ssl=1",
	},
	models.Bank{
		Name:    "Bank BNI",
		Code: "009",
		Logo:    "https://upload.wikimedia.org/wikipedia/id/thumb/5/55/BNI_logo.svg/1200px-BNI_logo.svg.png",
	},
}

//Load function
func Load(db *gorm.DB) {
	for h, _ := range bank {
		err1 := db.Debug().Model(&models.Bank{}).Create(&bank[h]).Error
		if err1 != nil {
			log.Fatalf("cannot seed bank table: %v", err1)
		}
	}
}