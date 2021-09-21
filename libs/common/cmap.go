// Copyright (c) 2020 The Meter.io developers
// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying

// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package common

import "sync"

// CMap is a goroutine-safe map
type CMap struct {
	m map[string]interface{}
	l sync.Mutex
}

func NewCMap() *CMap {
	return &CMap{
		m: make(map[string]interface{}),
	}
}

func (cm *CMap) Set(key string, value interface{}) {
	cm.l.Lock()
	cm.m[key] = value
	cm.l.Unlock()
}

func (cm *CMap) Get(key string) interface{} {
	cm.l.Lock()
	val := cm.m[key]
	cm.l.Unlock()
	return val
}

func (cm *CMap) Has(key string) bool {
	cm.l.Lock()
	_, ok := cm.m[key]
	cm.l.Unlock()
	return ok
}

func (cm *CMap) Delete(key string) {
	cm.l.Lock()
	delete(cm.m, key)
	cm.l.Unlock()
}

func (cm *CMap) Size() int {
	cm.l.Lock()
	size := len(cm.m)
	cm.l.Unlock()
	return size
}

func (cm *CMap) Clear() {
	cm.l.Lock()
	cm.m = make(map[string]interface{})
	cm.l.Unlock()
}

func (cm *CMap) Keys() []string {
	cm.l.Lock()

	keys := []string{}
	for k := range cm.m {
		keys = append(keys, k)
	}
	cm.l.Unlock()
	return keys
}

func (cm *CMap) Values() []interface{} {
	cm.l.Lock()
	items := []interface{}{}
	for _, v := range cm.m {
		items = append(items, v)
	}
	cm.l.Unlock()
	return items
}
