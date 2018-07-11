package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GetTimeByHourMin(t string) time.Time {
	hm := strings.Split(t, ":")
	h, _ := strconv.Atoi(hm[0])
	m, _ := strconv.Atoi(hm[1])

	return time.Date(
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		int(h),
		int(m),
		0,
		0,
		time.UTC,
	)
}

func GetCronTime(hrMn string) string {
	t := GetTimeByHourMin(hrMn)
	cronTime := t.Add(time.Minute * -10)

	return fmt.Sprintf(
		"%02v:%02v",
		cronTime.Hour(),
		cronTime.Minute(),
	)
}
