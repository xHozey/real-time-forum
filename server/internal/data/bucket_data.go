package data

import (
	"log"
	"time"
)

func (db *DataLayer) GiveBucket(ip string, tokens int, refill time.Duration, route string) error {
	_, err := db.DataDB.Exec("INSERT OR IGNORE INTO token_bucket (ip, tokens, maxTokens, refill, route) VALUES (?,?,?,?,?)", ip, tokens, tokens, refill, route)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (db *DataLayer) RefillTokens(tokensToAdd int, ip string) {
	db.DataDB.Exec("UPDATE token_bucket SET tokens = MIN(?, maxTokens), lastRefill = ? WHERE ip = ?", tokensToAdd, time.Now(), ip)
}

func (db *DataLayer) ExtractBucketDate(ip string) (int, time.Time) {
	var refill int
	var lastRefill time.Time
	db.DataDB.QueryRow("SELECT refill, lastRefill FROM token_bucket WHERE ip = ?", ip).Scan(&refill, &lastRefill)
	return refill, lastRefill
}

func (db *DataLayer) TakeToken(ip string) bool {
	result, err := db.DataDB.Exec(`
    UPDATE token_bucket 
    SET tokens = tokens - 1 
    WHERE ip = ? AND tokens > 0`, ip)
	if err != nil {
		return false
	}
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected > 0
}
