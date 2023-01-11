package dao

import (
	"fmt"
	"strings"
	"suning/model"
)

func SearchBalanceFromUid(uid string) (a model.Account, err error) {
	row := DB.QueryRow("select * from account where uid=?", uid)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&a.Username, &a.Balance, &a.Uid)
	return
}

func InsertAccount(a model.Account) (err error) {
	_, err = DB.Exec("insert into account (username,account_balance,uid) values(?,?,?)", a.Username, 0, a.Uid)
	return
}

func UpdateAccount(uid string, accounted int) (err error) {
	_, err = DB.Exec("update account set account_balance=? where uid=?", accounted, uid)
	return
}

func InsertInformation(username string, uid int) (err error) {
	_, err = DB.Exec("insert into information (username,uid) value (?,?)", username, uid)
	return
}

func SearchInformationByUid(uid string) (i model.Information, err error) {
	row := DB.QueryRow("select * from information where uid=?", uid)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&i.Username, &i.Nickname, &i.Gender, &i.PhoneNum, &i.Email, &i.Year, &i.Month, &i.Day, &i.Avatar, &i.Uid)
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
	sql.WriteString(" where uid =?")
	arg = append(arg, i.Uid)
	fmt.Println(sql.String())
	_, err = DB.Exec(sql.String(), arg...)
	return
}
