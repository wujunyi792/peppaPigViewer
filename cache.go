package main

var localCache map[string]string

func init() {
	localCache = make(map[string]string)
}

func setCache(k, v string) {
	localCache[k] = v
}

func getCache(k string) (v string, exist bool) {
	v, exist = localCache[k]
	return
}
