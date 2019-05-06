package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	golog "github.com/brightappsllc/golog"
	gologC "github.com/brightappsllc/golog/contracts"

	reflectionHelpers "github.com/brightappsllc/gohelpers/reflection"

	fileHelpers "github.com/brightappsllc/gohelpers/files"
)

var configInstance Config
var configOnce sync.Once

// LoadConfig -
func LoadConfig(config Config) Config {
	configOnce.Do(func() {
		configInstance = config.DefaultConfig()

		var configFileName = fileHelpers.GetConfigDir() + string(os.PathSeparator) + "config.json"
		if fileHelpers.FileOrFolderExists(configFileName) {
			file, err := ioutil.ReadFile(configFileName)
			if err != nil {
				golog.Instance().LogWarningWithFields(gologC.Fields{
					"method": reflectionHelpers.GetThisFuncName(),
					"error":  fmt.Sprintf("unable to load config file %s, using default", configFileName),
				})
			} else {
				err := json.Unmarshal(file, configInstance)
				if err != nil {
					golog.Instance().LogWarningWithFields(gologC.Fields{
						"method": reflectionHelpers.GetThisFuncName(),
						"error":  fmt.Sprintf("unable to load config file %s, using default", configFileName),
					})

					configInstance = config.DefaultConfig()
				}
			}
		}
	})

	return configInstance
}
