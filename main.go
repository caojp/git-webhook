package main

import (
	"flag"
	"fmt"
	"git-webhook/config"
	"git-webhook/git"
	"git-webhook/logger"
	"git-webhook/webhook"
	"net/http"
)

func main() {
	configPath := flag.String("config", "config.yml", "Path to configuration file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		return
	}

	if err := logger.Init(cfg.LogFilePath); err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		return
	}

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		valid, _, err := webhook.VerifySignature(r, cfg)
		if err != nil || !valid {
			fmt.Printf("Failed to verify signature: %v\n", err)
			http.Error(w, "Invalid signature", http.StatusForbidden)
			return
		}

		authStrategy, err := git.GetAuthStrategy(cfg)
		if err != nil || authStrategy.Apply() != nil {
			fmt.Printf("Failed to get auth strategy: %v\n", err)
			http.Error(w, "Authentication failed", http.StatusInternalServerError)
			return
		}

		gitOp, err := git.GetGitOperation(cfg)
		if err != nil || gitOp.Execute() != nil {
			fmt.Printf("Failed to get git operation: %v\n", err)
			http.Error(w, "Git operation failed", http.StatusInternalServerError)
			return
		}

		logger.Logger.Println("Webhook processed successfully")
		w.WriteHeader(http.StatusOK)
	})

	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil)
	if err != nil {
		logger.Logger.Fatal("Server start err: %v \n", err)
		return
	}
}
