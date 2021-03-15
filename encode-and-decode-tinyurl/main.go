package main

import "fmt"

type Codec struct {
	e map[string]string
	d map[string]string
}

func Constructor() Codec {
	return Codec{e: map[string]string{}, d: map[string]string{}}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	if v, ok := this.e[longUrl]; ok {
		return v
	}
	encoded := fmt.Sprintf("http://tinyurl.com/%d", len(this.e)+1)
	this.e[longUrl] = encoded
	this.d[encoded] = longUrl
	return encoded
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	if v, ok := this.d[shortUrl]; ok {
		return v
	}
	return ""
}
