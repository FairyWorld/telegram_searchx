package searchbot

type TSearchResults struct {
	Keyword string
	Results []*TSearchResult
	Took    int64
}

type TSearchResult struct {
	Seq        int
	ViewLink   string
	SenderName string
	SenderLink string
	ChatName   string
	Date       string
	Content    string
	GoLink     string
}

type TSearchView struct {
	MsgID      int
	ChatID     int64
	ChatType   string
	ChatName   string
	SenderID   int64
	SenderName string
	Date       string
	Content    string
}

type TSearchGoPrivate struct {
	PeerID int64
	MsgID  int64
}
