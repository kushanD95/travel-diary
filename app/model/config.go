package model

type EnvConfig struct {
	_ struct{}
	Base
	EnvURL  string `gorm:"type:varchar(500);column:ENV_URL"`
	EnvName string `gorm:"type:varchar(10);column:ENV_NAME"`
}

func (EnvConfig) TableName() string {
	return "ENV_CONFIG"
}
