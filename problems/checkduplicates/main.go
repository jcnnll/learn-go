package main

func main() {

}

func checkDuplicates(xi ...int) bool {
	if len(xi) < 2 {
		return false
	}
	fc := make(map[int]int)

	for _, v := range xi {
		if _, ok := fc[v]; ok {
			fc[v] += 1
			if fc[v] == 2 {
				return true
			}
			continue
		}
		fc[v] = 1
	}
	return false
}
