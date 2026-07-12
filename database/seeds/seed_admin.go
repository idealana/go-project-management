package seeds

import (
	"log"

	"github.com/idealana/go-project-management/config"
	"github.com/idealana/go-project-management/models"
	"github.com/idealana/go-project-management/utils"

	"github.com/google/uuid"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("@Dmin_123#")

	admin := models.User{
		Name: "Super Admin",
		Email: "admin@example.com",
		Password: password,
		Role: "admin",
		PublicID: uuid.New(),
	}

	if err := config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Println("failed to seed admin", err)
	} else {
		log.Println("admin user seeded")
	}
}
