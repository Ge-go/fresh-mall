package initialize

import "go.uber.org/zap"

func InitLogger() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./test.log",
		"stderr",
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("zap.NewDevelopment err:" + err.Error())
	}
	zap.ReplaceGlobals(logger) // 将logger注册到全局中  后续使用直接使用zap.S() or zap.L() 即可
}
