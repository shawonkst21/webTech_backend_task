package cmd

import (
	"fmt"
	"net/http"
	"taskManager/handlers"
     "taskManager/midleware"
	"taskManager/util"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("POST /Register",http.HandlerFunc(handlers.Register))
	mux.Handle("POST /Login",http.HandlerFunc(handlers.Login))
	mux.Handle("POST /Reset",http.HandlerFunc(handlers.RequestPasswordReset))
	mux.Handle("POST /ResetPassword",http.HandlerFunc(handlers.ResetPassword))




    mux.Handle("GET /Tasks",middleware.AuthMiddleware(http.HandlerFunc(handlers.GetTasks)))
	mux.Handle("POST /Tasks",middleware.AuthMiddleware(http.HandlerFunc(handlers.CreateTask)))
	mux.Handle("PUT /Tasks/{id}",middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateTask)))
	mux.Handle("DELETE /Tasks/{id}",middleware.AuthMiddleware(http.HandlerFunc(handlers.DeleteTask)))

	fmt.Println("server running on:8080...")
	err:=http.ListenAndServe(":8080",util.GlobalRouter(mux))
	if err != nil {
		fmt.Println("error starting the server", err)
	}
}