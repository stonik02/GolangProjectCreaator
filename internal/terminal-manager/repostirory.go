package terminalManager

import (
	"fmt"
	"os"
	"os/exec"

	cli "github.com/stonik02/GolangProjectCreator/internal/cli-manager"
	constants "github.com/stonik02/GolangProjectCreator/internal/const"

)

type terminalManager struct {
}

type TerminalManager interface {
	commandExecution(name, dir string, arg ...string) error
	CreatingGoModule() error
	GetRedisGo() error
	GetPgx() error
	GetPgxPool() error
	GetCleanEnv() error
}

func NewTerminalManager() TerminalManager {
	return &terminalManager{}
}

// CreatingGoModule implements TerminalManager.
func (t *terminalManager) CreatingGoModule() error {
	err := t.commandExecution(
		cli.PathToProjectAndName, "go", "mod", "init", cli.Module_name,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetPgx implements TerminalManager.
func (tm *terminalManager) GetPgx() error {
	err := tm.commandExecution(cli.PathToProjectAndName, "go", "get", constants.Pgx)
	if err != nil {
		fmt.Printf("Error GetPgxPool: %s", err)
		return err
	}
	return nil
}

// GetPgxPool implements TerminalManager.
func (tm *terminalManager) GetPgxPool() error {
	err := tm.commandExecution(cli.PathToProjectAndName, "go", "get", constants.Pgx_pool)
	if err != nil {
		fmt.Printf("Error GetPgxPool: %s", err)
		return err
	}
	return nil
}

// GetRedisGo implements TerminalManager.
func (tm *terminalManager) GetRedisGo() error {
	err := tm.commandExecution(cli.PathToProjectAndName, "go", "get", constants.Go_redis)
	if err != nil {
		fmt.Printf("Error GetRedisGo: %s", err)
		return err
	}
	return nil
}

// GetCleanEnv implements TerminalManager.
func (tm *terminalManager) GetCleanEnv() error {
	err := tm.commandExecution(cli.PathToProjectAndName, "go", "get", constants.CleanEnv)
	if err != nil {
		fmt.Printf("Error GetCleanEnv: %s", err)
		return err
	}
	return nil
}

// commandExecution implements TerminalManager.
func (*terminalManager) commandExecution(dir, name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
