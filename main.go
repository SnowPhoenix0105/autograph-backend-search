package main

import (
	"autograph-backend-search/logging"
	"autograph-backend-search/repository/filesave"
	"autograph-backend-search/repository/neograph"
	"autograph-backend-search/rpc"
	"autograph-backend-search/server"
	"github.com/sirupsen/logrus"
)

const DEBUG = true

func loggingConf() *logging.Config {
	return &logging.Config{
		FileLevel:      logrus.DebugLevel,
		ConsoleLevel:   logrus.InfoLevel,
		FileDir:        "logs",
		DisableConsole: false,
	}
}

func rpcConf() *rpc.Config {
	if DEBUG {
		return rpc.GenerateTestConfig()
	}

	// TODO
	return &rpc.Config{
		ControllerBaseUrl: "",
	}
}

func filesaveConf() *filesave.Config {
	if DEBUG {
		return filesave.GenerateTestConfig()
	}

	// TODO
	return &filesave.Config{
		Host:    "",
		Port:    "",
		TimeOut: 0,
	}
}

func neographConf() *neograph.Config {
	if DEBUG {
		return neograph.GenerateTestConfig()
	}

	// TODO
	return &neograph.Config{Neo4j: neograph.Neo4jConfig{
		Host: "",
		Port: 0,
		User: "",
		Pwd:  "",
	}}
}

func main() {
	logging.SetDefaultConfig(loggingConf())
	logger := logging.NewLogger()

	rpc.Init(rpcConf())

	filesave.Init(filesaveConf())

	neograph.Init(neographConf())
	defer neograph.Close()

	s := server.New(&server.Config{
		Host:      "",
		Port:      8004,
		DebugMode: DEBUG,
	})
	err := s.RunServer()
	if err != nil {
		logger.WithError(err).Errorf("run server error=\n%v", err)
	}
}
