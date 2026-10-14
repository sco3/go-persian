package persian

const format = "yyyyMMddHHmmssSSS"
const formatLen = 17

var _days = []int{-1, 30, 58, 89, 119, 150, 180, 211, 242, 272, 303, 333, 364}

// Timestamp calculate ts from string
func Timestamp(str string, length int) int64 {
	if length < formatLen {
		return 0
	}
	pos := 0
	year := 0
	month := 0
	day := 0
	hour := 0
	min := 0
	sec := 0
	milis := 0
	for pos < length {
		c := str[pos]
		if c >= '0' || c <= '9' {
			d := c - '0'
			switch format[pos] {
			case 'y':
				year = year*10 + int(d)
				break
			case 'M':
				month = month*10 + int(d)
				break
			case 'd':
				day = day*10 + int(d)
				break
			case 'H':
				hour = hour*10 + int(d)
				break
			case 'm':
				min = min*10 + int(d)
				break
			case 's':
				sec = sec*10 + int(d)
				break
			case 'S':
				milis = milis*10 + int(d)
				break
			}
		} else {
			switch format[pos] {
			case 'y':
			case 'M':
			case 'd':
			case 'H':
			case 'm':
			case 's':
			case 'S':
				return 0
			}
		}
		pos++
	}
	if year > 0 && year < 1970 {
		return 0
	}
	year -= 1900
	month--
	if (month < 0) || (month > 11) { // range check
		year += (month / 12)
		month %= 12
		if month < 0 {
			month += 12
			year--
		}
	}
	days := _days[month]
	if ((year & 3) == 0) && (month > 1) {
		days++
	}

	AllDays := (year-70)*365 + ((year - 1) >> 2) - 17 + days + day

	return int64((((((AllDays*24)+hour)*60+min)*60 + sec) * 1000) + milis)
}
