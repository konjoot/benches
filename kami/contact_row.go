package main

import "sync"

var contactPool = sync.Pool{
	New: func() interface{} {
		iface := make([]interface{}, 7)
		for i := range iface {
			iface[i] = new(interface{})
		}
		return iface
	},
}

// func NewContactRow() []interface{} {
// 	return contactPool.Get().([]interface{})
// }
