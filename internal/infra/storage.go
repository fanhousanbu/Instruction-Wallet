package infra

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbFile = "./mappings.db"

func getDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbFile))
	if err != nil {
		return nil, err
	}

	// 迁移 schema
	err = db.AutoMigrate(&PhoneNumberMapping{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

type PhoneNumberMapping struct {
	gorm.Model
	EncodedPhoneNumber string
	HashedPhoneNumber  string
}

// SavePhoneNumberMapping 保存手机号码映射关系
func SavePhoneNumberMapping(encodedPhoneNumber string, hashedPhoneNumber string) error {
	db, err := getDb()
	if err != nil {
		return nil
	}

	db.FirstOrCreate(&PhoneNumberMapping{
		EncodedPhoneNumber: encodedPhoneNumber,
		HashedPhoneNumber:  hashedPhoneNumber,
	})

	return nil
}

func GetHashedPhoneNumberByEncodedPhoneNumber(encodedPhoneNumber string) (*string, error) {
	db, err := getDb()
	if err != nil {
		return nil, err
	}

	var mapping PhoneNumberMapping
	if rows := db.First(&mapping, "EncodedPhoneNumber = ?", encodedPhoneNumber); rows.Error == nil {
		if rows.RowsAffected == 1 {
			return &mapping.HashedPhoneNumber, nil
		} else {
			return nil, nil
		}
	} else {
		return nil, rows.Error
	}
}
