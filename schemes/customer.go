package schemes

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

type (
	// Пользователь
	Customer struct {
		ID           uuid.UUID `json:"id,omitempty" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"` //Уникальный ключ
		ExternalID   string    `json:"external_id,omitempty"`                                                //Внешний ключ
		BARCode      float64   `json:"bar_code,omitempty"`                                                   //Штрих-код
		Activated    bool      `json:"activated,omitempty"`                                                  //Доступна ли бонусная
		Created      time.Time `json:"created,omitempty"`                                                    //Врея создания
		UpdateAT     time.Time `json:"update_at,omitempty"`                                                  //Последнее изменение
		FirstName    string    `json:"first_name,omitempty"`                                                 //Имя
		SecondName   string    `json:"second_name,omitempty"`                                                //Отчество
		LastName     string    `json:"last_name,omitempty"`                                                  //Фамилия
		Birthday     time.Time `json:"birthday,omitempty"`                                                   //Дата рождения
		SEX          string    `json:"sex,omitempty"`                                                        //Пол
		ReferrerCode string    `json:"referrer_code,omitempty"`                                              //Код
	}

	// Мета анные пользователя
	CustomerMeta struct {
		ID int64 `json:"id,omitempty" gorm:"primary_key"` //
		// TODO уникалья пара metatag+value
		MetaTAG string `json:"meta_tag,omitempty"` //Тэг
		Value   string `json:"value,omitempty"`    //
		//Customer   *Customer `json:"-" gorm:"foreignKey:ID;references:CustomerID"` //
		CustomerID uuid.UUID `json:"customer_id,omitempty" gorm:"type:uuid"` //Ключ покупателя
	}

	// Рефералы
	CustomerReferral struct {
		ReferralID uuid.UUID `json:"referral_id,omitempty" gorm:"type:uuid"` //Ключ реферала
		ReferrerID uuid.UUID `json:"referrer_id,omitempty" gorm:"type:uuid"` //Ключ реферре
	}
)

func (Customer) TableName() string {
	return "rest_customer"
}

func (CustomerMeta) TableName() string {
	return "rest_customermeta"
}

func (CustomerReferral) TableName() string {
	return "rest_customerreferral"
}
