package main

import (
	"archive/tar"
	"fmt"
	"os"
)

func main() {
	test_FileInfoHeader()
	test_reader()
	testWriter()
}

func test_FileInfoHeader() {
	fileinfo, err := os.Stat("/home/chenbaoke/test.go")
	if err != nil {
		fmt.Println(err)
	}
	h, err := tar.FileInfoHeader(fileinfo, "")
	h.Linkname = "haha"
	h.Gname = "test"
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(h.AccessTime, h.ChangeTime, h.Devmajor, h.Devminor, h.Gid, h.Gname, h.Linkname, h.ModTime, h.Mode, h.Name, h.Size, h.Typeflag, h.Uid, h.Uname, h.Xattrs)

}

/*
输出结果如下：
2015-08-28 21:26:03.636592126 +0800 CST 2015-08-28 21:26:03.092592112 +0800 CST 0 0 1000 test haha 2015-08-28 21:26:03.092592112 +0800 CST 33206 test.go 581
 48 1000  map[]

由此可见，通过fileinfoheader可以创建tar.header，并自动填写了tar.Header 中的大部分信息，当然，还有一些信息无法从 os.FileInfo 中获取，所以需要你自己去补充，
如Linkname,Gname等。
*/

/*
func NewReader(r io.Reader) *Reader// 从r中创建一个新的reader
func (tr *Reader) Next() (*Header, error)//该函数指向tar文件的下一个实体，在输入的最后返回io.EOF
func (tr *Reader) Read(b []byte) (n int, err error)//该函数读取在tar中当前实体，当读取到实体的结束位置时，返回io.EOF，当调用Next时，读取下一个实体。
*/
func test_reader() {
	f, err := os.Open("/home/chenbaoke/10.tar")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := tar.NewReader(f)
	for hdr, err := r.Next(); err != io.EOF; hdr, err = r.Next() {
		if err != nil {
			fmt.Println(err)
			return
		}
		fileinfo := hdr.FileInfo()
		fmt.Println(fileinfo.Name())
		f, err := os.Create("/home/chenbaoke/develop/" + fileinfo.Name())
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		_, err = io.Copy(f, r)
		if err != nil {
			fmt.Println(err)
		}
	}
}

/*
func NewWriter(w io.Writer) *Writer  //创建一个新的writer，向w中写入。
func (tw *Writer) Close() error //关闭tar归档文件，并将未写入的数据写入底层writer。
func (tw *Writer) Flush() error //完成写当前文件
func (tw *Writer) Write(b []byte) (n int, err error)
func (tw *Writer) WriteHeader(hdr *Header) error//该函数将hdr写入tar文件中，如果hdr不是第一个header，
                   该函数调用flush。在调用close之后在调用该函数就会报错ErrWriteAfterClose。
*/
func testWriter() {
	f, err := os.Create("/home/chenbaoke/10.tar") //创建一个tar文件
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	tw := tar.NewWriter(f)
	defer tw.Close()

	fileinfo, err := os.Stat("/home/chenbaoke/1.go") //获取文件相关信息
	if err != nil {
		fmt.Println(err)
	}
	hdr, err := tar.FileInfoHeader(fileinfo, "")
	if err != nil {
		fmt.Println(err)
	}

	err = tw.WriteHeader(hdr) //写入头文件信息
	if err != nil {
		fmt.Println(err)
		// return
	}

	f1, err := os.Open("/home/chenbaoke/1.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	m, err := io.Copy(tw, f1) //将文件1.go中信息写入压缩包中
	if err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Println(m)
}
