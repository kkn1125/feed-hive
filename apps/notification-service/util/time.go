package util

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"
)

type DateTime struct {
	time.Time
}

// MarshalJSON: time.Time → "YYYY-MM-DD HH:MM:SS" 형식 변환
func (ct DateTime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return json.Marshal(nil) // NULL일 경우 null 반환
	}
	return json.Marshal(ct.Time.Format("2006-01-02 15:04:05"))
}

func (ct DateTime) Value() (driver.Value, error) {
	return ct.Time, nil
}

func (ct *DateTime) Scan(value interface{}) error {
	if value == nil {
		*ct = DateTime{Time: time.Time{}}
		return nil
	}
	t, ok := value.(time.Time)
	if !ok {
		// return fmt.Errorf("can't convert %v to time.Time", value)
		return nil
	}
	*ct = DateTime{Time: t}
	return nil
}

// NullTime 커스텀 구조체
type NullTime struct {
	sql.NullTime
}

// JSON 직렬화 커스텀
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return json.Marshal(nil) // NULL 값일 때 null 반환
	}
	return json.Marshal(nt.Time.Format("2006-01-02 15:04:05")) // 원하는 형식으로 출력
}
