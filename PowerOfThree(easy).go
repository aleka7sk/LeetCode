package leetcode

// func isPowerOfThree(n int) bool {
//     i := 1
//     for i <= n {
//         if i == n {
//             return true
//         }
//         i *= 3
//     }
//     return false
// }

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	if 1162261467%n == 0 {
		return true
	} else {
		return false
	}
}
