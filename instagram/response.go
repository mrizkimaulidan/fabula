package instagram

type InstagramProfile struct {
	Users []struct {
		Position int `json:"position"`
		User     struct {
			Pk                         string        `json:"pk"`
			Username                   string        `json:"username"`
			FullName                   string        `json:"full_name"`
			IsPrivate                  bool          `json:"is_private"`
			PkID                       string        `json:"pk_id"`
			ProfilePicURL              string        `json:"profile_pic_url"`
			ProfilePicID               string        `json:"profile_pic_id"`
			IsVerified                 bool          `json:"is_verified"`
			HasAnonymousProfilePicture bool          `json:"has_anonymous_profile_picture"`
			HasHighlightReels          bool          `json:"has_highlight_reels"`
			HasOptEligibleShop         bool          `json:"has_opt_eligible_shop"`
			AccountBadges              []interface{} `json:"account_badges"`
			FriendshipStatus           struct {
				Following       bool `json:"following"`
				IsPrivate       bool `json:"is_private"`
				IncomingRequest bool `json:"incoming_request"`
				OutgoingRequest bool `json:"outgoing_request"`
				IsBestie        bool `json:"is_bestie"`
				IsRestricted    bool `json:"is_restricted"`
				IsFeedFavorite  bool `json:"is_feed_favorite"`
			} `json:"friendship_status"`
			LatestReelMedia    int         `json:"latest_reel_media"`
			LiveBroadcastID    interface{} `json:"live_broadcast_id"`
			ShouldShowCategory bool        `json:"should_show_category"`
		} `json:"user,omitempty"`
	}
}
