package pinyin

// VarName 中文转换拼音变量名
func VarName(name string, prefix string, firstUpper bool) string {
	py := NewPy(STYLE_NORMAL, NO_SEGMENT)

	ret := py.Convert(name)
	varName := prefix

	var newWord []rune

	for _, word := range ret {
		if len(word) <= 0 {
			continue
		}

		newWord = []rune(word)

		// 首字母大写
		if firstUpper {
			if 'a' <= newWord[0] && newWord[0] <= 'z' {
				newWord[0] -= 32
			}
		}

		varName += string(newWord)
	}

	return varName
}
