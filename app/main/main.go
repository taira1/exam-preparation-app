package main

import (
	"exam-preparation-app/app/infrastructure"
	"exam-preparation-app/app/routes"
	"flag"
	"log"
	"net/http"
)

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse()
	routes.RegisterTheHandler()
	log.Println("Webサーバーを開始します。ポート: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	defer infrastructure.InfrastructureOBJ.Close()
}
