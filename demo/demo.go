// Copyright (c) 2021 xin053
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/xin053/hsperfdata"
)

func main() {
	filePaths, err := hsperfdata.DataPathsByProcessName("java")
	if err != nil {
		log.Fatal(err)
	}

	for pid := range filePaths {
		entryMap, err := hsperfdata.ReadPerfData(filePaths[pid], true)
		if err != nil {
			log.Fatal("open fail", err)
		}

		var keys []string
		for k := range entryMap {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		for _, key := range keys {
			fmt.Printf("%s=%v\n", key, entryMap[key])
		}
	}
}
