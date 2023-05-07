package pkg

func Like(arr []string, arr2 []string) bool {
	if len(arr) != len(arr2) {
		return false
	}

	var check = make(map[string]struct{})
	for _, v := range arr {
		check[v] = struct{}{}
	}

	for _, v := range arr2 {
		if _, ok := check[v]; !ok {
			return false
		}
	}

	return true
}
