package main

import webview "gooey-app/webview"
import "embed"
import "fmt"
import "log"
import "io/fs"
import "net/http"
// import "os"
import "time"

//go:embed public/*
var EMBED_FS embed.FS

func main() {

	// XXX: Development Mode
	// fsys := os.DirFS("public")

	fsys, _ := fs.Sub(EMBED_FS, "public")
	fsrv := http.FileServer(http.FS(fsys))

	http.Handle("/", fsrv)

	fmt.Println("Listening on http://localhost:3000")

	go func() {

		time.Sleep(1 * time.Second)

		fmt.Println("Opening WebView...")

		view := webview.New(true)
		defer view.Destroy()

		view.SetTitle("gooey-boilerplate")
		view.SetSize(800, 600, webview.HintNone)
		view.Navigate("http://localhost:3000/index.html")
		view.Run()

	}()

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}

}
