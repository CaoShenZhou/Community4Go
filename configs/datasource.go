package configs

type Datasource struct {
	DriverName string // 驱动名
	Host       string // IP
	Port       string // 端口
	Database   string // 数据库
	Username   string // 用户名
	Password   string // 密码
	Config     string // 配置
}
