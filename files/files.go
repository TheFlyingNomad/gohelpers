package files

import (
	"fmt"
	"os"

	golog "github.com/brightappsllc/golog"
	gologC "github.com/brightappsllc/golog/contracts"

	reflectionHelpers "github.com/brightappsllc/gohelpers/reflection"
)

// FileOrFolderExists -
func FileOrFolderExists(name string) bool {
	_, error := os.Stat(name)
	return !os.IsNotExist(error)
}

// GetConfigDir -
func GetConfigDir() string {
	directoryName := "." + string(os.PathSeparator) + "config"
	if !FileOrFolderExists(directoryName) {
		err := os.MkdirAll(directoryName, 0755)
		if err != nil {
			golog.Instance().LogFatalWithFields(gologC.Fields{
				"method": reflectionHelpers.GetThisFuncName(),
				"error":  fmt.Sprintf("unable to create directory %s", directoryName),
			})

			panic(err)
		}
	}
	return directoryName
}
