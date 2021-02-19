package browserhistory

type BrowserHistory struct {
	cur   []string
	index int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		cur:   []string{homepage},
		index: 0,
	}
}

func (this *BrowserHistory) Visit(url string) {
	lenth := len(this.cur)
	if this.index == lenth-1 {
		this.cur = append(this.cur, url)
	} else if this.index < lenth-1 {
		this.cur = this.cur[:this.index+1]
		this.cur = append(this.cur, url)
	}
	this.index++
}

func (this *BrowserHistory) Back(steps int) string {
	if steps > this.index {
		this.index = 0
		return this.cur[0]
	}
	this.index -= steps
	return this.cur[this.index]
}

func (this *BrowserHistory) Forward(steps int) string {
	lenth := len(this.cur)
	if this.index == lenth-1 {
		return this.cur[this.index]
	} else if this.index+steps > lenth-1 {
		this.index = lenth - 1
		return this.cur[lenth-1]
	} else {
		this.index += steps
		return this.cur[this.index]
	}
}
