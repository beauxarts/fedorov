package litres_integration

type ArtsReviewsData struct {
	Id                      int               `json:"id"`
	Text                    string            `json:"text"`
	UserDisplayName         string            `json:"user_display_name"`
	UserAvatarUrl           *string           `json:"user_avatar_url"`
	UserIsVerified          bool              `json:"user_is_verified"`
	UserId                  int               `json:"user_id"`
	CreatedAt               string            `json:"created_at"`
	LikesRating             int               `json:"likes_rating"`
	LikesCount              int               `json:"likes_count"`
	DislikesCount           int               `json:"dislikes_count"`
	IsLikedByCurrentUser    bool              `json:"is_liked_by_current_user"`
	IsDislikedByCurrentUser bool              `json:"is_disliked_by_current_user"`
	IsRemovedByModerator    bool              `json:"is_removed_by_moderator"`
	ItemRating              *int              `json:"item_rating"`
	Replies                 []ArtsReviewsData `json:"replies"`
	RepliesCount            int               `json:"replies_count"`
	//Caption                 interface{}       `json:"caption"`
	//UserRole                interface{}       `json:"user_role"`
}

type ArtsReviews struct {
	Status  int         `json:"status"`
	Error   interface{} `json:"error"`
	Payload struct {
		Pagination ArtsPagination    `json:"pagination"`
		Data       []ArtsReviewsData `json:"data"`
	} `json:"payload"`
}
