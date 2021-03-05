package shortenurl

import (
	"math/rand"
	"strings"
)

const alpha = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Codec struct {
	maxValue   int64
	usedURL    map[string]string
	randomSize int
}

func Constructor() Codec {
	c := Codec{}
	c.usedURL = make(map[string]string)
	c.randomSize = 3
	c.maxValue = 1
	return c
}

func tosixtytwo(a int64) string {
	var sb strings.Builder
	for a > 0 {
		sb.WriteByte(alpha[a%62])
		a /= 62
	}
	return sb.String()
}

func (this *Codec) genprefix() string {
	var sb strings.Builder
	for i := 0; i < this.randomSize; i++ {
		sb.WriteByte(alpha[rand.Intn(61)])
	}
	return sb.String()
}

// 加密方法1： 使用随机数+自增ID 进行实现，但是实际上要执行62进制的算法
// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {

	suffix := tosixtytwo(this.maxValue)
	prefix := this.genprefix()
	full := prefix + suffix
	this.usedURL[full] = longUrl
	this.maxValue++
	return "http://tinyurl.com/" + full
}

// 加密方法2： 采用直接的随机数为来进行，通过不断累加后面的长度来进行处理
func (this *Codec) encodeNotUsedIncrId(longUrl string) string {

	var sb strings.Builder
	for i := 0; i < 7; i++ {
		sb.WriteByte(alpha[rand.Intn(61)])
	}
	full := sb.String()
	this.usedURL[full] = longUrl
	return "http://tinyurl.com/" + full
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	suffix := strings.Replace(shortUrl, "http://tinyurl.com/", "", -1)
	return this.usedURL[suffix]
}
