package gopitman

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

type (
	// Пользователь
	Customer struct {
		ID           uuid.UUID `json:"id,omitempty" db:"id"`                       //Уникальный ключ
		ExternalID   string    `json:"external_id,omitempty" db:"external_id"`     //Внешний ключ
		BARCode      float64   `json:"bar_code,omitempty" db:"bar_code"`           //Штрих-код
		Activated    bool      `json:"activated,omitempty" db:"activated"`         //Доступна ли бонусная
		Created      time.Time `json:"created,omitempty" db:"created"`             //Врея создания
		UpdateAT     time.Time `json:"update_at,omitempty" db:"update_at"`         //Последнее изменение
		FirstName    string    `json:"first_name,omitempty" db:"first_name"`       //Имя
		SecondName   string    `json:"second_name,omitempty" db:"second_name"`     //Отчество
		LastName     string    `json:"last_name,omitempty" db:"last_name"`         //Фамилия
		Birthday     time.Time `json:"birthday,omitempty" db:"birthday"`           //Дата рождения
		SEX          string    `json:"sex,omitempty" db:"sex"`                     //Пол
		ReferrerCode string    `json:"referrer_code,omitempty" db:"referrer_code"` //Код
	}

	// Мета анные пользователя
	CustomerMeta struct {
		ID         int64     `json:"id,omitempty" db:"id"`                   //
		MetaTAG    string    `json:"meta_tag,omitempty" db:"meta_tag"`       //Тэг
		Value      string    `json:"value,omitempty" db:"value"`             //
		CustomerID uuid.UUID `json:"customer_id,omitempty" db:"customer_id"` //Ключ покупателя
	}

	// Рефералы
	CustomerReferral struct {
		ReferralID uuid.UUID `json:"referral_id,omitempty" gorm:"type:uuid"` //Ключ реферала
		ReferrerID uuid.UUID `json:"referrer_id,omitempty" gorm:"type:uuid"` //Ключ реферре
	}

	UpdateCustomer struct {
		Activate     *string `json:"activated,omitempty" db:"activated"`
		FirstName    *string `json:"first_name,omitempty" db:"first_name"`
		SecondName   *string `json:"second_name,omitempty" db:"second_name"`
		LastName     *string `json:"last_name,omitempty" db:"last_name"`
		Birthday     *string `json:"birthday,omitempty" db:"birthday"`
		SEX          *string `json:"sex,omitempty" db:"sex"`
		ReferrerCode *string `json:"referrer_code,omitempty" db:"referrer_code"`
	}

	UpdateCustomerMeta struct {
		Value *string `json:"value,omitempty" db:"value"`
	}
)
