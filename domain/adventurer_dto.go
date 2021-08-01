package domain

// Adventurer (Adventurer info) response is exported, it models the data we receive.
type Adventurer struct {
	ClientID          string `json:"client_id"`
	Total             int    `json:"total"`
	ClaimableTotal    int    `json:"claimable_total"`
	LastClaimedItemAt string `json:"last_claimed_item_at"`
	ClaimedItemAt     string `json:"claimed_item_at"`
}
