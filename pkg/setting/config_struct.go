package setting

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

// 读取配置结构
func (s *Setting) ReadConfigStruct(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
