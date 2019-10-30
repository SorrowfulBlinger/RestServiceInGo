package dto

type LoginRequestDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseDto struct {
	StatusCode int8 `json:"statusCode"`
}
