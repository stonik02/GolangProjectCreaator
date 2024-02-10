package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	constants "github.com/stonik02/GolangProjectCreator/internal/const"
)

const (
	cmd            = "cmd"
	cmd_main       = "main"
	internal       = "internal"
	pkg            = "pkg"
	docker_compose = "docker-compose.yml"
	app_go         = "app.go"
	redis          = "redisclient.go"
	pg             = "pgclient.go"
	client         = "client"
	go_redis       = "github.com/go-redis/redis"
	pgx            = "github.com/jackc/pgx/v4"
	pgx_pool       = "github.com/jackc/pgx/v4/pgxpool"
)

var docker, docker_pg, docker_redis bool
var pg_client, redis_client bool
var path_to_new_project, project_name, module_name string
var pathToProjectAndName, pathToCmd, pathToMain,
	pathToInternal, pathToPkg, pathToAppFile, pathToClient string

func main() {
	// Определение флагов
	definingFlags()

	// Определение путей
	definingPaths()

	// Создание директорий
	creatingDirectory()

	// Создание go модуля
	creatingGoModule()

	// Создание app.go
	creatingAppGo()

	// Создание Docker-compose.yml
	creatingDockerCompose()

	// Создание клиентов
	creatingClients()

}

func creatingClients() {
	if redis_client || pg_client {
		// Создание pkg/client
		pathToClient = fmt.Sprintf("%s/%s", pathToPkg, client)
		err := os.Mkdir(pathToClient, 0777)
		if err != nil {
			fmt.Printf("Error create client: %s \n", err)
			return
		}
	}
	if redis_client {
		creatingRedisClient()
	}
	if pg_client {
		creatingPgClient()
	}
}

func creatingRedisClient() {

	// Создание client/Redis
	pathToRedisClient := fmt.Sprintf("%s/%s", pathToClient, "redis")
	err := os.Mkdir(pathToRedisClient, 0777)
	if err != nil {
		fmt.Printf("Error create redis/client: %s \n", err)
		return
	}

	// Создание redisclient.go
	pathToRedisClientFile := fmt.Sprintf("%s/%s", pathToRedisClient, redis)
	file, err := os.Create(pathToRedisClientFile)
	if err != nil {
		fmt.Printf("Error redisclient.go create: %s \n", err)
		return
	}

	defer file.Close()

	// Заполнение файла
	_, err = file.Write([]byte(constants.Redis_client_file))
	if err != nil {
		fmt.Printf("Error redisclient.go write: %s \n", err)
	}

	// Установка зависимостей
	cmd := exec.Command("go", "get", go_redis)
	cmd.Dir = pathToProjectAndName
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error downloading go-redis package: %s \n", err)
	}
}

func creatingPgClient() {

	// Создание client/psql
	pathToPgClient := fmt.Sprintf("%s/%s", pathToClient, "psql")
	err := os.Mkdir(pathToPgClient, 0777)
	if err != nil {
		fmt.Printf("Error create psql/client: %s \n", err)
		return
	}

	// Создание pgclient.go
	pathToPgClientFile := fmt.Sprintf("%s/%s", pathToPgClient, pg)
	file, err := os.Create(pathToPgClientFile)
	if err != nil {
		fmt.Printf("Error pgclient.go create: %s \n", err)
		return
	}

	defer file.Close()

	// Заполнение файла
	_, err = file.Write([]byte(constants.Pg_client_file))
	if err != nil {
		fmt.Printf("Error pgclient.go write: %s \n", err)
	}

	// Создание repeatable.go
	pathToRepeatableFile := fmt.Sprintf("%s/%s", pathToPgClient, "repeatable.go")
	repeatableFile, err := os.Create(pathToRepeatableFile)
	if err != nil {
		fmt.Printf("Error repeatable.go create: %s \n", err)
		return
	}

	defer repeatableFile.Close()

	// Заполнение repeatable.go файла
	_, err = repeatableFile.Write([]byte(constants.Repeatable_file))
	if err != nil {
		fmt.Printf("Error repeatable.go write: %s \n", err)
	}

	// Установка зависимостей
	cmd := exec.Command("go", "get", pgx_pool)
	cmd.Dir = pathToProjectAndName
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error downloading pgx package: %s \n", err)
	}

	cmd = exec.Command("go", "get", pgx)
	cmd.Dir = pathToProjectAndName
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error downloading pgx package: %s \n", err)
	}

}

func creatingDockerCompose() {
	// Создание docker-compose.yml
	if docker {
		pathToDockerFile := fmt.Sprintf("%s/%s", pathToProjectAndName, docker_compose)
		file, err := os.Create(pathToDockerFile)
		if err != nil {
			fmt.Printf("Error docker-compose create: %s \n", err)
			return
		}

		defer file.Close()

		_, err = file.Write([]byte(constants.Docker_header))
		if err != nil {
			fmt.Printf("Error docker-compose create: %s \n", err)
		}
		if docker_pg {
			_, err = file.Write([]byte(constants.Docker_pg))
			if err != nil {
				fmt.Printf("Error docker-compose pg create: %s \n", err)
			}
		}
		if docker_redis {
			_, err = file.Write([]byte(constants.Docker_redis))
			if err != nil {
				fmt.Printf("Error docker-compose redis create: %s \n", err)
			}
		}
	}
}

func creatingAppGo() {
	// Создание app.go
	file, err := os.Create(pathToAppFile)
	if err != nil {
		fmt.Printf("Error app.go create: %s \n", err)
	}

	defer file.Close()

	file.Write([]byte(constants.App_go))

}

func creatingDirectory() {
	// Создание папки проекта
	err := os.Mkdir(pathToProjectAndName, 0777)
	if err != nil {
		fmt.Printf("Error create project: %s \n", err)
		return
	}

	// Создание папок cmd
	err = os.Mkdir(pathToCmd, 0777)
	if err != nil {
		fmt.Printf("Error create project: %s \n", err)
		return
	}
	// Создание папок main
	err = os.Mkdir(pathToMain, 0777)
	if err != nil {
		fmt.Printf("Error create project: %s \n", err)
		return
	}
	// Создание папок internal
	err = os.Mkdir(pathToInternal, 0777)
	if err != nil {
		fmt.Printf("Error create project: %s \n", err)
		return
	}
	// Создание папок pkg
	err = os.Mkdir(pathToPkg, 0777)
	if err != nil {
		fmt.Printf("Error create project: %s \n", err)
		return
	}
}

func definingFlags() {
	// Определение флагов
	flag.StringVar(&module_name, "mn", "github.com", "Название модуля проекта")
	flag.StringVar(&path_to_new_project, "p", "./", "Путь до нового проекта")
	flag.StringVar(&project_name, "n", "go_project", "Название проекта")
	flag.BoolVar(&docker, "d", false, "Добавить docker-compose файл")
	flag.BoolVar(&docker_pg, "dpg", false, "Добавить в docker-compose файл postgreSQL")
	flag.BoolVar(&docker_redis, "dr", false, "Добавить в docker-compose файл redis")
	flag.BoolVar(&redis_client, "rc", false, "Создать redis client")
	flag.BoolVar(&pg_client, "pgc", false, "Создать postgreSQL client")
	flag.Parse()
}

func definingPaths() {
	// Определение путей для создание папок и файлов
	pathToProjectAndName = fmt.Sprintf("%s/%s", path_to_new_project, project_name)
	pathToCmd = fmt.Sprintf("%s/%s", pathToProjectAndName, cmd)
	pathToMain = fmt.Sprintf("%s/%s", pathToCmd, cmd_main)
	pathToInternal = fmt.Sprintf("%s/%s", pathToProjectAndName, internal)
	pathToPkg = fmt.Sprintf("%s/%s", pathToProjectAndName, pkg)
	pathToAppFile = fmt.Sprintf("%s/%s", pathToMain, app_go)
}

func creatingGoModule() {
	cmd := exec.Command("go", "mod", "init", module_name)
	cmd.Dir = pathToProjectAndName
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error create go module: %s \n", err)
	}

}
