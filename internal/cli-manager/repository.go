package climanager

import (
	"flag"
	"fmt"

	constants "github.com/stonik02/GolangProjectCreator/internal/const"
)

var PathToProjectAndName, PathToCmd, PathToMain,
	PathToInternal, PathToPkg, PathToAppFile, PathToClient, PathToRedisClient, PathToPgClient,
	PathToConfigGoFile, PathToConfigYml, PathToConfigModule string

var Docker, Docker_pg, Docker_redis bool
var Pg_client, Redis_client bool
var Config bool
var Path_to_new_project, Project_name, Module_name string

type cliManager struct {
}

type CliManager interface {
	// Определение флагов
	DefiningFlags()

	// Определение путей
	DefiningPaths()
}

func NewCliManager() CliManager {
	return &cliManager{}
}

// definingFlags implements CliManager.
// Определение флагов
func (cm *cliManager) DefiningFlags() {
	flag.StringVar(&Module_name, "mn", "github.com", "Название модуля проекта")
	flag.StringVar(&Path_to_new_project, "p", "./", "Путь до нового проекта")
	flag.StringVar(&Project_name, "n", "go_project", "Название проекта")
	flag.BoolVar(&Docker, "d", false, "Добавить docker-compose файл")
	flag.BoolVar(&Docker_pg, "dpg", false, "Добавить в docker-compose файл postgreSQL")
	flag.BoolVar(&Docker_redis, "dr", false, "Добавить в docker-compose файл redis")
	flag.BoolVar(&Redis_client, "rc", false, "Создать redis client")
	flag.BoolVar(&Pg_client, "pgc", false, "Создать postgreSQL client")
	flag.BoolVar(&Config, "cfg", false, "Создать config.go файл и подтянуть зависимости")
	flag.Parse()
}

// definingPaths implements CliManager.
// Определение путей для создание папок и файлов
func (cm *cliManager) DefiningPaths() {
	PathToProjectAndName = fmt.Sprintf("%s%s", Path_to_new_project, Project_name)
	fmt.Printf(" PathToProjectAndName = %s \n", PathToProjectAndName)
	PathToCmd = fmt.Sprintf("%s/%s", PathToProjectAndName, constants.Cmd)
	fmt.Printf(" PathToCmd = %s \n", PathToCmd)
	PathToMain = fmt.Sprintf("%s/%s", PathToCmd, constants.Cmd_main)
	fmt.Printf(" PathToMain = %s \n", PathToMain)
	PathToInternal = fmt.Sprintf("%s/%s", PathToProjectAndName, constants.Internal)
	fmt.Printf("PathToInternal = %s \n", PathToInternal)
	PathToPkg = fmt.Sprintf("%s/%s", PathToProjectAndName, constants.Pkg)
	fmt.Printf(" PathToPkg = %s \n", PathToPkg)
	PathToAppFile = fmt.Sprintf("%s/%s", PathToMain, constants.App_golang)
	fmt.Printf(" PathToAppFile = %s \n", PathToAppFile)
	PathToConfigModule = fmt.Sprintf("%s/%s", PathToInternal, "config")
	fmt.Printf(" PathToConfigModule = %s \n", PathToConfigModule)
	PathToConfigGoFile = fmt.Sprintf("%s/%s", PathToConfigModule, constants.Config)
	fmt.Printf(" PathToConfigGoFile = %s \n", PathToConfigGoFile)
	PathToConfigYml = fmt.Sprintf("%s/%s", Path_to_new_project, "config.yml")
	fmt.Printf(" PathToConfigYml = %s \n", PathToConfigYml)
}
