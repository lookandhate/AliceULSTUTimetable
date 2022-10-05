package alice

import (
	"AliceULSTUTimetable/sheets"
	"fmt"
	"github.com/azzzak/alice"
	"log"
	"net/http"
	"os"
	"strings"
)

func Test() {
	updates := alice.ListenForWebhook("/hook")
	go func() {

		port := fmt.Sprintf(":%s", os.Getenv("PORT"))
		fmt.Println(port)
		err := http.ListenAndServe(port, nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	updates.Loop(func(k alice.Kit) *alice.Response {
		req, resp := k.Init()
		if req.IsNewSession() {
			return resp.Text("привет")
		}
		fmt.Printf(req.Command())
		if strings.Contains(req.Command(), "расписание") || strings.Contains(req.Command(), "пары") {
			lessons := sheets.GetTodaySchedule()
			outputString := ""
			for i, lesson := range lessons {
				outputString += fmt.Sprintf("Пара %d: %s\n", i+1, lesson.Lesson)

			}
			return resp.Text(outputString)
		}
		return resp.Text(req.OriginalUtterance())

	})
}
