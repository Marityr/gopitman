package gopitman

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}
