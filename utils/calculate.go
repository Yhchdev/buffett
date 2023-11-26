package utils

// 经营，投资，筹资
func CorporatePortrait(operation, investment, financing float64) string {

	overall := operation + investment + financing

	if operation > 0 && investment < 0 && financing < 0 && overall > 0 {
		return "8-奶牛"
	}

	if operation > 0 && investment < 0 && financing < 0 && overall < 0 {
		return "7-穷奶牛"
	}

	if operation > 0 && investment > 0 && financing < 0 && overall > 0 {
		return "6-老母鸡"
	}

	if operation > 0 && investment > 0 && financing < 0 && overall < 0 {
		return "5-穷母鸡"
	}

	if operation > 0 && investment < 0 && financing > 0 && overall > 0 {
		return "4-蛮牛"
	}

	if operation > 0 && investment < 0 && financing > 0 && overall < 0 {
		return "3-疯牛"
	}

	if operation > 0 && investment > 0 && financing > 0 {
		return "2-妖精型"
	}

	if operation < 0 && investment > 0 && financing > 0 {
		return "0-骗吃骗喝型"
	}

	if operation < 0 && investment < 0 && financing < 0 {
		return "0-混吃等死型"
	}

	if operation < 0 && investment < 0 && financing > 0 {
		return "0-赌徒型"
	}

	if operation < 0 && investment < 0 && financing < 0 {
		return "0-大出血型"
	}
	return ""
}
