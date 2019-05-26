package db

import (
	"hash/fnv"
	"reflect"
	"strconv"
)

// return OID of object
func getOID(value interface{}) (res int32) {
	var s string

	if value == nil {
		value = "nil"
	}

	switch v := value.(type) {
	case string:
		s = v
	case []byte:
		s = string(v)
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', 2, 64)
	case float64:
		s = strconv.FormatFloat(v, 'f', 2, 64)
	case int:
		s = strconv.FormatInt(int64(v), 10)
	case int8:
		s = strconv.FormatInt(int64(v), 10)
	case int16:
		s = strconv.FormatInt(int64(v), 10)
	case int32:
		s = strconv.FormatInt(int64(v), 10)
	case int64:
		s = strconv.FormatInt(v, 10)
	case uint:
		s = strconv.FormatUint(uint64(v), 10)
	case uint8:
		s = strconv.FormatUint(uint64(v), 10)
	case uint16:
		s = strconv.FormatUint(uint64(v), 10)
	case uint32:
		s = strconv.FormatUint(uint64(v), 10)
	case uint64:
		s = strconv.FormatUint(v, 10)
	default:
		s = reflect.ValueOf(value).Type().String()
	}
	res = hashText(s)
	return
}

func hashText(s string) int32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	sum := int32(h.Sum32())
	if sum < 0 {
		sum = -sum
	}
	return sum
}
