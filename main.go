package main

import (
	"server/infrastructure"
	"server/middleware/errorhandler"
	"server/middleware/logger"
)

func main() {
	logger.InitLogger()
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("caught panic: %v", err)
		}
		_ = logger.LoggerSync()
	}()
	conn, err := infrastructure.Open()
	if err != nil {
		logger.Error(err)
	}
	defer conn.Close()
	r, err := infrastructure.SetupServer(conn)
	if err != nil {
		logger.Error(err)
	}

	if err := r.Router.Run(); err != nil {
		logger.Error(errorhandler.MsgErrServer, err)
	}
}
