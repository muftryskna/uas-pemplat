// models/models.go
package models

type User struct {
	ID       int64   `gorm:"primaryKey" json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`
	Nama     string `json:"name"`
	Email    string `json:"email"`
}

type Product struct {
	ID         int64    `gorm:"primaryKey" json:"id"`
	NamaProduk string  `json:"nama_produk"`
	Harga      float64 `json:"harga"`
}
