package dto

type EnvConfigReq struct {
	_   struct{}
	Env string `json:"env" validate:"required,max=10"`
}
