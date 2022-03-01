package model

type (
	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginResponses struct {
		AccessToken         string `json:"access_token"`
		RefreshToken        string `json:"refresh_token"`
		AccessTokenExpires  int64  `json:"token_expires"`
		RefreshTokenExpires int64  `json:"refresh_expires"`
		AccessTokenUUID     string `json:"access_uuid"`
		RefreshTokenUUID    string `json:"refresh_uuid"`
	}

	TokenArgs struct {
		ID    int64
		Name  string
		Email string
	}

	TokenResponse struct {
		AccessUUID  string
		RefreshUUID string
		ID          int64
		Fullname    string
		Email       string
		Type        string
	}
)
