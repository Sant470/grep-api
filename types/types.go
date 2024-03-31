package apptypes

type SearchReq struct {
	SearchKeyword string `json:"search_keyword"`
	From          int64  `json:"from"`
	To            int64  `json:"to"`
}

type Match struct {
	FilePath string `json:"filepath"`
	Line     int    `json:"line_no"`
	Text     string `json:"text"`
}

type SearchResult struct {
	Matches []Match `json:"matches"`
}
