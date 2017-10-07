package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Incident struct {
	id          int
	trigger     string
	start       string
	end         string
	assignee    string
	duration    string
	escalations int
	prio        string
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func main() {

	if len(os.Args) < 3 {
		log.Fatal("Not enough arguments")
		os.Exit(2)
	}
	fileName := os.Args[1]
	name := os.Args[2]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		if name == data[4] {

			alertstart, err := time.Parse("2006-01-02 15:04:05 MST", data[2])
			if err != nil {
				log.Fatal(err)
			}
			alertend, err := time.Parse("2006-01-02 15:04:05 MST", data[2])
			if err != nil {
				log.Fatal(err)
			}

			workdaystart, err := time.Parse("2006-01-02 15:04:05 -0700", alertstart.Format("2006-01-02")+" 10:00:00 +0200")
			if err != nil {
				log.Fatal(err)
			}
			workdayend, err := time.Parse("2006-01-02 15:04:05 -0700", alertstart.Format("2006-01-02")+" 18:00:00 +0200")
			if err != nil {
				log.Fatal(err)
			}
			_, w := alertstart.ISOWeek()
			startDay := alertstart.Weekday()
			endDay := alertend.Weekday()
			startIsWeekend := startDay == time.Saturday || startDay == time.Sunday
			endDayIsWeekend := endDay == time.Saturday || endDay == time.Sunday
			if (startIsWeekend || !inTimeSpan(workdaystart, workdayend, alertstart)) && (endDayIsWeekend || !inTimeSpan(alertend, workdayend, alertend)) {
				fmt.Println("Incident", data[0], "in CW", w, "at", alertstart.Format("02.01 15:04 MST"), "(", data[5], ") is outside working hours")
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}
