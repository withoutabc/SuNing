package dao

import "database/sql"

func CartPay(payProducts []string) (totalPrice float64, err error) {
	var pay float64
	for i := 0; i < len(payProducts); i++ {
		var row *sql.Row
		row = DB.QueryRow("select price from cart where name=?", payProducts[i])
		err = row.Scan(&pay)
		if err != nil {
			return 0, err
		}
		totalPrice += pay
	}
	return
}
