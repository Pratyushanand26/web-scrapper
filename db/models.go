package db

type User struct {
	Id       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Username string `gorm:"type:varchar(50);not null"`
	Password string `gorm:"type:varchar(100);not null"`
}
