package moving_average

type MovingAverage struct {
	Numbers []int
	Size    int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		Numbers: make([]int, 0),
		Size:    size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.Numbers) == this.Size {
		this.Numbers = this.Numbers[1:]
	}
	this.Numbers = append(this.Numbers, val)
	total := 0
	for _, i := range this.Numbers {
		total += i
	}

	return float64(total) / float64(len(this.Numbers))
}
