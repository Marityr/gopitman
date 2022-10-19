package repository

import (
	"gorm.io/gorm"
)

func DropView(db *gorm.DB) {
	sql := "DROP VIEW public.customer_view"
	db.Exec(sql)
}

func CreateView(db *gorm.DB) {
	CustomerViewSql := `
	CREATE OR REPLACE VIEW public.customer_view
	AS
	SELECT row_number() OVER (ORDER BY rest_transaction.customer_id) AS id,
		rest_transaction.customer_id,
		sum(rest_transaction.cost) AS cost,
		sum(rest_transaction.bonuses) AS bonuses
	FROM rest_transaction
	WHERE rest_transaction.status = 1
	GROUP BY rest_transaction.customer_id
	ORDER BY rest_transaction.customer_id;
	`
	db.Exec(CustomerViewSql)
}

func AutoMigrate(db *gorm.DB) {
	// TODO убрать автомиграции от GORM просмотреть другие решения
	db.AutoMigrate()
	DropView(db)
	CreateView(db)
}
