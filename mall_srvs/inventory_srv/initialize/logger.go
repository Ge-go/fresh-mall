package initialize

import "go.uber.org/zap"

// InitLogger init zap logger
func InitLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("zap.NewDevelopment err:" + err.Error())
	}
	zap.ReplaceGlobals(logger) // 将logger注册到全局中  后续使用直接使用zap.S() or zap.L() 即可
}
