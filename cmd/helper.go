package cmd

import (
	"fmt"
	"os"
	"sort"
)

func exists(dirName string) (bool, error) {
	_, err := os.Stat(dirName)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func mkDirs(dirName string) {
	status, err := exists(dirName)
	check(err)
	if status {
		fmt.Println("Specified directory already exists")
		os.Exit(0)
	} else {
		err = os.Mkdir(dirName, 0766)
		check(err)
	}
}

func sortedCopy(s []string) []string {
	c := make([]string, len(s))
	copy(c, s)
	sort.Slice(c, func(i, j int) bool {
		return c[i] < c[j]
	})
	return c
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
