package gmap

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func LoadLevelMapData(nameMap string) Map {
	jsonFile, err := os.Open("C:\\Users\\Sam\\Desktop\\go\\d_game\\maps\\" + nameMap + "\\" + nameMap + ".json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(jsonFile)

	data, err := io.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}

	var result Map

	jsonErr := json.Unmarshal(data, &result)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return result
}