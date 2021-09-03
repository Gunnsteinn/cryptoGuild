package domain

// ClaimInfo (Marketplace Axie infinity info) response is exported, it models the data we receive.
type ClaimInfo struct {
	Success           bool   `json:"success"`
	ClientID          string `json:"client_id"`
	ItemID            int    `json:"item_id"`
	Total             int    `json:"total"`
	BlockchainRelated struct {
		Signature struct {
			Signature string `json:"signature"`
			Amount    int    `json:"amount"`
			Timestamp int64  `json:"timestamp"`
		} `json:"signature"`
		Balance     int `json:"balance"`
		Checkpoint  int `json:"checkpoint"`
		BlockNumber int `json:"block_number"`
	} `json:"blockchain_related"`
	ClaimableTotal    int   `json:"claimable_total"`
	LastClaimedItemAt int64 `json:"last_claimed_item_at"`
	Item              struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ImageURL    string `json:"image_url"`
		UpdatedAt   int64  `json:"updated_at"`
		CreatedAt   int64  `json:"created_at"`
	} `json:"item"`
}
