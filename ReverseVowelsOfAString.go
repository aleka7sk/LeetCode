package leetcode

func reverseVowels(s string) string {
	arr := make([]int, 0)
	for index, char := range s {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
			arr = append(arr, index)
		}
	}
	changeableString := []byte(s)
	for i := 0; i < len(arr)/2; i++ {
		changeableString[arr[i]], changeableString[arr[len(arr)-1-i]] = changeableString[arr[len(arr)-1-i]], changeableString[arr[i]]
	}
	s = string(changeableString)
	return s
}
