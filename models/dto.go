package models

type Response struct {
	Msg   string `json:"string"` // 错误信息
	Total int64  `json:"total"`  // 分页数量

	Data interface{} `json:"data"` // 返回内容
}

// 配置文件读取
type Config struct {
	Ser   *Server
	Mysql *MysqlInI
	Redis *RedisInI
}

type Server struct {
	HttpPort  int
	PprofPort int
	// consul的健康检查的端口
	HealthPort int
}

type MysqlInI struct {
	MysqlStr string
}

type RedisInI struct {
	Addr string
	Port int
	Db   int
	Pwd  string
}

type ClientInfo struct {
	Addr string
	Port string
}
