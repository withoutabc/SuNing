package dao

import "suning/model"

func SearchBalanceFromUsername(username string) (a model.Account, err error) {
	row := DB.QueryRow("select * from account where username=?", username)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&a.Username, &a.Balance)
	return
}
func InsertAccount(a model.Account) (err error) {
	_, err = DB.Exec("insert into account (username,account_balance) values(?,?)", a.Username, 0)
	return
}
func UpdateAccount(username string, accounted int) (err error) {
	_, err = DB.Exec("update account set account_balance=? where username=?", accounted, username)
	return
}
