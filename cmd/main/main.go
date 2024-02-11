package main

import (
	cli "github.com/stonik02/GolangProjectCreator/internal/cli-manager"
	"github.com/stonik02/GolangProjectCreator/internal/creator"
	tm "github.com/stonik02/GolangProjectCreator/internal/terminal-manager"

)

func main() {
	// Создание флагов и заполнение путей
	cliManager := cli.NewCliManager()
	cliManager.DefiningFlags()
	cliManager.DefiningPaths()

	// Создание директорий проекта
	creator := creator.NewCreator()
	err := creator.CreateProjectDirectories()
	if err != nil {
		return
	}

	// Создание модуля проекта
	terminalManager := tm.NewTerminalManager()
	err = terminalManager.CreatingGoModule()
	if err != nil {
		return
	}

	// Создание app.go
	err = creator.CreateAppGo()
	if err != nil {
		return
	}

	// Создание Docker-compose.yml
	if cli.Docker {
		creator.CreateDockerComposeFile()
	}

	// Создание клиентов
	creatingClients(creator, terminalManager)

	if cli.Config {
		// Установка зависимостей
	terminalManager.GetCleanEnv()

	// Создание config.go файла
	creator.CreateConfigGoFile()

	// Создание config.yml файла
	creator.CreateConfigYmlFile()
	}

}

func creatingClients(creator creator.Creator, terminalMagager tm.TerminalManager) {
	if cli.Redis_client || cli.Pg_client {
		err := creator.CreateClientDirectory()
		if err != nil {
			return
		}
	}
	if cli.Redis_client {
		creator.CreateRedisClientDirectory()
		creator.CreateRedisClientFile()
		terminalMagager.GetRedisGo()
	}
	if cli.Pg_client {
		creator.CreatePgClientDirectory()
		creator.CreatePgClientFile()
		creator.CreateRepeatableFile()
		terminalMagager.GetPgx()
		terminalMagager.GetPgxPool()
	}
}
