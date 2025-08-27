package worker

import (
	"fmt"
	"log"
	"time"

	"github.com/matheushermes/FinGO/internal/database"
	"github.com/matheushermes/FinGO/internal/repository"
	"github.com/matheushermes/FinGO/pkg/utils"
)

func StartCryptoUpdater(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			start := time.Now()
			log.Println("[Worker] Starting crypto price update...")

			if err := updateAllCryptos(start); err != nil {
				log.Printf("[Worker] error updating cryptos: %v\n", err)
			}
		}
	}()
}

func updateAllCryptos(start time.Time) error {
	db, err := database.ConnectDB()
	if err != nil {
		return fmt.Errorf("[Worker] error connecting to DB: %w", err)
	}
	defer db.Close()

	repo := repository.NewCryptosRepository(db)
	cryptos, err := repo.GetAllCryptosAllUsers()
	if err != nil {
		return err
	}

	log.Printf("[Worker] Fetched %d cryptos to update\n", len(cryptos))

	successCount := 0
	failCount := 0

	for i := range cryptos {
		if err := utils.EnrichCryptoWithPrice(&cryptos[i]); err != nil {
			log.Printf("[Worker] error enriching %s (user %d): %v",
				cryptos[i].Symbol, cryptos[i].UserID, err)
			failCount++
			continue
		}

		if err := repo.UpdateCrypto(cryptos[i]); err != nil {
			log.Printf("[Worker] error updating %s (user %d) in DB: %v",
				cryptos[i].Symbol, cryptos[i].UserID, err)
			failCount++
			continue
		}

		successCount++
	}

	duration := time.Since(start)
	log.Printf("[Worker] Finished updating cryptos: %d succeeded, %d failed in %v\n",
		successCount, failCount, duration)

	return nil
}