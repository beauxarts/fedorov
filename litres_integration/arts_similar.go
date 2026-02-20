package litres_integration

type ArtsSimilar struct {
	Status  int `json:"status"`
	Error   any `json:"error"`
	Payload struct {
		Pagination ArtsPagination    `json:"pagination"`
		Data       []ArtsDetailsData `json:"data"`
	} `json:"payload"`
}
