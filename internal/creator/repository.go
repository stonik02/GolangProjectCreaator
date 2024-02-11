package creator

import (
	"fmt"
	"os"

	cli "github.com/stonik02/GolangProjectCreator/internal/cli-manager"
	constants "github.com/stonik02/GolangProjectCreator/internal/const"
)

type creator struct {
}

type Creator interface {
	// Создание папок
	createDirectory(path string) error

	// Создание папок проекта
	CreateProjectDirectories() error

	CreateClientDirectory() error

	CreatePgClientDirectory() error

	CreateRedisClientDirectory() error

	// Создание файлов
	createFile(path string) (*os.File, error)

	// Создание app.go
	CreateAppGo() error

	// Создание docker-compose.yml
	CreateDockerComposeFile() error

	// Создание файла redisclient.go
	CreateRedisClientFile() error

	// Создание файла pgclient.go
	CreatePgClientFile() error

	// Создание repeatable.go
	CreateRepeatableFile() error

	// Создание config.go
	CreateConfigGoFile() error

	// Создание config.yml
	CreateConfigYmlFile() error
}

func NewCreator() Creator {
	return &creator{}
}

// CreateProjectDirectories implements Creator.
func (c *creator) CreateProjectDirectories() error {
	err := c.createDirectory(cli.PathToProjectAndName)
	if err != nil {
		fmt.Printf("Eror creatingDirectory: %s", err)
		return err
	}
	// Создание папок cmd
	err = c.createDirectory(cli.PathToCmd)
	if err != nil {
		fmt.Printf("Eror creatingDirectory: %s", err)
		return err
	}
	// Создание папок main
	err = c.createDirectory(cli.PathToMain)
	if err != nil {
		fmt.Printf("Error create project: %s \n", err)
		return err
	}
	// Создание папок internal
	err = c.createDirectory(cli.PathToInternal)
	if err != nil {
		fmt.Printf("Error create project: %s \n", err)
		return err
	}
	// Создание папок pkg
	err = c.createDirectory(cli.PathToPkg)
	if err != nil {
		fmt.Printf("Error create project: %s \n", err)
		return err
	}
	return nil
}

// CreateDirectory implements Creator.
// Создает папку с правами 0777
func (*creator) createDirectory(path string) error {
	err := os.Mkdir(path, 0777)
	if err != nil {
		return err
	}
	return nil
}

// CreateAppGo implements Creator.
// Создание и заполнение app.go файла
func (c *creator) CreateAppGo() error {
	file, err := c.createFile(cli.PathToAppFile)
	if err != nil {
		fmt.Printf("Error CreateAppGo : %s \n", err)
		return err
	}
	file.Write([]byte(constants.App_go))
	return nil
}

// CreateDockerComposeFile implements Creator.
// Создание и заполнение docker-compose.yml
func (c *creator) CreateDockerComposeFile() error {

	pathToDockerFile := fmt.Sprintf("%s/%s", cli.PathToProjectAndName, constants.Docker_compose)
	file, err := c.createFile(pathToDockerFile)
	if err != nil {
		fmt.Printf("Error CreateAppGo : %s \n", err)
		return nil
	}

	defer file.Close()
	_, err = file.Write([]byte(constants.Docker_header))
	if err != nil {
		fmt.Printf("Error CreateAppGo write : %s \n", err)
		fmt.Printf("Error docker-compose create: %s \n", err)
	}
	if cli.Docker_pg {
		_, err = file.Write([]byte(constants.Docker_pg))
		if err != nil {
			fmt.Printf("Error CreateAppGo write : %s \n", err)
			fmt.Printf("Error docker-compose pg create: %s \n", err)
		}
	}
	if cli.Docker_redis {
		_, err = file.Write([]byte(constants.Docker_redis))
		if err != nil {
			fmt.Printf("Error CreateAppGo write : %s \n", err)
			fmt.Printf("Error docker-compose redis create: %s \n", err)
		}
	}

	return nil
}

// CreateClientDirectory implements Creator.
func (c *creator) CreateClientDirectory() error {
	cli.PathToClient = fmt.Sprintf("%s/%s", cli.PathToPkg, constants.Client)
	err := c.createDirectory(cli.PathToClient)
	if err != nil {
		fmt.Printf("Error create client: %s \n", err)
		return err
	}
	return nil
}

// CreatePgClientDirectory implements Creator.
func (c *creator) CreatePgClientDirectory() error {
	cli.PathToPgClient = fmt.Sprintf("%s/%s", cli.PathToClient, "psql")
	err := c.createDirectory(cli.PathToPgClient)
	if err != nil {
		fmt.Printf("Error create psql/client: %s \n", err)
		return err
	}
	return nil
}

// CreatePgClientFile implements Creator.
func (c *creator) CreatePgClientFile() error {
	pathToPgClientFile := fmt.Sprintf("%s/%s", cli.PathToPgClient, constants.Pg)
	file, err := c.createFile(pathToPgClientFile)
	if err != nil {
		fmt.Printf("Error pgclient.go create: %s \n", err)
		return err
	}

	defer file.Close()

	// Заполнение файла
	_, err = file.Write([]byte(constants.Pg_client_file))
	if err != nil {
		fmt.Printf("Error pgclient.go write: %s \n", err)
		return err
	}
	return nil
}

// CreateRepeatableFile implements Creator.
func (c *creator) CreateRepeatableFile() error {
	// Создание repeatable.go
	pathToRepeatableFile := fmt.Sprintf("%s/%s", cli.PathToPgClient, "repeatable.go")
	repeatableFile, err := c.createFile(pathToRepeatableFile)
	if err != nil {
		fmt.Printf("Error repeatable.go create: %s \n", err)
		return err
	}

	defer repeatableFile.Close()

	// Заполнение repeatable.go файла
	_, err = repeatableFile.Write([]byte(constants.Repeatable_file))
	if err != nil {
		fmt.Printf("Error repeatable.go write: %s \n", err)
	}
	return nil
}

// CreateRedisClientDirectory implements Creator.
func (c *creator) CreateRedisClientDirectory() error {
	cli.PathToRedisClient = fmt.Sprintf("%s/%s", cli.PathToClient, "redis")
	err := c.createDirectory(cli.PathToRedisClient)
	if err != nil {
		fmt.Printf("Error create redis/client: %s \n", err)
		return err
	}
	return nil
}

// CreateRedisClientFile implements Creator.
func (c *creator) CreateRedisClientFile() error {
	pathToRedisClientFile := fmt.Sprintf("%s/%s", cli.PathToRedisClient, constants.Redis)
	file, err := c.createFile(pathToRedisClientFile)
	if err != nil {
		fmt.Printf("Error redisclient.go create: %s \n", err)
		return err
	}

	defer file.Close()

	// Заполнение файла
	_, err = file.Write([]byte(constants.Redis_client_file))
	if err != nil {
		fmt.Printf("Error redisclient.go write: %s \n", err)
		return err
	}
	return nil
}

// CreateConfigGoFile implements Creator.
func (c *creator) CreateConfigGoFile() error {
	err := c.createDirectory(cli.PathToConfigModule)
	if err != nil {
		fmt.Printf("Error internal/config create: %s \n", err)
		return err
	}

	file, err := c.createFile(cli.PathToConfigGoFile)
	if err != nil {
		fmt.Printf("Error config.go create: %s \n", err)
		return err
	}

	defer file.Close()

	// Заполнение файла
	_, err = file.Write([]byte(constants.Config_Go))
	if err != nil {
		fmt.Printf("Error config.go write: %s \n", err)
		return err
	}
	return nil
}

// CreateConfigYmlFile implements Creator.
func (c *creator) CreateConfigYmlFile() error {
	file, err := c.createFile(cli.PathToConfigYml)
	if err != nil {
		fmt.Printf("Error config.yml create: %s \n", err)
		return err
	}

	defer file.Close()

	return nil
}

// CreateFile implements Creator.
func (*creator) createFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error create file : %s \n", err)
		return nil, err
	}

	return file, nil
}
