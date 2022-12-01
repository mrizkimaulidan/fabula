package instagram

type InstagramProfile struct {
	Users []struct {
		Position int `json:"position"`
		User     struct {
			Pk                         string        `json:"pk"`
			Username                   string        `json:"username"`
			ProfilePicURL              string        `json:"profile_pic_url"`
			ProfilePicID               string        `json:"profile_pic_id"`
			IsVerified                 bool          `json:"is_verified"`
			IsPrivate                  bool          `json:"is_private"`
			PkID                       string        `json:"pk_id"`
			FullName                   string        `json:"full_name"`
			HasAnonymousProfilePicture bool          `json:"has_anonymous_profile_picture"`
			HasHighlightReels          bool          `json:"has_highlight_reels"`
			HasOptEligibleShop         bool          `json:"has_opt_eligible_shop"`
			AccountBadges              []interface{} `json:"account_badges"`
			LatestReelMedia            int           `json:"latest_reel_media"`
			LiveBroadcastID            interface{}   `json:"live_broadcast_id"`
			ShouldShowCategory         bool          `json:"should_show_category"`
		} `json:"user,omitempty"`
	} `json:"users"`
	Places           []interface{} `json:"places"`
	Hashtags         []interface{} `json:"hashtags"`
	HasMore          bool          `json:"has_more"`
	RankToken        string        `json:"rank_token"`
	ClearClientCache interface{}   `json:"clear_client_cache"`
	Status           string        `json:"status"`
}
