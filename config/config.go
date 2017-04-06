package config

import (
	"encoding/json"
	"io/ioutil"
)

/* config.json from https://gist.github.com/border/775526
ConfigObject
*/

// Object comment
type Object struct {
	Config struct {
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
		Databases struct {
			Elastic01 struct {
				Host  string `json:"host"`
				Port  string `json:"port"`
				User  string `json:"user"`
				Pass  string `json:"pass"`
				Type  string `json:"type"`
				Name  string `json:"name"`
				Index []struct {
					Name               string `json:"name"`
					MappingDefEncoding string `json:"mapping_def_encoding"`
					MappingName        string `json:"mapping_name"`
					MappingDef         string `json:"mapping_def"`
				} `json:"Index"`
			} `json:"Elastic01"`
			Mysql01 struct {
				Host   string `json:"host"`
				Port   string `json:"port"`
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
			} `json:"Mysql01"`
		} `json:"Databases"`
	} `json:"config"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//GetConfig ConfigObject comment
func GetConfig() *Object {
	file, e := ioutil.ReadFile("./config.json")
	check(e)

	var jsontype Object
	json.Unmarshal(file, &jsontype)
	return &jsontype
}
