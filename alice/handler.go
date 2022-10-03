package alice

import (
	"fmt"
	"github.com/azzzak/alice"
	"log"
	"net/http"
	"os"
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
		return resp.Text(req.OriginalUtterance())
	})
}
