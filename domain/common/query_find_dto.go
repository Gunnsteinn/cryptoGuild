package common

type QueryFind struct {
	QueryFilters     string `json:"query_filters,omitempty"`
	ProjectionFilter string `json:"projection_filter,omitempty"`
}
