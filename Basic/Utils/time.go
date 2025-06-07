package Utils


import "fmt"
import "time"
import "strings"

func Time() {

	currentTime := time.Now()
	fmt.Println(currentTime)

	// The date is must to be 02-01-2006 any other values than it will result into some wrong date
	// dd-mm-yyyy
	format1 := currentTime.Format("02-01-2006")
	fmt.Println(format1)

	// yyyy-mm-dd
	format2 := currentTime.Format("2006-01-02")
	fmt.Println(format2);

	// This is fixed for day - Monday
	// get current day 
	format3 := currentTime.Format("Monday")
	fmt.Println(format3)

	// get date and day
	format4 := currentTime.Format("02-01-2006, Monday")
	fmt.Println(format4)

	// get current time 
	format5 := currentTime.Format("15-04-05")
	fmt.Println(format5)

	// get date and time
	format6 := currentTime.Format("02-01-2006, 15-04-05")
	fmt.Println(format6)

	// This is fixed for time 15-04-05
	// get date day and time
	format7 := currentTime.Format("02-01-2006, 15-04-05, Monday")
	fmt.Println(format7)

	// get date day and time
	format8 := currentTime.Format("02-01-2006, Monday, 15-04-05")
	fmt.Println(format8)

	// get date day and time
	format9 := currentTime.Format("02/01/2006, Monday, 15:04:05")
	fmt.Println(format9)

	// get am or pm 
	format10 := currentTime.Format("03:04 PM")
	fmt.Println(format10)
	fmt.Println(strings.Split(format10," ")[1])

	// string to time 
	// layout is in yyyy-mm-dd format
	layout_string := "2006-01-02"
	str := "2025-05-21"
	format11,_ := time.Parse(layout_string,str)
	fmt.Println(format11)

	// add time to current data to produce new time
	new_time := currentTime.Add(24*time.Hour)
	fmt.Println(new_time)
	format12 := new_time.Format("02-01-2006 Monday")
	fmt.Println(format12)
}