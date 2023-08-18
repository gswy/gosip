package headers

import (
	"strings"
	"time"
)

type Date struct {
	Value time.Time
}

func (d *Date) String() string {
	return d.Value.Format("2006-01-02T15:04:05.000")
}

func ParseDate(date string) *Date {
	field := ScanHeaderField(date, "Date:")
	if field == nil {
		return nil
	}
	value := strings.Replace(*field, "Date:", "", 1)
	value = strings.TrimSpace(value)
	parse, err := time.Parse("2006-01-02T15:04:05.000", date)
	if err != nil {
		parse = time.Now()
	}
	return &Date{
		Value: parse,
	}

}
