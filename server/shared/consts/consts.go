package consts

const (
	ApiConfigPath  = "server/cmd/api/config.yaml"
	UserConfigPath = "server/cmd/user/config.yaml"
	GameConfigPath = "server/cmd/game/config.yaml"
	DataConfigPath = "server/cmd/data/config.yaml"

	Issuer = "tank_war"

	UserID = "uid"
	User   = "user"
	Room   = "room"

	GameServer = "game_srv"

	HlogFilePath = "./tmp/hlog/logs/"
	KlogFilePath = "./tmp/klog/logs/"

	FreePortAddress = "localhost:0"

	ConsulCheckInterval                       = "7s"
	ConsulCheckTimeout                        = "5s"
	ConsulCheckDeregisterCriticalServiceAfter = "15s"
	ConsulCheckTTL                            = "15s"

	IPFlagName  = "ip"
	IPFlagValue = "0.0.0.0"
	IPFlagUsage = "address"

	PortFlagName  = "port"
	PortFlagUsage = "port"

	MysqlDSN    = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	RabbitMqUrl = "amqp://%s:%s@%s:%d/"
	GameNone    = 0
	GameStart   = 1
	GameOver    = 2
)
