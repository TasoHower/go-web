package cache

import (
	gc "github.com/patrickmn/go-cache"
)

var c *gc.Cache

func InitMMCache() {
	c = gc.New(-1, -1)
}

func GetCache() *gc.Cache {
	return c
}
