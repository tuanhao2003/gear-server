package domain

type User struct {
	ID           string  `gorm:"primaryKey"`
	Username     string  `gorm:"uniqueIndex;not null"`
	Email        string  `gorm:"uniqueIndex;not null"`
	HashPassword string  `gorm:"not null" json:"-"`
	Profile      Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
