package mystrings

const (
	LowerCaseA = 97
	UpperCaseZ = 90
)

func Lower(s string) string {
	input := []rune(s)
	out := make([]rune, len(input))
	for i, r := range input {
		if r <= UpperCaseZ {
			out[i] = r + 32
		} else {
			out[i] = r
		}
	}
	return string(out)
}

func Upper(s string) string {
	input := []rune(s)
	out := make([]rune, len(input))
	for i, r := range input {
		if r >= LowerCaseA {
			out[i] = r - 32
		} else {
			out[i] = r
		}
	}
	return string(out)
}
