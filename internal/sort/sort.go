package sort

import (
	"fmt"
	"gosort/internal/config"
	"os"
	"strings"
)

func SortFiles(entries []os.DirEntry, folder string) {

	for _, entry := range entries {

		var path = folder + entry.Name()

		if entry.Type().IsRegular() {

			var extWithPrefix = strings.LastIndex(path, ".")

			if extWithPrefix == -1 {
				continue
			}

			var final = strings.Split(strings.ToLower(path[extWithPrefix:]), ".")

			for p, e := range config.MyDict {
				// p = path/chemin
				// e = extensions
				for i := range e {

					if final[1] == e[i] {
						fmt.Println("File = " + entry.Name())
						os.Rename(path, p+"/"+entry.Name())
					}
				}
			}
		}

	}
}
