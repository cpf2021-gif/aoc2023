package main

import (
	"fmt"
	"os"
	"text/template"
	"flag"
)

func main() {
	var dayNamePrefix = "Day"
	var fileNamePrefix = "day"
	var dirNamePrefix = "../day"

	var num string
	flag.StringVar(&num, "n", "", "day number")
	flag.Parse()

	data := struct {
		Day string
		Filename string
	}{
		Day: dayNamePrefix + num,
		Filename: fileNamePrefix + num,
	}

	testdata := struct {
		Day string
	}{
		Day: dayNamePrefix + num,
	}

	// 创建文件夹
	dirName := dirNamePrefix + "/" + data.Filename
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// dayx.go生成
	templateFile := "day.tpl"
	teml, err := template.ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}


	mainFilename := data.Filename + ".go"
	newFile, err := os.Create(dirName + "/" + mainFilename)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	err = teml.Execute(newFile, data)
	if err != nil {
		panic(err)
	}

	// dayx_test.go生成
	templateFile = "test.tpl"
	teml, err = template.ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}

	testFilename := data.Filename + "_test.go"
	newFile, err = os.Create(dirName + "/" + testFilename)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	err = teml.Execute(newFile, testdata)
	if err != nil {
		panic(err)
	}

	// 生成data文件
	datadir := dirName + "/data"
	err = os.Mkdir(datadir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	datafilenames := []string{data.Filename + "_input.txt", data.Filename + "_q1_test.txt", data.Filename + "_q2_test.txt"}
	for _, n := range datafilenames {
		func() {
			newFile, err = os.Create(datadir + "/" + n)
			if err != nil {
				panic(err)
			}
			defer newFile.Close()
		}()
	}

	fmt.Println("done!!")
}