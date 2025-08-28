package cmd

import (
	"fmt"
	"net/http"
	"taskManager/handlers"
	"taskManager/util"
)

func Serve() {
	mux := http.NewServeMux()
	
    mux.Handle("GET /Tasks",http.HandlerFunc(handlers.GetTasks))
	mux.Handle("POST /Tasks",http.HandlerFunc(handlers.CreateTask))
	mux.Handle("PUT /Tasks/{id}",http.HandlerFunc(handlers.UpdateTask))
	mux.Handle("DELETE /Tasks/{id}",http.HandlerFunc(handlers.DeleteTask))

	fmt.Println("server running on:8080...")
	err:=http.ListenAndServe(":8080",util.GlobalRouter(mux))
	if err != nil {
		fmt.Println("error starting the server", err)
	}
}