package main

import (
	"expertExcel/lib"
	"expertExcel/modules"
	"feifancron/db"
	"fmt"
)
//TODO 生产的 class id 和测试不一致，生产导入的时候需要确认 class id。 并且需要修改数据库的配置为正式的 mysql 账号
//TODO 改功能只一次性导入数据 不清除原有的数据
func main() {
	excel := new(lib.Excel)
	excel.Path = "./excel/expert.xlsx"
	excel.Sheet = "Sheet1"
	experts := excel.Read()
	classId := 124
	count := len(experts)
	fmt.Printf("%d 条数据正在自动导入中...\n", count)
	for k, v := range experts {
		index := new(modules.AhgEcmsFfExpertIndex)
		data := new(modules.AhgEcmsFfExpertData)
		c, _ := Dbs.GetDb().Begin()
		id, err := index.Insert(classId)
		if err != nil {
			c.Rollback()
			panic(err)
		}
		err = data.Insert(classId, id)
		if err != nil {
			c.Rollback()
			panic(err)
		}
		err = v.Insert(classId, id)
		if err != nil {
			c.Rollback()
			panic(err)
		}
		c.Commit()
		fmt.Printf("第 %d 条数据导入完成！\n", k+1)
	}
	fmt.Println("导入完成！")
}
