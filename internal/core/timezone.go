package core

import "time"

func ConvertTimezone(t time.Time, fromTimeZone, toTimeZone string) (time.Time, error){
	location, err := time.LoadLocation(toTimeZone)
	if err != nil {
		return time.Time{}, err
	}
	return t.In(location), nil
}