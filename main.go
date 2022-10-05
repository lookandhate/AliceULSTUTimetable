package main

import (
	"AliceULSTUTimetable/sheets"
	"fmt"
	"os"
)
import "AliceULSTUTimetable/alice"

func main() {
	shouldStartAlice := os.Getenv("START_ALICE_SERVICE") != "0"
	sheets.Test()
	if shouldStartAlice {
		fmt.Printf("Starting Alice service")
		alice.Test()
	}
}
