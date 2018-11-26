package lightGoExcel

import (
	"log"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type LightExecl struct {
	Xlsx     *excelize.File
	hasTitle bool
	Sheet    string
}

//初始化
func (excel *LightExecl) Init() {
	excel.Xlsx = excelize.NewFile()
	excel.Sheet = "Sheet1"
}

//添加标题
func (excel *LightExecl) AddTitle(title []string) {
	excel.hasTitle = true
	tdNum := 1
	for _, val := range title {
		excel.Xlsx.SetCellValue(excel.Sheet, TdIndex(tdNum)+"1", val)
		tdNum++
	}
}

//生成excel表
func (excel *LightExecl) SaveFile(filename string, data interface{}) {
	var trNum, tdNum int
	//如果有第一行标题
	if excel.hasTitle {
		trNum++
	}
	switch info := data.(type) {
	case map[string]string:
		for _, v1 := range info {
			tdNum = 0
			trNum++
			for _, v2 := range v1 {
				tdNum++
				excel.Xlsx.SetCellValue(excel.Sheet, TdIndex(tdNum)+strconv.Itoa(trNum), v2)
			}

		}
	case [][]string:
		for _, v1 := range info {
			tdNum = 0
			trNum++
			for _, v2 := range v1 {
				tdNum++
				excel.Xlsx.SetCellValue(excel.Sheet, TdIndex(tdNum)+strconv.Itoa(trNum), v2)
			}
		}
	default:
		log.Fatal("不支持的数据类型")
	}
	err := excel.Xlsx.SaveAs(filename)
	if err != nil {
		log.Fatal(err)
	}
}
