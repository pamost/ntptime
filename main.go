package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

func main() {

	fmt.Printf("Current system time: %v\n", time.Now())
	fmt.Printf("Formatted current system time: %v\n", time.Now().Format("2006-01-02 15:04:05"))

	// Time server https://www.ntp-servers.net/servers.html
	ntpTime, err := ntp.Time("ntp3.stratum2.ru")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Exact time using the NTP library: %v\n", ntpTime)

	formatNtpTime := ntpTime.Format("2006-01-02 15:04:05")
	fmt.Printf("Formatted time using the NTP library: %v\n", formatNtpTime)
}
