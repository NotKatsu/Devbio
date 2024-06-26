package main

import (
	"devbio/database"
	"devbio/endpoints"
	ReturnModule "devbio/modules/return_module"
	"net/http"

	"github.com/pterm/pterm"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pterm.Success.Println(r.Method, "Request received at", r.URL.Path)

		isRateLimited, _ := database.IsRatelimited(r.Header.Get("session"))

		if isRateLimited {
			ReturnModule.MethodNotAllowed(w, r)
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func main() {
	if database.InitializeStorageDatabase() && database.CreateStorageTables() {
		pterm.Success.Println("Storage database has been initialized successfully.")

		if database.InitializeDatabase() && database.CreateTables() {
			pterm.Success.Println("Database has been initialized successfully.")

			http.HandleFunc("/api/", endpoints.ManageApi)

			http.HandleFunc("/api/explore", endpoints.ManageExplore)

			http.HandleFunc("/api/account", endpoints.ManageAccounts)
			http.HandleFunc("/api/account/session", endpoints.ManageSessions)
			http.HandleFunc("/api/account/update", endpoints.ManageUpdate)
			http.HandleFunc("/api/account/statistics", endpoints.ManageStatistics)
			http.HandleFunc("/api/account/connections", endpoints.ManageConnections)
			http.HandleFunc("/api/account/connections/github", endpoints.ManageGithubConnectionsCallback)
			http.HandleFunc("/api/account/connections/callback", endpoints.ManageConnectionsCallback)

			http.HandleFunc("/api/notifications", endpoints.ManageNotifications)

			http.HandleFunc("/api/storage/profile/icon/", endpoints.ManageIcon)
			http.HandleFunc("/api/storage/profile/banner/", endpoints.ManageBanner)

			http.HandleFunc("/api/subscriptions/subscribe", endpoints.ManageSubscribe)

			pterm.Error.Println(http.ListenAndServe(":6969", corsMiddleware(http.DefaultServeMux)))
		} else {
			pterm.Fatal.WithFatal(true).Println("Failed to initialize database successfully.")
		}
	} else {
		pterm.Fatal.WithFatal(true).Println("Failed to initialize storage database successfully.")
	}
}
