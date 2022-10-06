package alice

import (
	"AliceULSTUTimetable/sheets"
	"fmt"
	"github.com/azzzak/alice"
	"strings"
)

func AliceHandler(k alice.Kit) *alice.Response {

	req, resp := k.Init()
	if req.IsNewSession() {
		return resp.Text("Привет! Я подскажу тебе расписание. Напиши мне \"Расписание\" или \"Пары\" " +
			"и скажи день, на который ты хочешь узнать расписание")
	}

	var lessons []sheets.Lesson

	if strings.Contains(req.Command(), "расписание") || strings.Contains(req.Command(), "пары") {
		if strings.Contains(req.Command(), "сегодня") {
			lessons = sheets.GetTodaySchedule()
		}

		if strings.Contains(req.Command(), "завтра") {
			lessons = sheets.GetTomorrowSchedule()
		}
		outputString := ""
		for i, lesson := range lessons {
			outputString += fmt.Sprintf("Пара %d: %s\n", i+1, lesson.Lesson)
		}
		return resp.Text(outputString)
	}
	return resp.Text("Я не поняла тебя. Напиши \"Расписание\" или \"Пары\" ")

}
