package dto

type RegisterUserResponseDataDto struct{
	UserId string `json:"user_id"`
	UserEmail string `json:"user_email"`
	UserUsername string `json:"user_username"`
}
type RegisterUserResponseDto struct{
	BaseResponseDTO
	Data RegisterUserResponseDataDto
}