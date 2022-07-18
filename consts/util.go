package consts

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
	"sync"
)

var wait sync.WaitGroup
var lo sync.RWMutex

type Util struct {
}
type Result struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    any    `json:"data"`
}

func (r Result) ToBytes() []byte {
	buf := new(bytes.Buffer)
	if s, err := json.Marshal(&r); err != nil {
		log.Printf("解析数据异常--->>>%v \n", s)
		return nil
	} else {
		if err := binary.Write(buf, binary.BigEndian, s); err != nil {
			log.Printf("err --->>>> %v \n", err.Error())
			return nil
		}
		return buf.Bytes()
	}
}

type ConcurrentMap map[string]any

func (c ConcurrentMap) GetValue(key string) any {
	lo.Lock()
	defer lo.Unlock()
	return c[key]
}

func (c ConcurrentMap) SetValue(key string, obj any) {
	lo.Lock()
	defer lo.Unlock()
	c[key] = obj
}

func (c ConcurrentMap) DeleteValue(key string) {
	lo.Lock()
	defer lo.Unlock()
	delete(c, key)
}
