package DpkgList

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jtfogarty/dpkg-es/config"
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

// ListPkgs is commented
func ListPkgs(text string, configObj config.Object) {
	type LinesOfText [][]byte // A slice of byte slices.

	//	var text = "ii  adduser                                  3.113+nmu3ubuntu4                   all          add and remove users and groups"
	var esUser = configObj.Config.Databases.Elastic01.User
	var esPW = configObj.Config.Databases.Elastic01.Pass
	var esHost = configObj.Config.Databases.Elastic01.Host
	var esPort = configObj.Config.Databases.Elastic01.Port
	var esIndex = configObj.Config.Databases.Elastic01.Index[0].Name
	var esMap = configObj.Config.Databases.Elastic01.Index[0].MappingName
	var curlU = esUser + ":" + esPW
	var curlURL = esHost + ":" + esPort + "/" + esIndex + "/" + esMap + "/"

	s := string(Parse(text))
	c := exec.Command("curl", "-u", curlU, curlURL, "-d", s)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
}
