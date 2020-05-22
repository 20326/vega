package admission

type Config struct {
	CasbinModel string
	TablePrefix string
	LogMode     bool
}

// New db connect to a database
func NewConfig(casbinModel, tablePrefix string, logMode bool) Config {
	return Config{
		CasbinModel: casbinModel,
		TablePrefix: tablePrefix,
		LogMode:     logMode,
	}
}
