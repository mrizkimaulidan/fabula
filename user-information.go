package main

type UserInformation struct {
	Result struct {
		User struct {
			FollowerCount                                int  `json:"follower_count"`
			MediaCount                                   int  `json:"media_count"`
			FollowingCount                               int  `json:"following_count"`
			PublicEmail                                string  `json:"public_email"`
			Pk                                         string  `json:"pk"`
			Username                                   string  `json:"username"`
			FullName                                   string  `json:"full_name"`
		} `json:"user"`
		Status string `json:"status"`
	} `json:"result"`
}