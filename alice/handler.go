package alice

import (
	"github.com/azzzak/alice"
	"log"
	"net/http"
)

func Test() {
	updates := alice.ListenForWebhook("/hook")
	go func() {
		err := http.ListenAndServe(":5000", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	updates.Loop(func(k alice.Kit) *alice.Response {
		req, resp := k.Init()
		if req.IsNewSession() {
			return resp.Text("привет")
		}
		return resp.Text(req.OriginalUtterance())
	})
}
