package main

func maximumTime(time string) string {
	hh := time[0:2]
	mm := time[3:]
	hour, minute := make([]byte, 2, 2), make([]byte, 2, 2)

	if hh[0] == '?' && hh[1] == '?' {
		hour[0] = '2'
		hour[1] = '3'
	} else if hh[0] == '?' {
		if hh[1] < '4' {
			hour[0] = '2'
			hour[1] = hh[1]
		} else {
			hour[0] = '1'
			hour[1] = hh[1]
		}
	} else if hh[1] == '?' {
		if hh[0] == '2' {
			hour[0] = '2'
			hour[1] = '3'
		} else {
			hour[0] = hh[0]
			hour[1] = '9'
		}
	} else {
		hour[0] = hh[0]
		hour[1] = hh[1]
	}

	if mm[0] == '?' {
		minute[0] = '5'
	} else {
		minute[0] = mm[0]
	}
	if mm[1] == '?' {
		minute[1] = '9'
	} else {
		minute[1] = mm[1]
	}
	return string(hour) + ":" + string(minute)
}
