package app

import (
	"github.com/dylanh/library-app/app/helper"
	"github.com/dylanh/library-app/model"
	"os"
)

// allowed app env name
const (
	EnvProd = "prod"
	EnvPre  = "pre"
	EnvTest = "test"
	EnvDev  = "dev"
)

// for application
const (
	Timezone = "PRC"
	BaseDate = "2006-01-02 15:04:05"

	Timeout     = 10
	PageSize    = 20
	PageSizeStr = "20"
	MaxPageSize = 100

	ConfigSuffix = ".toml"
)

// application info
var (
	EnvName = "dev"
	Name    = "library-app"

	Debug bool

	Hostname string
	RootPath string
	GitInfo  model.GitInfo

	HttpPort = 9440
	// AbsPath always return abs path.
	AbsPath = helper.GetRootPath()
)

// the app work dir path
var WorkDir, _ = os.Getwd()

// IsEnv current env name check
func IsEnv(env string) bool {
	return env == EnvName
}
