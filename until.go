package lightGoExcel

var alphabet []string = []string{"Z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

//翻转字符串
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//根据行数字获取execl对应的英文索引
func TdIndex(tdNum int) string {
	var res string
	for tdNum != 0 {
		if tdNum%26 == 0 {
			res += alphabet[0]
			tdNum = tdNum/26 - 1
			continue
		}
		res += alphabet[tdNum%26]
		tdNum /= 26
	}
	return reverse(res)
}
