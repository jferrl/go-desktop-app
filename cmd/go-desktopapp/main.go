package main

import (
	"go-desktopapp/api"
	"log"

	"github.com/zserge/webview"
)

func main() {
	debug := true
	w := webview.New(debug)
	s, err := api.NewServer()
	if err != nil {
		log.Print("Error bootstraping http server")
	}
	go s.Run()
	defer w.Destroy()
	w.SetTitle("Golang desktop app")
	w.SetSize(800, 600, webview.HintNone)

	err = w.Bind("add", func(a, b int) int {
		return a + b
	})
	if err != nil {
		log.Print("Error binding func")
	}
	err = w.Bind("printObject", func(obj struct{ ID string }) {
		log.Print(obj)
	})
	if err != nil {
		log.Print("Error binding func")
	}

	w.Navigate("http://localhost" + s.Port() + "/")
	w.Run()
}
