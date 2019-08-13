package lib

import (
	"expertExcel/modules"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type Excel struct {
	Path string
	Sheet string
}

//读取 excel 转换成 expert 模型
func (e *Excel) Read() []modules.AhgEcmsFfExpert {
	f, err := excelize.OpenFile(e.Path)
	if err != nil {
		panic(err)
	}
	rows, err := f.GetRows(e.Sheet)
	if err != nil {
		panic(err)
	}
	experts := make([]modules.AhgEcmsFfExpert, 0)
	for _, row := range rows {
		expert := new(modules.AhgEcmsFfExpert)
		for i, colCell := range row {
			switch i {
			case 0:
				expert.Title = colCell
			case 1:
				expert.Head = colCell
			case 2:
				expert.Post = colCell
			case 3:
				expert.Type = colCell
			case 4:
				expert.ExpertTitle = colCell
			case 5:
				expert.Tag = colCell
			case 6:
				expert.Brief = colCell
			}
		}
		experts = append(experts, *expert)
	}
	return experts
}