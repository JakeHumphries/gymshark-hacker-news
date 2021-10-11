package models

// Item models the items coming back from the hackernews api
type Item struct {
	By          string `bson:"by" json:"by"`
	Title       string `bson:"title" json:"title"`
	Url         string `bson:"url" json:"url"`
	Text        string `bson:"text" json:"text"`
	ItemType    string `bson:"itemType" json:"type"`
	Descendants int    `bson:"descendants" json:"descendants"`
	Id          int    `bson:"id" json:"id"`
	Score       int    `bson:"score" json:"score"`
	Time        int    `bson:"time" json:"time"`
	Parent      int    `bson:"parent" json:"parent"`
	Poll        int    `bson:"poll" json:"poll"`
	Kids        []int  `bson:"kids" json:"kids"`
	Parts       []int  `bson:"parts" json:"parts"`
	Deleted     bool   `bson:"deleted" json:"deleted"`
	Dead        bool   `bson:"dead" json:"dead"`
}
