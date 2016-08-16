package main

import (
	"archive/zip"
	//	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func ExampleWriter() {

	fw, err := os.Create("/home/xxx/mem_sub/go/z.zip")
	if err != nil {
		//panic(err)
		fmt.Println(err)
		return
	}
	defer fw.Close()

	// Create a buffer to write our archive to.
	//	buf := new(bytes.Buffer)
	// Create a new zip archive.
	w := zip.NewWriter(fw)
	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}
	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}
}
func ExampleReader() {
	// Open a zip archive for reading.
	r, err := zip.OpenReader("ss.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
	// Output:
	// Contents of README:
	// This is the source code repository for the Go programming language.
}

func CompressZip() {
	const dir = "../img-50/"
	//获取源文件列表
	f, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	fzip, _ := os.Create("img-50.zip")
	w := zip.NewWriter(fzip)
	defer w.Close()
	for _, file := range f {
		fw, _ := w.Create(file.Name())
		filecontent, err := ioutil.ReadFile(dir + file.Name())
		if err != nil {
			fmt.Println(err)
		}
		n, err := fw.Write(filecontent)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	}
}

func DeCompressZip() {
	const File = "img-50.zip"
	const dir = "img/"
	os.Mkdir(dir, 0777) //创建一个目录

	cf, err := zip.OpenReader(File) //读取zip文件
	if err != nil {
		fmt.Println(err)
	}
	defer cf.Close()
	for _, file := range cf.File {
		rc, err := file.Open()
		if err != nil {
			fmt.Println(err)
		}

		f, err := os.Create(dir + file.Name)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		n, err := io.Copy(f, rc)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	}

}
func main() {
	ExampleWriter()
	//ExampleReader()
}
