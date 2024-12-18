package schoolwscomp

import (
	"fmt"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

func GetAllProductInfo(
	schoolId int,
	conn *websocket.Conn) []globalObject.StoreProducts {

	schoolProducts := []globalObject.StoreProducts{}
	schoolProductstmp := globalObject.StoreProducts{}

	// sql connection

	// query
	getManagerQuery := "SELECT * FROM store_products where school_id = ?;"
	// sql connection
	connObj := sql.InitConnection()
	// query manager info
	rows, sqlError := connObj.Connections.Query(getManagerQuery, schoolId)

	if sqlError != nil {
		libErrors.ReturnError(conn)
	}

	defer rows.Close()
	defer connObj.Connections.Close()

	for rows.Next() {
		err := rows.Scan(
			&schoolProductstmp.StoreProductsId,
			&schoolProductstmp.StoreProductName,
			&schoolProductstmp.StoreProductPath,
			&schoolProductstmp.StoreProductPrice,
			&schoolProductstmp.StoreProductImagePath,
			&schoolProductstmp.SchoolId,
			&schoolProductstmp.ProductListId)
		if err != nil {
      fmt.Println("bus error")
			libErrors.ReturnError(conn)
		}

		schoolProducts = append(schoolProducts, globalObject.StoreProducts{
			schoolProductstmp.StoreProductsId,
			schoolProductstmp.StoreProductName,
			schoolProductstmp.StoreProductPath,
			schoolProductstmp.StoreProductPrice,
			schoolProductstmp.StoreProductImagePath,
			schoolProductstmp.SchoolId,
			schoolProductstmp.ProductListId})

	}
	return schoolProducts
}
