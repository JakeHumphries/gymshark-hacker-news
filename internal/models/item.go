package models

// Item models the items coming back from the hackernews api
type Item struct {
	By          string  `bson:"by" json:"by"`
	Title       string  `bson:"title" json:"title"`
	Url         string  `bson:"url" json:"url"`
	Text        string  `bson:"text" json:"text"`
	Type        string  `bson:"type" json:"type"`
	Descendants int32   `bson:"descendants" json:"descendants"`
	Id          int32   `bson:"id" json:"id"`
	Score       int32   `bson:"score" json:"score"`
	Time        int32   `bson:"time" json:"time"`
	Parent      int32   `bson:"parent" json:"parent"`
	Poll        int32   `bson:"poll" json:"poll"`
	Kids        []int32 `bson:"kids" json:"kids"`
	Parts       []int32 `bson:"parts" json:"parts"`
	Deleted     bool    `bson:"deleted" json:"deleted"`
	Dead        bool    `bson:"dead" json:"dead"`
}
