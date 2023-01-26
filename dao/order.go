package dao

import "database/sql"

func Payment(userId string, payProducts []string) (totalPrice float64, err error) {
	var pay float64
	for i := 0; i < len(payProducts); i++ {
		var row *sql.Row
		row = DB.QueryRow("select price from cart where user_id=? and name=?", userId, payProducts[i])
		err = row.Scan(&pay)
		if err != nil {
			return 0, err
		}
		totalPrice += pay
	}
	return
}

func DeleteCartName(userId string, payProducts []string) (err error) {
	for i := 0; i < len(payProducts); i++ {
		_, err = DB.Exec("delete * from cart where user_id=? and name=? ", userId, payProducts[i])
		if err != nil {
			return err
		}
	}
	return
}
