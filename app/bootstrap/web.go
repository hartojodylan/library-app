package bootstrap

import (
	"github.com/dylanh/library-app/app"
	"github.com/dylanh/library-app/app/listener"
)

// Web Bootstrap http application
func Web() {
	l := &Launcher{}
	l.Add(
		BootFunc(initEnv),
		BootFunc(initConfig),
		BootFunc(initApp),
		BootFunc(func() error {
			initAppInfo()
			return nil
		}),
		BootFunc(app.InitLogger),
		//BootFunc(initI18n),
		// init cache redis connection pool
		//BootFunc(cache.InitCache),
		//BootFunc(mongo.InitMongo),
	)

	l.Run()

	// initEurekaService()

	// listen exit signal
	listener.ListenSignals(onExit)
}
