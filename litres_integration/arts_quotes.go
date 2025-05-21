package litres_integration

type ArtsPagination struct {
	NextPage     string  `json:"next_page"`
	PreviousPage *string `json:"previous_page"`
}

type ArtsQuotesData struct {
	ArtsUUId
	Nickname      string `json:"nickname"`
	UserId        int    `json:"user_id"`
	UserUrl       string `json:"user_url"`
	SelectionUrl  string `json:"selection_url"`
	SelectionHtml string `json:"selection_html"`
	VotesGood     int    `json:"votes_good"`
	VotesBad      int    `json:"votes_bad"`
	UserReaction  string `json:"user_reaction"`
}

type ArtsQuotes struct {
	Status  int         `json:"status"`
	Error   interface{} `json:"error"`
	Payload struct {
		Pagination ArtsPagination   `json:"pagination"`
		Data       []ArtsQuotesData `json:"data"`
	} `json:"payload"`
}
