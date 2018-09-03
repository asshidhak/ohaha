package host

type Host struct {
	Env int				`json:"env"`
	Region    int		`json:"region"`
	PrivateIP string	`json:"privateIp"`
}
