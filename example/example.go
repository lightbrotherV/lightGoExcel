package main

import "github.com/lightGoExcel"

func main() {
	e := lightGoExcel.LightExecl{}
	e.Init()
	e.AddTitle([]string{"t1", "t2", "t3"})
	e.SaveFile("test.xlsx", [][]string{
		{"1", "2"},
		{"2", "3"},
		{"3", "4"},
	})
}
