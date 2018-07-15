// Often we want to recursively walk a folder (read the folder's contents, all
// the sub-folders, all the sub-sub-folders, …). To make this easier there's a
// Walk function provided in the path/filepath package:

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}

// The function you pass to Walk is called for every file and folder in the
// root folder. (in this case .)
