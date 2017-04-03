package config

import (
	"encoding/json"
	"io/ioutil"
)

/* config.json from https://gist.github.com/border/775526

 */

type ConfigObject struct {
	Object struct {
		BufferSize int `json:"buffer_size"`
		DDA        struct {
			Definition struct {
				Version  int    `json:"version"`
				Filename string `json:"filename"`
				Path     string `json:"path"`
			} `json:"Definition"`
			Source struct {
				Path        string   `json:"path"`
				Filepattern string   `json:"filepattern"`
				Filelist    []string `json:"filelist"`
				Dirorlist   string   `json:"dirorlist"`
			} `json:"Source"`
			Target struct {
				Path        string `json:"path"`
				Name        string `json:"name"`
				Fileorkafka string `json:"fileorkafka"`
				Kafka       struct {
					Server string `json:"server"`
					Topic  string `json:"topic"`
					Port   int    `json:"port"`
				} `json:"kafka"`
			} `json:"Target"`
		} `json:"DDA"`
		Databases []struct {
			Host   string `json:"host"`
			User   string `json:"user"`
			Pass   string `json:"pass"`
			Type   string `json:"type"`
			Name   string `json:"name"`
			Tables []struct {
				Name     string `json:"name"`
				Statment string `json:"statment"`
				Regex    string `json:"regex"`
				Types    []struct {
					ID    string `json:"id"`
					Value string `json:"value"`
				} `json:"Types"`
			} `json:"Tables"`
		} `json:"Databases"`
	} `json:"object"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func GetConfig() *ConfigObject {
	file, e := ioutil.ReadFile("./config.json")
	check(e)

	var jsontype ConfigObject
	json.Unmarshal(file, &jsontype)
	//fmt.Printf("Results: %v\n", jsontype)
	return &jsontype
}
