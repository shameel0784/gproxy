package main

type GProxyConf struct {
	Log	   string 		`json:"log"`
	Http   HttpConf		`json:"http"`
}

type HttpConf struct {
	Server		ServerConf     `json:"server"`
	Connection	ConnectionConf `json:"connection"`
}

type ServerConf struct {
	Port 			int				`json:"port"`
	ReadTimeout 	int64			`json:"read_timeout"`
	WriteTimeout	int64			`json:"write_timeout"`
	Locations		[]LocationConf	`json:"location"`
}

type LocationConf struct {
	Path			string	`json:"path"`
	Destination		string	`json:"destination,omitempty"`
	StaticRoot		string	`json:"root,omitempty"`
}

type ConnectionConf struct {
	Timeout				 int64	`json:"timeout"`
	MaxIdleConns        int		`json:"max_idle_conns"`
	IdleConnTimeout     int		`json:"idle_conn_timeout"`
}