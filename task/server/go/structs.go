package openapi

import "time"

type User struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	Username     string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	Balance      int    `gorm:"default:1000;check:balance >= 0"`
}

type Inventory struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	UserID   int    `gorm:"not null"`
	ItemType string `gorm:"not null"`
	Quantity int32  `gorm:"default:1;check:quantity >= 0"`
	User     User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type CoinHistory struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	FromUserID int       `gorm:"not null"`
	ToUserID   int       `gorm:"not null"`
	Amount     int       `gorm:"check:amount > 0"`
	Timestamp  time.Time `gorm:"autoCreateTime"`
}

type Products struct {
	ID    int    `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"not null; unique"`
	Price int    `gorm:"check:price > 0"`
}
