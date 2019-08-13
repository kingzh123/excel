package modules

import (
	"feifancron/db"
)

type AhgEcmsFfExpertData struct {
	Id         int    `orm:"id" json:"id"`
	Classid    int    `orm:"classid" json:"classid"`
	Keyid      string `orm:"keyid" json:"keyid"`
	Dokey      int    `orm:"dokey" json:"dokey"`
	Newstempid int    `orm:"newstempid" json:"newstempid"`
	Closepl    int    `orm:"closepl" json:"closepl"`
	Haveaddfen int    `orm:"haveaddfen" json:"haveaddfen"`
	Infotags   string `orm:"infotags" json:"infotags"`
}

func (*AhgEcmsFfExpertData) GetTable() string {
	return "ahg_ecms_ff_expert_data_1"
}

func (e *AhgEcmsFfExpertData) Insert(classId int, id int64) error {
	sql := "INSERT INTO " + e.GetTable() + " (id,classid) VALUES (?,?)"
	stmt, err := Dbs.GetDb().Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id,classId)
	if err != nil {
		return err
	}
	return nil
}
