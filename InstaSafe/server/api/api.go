package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Transaction struct {
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

var TotalTransactions []Transaction

func HandleTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		fmt.Println("Nowwwww", time.Now().UTC())
		var transaction Transaction

		// if json is inValide
		err := json.NewDecoder(r.Body).Decode(&transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// check if the transaction is in the future
		if transaction.Timestamp.After(time.Now().UTC()) {
			http.Error(w, "Transaction date is in the future", http.StatusUnprocessableEntity)
			return
		}

		//check if the transaction is older than 60 seconds
		if time.Now().Sub(transaction.Timestamp.UTC()) > 60*time.Second {
			http.Error(w, "Transaction is older than 60 seconds", http.StatusNoContent)
			return
		}

		// success case
		TotalTransactions = append(TotalTransactions, transaction)
		fmt.Println("Transactions", TotalTransactions)
		w.WriteHeader(http.StatusOK)
		return
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

type Statistics struct {
	Sum   float64 `json:"sum"`
	Avg   float64 `json:"avg"`
	Max   float64 `json:"max"`
	Min   float64 `json:"min"`
	Count int     `json:"count"`
}

func getStatistics() Statistics {
	stats := Statistics{}
	var sum, avg, max, min float64
	var count int
	for i := 0; i < len(TotalTransactions); i++ {
		//Sum
		sum = sum + TotalTransactions[i].Amount
		count++
		// if time.Now().Sub(TotalTransactions[i].Timestamp).Seconds() <=60{
		// 	sum = sum + TotalTransactions[i].Amount
		// }

		//Avg
		avg = sum / float64(len(TotalTransactions))

		if count == 1 {
			max = TotalTransactions[i].Amount
			min = TotalTransactions[i].Amount
		} else {
			if TotalTransactions[i].Amount > max {
				max = TotalTransactions[i].Amount
			}

			if TotalTransactions[i].Amount < min {
				min = TotalTransactions[i].Amount
			}
		}
		stats.Sum = sum
		stats.Avg = avg
		stats.Max = max
		stats.Min = min
		stats.Count = len(TotalTransactions)
	}
	return stats
}
func HandleStatistics(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		// Check if user location matches requested location
		location := r.URL.Query().Get("location")
		// fmt.Println("Loc", location)

		if userLocation.City != "" {
			if location != userLocation.City {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			stats := getStatistics()
			statsJSON, _ := json.Marshal(stats)
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(statsJSON))
			return
		} else {
			stats := getStatistics()
			statsJSON, _ := json.Marshal(stats)
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(statsJSON))
			return
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
func DeleteTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		TotalTransactions = nil
		w.WriteHeader(http.StatusNoContent)
		return
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

type Location struct {
	City string `json:"city"`
}

var userLocation Location

func SetLocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var location Location
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userLocation = location

	w.WriteHeader(http.StatusCreated)
}
func ResetLocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userLocation = Location{}

	w.WriteHeader(http.StatusNoContent)
}
