package litres_integration

type ArtsId struct {
	Id int `json:"id"`
}

type ArtsUUId struct {
	UUId string `json:"uuid"`
}

type ArtsName struct {
	Name string `json:"name"`
}

type ArtsUrl struct {
	Url string `json:"url"`
}

type ArtsPrice struct {
	FinalPrice              float64  `json:"final_price"`
	InappPrice              *float64 `json:"inapp_price"`
	InappBasePrice          *float64 `json:"inapp_base_price"`
	FullPrice               float64  `json:"full_price"`
	DiscountPrice           *float64 `json:"discount_price"`
	LitcoinCount            int      `json:"litcoin_count"`
	InappName               *string  `json:"inapp_name"`
	InappBaseName           *string  `json:"inapp_base_name"`
	Currency                string   `json:"currency"`
	DiscountPercent         *float64 `json:"discount_percent"`
	CostWithBonus           float64  `json:"cost_with_bonus"`
	BonusMoney              float64  `json:"bonus_money"`
	AccountMoney            float64  `json:"account_money"`
	CanBePaidByAccountMoney bool     `json:"can_be_paid_by_account_money"`
}

type ArtsPerson struct {
	ArtsId
	ArtsUUId
	ArtsUrl
	FullName  string `json:"full_name"`
	FullRodit string `json:"full_rodit"`
	Role      string `json:"role"`
}

type ArtsRating struct {
	UserRating      *int    `json:"user_rating"`
	Rated1Count     int     `json:"rated_1_count"`
	Rated2Count     int     `json:"rated_2_count"`
	Rated3Count     int     `json:"rated_3_count"`
	Rated4Count     int     `json:"rated_4_count"`
	Rated5Count     int     `json:"rated_5_count"`
	RatedAvg        float64 `json:"rated_avg"`
	RatedTotalCount int     `json:"rated_total_count"`
}

type ArtsIdLinkType struct {
	ArtsId
	LinkType string `json:"link_type"`
}

type ArtsSeries struct {
	ArtsId
	ArtsName
	ArtsUrl
	ArtsUUId
	ArtsCount int `json:"arts_count"`
	ArtOrder  int `json:"art_order"`
	//TopArts   []interface{} `json:"top_arts"`
}

type ArtsLabels struct {
	IsBestseller      bool `json:"is_bestseller"`
	IsLitresExclusive bool `json:"is_litres_exclusive"`
	IsNew             bool `json:"is_new"`
	IsSalesHit        bool `json:"is_sales_hit"`
}

type ArtsTag struct {
	ArtsId
	ArtsName
	ArtsUrl
	ArtsUUId
	IsRoot bool `json:"is_root"`
}

type ArtsCopyright struct {
	ArtsId
	ArtsName
	ArtsUrl
}

type ArtsAdditionalInfo struct {
	RegisteredAt           string `json:"registered_at"`
	TranslatedAt           string `json:"translated_at"`
	CurrentPagesOrSeconds  int    `json:"current_pages_or_seconds"`
	ExpectedPagesOrSeconds *int   `json:"expected_pages_or_seconds"`
}

type ArtsSynchronization struct {
	ArtsId
	TextArtId  int `json:"text_art_id"`
	AudioArtId int `json:"audio_art_id"`
}

type ArtsYoutubeVideo struct {
	ArtsUrl
	Type string `json:"type"`
}

type ArtsArt struct {
	ArtsId
	ArtsUUId
	ArtsUrl
	CoverUrl                    string       `json:"cover_url"`
	Title                       string       `json:"title"`
	CoverRatio                  float64      `json:"cover_ratio"`
	IsDraft                     bool         `json:"is_draft"`
	ArtType                     int          `json:"art_type"`
	Prices                      ArtsPrice    `json:"prices"`
	IsAutoSpeechGift            bool         `json:"is_auto_speech_gift"`
	Availability                int          `json:"availability"`
	IsFree                      bool         `json:"is_free"`
	Persons                     []ArtsPerson `json:"persons"`
	MyArtStatus                 int          `json:"my_art_status"`
	IsSubscriptionArt           bool         `json:"is_subscription_art"`
	IsAvailableWithSubscription bool         `json:"is_available_with_subscription"`
	ReleaseFileId               int          `json:"release_file_id"`
	PreviewFileId               *int         `json:"preview_file_id"`
	FirstTimeSaleAt             string       `json:"first_time_sale_at"`
	LinkedArtGroup              string       `json:"linked_art_group"`
	IsAutoSpeech                bool         `json:"is_auto_speech"`
	CanBePreOrdered             bool         `json:"can_be_preordered"`
	IsHidden                    bool         `json:"is_hidden"`
	//LibraryInformation          interface{}  `json:"library_information"`
}

type ArtsDetailsData struct {
	ArtsArt
	Subtitle                          string                `json:"subtitle"`
	MinAge                            int                   `json:"min_age"`
	IsAdultContent                    bool                  `json:"is_adult_content"`
	CoverHeight                       int                   `json:"cover_height"`
	CoverWidth                        int                   `json:"cover_width"`
	ForeignPublisherId                *string               `json:"foreign_publisher_id"`
	LanguageCode                      string                `json:"language_code"`
	SymbolsCount                      int                   `json:"symbols_count"`
	ExpectedSymbolsCount              *int                  `json:"expected_symbols_count"`
	PodcastSerialNumber               *int                  `json:"podcast_serial_number"`
	LastUpdatedAt                     string                `json:"last_updated_at"`
	LastReleasedAt                    string                `json:"last_released_at"`
	ReadPercent                       int                   `json:"read_percent"`
	IsFinished                        bool                  `json:"is_finished"`
	IsPromo                           bool                  `json:"is_promo"`
	IsAvailableWithLitresSubscription bool                  `json:"is_available_with_litres_subscription"`
	ArtSubscriptionStatusForUser      string                `json:"art_subscription_status_for_user"`
	IsAbonementArt                    bool                  `json:"is_abonement_art"`
	IsAvailableWithAbonement          bool                  `json:"is_available_with_abonement"`
	IsPreorderNotifiedForUser         bool                  `json:"is_preorder_notified_for_user"`
	IsExclusiveAbonement              bool                  `json:"is_exclusive_abonement"`
	IsDrm                             bool                  `json:"is_drm"`
	InGifts                           int                   `json:"in_gifts"`
	IsFourthArtGift                   bool                  `json:"is_fourth_art_gift"`
	PodcastLeftToBuy                  *bool                 `json:"podcast_left_to_buy"`
	AvailableFrom                     string                `json:"available_from"`
	IsLiked                           bool                  `json:"is_liked"`
	Rating                            ArtsRating            `json:"rating"`
	LinkedArts                        []ArtsArt             `json:"linked_arts"`
	Series                            []ArtsSeries          `json:"series"`
	DateWrittenAt                     string                `json:"date_written_at"`
	Labels                            ArtsLabels            `json:"labels"`
	IsArchived                        bool                  `json:"is_archived"`
	InteractedAt                      string                `json:"interacted_at"`
	ReadAt                            string                `json:"read_at"`
	PurchasedAt                       string                `json:"purchased_at"`
	SynchronizedArts                  []ArtsId              `json:"synchronized_arts"`
	Synchronizations                  []ArtsSynchronization `json:"synchronizations"`
	AlternativeVersion                ArtsIdLinkType        `json:"alternative_version"`
	InUserCartStatus                  string                `json:"in_user_cart_status"`
	IsDraftFullFree                   bool                  `json:"is_draft_full_free"`
	IsSubscribedToPodcast             bool                  `json:"is_subscribed_to_podcast"`
	IsSubscribedToAudioDraft          bool                  `json:"is_subscribed_to_audio_draft"`
	IsTrustedCover                    bool                  `json:"is_trusted_cover"`
	ArtRefreshInDays                  int                   `json:"art_refresh_in_days"`
	ImagesCount                       int                   `json:"images_count"`
	IsFilesPendingForDownload         bool                  `json:"is_files_pending_for_download"`
	IsEpubPendingForDownload          bool                  `json:"is_epub_pending_for_download"`
	IsAbonementForbidden              bool                  `json:"is_abonement_forbidden"`
	HTMLAnnotation                    string                `json:"html_annotation"`
	HTMLAnnotationLitres              string                `json:"html_annotation_litres"`
	IsAzimytTrial                     bool                  `json:"is_azimyt_trial"`
	IsPdwTrial                        bool                  `json:"is_pdw_trial"`
	IsDiscountForbidden               bool                  `json:"is_discount_forbidden"`
	ReviewsCount                      int                   `json:"reviews_count"`
	LivelibRatedCount                 int                   `json:"livelib_rated_count"`
	LivelibRatedAvg                   float64               `json:"livelib_rated_avg"`
	Tags                              []ArtsTag             `json:"tags"`
	Genres                            []ArtsTag             `json:"genres"`
	ISBN                              string                `json:"isbn"`
	PublicationDate                   string                `json:"publication_date"`
	QuotesCount                       int                   `json:"quotes_count"`
	YoutubeVideos                     []ArtsYoutubeVideo    `json:"youtube_videos"`
	ContentsUrl                       string                `json:"contents_url"`
	IsProhibitedForSale               bool                  `json:"is_prohibited_for_sale"`
	AdditionalInfo                    ArtsAdditionalInfo    `json:"additional_info"`
	Publisher                         *ArtsCopyright        `json:"publisher"`
	Copyrighter                       ArtsCopyright         `json:"copyrighter"`
	Rightholders                      []ArtsCopyright       `json:"rightholders"`
	UpdateEvents                      []string              `json:"update_events"`
	//InFolders                         interface{}           `json:"in_folders"`
	//AvailabilityWithPartnerSubscriptions []interface{}         `json:"availability_with_partner_subscriptions"`
	//ParentPodcastName                    *string               `json:"parent_podcast_name"`
	//ParentPodcastId                      *string               `json:"parent_podcast_id"`
	//ParentPodcastUrl                     *string               `json:"parent_podcast_url"`
	//PodcastEpisodesCount                 *int                  `json:"podcast_episodes_count"`
	//IsPodcastComplete   bool               `json:"is_podcast_complete"`
	//IsReferrable                         bool                  `json:"is_referrable"`
	//Referral                             interface{}           `json:"referral"`
	//AlisaUrl                             interface{}           `json:"alisa_url"`
}

type ArtsDetails struct {
	Status  int         `json:"status"`
	Error   interface{} `json:"error"`
	Payload struct {
		Data ArtsDetailsData `json:"data"`
	} `json:"payload"`
}
