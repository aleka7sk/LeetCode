package leetcode

type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	return *new(MinStack)
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	min := this.stack[0]
	for i := 1; i < len(this.stack); i++ {
		if min > this.stack[i] {
			min = this.stack[i]
		}
	}
	return min
}
