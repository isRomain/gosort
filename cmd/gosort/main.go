package main

import (
	"fmt"
	"log"
	"os"
	"gosort/internal/design"
	"gosort/internal/sort"
)

var Folder = "/Users/romain/Downloads/"

func main() {

	fmt.Print(design.Banner)
	var answer string

	for {

		fmt.Println("Are you sure ? (" + design.Red + "y"+ design.Reset +"/"+ design.Blue +"n" + design.Reset + ")")
		fmt.Scanf("%s\n", &answer)

		if "y" == answer||"yes" == answer { break }
		if "n" == answer||"no"  == answer { os.Exit(0) }
	}

	fmt.Println()

	entries, err := os.ReadDir(Folder)
	if err != nil {
		log.Fatal(err)
	}

	sort.SortFiles(entries, Folder)

}
