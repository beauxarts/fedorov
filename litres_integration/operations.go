package litres_integration

type OperationArt struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	ArtType int    `json:"art_type"`
}

type OperationMoney struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type Operations struct {
	Status  int         `json:"status"`
	Error   interface{} `json:"error"`
	Payload struct {
		Pagination struct {
			NextPage     *string `json:"next_page"`
			PreviousPage *string `json:"previous_page"`
		} `json:"pagination"`
		Data []struct {
			Date              string `json:"date"`
			MoneyMovementId   int64  `json:"money_movement_id"`
			MoneyMovementType int    `json:"money_movement_type"`
			SpecificData      struct {
				OperationType      string         `json:"operation_type"`
				Money              OperationMoney `json:"money,omitempty"`
				HasReceipt         bool           `json:"has_receipt"`
				PaymentSystemId    int            `json:"payment_system_id,omitempty"`
				IsAppliedPromoCode bool           `json:"is_applied_promocode,omitempty"`
				IsGift             bool           `json:"is_gift,omitempty"`
				IsSupport          bool           `json:"is_support,omitempty"`
				Arts               []OperationArt `json:"arts,omitempty"`
				Spendings          []struct {
					IsLitresBonuses bool           `json:"is_litres_bonuses"`
					Money           OperationMoney `json:"money"`
				} `json:"spendings,omitempty"`
				Product *string `json:"product,omitempty"`
			} `json:"specific_data"`
		} `json:"data"`
	} `json:"payload"`
}
