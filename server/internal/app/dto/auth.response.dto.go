package dto


type ObtainTokenResponseDataDTO struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type ObtainTokenResponseDTO struct {
	BaseResponseDTO
	Data ObtainTokenResponseDataDTO
	}