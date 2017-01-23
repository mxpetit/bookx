package response

// Multiple represents a response for multiple results, with links to other
// ressources if needed.
type Multiple struct {
	Links   string        `json:"_link"`
	Results []interface{} `json:"results"`
	Length  int           `json:"length"`
}
