package util

func ReverseSrting(s string) string {
	var res string
	for _, v := range s {
		res = string(v) + res
	}
	return res
}
