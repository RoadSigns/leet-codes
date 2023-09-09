package recent_calls

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	for this.queue[0] < t-3000 {
		this.queue = this.queue[1:]
	}
	return len(this.queue)
}
