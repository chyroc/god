package god_list

func IsStringIn(elem string, elemList []string) bool {
	for _, v := range elemList {
		if v == elem {
			return true
		}
	}
	return false
}

func IsBoolIn(elem bool, elemList []bool) bool {
	for _, v := range elemList {
		if v == elem {
			return true
		}
	}
	return false
}

func IsIntIn(elem int, elemList []int) bool {
	for _, v := range elemList {
		if v == elem {
			return true
		}
	}
	return false
}

func IsInt8In(elem int8, elemList []int8) bool {
	for _, v := range elemList {
		if v == elem {
			return true
		}
	}
	return false
}

func IsInt16In(elem int16, elemList []int16) bool {
	for _, v := range elemList {
		if v == elem {
			return true
		}
	}
	return false
}

func IsInt32In(elem int32, elemList []int32) bool {
	for _, v := range elemList {
		if v == elem {
			return true
		}
	}
	return false
}

func IsInt64In(elem int64, elemList []int64) bool {
	for _, v := range elemList {
		if v == elem {
			return true
		}
	}
	return false
}

func IsUintIn(elem uint, elemList []uint) bool {
	for _, v := range elemList {
		if v == elem {
			return true
		}
	}
	return false
}

func IsUint8In(elem uint8, elemList []uint8) bool {
	for _, v := range elemList {
		if v == elem {
			return true
		}
	}
	return false
}

func IsUint16In(elem uint16, elemList []uint16) bool {
	for _, v := range elemList {
		if v == elem {
			return true
		}
	}
	return false
}

func IsUint32In(elem uint32, elemList []uint32) bool {
	for _, v := range elemList {
		if v == elem {
			return true
		}
	}
	return false
}

func IsUint64In(elem uint64, elemList []uint64) bool {
	for _, v := range elemList {
		if v == elem {
			return true
		}
	}
	return false
}
