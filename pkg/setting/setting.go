package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func ReadSetting() (*Setting, error) {
	vp := viper.New()
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	vp.SetConfigName("config")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
