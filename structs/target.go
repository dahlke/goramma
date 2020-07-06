package structs

// InstagramMedia is the simple struct I put all the information useful to me into.
type InstagramMedia struct {
	ShortCode string `json:"shortcode"`
	Timestamp int    `json:"timestamp"`
	Location  string `json:"location"`
	URL       string `json:"url"`
}
