package main

import (
	"enhance/mux"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
)

/*
*初始化 cpu数量 log文件
 */
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	file, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Error opening process log file: %v", err)
	}
	log.SetOutput(file)
}

func main() {

	log.Println("System Starting now")

	myroute := mux.NewRouter()
	//routes defined at router.go line 92 and handlerfun() at line 35
	for _, route := range routes {

		myroute.HandleFunc(route.URL, handlerFun(route))
	}

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	http.Handle("/", myroute)

	var srv http.Server

	go func() {
		log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(red)))
	}()
	go func() {
		log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
	}()
	select {}
}
