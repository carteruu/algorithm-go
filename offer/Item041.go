package offer

type MovingAverage struct {
	size  int
	total int
	q     []int
	idx   int
}

/** Initialize your data structure here. */
func NewMovingAverage(size int) MovingAverage {
	return MovingAverage{
		size: size,
		q:    make([]int, 0, size),
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.q) < this.size {
		this.q = append(this.q, val)
		this.total += val
	} else {
		this.total = this.total - this.q[this.idx] + val
		this.q[this.idx] = val
		this.idx = (this.idx + 1) % this.size
	}
	return float64(this.total) / float64(len(this.q))
}
