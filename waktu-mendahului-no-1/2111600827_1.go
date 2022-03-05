// Golang program to illujaraktrate the usage of

// Including the main package
package main

// Importing fmt
import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Calling main
func main() {

	// Declaring some variable
	jarak := 0.0
	kecepatanali := 2.0
	kecepatanbadu := 3.0
	waktutempuh := 0
	waktumenyusul := 0.0
	selisihjarak := 0.0
	kecepatankedua := 0.0

	fmt.Println("+===========================================+")
	fmt.Println("| NIM: 2111600827 ")
	fmt.Println("| Nama: Mohd Amiruddin Saddam ")
	fmt.Println("| Program untuk mencetak Badu mendahului Ali ")
	fmt.Println("+===========================================+")
	// Calling Scanf() function for
	// scanning and reading the input
	// texts given in standard input
	for jarak < 1000 {
		jarak = jarak + kecepatanali
		waktutempuh = waktutempuh + 1
		if waktutempuh%10 == 0 {
			kecepatanali = kecepatanali + 0.1
		}
		if waktutempuh == 60 {
			selisihjarak = 60
			kecepatankedua = kecepatanbadu - kecepatanali
			waktumenyusul = selisihjarak / kecepatankedua
		}
	}

	timeHoursFinalStr, timeMinuteFinalStr, timeSecondFinalStr := convertToTime("08:00:00", int(waktumenyusul))

	fmt.Println(fmt.Sprintf("Badu Akan Mendahului Pada %s:%s:%s ", timeHoursFinalStr, timeMinuteFinalStr, timeSecondFinalStr))

}

//untuk convert dari detik ke waktu
func convertToTime(timeLeave string, secondReq int) (string, string, string) {

	// Declaring some variables
	var hours, minute, secondResp int
	var timeHours, timeMinute, timeSecond, timeHoursFinal, timeMinuteFinal, timeSecondFinal, totalTimeHours, totalTimeMinute, totalTimeSecond int

	// save process result into variabel
	hours = secondReq / 3600
	modMinuteSec := secondReq % 3600
	minute = modMinuteSec / 60
	secondResp = modMinuteSec % 60
	strTime := []string{}
	//convert to time
	if len(strings.Split(timeLeave, ":")) > 0 {
		for i := range strings.Split(timeLeave, ":") {
			convInt, _ := strconv.Atoi(strings.Split(timeLeave, ":")[i])
			if convInt < 10 {
				strings.Split(timeLeave, ":")[i] = strings.TrimLeft(strings.Split(timeLeave, ":")[i], "0")
			}
			strTime = append(strTime, strings.Split(timeLeave, ":")[i])
		}
		timeHours, _ = strconv.Atoi(strTime[0])
		timeMinute, _ = strconv.Atoi(strTime[1])
		timeSecond, _ = strconv.Atoi(strTime[2])

	}

	diffMinute := timeMinute
	diffSecond := timeSecond
	// final time
	totalTimeHours = timeHours + hours
	totalTimeMinute = timeMinute + minute
	totalTimeSecond = timeSecond + secondResp

	timeHoursFinal = totalTimeHours
	timeMinuteFinal = totalTimeMinute
	timeSecondFinal = totalTimeSecond

	if totalTimeSecond > 60 {
		diffSecond = totalTimeSecond - 60
		timeMinuteFinal = totalTimeMinute + int(math.Floor(float64(totalTimeSecond)/60))
		timeSecondFinal = diffSecond
	}

	if totalTimeMinute > 60 {
		diffMinute = timeMinuteFinal - 60
		timeHoursFinal = totalTimeHours + int(math.Floor(float64(totalTimeMinute)/60))
		timeMinuteFinal = diffMinute
	}

	timeHoursFinalStr := strconv.Itoa(timeHoursFinal)
	timeMinuteFinalStr := strconv.Itoa(timeMinuteFinal)
	timeSecondFinalStr := strconv.Itoa(timeSecondFinal)

	if timeHoursFinal < 10 {
		timeHoursFinalStr = "0" + strconv.Itoa(timeHoursFinal)
	}
	if timeMinuteFinal < 10 {
		timeMinuteFinalStr = "0" + strconv.Itoa(timeMinuteFinal)
	}
	if timeSecondFinal < 10 {
		timeSecondFinalStr = "0" + strconv.Itoa(timeSecondFinal)
	}

	return timeHoursFinalStr, timeMinuteFinalStr, timeSecondFinalStr
}
