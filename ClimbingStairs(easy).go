package leetcode

func climbStairs(n int) int {
	arr := []int{0, 1}
	for i := 0; i < n; i++ {
		arr = append(arr, arr[i]+arr[i+1])
	}
	return arr[len(arr)-1]

}

// func climbStairs(n int) int {
//     if n == 2 {
//         return 2
//     } else if n == 1 {
//         return 1
//     }
//     if n <= 45 {
//          return climbStairs(n-2) + climbStairs(n-1)
//     }
//     return 0
// }
