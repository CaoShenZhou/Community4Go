package initialize

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

// 读取配置
func ReadConfig() (*Setting, error) {
	vp := viper.New()
	vp.AddConfigPath("./")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

// 读取配置结构
func (s *Setting) ReadConfigStruct(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
