package database

import (
	"log"
	"productsapi/app/models"

	"github.com/jmoiron/sqlx"
)


type ProductDb interface {
	Open() error
	Close() error
	SaveProduct(p *models.Product) error
	GetProducts() ([]*models.Product, error)
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error {

	mysql, err := sqlx.Open("mysql", mySqlConnStr)
	if err != nil {
		return err
	}

	log.Println("Connected to database ...")

	mysql.MustExec(createProductSchema)

	d.db = mysql
	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) SaveProduct(p *models.Product) error {

	res, err := d.db.Exec(insertProduct, p.Label, p.Price, p.Quantity)

	if err != nil {
		return err
	}

	res.LastInsertId()

	return err
}

func (d *DB) GetProducts() ([]*models.Product, error) {

	var products []*models.Product

	err := d.db.Select(&products, "SELECT * FROM PRODUCTS")
	if err != nil {
		return products, err
	}

	return products, nil

}