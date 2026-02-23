package database

import (
	"kura/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SetupInitialUserData(userID uuid.UUID) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		for _, catInput := range model.DefaultCategories {
			category := model.Category{
				UserID:        userID,
				CategoryInput: catInput,
			}

			if err := DB.Create(&category).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
