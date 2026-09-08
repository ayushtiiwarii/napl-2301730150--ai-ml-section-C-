package strop

func CountVowels(s string) int {
	count := 0

	for _, H := range s {
		if H == 'a' || H == 'e' || H == 'i' || H == 'o' || H == 'u' ||
			H == 'A' || H == 'E' || H == 'I' || H == 'O' || H == 'U' {
			count++
		}
	}

	return count
}
