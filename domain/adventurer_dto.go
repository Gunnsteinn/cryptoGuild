package domain

// Adventurer (Adventurer info) response is exported, it models the data we receive.
type Adventurer struct {
	ClientID          string  `json:"client_id"`
	Total             int     `json:"total"`
	Performance       float64 `json:"performance"`
	ClaimableTotal    int     `json:"claimable_total"`
	LastClaimedItemAt string  `json:"last_claimed_item_at"`
	PlayedDays        float64 `json:"played_days"`
	ClaimedItemAt     string  `json:"claimed_item_at"`
}

type ListOfAdventurer struct {
	User string `json:"user"`
	Data *Adventurer
}

type BodyRequestListOfAdventurer []struct {
	User     string `json:"user"`
	ClientID string `json:"client_id"`
}
