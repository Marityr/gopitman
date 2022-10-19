package schemes

import (
	"time"

	"gorm.io/datatypes"
)

type (
	//Пользователь
	User struct {
		Id          int            `json:"id,omitempty" gorm:"primary_key"`
		Password    string         `json:"password,omitempty"`
		LastLogin   time.Time      `json:"lastlogin,omitempty"`
		IsSuperuser bool           `json:"is_superuser"`
		UserName    string         `json:"username,omitempty"`
		LastName    string         `json:"lastname,omitempty"`
		FirstName   string         `json:"firstname,omitempty"`
		Email       string         `json:"email,omitempty"`
		IsStaff     bool           `json:"is_staff"`
		IsActive    bool           `json:"is_active"`
		DateJoined  datatypes.Date `json:"date_joined,omitempty" gorm:"default:null;"`
		Group       Group          `json:"group,omitempty" gorm:"foreignkey:Group"`
	}

	// Группа
	Group struct {
		Id          int           `json:"id,omitempty" gorm:"primary_key"`
		Name        string        `json:"name,omitempty"`
		Permissions []Permissions `json:"permission,omitempty" gorm:"many2many:group_permissions"`
	}

	// Доступы
	Permissions struct {
		Id     int      `json:"id,omitempty" gorm:"primary_key"`
		Mark   string   `json:"mark,omitempty"`
		Group  []Group  `json:"group,omitempty" gorm:"many2many:group_permissions"`
		Models []Models `json:"models,omitempty" gorm:"many2many:permissions_models"`
	}

	// Структуры
	Models struct {
		Id          int           `json:"id,omitempty" gorm:"primary_key"`
		Name        string        `json:"name,omitempty"`
		Permissions []Permissions `json:"permission,omitempty" gorm:"many2many:permissions_models"`
	}

	// Данные пользователя
	UserData struct {
		Id      int  `json:"id,omitempty" gorm:"primary_key"`
		Users   User `json:"user,omitempty" gorm:"foreignkey:Users"`
		Count   int  `json:"count,omitempty"`
		Payment bool `json:"payment,omitempty"`
	}

	Login struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
)
