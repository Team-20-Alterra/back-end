package seeder

import (
	"geinterra/models"
	"log"

	"gorm.io/gorm"
)

var bank = []models.Bank{
	models.Bank{
		Name: "Bank BRI",
		Code: "002",
		Logo: "https://i2.wp.com/febi.uinsaid.ac.id/wp-content/uploads/2020/11/Logo-BRI-Bank-Rakyat-Indonesia-PNG-Terbaru.png?ssl=1",
	},
	models.Bank{
		Name: "Bank BNI",
		Code: "009",
		Logo: "https://upload.wikimedia.org/wikipedia/id/thumb/5/55/BNI_logo.svg/1200px-BNI_logo.svg.png",
	},
}

var payment = []models.PaymentMethod{
	models.PaymentMethod{
		Body:   "1. Pilih <b>Transfer > Virtual Account Billing</b><br>2. Pilih <b>Rekening Debet ></b> Masukkan <b>nomor Virtual Account</b> 1209303718200192 pada <b>menu Input Baru</b><br>3. Tagihan yang harus dibayar akan muncul pada layar konfirmasi",
		BankID: 1,
	},
	models.PaymentMethod{
		Body:   "1. Pilih <b>Transfer > Virtual Account Billing</b><br>2. Pilih <b>Rekening Debet ></b> Masukkan <b>nomor Virtual Account</b> 1209303718200192 pada <b>menu Input Baru</b><br>3. Tagihan yang harus dibayar akan muncul pada layar konfirmasi",
		BankID: 2,
	},
}

// Load function
func Load(db *gorm.DB) {
	for h, _ := range bank {
		err1 := db.Debug().Model(&models.Bank{}).Create(&bank[h]).Error
		if err1 != nil {
			log.Fatalf("cannot seed bank table: %v", err1)
		}
	}

	for key, _ := range payment {
		err1 := db.Debug().Model(&models.PaymentMethod{}).Create(&payment[key]).Error
		if err1 != nil {
			log.Fatalf("cannot seed payment table: %v", err1)
		}
	}
}
