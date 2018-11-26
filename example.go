package main

import "lightGoExcel"

func main() {
	e := lightGoExcel.LightExecl{}
	e.Init()
	e.SaveFile("test.xlsx", [][]string{
		{"1", "2"},
		{"2", "3"},
		{"3", "4"},
	})
}
