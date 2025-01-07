package data

import "time"

func (db *DataLayer) GiveBucket(ip string, tokens int, refill time.Duration) {
	db.DataDB.Exec("INSERT OR IGNORE INTO bucketToken (ip, tokens, maxTokens, refill) VALUES (?,?,?,?)", ip, tokens, tokens, refill)
}

func (db *DataLayer) RefillTokens(tokensToAdd int, ip string) {
	db.DataDB.Exec("UPDATE tokenBucket SET tokens = MIN(?, maxTokens) WHERE ip = ?", tokensToAdd, ip)
}

func (db *DataLayer) ExtractBucketDate(ip string) (time.Duration, time.Time) {
	var refill time.Duration
	var lastRefill time.Time
	db.DataDB.QueryRow("SELECT refill, lastRefill FROM bucketToken WHERE ip = ?", ip).Scan(&refill, &lastRefill)
	return refill, lastRefill
}

func (db *DataLayer) TakeToken(ip string) bool {
	result, err := db.DataDB.Exec(`
    UPDATE tokenBucket 
    SET tokens = tokens - 1 
    WHERE ip = ? AND tokens > 0`, ip)
	if err != nil {
		return false
	}
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected > 0
}
