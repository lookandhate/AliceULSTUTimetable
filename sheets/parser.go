package sheets

import (
	"fmt"
	"gopkg.in/Iwark/spreadsheet.v2"
	"log"
	"strconv"
	"time"
)

func getSheet() *spreadsheet.Sheet {
	service, err := spreadsheet.NewService()
	spreadsheetID := "1-3opnTX8_ZLIm83FTA__9iyMWSIqbYkJwvGhkfxpnsg"
	if err != nil {
		log.Fatal(err)
	}
	spreadsheet, err := service.FetchSpreadsheet(spreadsheetID)
	if err != nil {
		log.Fatal(err)
	}
	sheet, err := spreadsheet.SheetByID(0)
	if err != nil {
		log.Fatal(err)
	}
	return sheet
}

func getCurrentWeekNum() int {
	sheet := getSheet()

	weekNum, err := strconv.Atoi(sheet.Rows[1][10].Value)
	if err != nil {
		log.Fatal(err)
	}
	return weekNum

}

func GetTodaySchedule() {
	rowToSearch := getTodayRow()
	sheet := getSheet()

	for i := 1; i <= 8; i++ {
		fmt.Printf("Пара %d: %s\n", i, sheet.Columns[i][rowToSearch].Value)

	}

}

func getTodayRow() int {
	day := int(time.Now().Weekday())
	weekNum := getCurrentWeekNum()
	daysInWeek := 6
	OFFSET := 2

	row := OFFSET + day + daysInWeek*(weekNum-1) - 1
	return row
}

func Test() {
	fmt.Printf("Row is %d", getTodayRow())
	GetTodaySchedule()
}
