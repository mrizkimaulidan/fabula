package instagram

type InstagramInterface interface {
	SetProfileID(profileID string) *Instagram
}

type Instagram struct {
	ProfileID string
}

func SetProfileID(profileID string) *Instagram {
	return &Instagram{
		ProfileID: profileID,
	}
}
