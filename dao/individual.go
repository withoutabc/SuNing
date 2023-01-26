package dao

import (
	"database/sql"
	"fmt"
	"strings"
	"suning/model"
)

func SearchBalanceFromUserId(userId string) (a model.Account, err error) {
	row := DB.QueryRow("select * from account where user_id=?", userId)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&a.Username, &a.Balance, &a.UserId)
	return
}

func InsertAccount(a model.Account) (err error) {
	_, err = DB.Exec("insert into account (username,balance,user_id) values(?,?,?)", a.Username, 0, a.UserId)
	return
}

func UpdateAccount(userId string, accounted float64) (err error) {
	_, err = DB.Exec("update account set balance=? where user_id=?", accounted, userId)
	return
}

func DecreaseBalance(userId string, price float64) (err error) {
	_, err = DB.Exec("update account set balance=balance-? where user_id=?", price, userId)
	return
}

func InsertInformation(username string, userId int) (err error) {
	_, err = DB.Exec("insert into information (username,user_id) value (?,?)", username, userId)
	return
}

func SearchInformationByUserId(userId string) (i model.Information, err error) {
	row := DB.QueryRow("select * from information where user_id=?", userId)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&i.Username, &i.Nickname, &i.Gender, &i.PhoneNum, &i.Email, &i.Year, &i.Month, &i.Day, &i.Avatar, &i.UserId)
	return
}

func UpdateInformation(i model.Information) (err error) {
	var Sql strings.Builder
	var arg []interface{}
	Sql.WriteString("update information set")
	if i.Nickname != "" {
		if len(arg) > 0 {
			Sql.WriteString(",")
		}
		Sql.WriteString(" nickname=?")
		arg = append(arg, i.Nickname)
	}
	if i.Gender != "" {
		if len(arg) > 0 {
			Sql.WriteString(",")
		}
		Sql.WriteString(" gender=?")
		arg = append(arg, i.Gender)
	}
	if i.PhoneNum != "" {
		if len(arg) > 0 {
			Sql.WriteString(",")
		}
		Sql.WriteString(" phoneNum=?")
		arg = append(arg, i.PhoneNum)
	}
	if i.Email != "" {
		if len(arg) > 0 {
			Sql.WriteString(",")
		}
		Sql.WriteString(" email=?")
		arg = append(arg, i.Email)
	}
	if i.Year != "" {
		if len(arg) > 0 {
			Sql.WriteString(",")
		}
		Sql.WriteString(" year=?")
		arg = append(arg, i.Year)
	}
	if i.Month != "" {
		if len(arg) > 0 {
			Sql.WriteString(",")
		}
		Sql.WriteString(" month=?")
		arg = append(arg, i.Month)
	}
	if i.Day != "" {
		if len(arg) > 0 {
			Sql.WriteString(",")
		}
		Sql.WriteString(" day=?")
		arg = append(arg, i.Day)
	}
	if i.Avatar != "" {
		if len(arg) > 0 {
			Sql.WriteString(",")
		}
		Sql.WriteString(" avatar=?")
		arg = append(arg, i.Avatar)
	}
	Sql.WriteString(" where user_id =?")
	arg = append(arg, i.UserId)
	fmt.Println(Sql.String())
	_, err = DB.Exec(Sql.String(), arg...)
	return
}

func AddAddress(a model.Address) (err error) {
	_, err = DB.Exec("insert into address (user_id,recipient_name,recipient_phone,province,city,street_or_community) values (?,?,?,?,?,?)", a.UserId, a.RecipientName, a.RecipientName, a.Province, a.City, a.StateOrCommunity)
	return
}

func SearchAddress(userId string) (addresses []model.Address, err error) {
	var rows *sql.Rows
	rows, err = DB.Query("select * from address where user_id=?", userId)
	for rows.Next() {
		var address model.Address
		if err = rows.Scan(&address.AddressId, &address.UserId, &address.RecipientName, &address.RecipientPhone, &address.Province, &address.City, &address.StateOrCommunity); err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}
	return
}

func UpdateAddress(a model.Address) (err error) {
	_, err = DB.Exec("update address set recipient_name=?,recipient_phone=?,province=?,city=?,street_or_community=? where user_id=?", a.RecipientName, a.RecipientPhone, a.Province, a.City, a.StateOrCommunity, a.UserId)
	return
}

func DeleteAddress(addressId, userId string) (err error) {
	_, err = DB.Exec("delete * from address where user_id=? and address_id=?", userId, addressId)
	return
}
