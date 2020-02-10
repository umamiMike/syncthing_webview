package main

import (
	"github.com/zserge/webview"
)

func main() {
	// Open wikipedia in a 800x600 resizable window
	webview.Open("Minimal webview example", "http://localhost:8384", 800, 600, true)

	//ln, err := net.Listen("tcp", "127.0.0.1:0")
	//if err != nil {
	//log.Fatal(err)
	//}
	//defer ln.Close()
	//go func() {
	//// Set up your http server here
	//log.Fatal(http.Serve(ln, nil))
	//}()
	//webview.Open("Hello", "http://"+ln.Addr().String(), 400, 300, false)
}
