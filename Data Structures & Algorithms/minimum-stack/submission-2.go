type MinStack struct {
	ValStack 	[]int
	MinValStack []int
	CurrMin		int
}

func Constructor() MinStack {
	return MinStack{
		CurrMin: math.MaxInt,
	}
}

func (this *MinStack) Push(val int) {
	this.ValStack = append(this.ValStack, val)
	if val < this.CurrMin {
		this.CurrMin = val
	}
	this.MinValStack = append(this.MinValStack, this.CurrMin)
}

func (this *MinStack) Pop() {
	this.ValStack = this.ValStack[:len(this.ValStack)-1]
	this.MinValStack = this.MinValStack[:len(this.MinValStack)-1]

	if len(this.MinValStack) == 0 {
        this.CurrMin = math.MaxInt
    } else {
        this.CurrMin = this.MinValStack[len(this.MinValStack)-1]
    }
}

func (this *MinStack) Top() int {
	return this.ValStack[len(this.ValStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinValStack[len(this.MinValStack)-1]
}
