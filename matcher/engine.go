package matcher

import (
	"go.uber.org/zap"
)

// func main() {
// 	logger, _ := zap.NewProduction()
// 	defer logger.Sync()

// 	// 构造异常恢复环境

// }

type MatcherEngine struct {
}

func (me *MatcherEngine) run() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
}
