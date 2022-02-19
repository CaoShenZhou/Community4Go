package setting

import "time"

type ServerSetting struct {
	Port string
}

type DatasourceSetting struct {
	DriverName string
	Host       string
	Port       string
	Database   string
	Username   string
	Password   string
	Charset    string
}

type JwtSetting struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type RedisSetting struct {
	Host     string
	Port     string
	Database string
	Password string
}

type EmailSetting struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

// 读取配置结构
func (s *Setting) ReadConfigStruct(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
