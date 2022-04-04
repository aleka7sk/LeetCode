package leetcode

// func addDigits(num int) int {
//     for true{
//         var sum int
//         for num != 0 {
//             sum += num % 10
//             num = num / 10
//         }
//         if sum < 10 {
//             return sum
//         }
//         num = sum
//     }
//     return 0
// }

func addDigits(num int) int {
	if num == 0 {
		return 0
	} else if num%9 == 0 {
		return 9
	} else {
		return num % 9
	}
}
