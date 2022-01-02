package hackernews

// List of current top stories
type TopIds []int

// Stuct for importing story for hackernews api
type Story struct {
	By          string `json:"by" bson:"by"`
	Descendants int    `json:"descendants" bson:"descendants"`
	ID          int    `json:"id" bson:"id"`
	Kids        []int  `json:"kids" bson:"kids"`
	Score       int    `json:"score" bson:"score"`
	Time        int    `json:"time" bson:"time"`
	Title       string `json:"title" bson:"title"`
	Type        string `json:"type" bson:"type"`
	URL         string `json:"url" bson:"url"`
}
