package DpkgList

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// PKGDetail is commented
type PKGDetail struct {
	HostName     string `json:"HostName"`
	Name         string `json:"Name"`
	Version      string `json:"Version"`
	Architecture string `json:"Architecture"`
	Description  string `json:"Description"`
}

type pkgDef struct {
	PkgU []struct {
		FieldName string `json:"fieldName"`
		Length    int    `json:"Length"`
	} `json:"pkgU"`
}

func parseField(start int, len int, line string) string {
	return strings.TrimSpace(line[start : start+len])
}

//Parse is commented
func Parse(line string) string { //*[]PKGDetail
	dpkgStruct := `{"PkgU":[{"fieldName":"prefix", "Length": 4},{"fieldName":"Name","Length":41} ,{"fieldName": "Version","Length": 36} ,{"fieldName": "Architecture","Length": 13} ,{"fieldName": "Description","Length": 96}]}`
	var host string
	host, _ = os.Hostname()

	var p pkgDef
	bytes := []byte(dpkgStruct)
	//	var start int
	err := json.Unmarshal(bytes, &p)
	if err != nil {
		panic(err)
	}

	dt := &PKGDetail{
		HostName:     host,
		Name:         parseField(4, 41, line),
		Version:      parseField(45, 36, line),
		Architecture: parseField(81, 13, line),
		Description:  parseField(94, len(line)-94, line),
	}
	dt2, _ := json.Marshal(dt)
	return string(dt2)
}

func ListPkgs(text string) {
	type LinesOfText [][]byte // A slice of byte slices.

	//	var text = "ii  adduser                                  3.113+nmu3ubuntu4                   all          add and remove users and groups"
	s := string(Parse(text))
	c := exec.Command("curl", "-u", "Admin:Admin", "http://192.168.1.107:9200/dpkg/dpkg_list/", "-d", s)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
}
