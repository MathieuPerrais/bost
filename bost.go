package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

func slugify(str string) (slug string) {
	str = strings.ToLower(str)
	quoteless := strings.Replace(str, "'", "", -1)
	return strings.Replace(quoteless, " ", "-", -1)
}

const extension = "md"

const template = "---\nlayout: post\ntitle: %s\npermalink: %s\ntags:\n---\n"

var directory string

func main() {
	app := cli.NewApp()
	app.Name = "bost"
	app.Version = "0.1.0"
	app.Usage = "interact with jekyll posts"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "directory, d",
			Value:       "_posts",
			Usage:       "directory where to look for posts",
			Destination: &directory,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "create a blog post",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "force, f",
					Usage: "force the file creation",
				},
			},
			Action: func(c *cli.Context) {
				const layout = "2006-01-02"
				now := time.Now()
				today := now.Format(layout)
				title := c.Args().First()
				slug := slugify(title)
				if len(c.Args()) == 2 {
					title = c.Args()[1]
				}
				filename := fmt.Sprintf("%s-%s.%s", today, slug, extension)
				path := strings.Join([]string{directory, filename}, "/")
				if _, err := os.Stat(path); os.IsNotExist(err) || c.Bool("force") {
					ioutil.WriteFile(path, []byte(fmt.Sprintf(template, title, slug)), 0644)
					println(path)
				} else {
					println(path, ": file already exists")
				}
			},
		},
		{
			Name:    "open",
			Aliases: []string{"o"},
			Usage:   "open blog post lazily",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "editor, e",
					Usage:  "editor used to open a post",
					EnvVar: "EDITOR",
				},
			},
			Action: func(c *cli.Context) {
				query := c.Args().First()
				cmd := "find"
				args := []string{fmt.Sprintf("%s/", directory), "-name", fmt.Sprintf("*%s*", query)}
				out, err := exec.Command(cmd, args...).Output()
				if err != nil {
					os.Exit(1)
				}
				files := strings.Split(string(out), "\n")
				vi := exec.Command(c.String("editor"), files[0])
				vi.Stdin = os.Stdin
				vi.Stdout = os.Stdout
				vi.Stderr = os.Stderr
				vi.Run()
			},
		},
	}

	app.Run(os.Args)
}
