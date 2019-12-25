package reader

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

 var configMap=make(map[string] interface{})

func Reader(path string) {
	configFileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}


	if configFileInfo.IsDir() {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		fileInfoArr, _ := file.Readdir(0)
		for index := range fileInfoArr {
			fileInfo := fileInfoArr[index]
			childPath:=file.Name() + string(filepath.Separator) + fileInfo.Name()
			reader(childPath)
		}
	}else {
		content,err:= ioutil.ReadFile(path)
		if err!=nil {
			log.Fatal(err)
		}
		 var mp =make(map[string] interface{})
		err=yaml.Unmarshal(content, &mp)
		if err!=nil {
			log.Fatal(err)
		}
		for key,value :=range mp {
			configMap[key]=value
		}
	}
}
