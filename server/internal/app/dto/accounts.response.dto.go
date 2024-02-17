package dto


type RegisterUserResponseDataDTO struct{
	UserId string `json:"user_id"`
	UserEmail string `json:"user_email"`
	UserUsername string `json:"user_username"`
}
type RegisterUserResponseDTO struct{
	BaseResponseDTO
	Data RegisterUserResponseDataDTO
}

type ChangePasswordResponseDTO struct{
	BaseResponseDTO
}

type GetAllAccountsRespondDTO struct {
	BaseResponseDTO
}
type GetAuthenticatedUserProfile struct {
	BaseResponseDTO
}
type UpdateProfileResponseDTO struct {
	BaseResponseDTO
}