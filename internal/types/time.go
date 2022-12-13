package types

import "time"

const (
	timeFormat = "2006-01-02 15:04:05"
	dateFormat = "2006-01-02"
)

// JSONTime 格式化输出时间
type JSONTime time.Time

// MarshalJSON json解码
func (j JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(j).Local().Format(timeFormat) + `"`), nil
}

// UnmarshalJSON JSONDate反序列化
func (j *JSONTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*j = JSONTime(now)
	return
}

// JSONDate 格式化输出时间
type JSONDate time.Time

// MarshalJSON json解码
func (j JSONDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(j).Local().Format(dateFormat) + `"`), nil
}

// UnmarshalJSON JSONDate反序列化
func (j *JSONDate) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+dateFormat+`"`, string(data), time.Local)
	*j = JSONDate(now)
	return
}
