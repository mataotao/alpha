package slice

func Uint64SliceEqualBCE(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

//去重
func RemoveDuplicateElementUint64(s []uint64) []uint64 {
	result := make([]uint64, 0, len(s))
	temp := map[uint64]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func RemoveDuplicateElementString(s []string) []string {
	result := make([]string, 0, len(s))
	temp := map[string]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
func RemoveDuplicateElementint(s []int) []int {
	result := make([]int, 0, len(s))
	temp := map[int]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

//并集 返回 sSort 在s里的 并且按sSort本身的顺序返回
func IntersectionString(sSort, s []string) []string {
	m := make(map[string]struct{})
	filterS := make([]string, 0, len(sSort))
	for i := range s[:] {
		m[s[i]] = struct{}{}
	}
	for i := range sSort[:] {
		if _, ok := m[sSort[i]]; ok {
			filterS = append(filterS, sSort[i])
		}
	}
	return filterS
}

func PaginationString(s []string, page, limit int) []string {
	start := (page - 1) * limit
	stop := start + limit
	l := len(s)
	if stop > l {
		stop = l
	}
	r := s[start:stop]
	return r
}
