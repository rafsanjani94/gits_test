package conv

import (
	"strconv"
)

func StringToFloat32(buf string, dval float32) float32 {
	res, err := strconv.ParseFloat(buf, 32)

	if err == nil {
		return float32(res)
	}

	return dval
}

func Float32ToString(val float32) string {
	return strconv.FormatFloat(
		float64(val),
		'f',
		-1,
		32,
	)
}

func StringToFloat64(buf string, dval float64) float64 {
	res, err := strconv.ParseFloat(buf, 64)

	if err == nil {
		return res
	}

	return dval
}

func Float64ToString(val float64) string {
	return strconv.FormatFloat(
		val,
		'f',
		-1,
		64,
	)
}

func StringToInt(buf string, dval int) int {
	res, err := strconv.ParseInt(buf, 10, 0)

	if err == nil {
		return int(res)
	}

	return dval
}

func IntToString(val int) string {
	return strconv.FormatInt(
		int64(val),
		10,
	)
}

func StringToInt8(buf string, dval int8) int8 {
	res, err := strconv.ParseInt(buf, 10, 8)

	if err == nil {
		return int8(res)
	}

	return dval
}

func Int8ToString(val int8) string {
	return strconv.FormatInt(
		int64(val),
		10,
	)
}

func StringToInt16(buf string, dval int16) int16 {
	res, err := strconv.ParseInt(buf, 10, 16)

	if err == nil {
		return int16(res)
	}

	return dval
}

func Int16ToString(val int16) string {
	return strconv.FormatInt(
		int64(val),
		10,
	)
}

func StringToInt32(buf string, dval int32) int32 {
	res, err := strconv.ParseInt(buf, 10, 32)

	if err == nil {
		return int32(res)
	}

	return dval
}

func Int32ToString(val int32) string {
	return strconv.FormatInt(
		int64(val),
		10,
	)
}

func StringToInt64(buf string, dval int64) int64 {
	res, err := strconv.ParseInt(buf, 10, 64)

	if err == nil {
		return res
	}

	return dval
}

func Int64ToString(val int64) string {
	return strconv.FormatInt(val, 10)
}

func StringToUint(buf string, dval uint) uint {
	res, err := strconv.ParseUint(buf, 10, 0)

	if err == nil {
		return uint(res)
	}

	return dval
}

func UintToString(val uint) string {
	return strconv.FormatUint(
		uint64(val),
		10,
	)
}

func StringToUint8(buf string, dval uint8) uint8 {
	res, err := strconv.ParseUint(buf, 10, 8)

	if err == nil {
		return uint8(res)
	}

	return dval
}

func Uint8ToString(val uint8) string {
	return strconv.FormatUint(
		uint64(val),
		10,
	)
}

func StringToUint16(buf string, dval uint16) uint16 {
	res, err := strconv.ParseUint(buf, 10, 16)

	if err == nil {
		return uint16(res)
	}

	return dval
}

func Uint16ToString(val uint16) string {
	return strconv.FormatUint(
		uint64(val),
		10,
	)
}

func StringToUint32(buf string, dval uint32) uint32 {
	res, err := strconv.ParseUint(buf, 10, 32)

	if err == nil {
		return uint32(res)
	}

	return dval
}

func Uint32ToString(val uint32) string {
	return strconv.FormatUint(
		uint64(val),
		10,
	)
}

func StringToUint64(buf string, dval uint64) uint64 {
	res, err := strconv.ParseUint(buf, 10, 64)

	if err == nil {
		return res
	}

	return dval
}

func Uint64ToString(val uint64) string {
	return strconv.FormatUint(val, 10)
}

func StringToBool(val string) bool {
	out, _ := strconv.ParseBool(val)
	return out
}

func BoolToInt(val bool) (res int8) {
	res = 0
	if val {
		res = 1
	}

	return res
}

func Int8ToBool(val int8) (res bool) {
	res = false
	if val == 1 {
		res = true
	}

	return res
}

func Float64ToBool(val float64) (res bool) {
	res = false
	if val == 1 {
		res = true
	}

	return res
}
