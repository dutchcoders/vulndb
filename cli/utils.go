package cli

import (
	"fmt"
	"os"
	"os/user"
	"path"
)

func defaultDbFile() string {
	return path.Join(defaultBaseDir(), "db.bleve")

}

func defaultBaseDir() string {
	user, err := user.Current()
	check(err)

	return path.Join(user.HomeDir, ".vulndb")
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
