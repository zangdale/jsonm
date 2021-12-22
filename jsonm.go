package jsonm

import (
	"encoding/json"
)

/*
	获取 json[]byte 类型数据中的指定位置的 key 所表示的值 信息
*/

type JsonM map[string]interface{}

func NewJsonM(data []byte) (JsonM, error) {
	jm := make(JsonM)
	return jm, json.Unmarshal(data, &jm)
}

func (j JsonM) Map(key string) (interface{}, bool) {
	body, ok := j[key]
	return body, ok
}

func (j JsonM) JsonM(key string) (JsonM, bool) {
	jm, ok := j[key]
	if ok {
		body, ok := jm.(map[string]interface{})
		return body, ok
	}
	return nil, ok
}

func (j JsonM) Maps(keys ...string) (interface{}, bool) {

	n := len(keys)
	if n == 0 {
		return j, false
	}
	if n == 1 {
		body, ok := j[keys[0]]
		return body, ok
	}

	temp := j

	m := n - 1
	for i := 0; i < m; i++ {
		jm, ok := temp[keys[i]]
		if ok {
			body, ok := jm.(map[string]interface{})
			if ok {
				temp = body
				continue
			}
			return jm, ok
		}
		return jm, ok
	}
	body, ok := temp[keys[m]]
	return body, ok
}

func JsonMap(data []byte, key string) (interface{}, bool, error) {
	jm, err := NewJsonM(data)
	if err != nil {
		return nil, false, err
	}
	body, ok := jm.Map(key)
	return body, ok, nil
}

func JsonMaps(data []byte, keys ...string) (interface{}, bool, error) {
	jm, err := NewJsonM(data)
	if err != nil {
		return nil, false, err
	}
	body, ok := jm.Maps(keys...)
	return body, ok, nil
}
