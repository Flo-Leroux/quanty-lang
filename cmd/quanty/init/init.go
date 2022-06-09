package init

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/urfave/cli/v2"
)

const MAIN_TPL = `module Main

component Main {
    main {
        h1 {
            "{{ . }} - Home page!"
        }
    }
}
`

const GITIGNORE_TPL = `
.quanty
`

func Register() *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: "generate new Quanty application named `NAME`",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "no-git",
				Usage: "No initialize application without git",
				Value: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			name := ctx.Args().First()
			if name == "" {
				return errors.New("name not set")
			}

			dirName := createDir(name)
			accessToDir(dirName)

			createMainFile(name)

			noGit := ctx.Bool("no-git")
			if !noGit {
				createGitignore()
				initGit()
			}

			return nil
		},
	}
}

func createDir(name string) string {
	dirName := strings.ToLower(strings.ReplaceAll(name, " ", "-"))

	err := os.Mkdir(dirName, os.ModePerm)
	check(err)

	return dirName
}

func createMainFile(name string) {
	mainPath := filepath.Join("main.qy")

	file, err := os.Create(mainPath)
	check(err)
	defer file.Close()

	mainTpl, err := template.New("main").Parse(MAIN_TPL)
	check(err)

	err = mainTpl.Execute(file, name)
	check(err)
}

func createGitignore() {
	gitignorePath := filepath.Join(".gitignore")

	file, err := os.Create(gitignorePath)
	check(err)
	defer file.Close()

	gitignoreTpl, err := template.New("gitignore").Parse(GITIGNORE_TPL)
	check(err)

	err = gitignoreTpl.Execute(file, nil)
	check(err)
}

func accessToDir(dirName string) {
	err := os.Chdir(dirName)
	check(err)
}

func initGit() {
	gitCmd := exec.Command("git", "init")

	err := gitCmd.Run()
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}
