package shared

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

func ToFile(text string, filePath string) {

	//write to (new) file
	file, err1 := os.Create(filePath)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer file.Close()
	_, err2 := file.WriteString(text)
	if err2 != nil {
		log.Fatal(err2)
	}
}

//implement sort interface for lines in file
type sortfile [][]byte

func (line sortfile) Swap(i, j int)      { line[i], line[j] = line[j], line[i] }
func (line sortfile) Len() int           { return len(line) }
func (line sortfile) Less(i, j int) bool { return bytes.Compare(line[i], line[j]) < 0 }

func SortFile(filePath string) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	//sort using interface
	lines := bytes.Split(file, []byte{'\n'})
	sort.Sort(sortfile(lines))
	//create file again
	file = bytes.Join(lines, []byte{'\n'})
	return ioutil.WriteFile(filePath, file, 0644)
}

func CmpFiles(filePath1 string, filePath2 string) bool {
	file1, err1 := ioutil.ReadFile(filePath1)
	if err1 != nil {
		log.Fatal(err1)
	}

	file2, err2 := ioutil.ReadFile(filePath2)
	if err2 != nil {
		log.Fatal(err2)
	}
	//string() == string()
	return bytes.Equal(file1, file2)
}
