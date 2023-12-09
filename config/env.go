package config

const (
	EnvKeyStorageRoot     = "STORAGE_ROOT"
	EnvKeyListenAddress   = "LISTEN_ADDRESS"
	EnvKeyRabbitMQAddress = "RABBITMQ_SERVER"
	DirPath               = "/objects/"
	RouterDataObjects     = "/objects/"
	RouterApiObjects      = "/objects/"
	RouterApiLocate       = "/locate/"
	RouterApiVersions     = "/versions/"

	RabbitExchangeKeyApiServer  = "apiServers"
	RabbitExchangeKeyDataServer = "dataServers"
)
