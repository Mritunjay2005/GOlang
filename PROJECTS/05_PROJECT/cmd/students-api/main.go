package main

import (
	"net/http"
	"log"
	"github.com/Mritunjay2005/students-api/internal/config"
)

func main() {
     //load config
	cfg := config.MustLoad()
	//setip router
	router :=http.NewServeMux()
	router.HandleFunc("GET /",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("wlecome to student api"))
	})
	//setup server
	server:= http.Server{
		Addr :cfg.Addr,
		Handler: router,
	}
	log.Printf("Server started on %s", cfg.HTTPServer.Addr)

	err := server.ListenAndServe()
	
	if err!=nil{
          log.Fatalf("Failed to start server")
	}
	
	//setup database
	
}
