package strptime

import (
	"testing"
	"time"
)

func TestDatePaths(t *testing.T) {
	var err error
	var tm time.Time
	datepath_format1 := "/path/with/dates/%Y/%m/%d"
	datepath_format2 := "/blog/%Y/%m/%d/%-"
	datepath_format3 := "/blog/%Y/%m/%d/%-/page2"
	datepath_format4 := "/random/02/03/-0700/numbers/%Y/%m/%d"
	datepath_format5 := "/random/%-/numbers/%Y/%m/%d"

	tm, err = Strptime("/path/with/dates/2012/04/28", datepath_format1)
	if err != nil {
		t.Error(err)
	} else if !checkDate(2012, time.April, 28, tm) {
		t.Error("Invalid date parsed")
	}

	tm, err = Strptime("/unmatched/path/with/dates/2012/04/28", datepath_format1)
	if err == nil {
		t.Error("Should fail with invalid format")
	}

	tm, err = Strptime("/blog/2012/04/22/a-post-slug", datepath_format2)
	if err != nil {
		t.Error(err)
	} else if !checkDate(2012, time.April, 22, tm) {
		t.Error("Invalid date parsed")
	}

	tm, err = Strptime("/badblog/2012/04/28/a-post-slug", datepath_format2)
	if err == nil {
		t.Error("Should fail with invalid format")
	}

	tm, err = Strptime("/blog/2012/04/28/a-post-slug/page2", datepath_format2)
	if err != nil {
		t.Error(err)
	} else if !checkDate(2012, time.April, 28, tm) {
		t.Error("Invalid date parsed")
	}

	tm, err = Strptime("/blog/2012/04/22/a-post-slug/page2", datepath_format3)
	if err != nil {
		t.Error(err)
	} else if !checkDate(2012, time.April, 22, tm) {
		t.Error("Invalid date parsed")
	}

	tm, err = Strptime("/blog/2012/04/28/a-post-slug", datepath_format3)
	if err == nil {
		t.Error("Should fail with invalid format")
	}

	tm, err = Strptime("/badblog/2012/04/28/a-post-slug/page2", datepath_format3)
	if err == nil {
		t.Error("Should fail with invalid format")
	}

	tm, err = Strptime("/random/02/03/-0700/numbers/2012/04/28", datepath_format4)
	if err != nil {
		t.Error(err)
	} else if !checkDate(2012, time.April, 28, tm) {
		t.Error("Invalid date parsed")
	}

	tm, err = Strptime("/random/04/28/-0500/numbers/2012/04/28", datepath_format4)
	if err == nil {
		t.Error("Should fail with invalid format")
	}

	tm, err = Strptime("/random/02/03/-0700/numbers/2012/04/22", datepath_format5)
	if err != nil {
		t.Error(err)
	} else if !checkDate(2012, time.April, 22, tm) {
		t.Error("Invalid date parsed")
	}

	tm, err = Strptime("/random/04/28/-0500/numbers/2012/04/28", datepath_format5)
	if err != nil {
		t.Error(err)
	} else if !checkDate(2012, time.April, 28, tm) {
		t.Error("Invalid date parsed")
	}
}

func TestCrazyDates(t *testing.T) {
	var err error
	var tm time.Time
	date_format1 := "%Y%m%d"
	date_format2 := "%b%d%Y"
	date_format3 := "%b%d"

	tm, err = Strptime("20120428", date_format1)
	if err != nil {
		t.Error(err)
	} else if !checkDate(2012, time.April, 28, tm) {
		t.Error("Invalid date parsed")
	}
	tm, err = Strptime("20121111", date_format1)
	if err != nil {
		t.Error(err)
	} else if !checkDate(2012, time.November, 11, tm) {
		t.Error("Invalid date parsed")
	}

	tm, err = Strptime("2012111", date_format1)
	if err == nil {
		t.Error("Should fail with invalid format")
	}
	tm, err = Strptime("20120440", date_format1)
	if err == nil {
		t.Error("Should fail with invalid date", tm)
	}

	///

	tm, err = Strptime("Apr152012", date_format2)
	if err != nil {
		t.Error(err)
	} else if !checkDate(2012, time.April, 15, tm) {
		t.Error("Invalid date parsed")
	}
	tm, err = Strptime("Apr201211", date_format2)
	if err != nil {
		t.Error(err)
	} else if !checkDate(1211, time.April, 20, tm) {
		t.Error("Invalid date parsed")
	}

	tm, err = Strptime("2012111", date_format2)
	if err == nil {
		t.Error("Should fail with invalid format")
	}
	tm, err = Strptime("04042012", date_format2)
	if err == nil {
		t.Error("Should fail with invalid format")
	}

	///

	tm, err = Strptime("Apr15", date_format3)
	if err != nil {
		t.Error(err)
	} else if !checkDate(0, time.April, 15, tm) {
		t.Error("Invalid date parsed")
	}
	tm, err = Strptime("Apr20", date_format3)
	if err != nil {
		t.Error(err)
	} else if !checkDate(0, time.April, 20, tm) {
		t.Error("Invalid date parsed")
	}

	tm, err = Strptime("Apr2012", date_format3)
	if err == nil {
		t.Error("Should fail with invalid date", tm)
	}
	tm, err = Strptime("Apr00", date_format3)
	if err == nil {
		t.Log("Really should fail with invalid date", tm)
	}
}

func TestTimes(t *testing.T) {
	var err error
	var tm time.Time
	time_format1 := "%H:%M:%S.%f"
	time_format2 := "%I:%M%p"

	tm, err = Strptime("20:42:15.98", time_format1)
	if err != nil {
		t.Error(err)
	} else if !checkTime(20, 42, 15, tm) {
		t.Error("Invalid time parsed")
	}
	tm, err = Strptime("02:42:15.4", time_format1)
	if err != nil {
		t.Error(err)
	} else if !checkTime(2, 42, 15, tm) {
		t.Error("Invalid time parsed")
	}
	tm, err = Strptime("32:42:15.4", time_format1)
	if err == nil {
		t.Error("Should fail with invalid time", tm)
	}

	tm, err = Strptime("11:42am", time_format2)
	if err != nil {
		t.Error(err)
	} else if !checkTime(11, 42, 0, tm) {
		t.Error("Invalid time parsed")
	}
	tm, err = Strptime("11:42pm", time_format2)
	if err != nil {
		t.Error(err)
	} else if !checkTime(23, 42, 0, tm) {
		t.Error("Invalid time parsed")
	}

	tm, err = Strptime("12:02am", time_format2)
	if err != nil {
		t.Error(err)
	} else if !checkTime(0, 2, 0, tm) {
		t.Error("Invalid time parsed")
	}
	tm, err = Strptime("12:02pm", time_format2)
	if err != nil {
		t.Error(err)
	} else if !checkTime(12, 2, 0, tm) {
		t.Error("Invalid time parsed")
	}

	tm, err = Strptime("00:02am", time_format2)
	if err == nil {
		t.Log("Really should fail with invalid time", tm)
	}
	tm, err = Strptime("00:02pm", time_format2)
	if err == nil {
		t.Log("Really should fail with invalid time", tm)
	}
}

func checkDate(year int, month time.Month, day int, tm time.Time) bool {
	y, m, d := tm.Date()
	if y != year || m != month || d != day {
		return false
	}
	return true
}

func checkTime(hour, minute, sec int, tm time.Time) bool {
	if tm.Hour() != hour || tm.Minute() != minute || tm.Second() != sec {
		return false
	}
	return true
}
