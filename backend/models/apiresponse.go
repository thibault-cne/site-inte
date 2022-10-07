package models

type GoogleApiResponse struct {
	Iss            string `json:"iss"`
	Nbf            string `json:"nbf"`
	Aud            string `json:"aud"`
	Sub            string `json:"sub"`
	Hd             string `json:"hd"`
	Email          string `json:"email"`
	Email_verified bool   `json:"email_verified"`
	Azp            string `json:"azp"`
	Name           string `json:"name"`
	Picture        string `json:"picture"`
	Given_name     string `json:"given_name"`
	Family_name    string `json:"family_name"`
	Iat            string `json:"iat"`
	Exp            string `json:"exp"`
	Jti            string `json:"jti"`
	Alg            string `json:"alg"`
	Kid            string `json:"kid"`
	Typ            string `json:"typ"`
}

type LoginApiResponse struct {
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}

type ProfileDataResponse struct {
	Username    string   `json:"username"`
	Points      int      `json:"points"`
	Color       string   `json:"color"`
	Users_stars []*Stars `json:"users_stars"`
}

type StarsResponse struct {
	Created_at        int64  `json:"created_at"`
	Id                int    `json:"id"`
	Giver_name        string `json:"giver_name"`
	Receiver_name     string `json:"receiver_name"`
	Type              int    `json:"type"`
	Message           string `json:"message"`
	Moderation_status bool   `json:"moderation_status"`
}

type AllUsersResponse struct {
	Name string `json:"name"`
}

type AllUserWithPointsResponse struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

func NewProfileDataResponse(user *User) *ProfileDataResponse {
	stars, err := GetStars(user.ID)

	if err != nil {
		panic(err)
	}

	return &ProfileDataResponse{
		Username:    user.Name,
		Points:      user.Points,
		Color:       user.Color,
		Users_stars: stars,
	}
}

func NewStarsResponse(stars []*Stars) []*StarsResponse {
	out := make([]*StarsResponse, len(stars))

	for i, star := range stars {
		giver, _ := GetUser(star.Giver_id)
		receiver, _ := GetUser(star.Receiver_id)

		out[i] = &StarsResponse{
			Created_at:        star.CreatedAt.Unix(),
			Id:                star.ID,
			Giver_name:        giver.Name,
			Receiver_name:     receiver.Name,
			Type:              star.Type,
			Message:           star.Message,
			Moderation_status: star.Moderation_status,
		}
	}

	return out
}

func NewAllUsersResponse(users []*User) []*AllUsersResponse {
	out := make([]*AllUsersResponse, len(users))

	for i, user := range users {
		out[i] = &AllUsersResponse{
			Name: user.Name,
		}
	}

	return out
}

func NewAllUsersWithPointsResponse(users []*User) []*AllUserWithPointsResponse {
	out := make([]*AllUserWithPointsResponse, len(users))

	for i, user := range users {
		out[i] = &AllUserWithPointsResponse{
			Name:   user.Name,
			Points: user.Points,
		}
	}

	return out
}
