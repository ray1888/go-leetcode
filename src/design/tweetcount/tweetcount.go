package tweetcount

import "sort"

/*
	总体逻辑：
	1. 插入的时候，直接根据用户在哈希表中进行查询，然后再append即可
	2. 查询的时候，先找到时间上下界的地方，然后加入返回值中

	查询的话：
	1. 目前版本是直接全部遍历数据来继续宁
	2. 更优化的版本是，通过平衡二叉树来实现这个数据结构，使得可以通过再log(n)的时间复杂度内找到上下界的所属，然后可以直接把那段数据进行截取。

*/

type TweetCounts struct {
	tweets map[string][]int
}

func Constructor() TweetCounts {
	tc := TweetCounts{
		tweets: map[string][]int{},
	}
	return tc
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	this.tweets[tweetName] = append(this.tweets[tweetName], time)
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) (rst []int) {
	sort.Ints(this.tweets[tweetName])
	var freqSecond = 60
	switch freq {
	case "minute":
		break
	case "hour":
		freqSecond = 3600
	case "day":
		freqSecond = 24 * 3600
	}

	var tweets = this.tweets[tweetName]
	rst = make([]int, (endTime-startTime)/freqSecond+1)
	for i := 0; i < len(tweets); i++ {
		if tweets[i] < startTime {
			continue
		}
		if tweets[i] > endTime {
			return rst

		}
		rst[(tweets[i]-startTime)/freqSecond]++
	}
	return rst
}
