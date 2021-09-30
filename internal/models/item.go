package models

// Item models the items coming back from the hackernews api
type Item struct {
	By          string `bson:"by"`
	Title       string `bson:"title"`
	Url         string `bson:"url"`
	Text        string `bson:"text"`
	ItemType    string `bson:"itemType" json:"type"`
	Descendants int    `bson:"descendants"`
	Id          int    `bson:"id"`
	Score       int    `bson:"score"`
	Time        int    `bson:"time"`
	Parent      int    `bson:"parent"`
	Poll        int    `bson:"poll"`
	Kids        []int  `bson:"kids"`
	Parts       []int  `bson:"parts"`
	Deleted     bool   `bson:"deleted"`
	Dead        bool   `bson:"dead"`
}
