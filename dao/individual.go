package dao

import (
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
	var sql strings.Builder
	var arg []interface{}
	sql.WriteString("update information set")
	if i.Nickname != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" nickname=?")
		arg = append(arg, i.Nickname)
	}
	if i.Gender != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" gender=?")
		arg = append(arg, i.Gender)
	}
	if i.PhoneNum != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" phoneNum=?")
		arg = append(arg, i.PhoneNum)
	}
	if i.Email != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" email=?")
		arg = append(arg, i.Email)
	}
	if i.Year != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" year=?")
		arg = append(arg, i.Year)
	}
	if i.Month != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" month=?")
		arg = append(arg, i.Month)
	}
	if i.Day != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" day=?")
		arg = append(arg, i.Day)
	}
	if i.Avatar != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" avatar=?")
		arg = append(arg, i.Avatar)
	}
	sql.WriteString(" where user_id =?")
	arg = append(arg, i.UserId)
	fmt.Println(sql.String())
	_, err = DB.Exec(sql.String(), arg...)
	return
}
