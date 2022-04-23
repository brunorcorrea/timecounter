package main

import (
	"fmt"
	"strconv"
)

func main() {
	var time string
	var timeInSeconds int

	for time != "x" {
		fmt.Scanf("%s", &time)
		length := len(time)
		switch length {
		case 2:
			seconds := time[0:2]
			secondsInt, err := strconv.Atoi(seconds)
			if err != nil {
				return
			}
			timeInSeconds += secondsInt
		case 4:
			minutes := time[0:2]
			seconds := time[2:4]

			minutesInt, err := (strconv.Atoi(minutes))
			if err != nil {
				return
			}

			secondsInt, err := strconv.Atoi(seconds)
			if err != nil {
				return
			}

			timeInSeconds += (secondsInt + minutesInt*60)
		case 6:
			hours := time[0:2]
			minutes := time[2:4]
			seconds := time[4:6]

			hourInt, err := (strconv.Atoi(hours))
			if err != nil {
				return
			}

			minutesInt, err := (strconv.Atoi(minutes))
			if err != nil {
				return
			}

			secondsInt, err := strconv.Atoi(seconds)
			if err != nil {
				return
			}

			timeInSeconds += (secondsInt + hourInt*3600 + minutesInt*60)
		default:
			time = "x"
		}
	}

	totalTime := strconv.Itoa(timeInSeconds/3600) + " hours " + strconv.Itoa((timeInSeconds/60)-(timeInSeconds/3600)*60) + " minutes " + strconv.Itoa(timeInSeconds%60) + " seconds."

	fmt.Println("The total time of the videos is " + totalTime)
}
