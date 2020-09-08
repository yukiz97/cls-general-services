package lcservices

import (
	"github.com/yukiz97/cls-general-services/models"
	"github.com/yukiz97/utils/dbcon"
)

var strDBConnect string

//GetProductList get customer list by keyword
func GetProductList() []models.Product {
	listProduct := make([]models.Product, 0)

	db := dbcon.InitDBMySQL(strDBConnect)
	defer db.Close()

	selectQuery, err := db.Prepare("SELECT * FROM Product")
	if err != nil {
		panic(err.Error())
	}
	result, _ := selectQuery.Query()
	for result.Next() {
		modelProduct := models.Product{}

		result.Scan(&modelProduct.ID, &modelProduct.Code, &modelProduct.Name)
		listProduct = append(listProduct, modelProduct)
	}

	return listProduct
}

//InitLocalServices init value for database functions
func InitLocalServices(host string, userName string, password string, db string) {
	strDBConnect = dbcon.GetMySQLOpenConnectString(host, userName, password, db)
}
