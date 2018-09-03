package app

var Cfg *runtimConfig

type runtimConfig struct {
	Listen string `json:"listen"`
	Mysql  struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Db       string `json:"database"`
	}
}