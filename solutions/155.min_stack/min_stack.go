package minstack

type MinStack struct {
	elements []int
	minimuns []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.elements = append(this.elements, val)

	if len(this.minimuns) == 0 {
		this.minimuns = append(this.minimuns, val)
	} else if val <= this.minimuns[len(this.minimuns)-1] {
		this.minimuns = append(this.minimuns, val)
	}
}

func (this *MinStack) Pop() {
	if this.Top() == this.minimuns[len(this.minimuns)-1] {
		this.minimuns = this.minimuns[:len(this.minimuns)-1]
	}

	if len(this.elements) > 0 {
		this.elements = this.elements[:len(this.elements)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.elements) > 0 {
		return this.elements[len(this.elements)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	return this.minimuns[len(this.minimuns)-1]
}
