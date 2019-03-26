package config

// Config - Generate entropy from different sources
type Config interface {
	CreateOrReturnInstance() Config
	DefaultConfig() Config
	String() string
}
