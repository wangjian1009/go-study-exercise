package conf

type Conf struct {
	Mysql struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Db       string `yaml:"db"`
	}
	Redis struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
	Token struct {
		PublicKey string `yaml:"publicKey"`
	}
	Kafka struct {
		Host string `yaml:"host"`
	}
}
