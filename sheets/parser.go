package sheets

import (
	"fmt"
	"gopkg.in/Iwark/spreadsheet.v2"
	"log"
	"strconv"
	"strings"
	"time"
)

type Lesson struct {
	lessonNum int
	Lesson    string
}

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

func GetTodaySchedule() []Lesson {
	rowToSearch := getTodayRow()
	sheet := getSheet()
	lessonsArray := make([]Lesson, 0)

	for i := 1; i <= 8; i++ {
		lessons := strings.Split(sheet.Columns[i][rowToSearch].Value, ":")
		if len(lessons[0]) > 0 {
			fmt.Printf("Пара %d: %s\n", i, lessons[0])
			lessonsArray = append(lessonsArray, Lesson{i, lessons[0]})
		}
	}
	return lessonsArray

}

func GetTomorrowSchedule() []Lesson {
	rowToSearch := getTomorrowRow()
	sheet := getSheet()
	lessonsArray := make([]Lesson, 0)

	for i := 1; i <= 8; i++ {
		lessons := strings.Split(sheet.Columns[i][rowToSearch].Value, ":")
		if len(lessons[0]) > 0 {
			fmt.Printf("Пара %d: %s\n", i, lessons[0])
			lessonsArray = append(lessonsArray, Lesson{i, lessons[0]})
		}
	}
	return lessonsArray
}

func getTodayRow() int {
	day := int(time.Now().Weekday())
	weekNum := getCurrentWeekNum()
	daysInWeek := 6
	OFFSET := 2

	row := OFFSET + day + daysInWeek*(weekNum-1) - 1
	return row
}

func getTomorrowRow() int {
	todayRow := getTodayRow()
	if todayRow > 14 {
		return 2
	}
	return todayRow + 1
}

func Test() {
	fmt.Printf("Row is %d", getTodayRow())
	GetTodaySchedule()
}
