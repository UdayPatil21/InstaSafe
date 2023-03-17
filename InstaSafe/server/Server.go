package main

import (
	"GoLang/src/InstaSafe/server/api"
	"fmt"
	"net/http"
)

func main() {
	// http.
	http.HandleFunc("/transaction", api.HandleTransactions)
	http.HandleFunc("/statistics", api.HandleStatistics)
	http.HandleFunc("/deletetransaction", api.DeleteTransactions)
	http.HandleFunc("/location", api.SetLocationHandler)
	http.HandleFunc("/location/reset", api.ResetLocationHandler)
	fmt.Println("Server --->> 8085")
	http.ListenAndServe(":8085", nil)
}
