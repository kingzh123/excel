package modules

import (
	"feifancron/db"
	"time"
)

type AhgEcmsFfExpertIndex struct {
	Id         int `orm:"id" json:"id"`
	Classid    int `orm:"classid" json:"classid"`
	Checked    int `orm:"checked" json:"checked"`
	Newstime   int `orm:"newstime" json:"newstime"`
	Truetime   int `orm:"truetime" json:"truetime"`
	Lastdotime int `orm:"lastdotime" json:"lastdotime"`
	Havehtml   int `orm:"havehtml" json:"havehtml"`
}

func (*AhgEcmsFfExpertIndex) GetTable() string {
	return "ahg_ecms_ff_expert_index"
}

//插入 ahg_ecms_ff_expert_index 数据库
func (e *AhgEcmsFfExpertIndex) Insert(classId int) (int64, error) {
	sql := "INSERT INTO " + e.GetTable() + " (classid,truetime,lastdotime,newstime,checked,havehtml) VALUES (?,?,?,?,?,?)"
	stmt, err := Dbs.GetDb().Prepare(sql)
	if err != nil {
		return 0, err
	}
	time := time.Now().Unix()
	res, err := stmt.Exec(classId,time,time,time,1,0)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}