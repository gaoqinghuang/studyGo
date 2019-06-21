package main
import (
	"fmt"
)

type Dictionary struct {
	data map[interface{}] interface{}
}

func (d *Dictionary) Get(key interface{}) interface{}{
	return d.data[key]
}

func (d *Dictionary) Set(key interface{}) interface{}{
	d.data[key] = value
}

func (d *Dictionary) Visit(callback)