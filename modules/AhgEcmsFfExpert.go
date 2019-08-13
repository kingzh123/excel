package modules

import (
	"feifancron/db"
	"time"
)

type AhgEcmsFfExpert struct {
	Id            int    `orm:"id" json:"id"`
	Classid       int    `orm:"classid" json:"classid"`
	Ttid          int    `orm:"ttid" json:"ttid"`
	Onclick       int    `orm:"onclick" json:"onclick"`
	Plnum         int    `orm:"plnum" json:"plnum"`
	Totaldown     int    `orm:"totaldown" json:"totaldown"`
	Newspath      string `orm:"newspath" json:"newspath"`
	Filename      string `orm:"filename" json:"filename"`
	Userid        int    `orm:"userid" json:"userid"`
	Username      string `orm:"username" json:"username"`
	Firsttitle    int    `orm:"firsttitle" json:"firsttitle"`
	Isgood        int    `orm:"isgood" json:"isgood"`
	Ispic         int    `orm:"ispic" json:"ispic"`
	Istop         int    `orm:"istop" json:"istop"`
	Isqf          int    `orm:"isqf" json:"isqf"`
	Ismember      int    `orm:"ismember" json:"ismember"`
	Isurl         int    `orm:"isurl" json:"isurl"`
	Truetime      int    `orm:"truetime" json:"truetime"`
	Lastdotime    int    `orm:"lastdotime" json:"lastdotime"`
	Havehtml      int    `orm:"havehtml" json:"havehtml"`
	Groupid       int    `orm:"groupid" json:"groupid"`
	Userfen       int    `orm:"userfen" json:"userfen"`
	Titlefont     string `orm:"titlefont" json:"titlefont"`
	Titleurl      string `orm:"titleurl" json:"titleurl"`
	Stb           int    `orm:"stb" json:"stb"`
	Fstb          int    `orm:"fstb" json:"fstb"`
	Restb         int    `orm:"restb" json:"restb"`
	Keyboard      string `orm:"keyboard" json:"keyboard"`
	Title         string `orm:"title" json:"title"`
	Newstime      int    `orm:"newstime" json:"newstime"`
	Titlepic      string `orm:"titlepic" json:"titlepic"`
	Eckuid        int    `orm:"eckuid" json:"eckuid"`
	Post          string `orm:"post" json:"post"`
	Head          string `orm:"head" json:"head"`
	Type          string `orm:"type" json:"type"`
	IsIndex       string `orm:"is_index" json:"is_index"`
	IsExpertIndex string `orm:"is_expert_index" json:"is_expert_index"`
	Hot           string `orm:"hot" json:"hot"`
	Category      string `orm:"category" json:"category"`
	Tag           string `orm:"tag" json:"tag"`
	TradeNum      string `orm:"trade_num" json:"trade_num"`
	Praise        string `orm:"praise" json:"praise"`
	FollowNum     string `orm:"follow_num" json:"follow_num"`
	Brief         string `orm:"brief" json:"brief"`
	IndexSort     int    `orm:"index_sort" json:"index_sort"`
	Dynamic       string `orm:"dynamic" json:"dynamic"`
	ExpertTitle   string `orm:"expert_title" json:"expert_title"`
}

func (*AhgEcmsFfExpert) GetTable() string {
	return "ahg_ecms_ff_expert"
}

func (e *AhgEcmsFfExpert) Insert(classId int, id int64) error {
	sql := "INSERT INTO " + e.GetTable() + " (id,classid,title,expert_title,newstime,lastdotime,truetime,post,head,type,tag,brief,username) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := Dbs.GetDb().Prepare(sql)
	if err != nil {
		return err
	}
	time := time.Now().Unix()
	_, err = stmt.Exec(id, classId, e.Title, e.ExpertTitle, time, time, time, e.Post, e.Head, e.Type, e.Tag, e.Brief, "Auto")
	if err != nil {
		return err
	}
	return nil
}