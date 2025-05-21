package litres_integration

type UserStats struct {
	Status  int         `json:"status"`
	Error   interface{} `json:"error"`
	Payload struct {
		Data struct {
			InProgress     int `json:"in_progress"`
			Wishlist       int `json:"wishlist"`
			Archive        int `json:"archive"`
			Folders        int `json:"folders"`
			Received       int `json:"received"`
			Finished       int `json:"finished"`
			InLibrary      int `json:"in_library"`
			BySubscription int `json:"by_subscription"`
			InCart         int `json:"in_cart"`
			Uploaded       int `json:"uploaded"`
			Authors        int `json:"authors"`
		} `json:"data"`
	} `json:"payload"`
}

func (us *UserStats) Received() int {
	if us.Status == 200 {
		return us.Payload.Data.Received
	}
	return 0
}
