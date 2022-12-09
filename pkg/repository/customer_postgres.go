package repository

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Marityr/gopitman"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type CustomerPostgres struct {
	db *sqlx.DB
}

func NewCustomerPostgres(db *sqlx.DB) *CustomerPostgres {
	return &CustomerPostgres{db: db}
}

func (r *CustomerPostgres) Create(firstName, lastName, birthday, referrerCode, phone, email string) (string, error) {
	var datauuid []byte
	birthdayDate, err := time.Parse("02-01-2006", birthday)
	if err != nil && birthday != "" {
		log.Println(err)
		return "", err
	}

	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return "", err
	}

	createCustomer := fmt.Sprintf("INSERT INTO %s ( id, first_name, last_name, birthday, referrer_code, bar_code,activated,created,update_at) VALUES (uuid_generate_v4(), $1, $2, $3, $4, 12, true, current_timestamp, current_timestamp ) RETURNING id", "customer")
	row := tx.QueryRow(createCustomer, firstName, lastName, birthdayDate, referrerCode)
	if err := row.Scan(&datauuid); err != nil {
		tx.Rollback()
		log.Println(err)
		return "", err
	}

	uuidmeta := uuid.Must(uuid.FromString(string(datauuid)))
	if phone != "" {
		createCustomerMeta := fmt.Sprintf("INSERT INTO %s ( customer_id, meta_tag, value) VALUES ($1, $2, $3)", "customermeta")
		_, err = tx.Exec(createCustomerMeta, uuidmeta, "phone", phone)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return "", err
		}
	}
	if email != "" {
		createCustomerMeta := fmt.Sprintf("INSERT INTO %s ( customer_id, meta_tag, value) VALUES ($1, $2, $3)", "customermeta")
		_, err = tx.Exec(createCustomerMeta, uuidmeta, "email", email)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return "", err
		}
	}

	return string(datauuid), tx.Commit()
}

func (r *CustomerPostgres) GetAll(page, limit int) ([]gopitman.Customer, error) {
	var lists []gopitman.Customer

	query := "SELECT id, COALESCE(external_id, '') as external_id, bar_code, activated,created, update_at, COALESCE(second_name, '') as second_name, last_name, first_name, birthday, COALESCE(sex, '') sex, referrer_code FROM customer"
	err := r.db.Select(&lists, query)
	if err != nil {
		log.Println(err)
	}

	return lists, err
}

func (r *CustomerPostgres) GetById(id uuid.UUID) (gopitman.Customer, error) {
	var lists gopitman.Customer

	query := "SELECT id, COALESCE(external_id, '') as external_id, bar_code, activated,created, update_at, COALESCE(second_name, '') as second_name, last_name, first_name, birthday, COALESCE(sex, '') sex, referrer_code FROM customer WHERE id=$1"
	err := r.db.Get(&lists, query, id)
	if err != nil {
		log.Println(err)
	}
	return lists, err
}

func (r *CustomerPostgres) Delete(id uuid.UUID) error {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	query := "DELETE FROM customermeta WHERE customer_id = $1"
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	query = "DELETE FROM customer WHERE id = $1"
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	return tx.Commit()
}

func (r *CustomerPostgres) Update(id uuid.UUID, input gopitman.UpdateCustomer) error {
	setvalues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Activate != nil {
		setvalues = append(setvalues, fmt.Sprintf("activated=$%d", argId))
		if *input.Activate == "true" || *input.Activate == "false" {
			args = append(args, *input.Activate)
			argId++
		}

	}
	if input.FirstName != nil {
		setvalues = append(setvalues, fmt.Sprintf("first_name=$%d", argId))
		args = append(args, *input.FirstName)
		argId++
	}
	if input.SecondName != nil {
		setvalues = append(setvalues, fmt.Sprintf("second_name=$%d", argId))
		args = append(args, *input.SecondName)
		argId++
	}
	if input.LastName != nil {
		setvalues = append(setvalues, fmt.Sprintf("last_name=$%d", argId))
		args = append(args, *input.LastName)
		argId++
	}
	if input.Birthday != nil {
		setvalues = append(setvalues, fmt.Sprintf("birthday=$%d", argId))
		args = append(args, *input.Birthday)
		argId++
	}
	if input.SEX != nil {
		setvalues = append(setvalues, fmt.Sprintf("sex=$%d", argId))
		args = append(args, *input.SEX)
		argId++
	}
	if input.ReferrerCode != nil {
		setvalues = append(setvalues, fmt.Sprintf("referrer_code=$%d", argId))
		args = append(args, *input.ReferrerCode)
		argId++
	}

	setQuery := strings.Join(setvalues, ", ")

	query := fmt.Sprintf("UPDATE customer SET %s WHERE id='%s'", setQuery, id.String())

	_, err := r.db.Exec(query, args...)
	if err != nil {
		log.Println(err)
	}
	return err
}
