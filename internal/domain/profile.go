package domain

type Profile struct {
	ID          string  `gorm:"primaryKey"`
	FullName    string  `gorm:"not null"`
	PhoneNumber string  `gorm:"not null; uniqueIndex"`
	Address     *string `gorm:"default:null"`
	AvatarUrl   *string `gorm:"default:null"`
	UserId      string  `gorm:"uniqueIndex; not null"`
}
