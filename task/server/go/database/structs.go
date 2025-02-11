package database

import "time"

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique;not null"`
	Balance  int    `gorm:"default:0;check:balance >= 0"`
}

type Inventory struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	UserID   int    `gorm:"not null"`
	ItemType string `gorm:"not null"`
	Quantity int    `gorm:"default:1;check:quantity >= 0"`
	User     User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type CoinHistory struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	FromUserID int       `gorm:"not null"`
	ToUserID   int       `gorm:"not null"`
	Amount     int       `gorm:"check:amount > 0"`
	Timestamp  time.Time `gorm:"autoCreateTime"`
}
