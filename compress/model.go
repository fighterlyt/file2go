package compress

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

type Model struct {
	data []byte
}

var (
	fileFmt = `package %s

import (
	"bytes"
)
const (
	%s=\U0027%s\U0027
)
type File2GoModel struct{
	buffer *bytes.Buffer
}
func NewFile2GoModel() File2GoModel{
	return File2GoModel{
		buffer:bytes.NewBuffer([]byte(%s)),
	}
}
func(m File2GoModel)Read(p []byte) (n int, err error){
	return m.buffer.Read(p)
}
`
)

func NewModel(filepath, packageName, name string) error {
	fileFmt=strings.Replace(fileFmt,"\\U0027","`",-1)
	if data, err := ioutil.ReadFile(filepath); err != nil {
		return err
	} else {
		if target, err := os.Create(name); err != nil {
			return err
		} else {
			defer target.Close()
			varName := nextSuffix()
			fmt.Fprintf(target, fileFmt, packageName, varName, string(data), varName)
			return nil
		}

	}
}
