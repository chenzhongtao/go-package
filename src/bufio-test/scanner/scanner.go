package main // Scanner 提供了一个方便的接口来读取数据，例如读取一个多行文本
// 连续调用 Scan 方法将扫描数据中的“指定部分”，跳过各个“指定部分”之间的数据
// Scanner 使用了缓存，所以“指定部分”的长度不能超出缓存的长度
// Scanner 需要一个 SplitFunc 类型的“切分函数”来确定“指定部分”的格式
// 本包中提供的“切分函数”有“行切分函数”、“字节切分函数”、“UTF8字符编码切分函数”
// 和“单词切分函数”，用户也可以自定义“切分函数”
// 默认的“切分函数”为“行切分函数”，用于获取数据中的一行数据（不包括行尾符）
//
// 扫描在遇到下面的情况时会停止：
// 1、数据扫描完毕，遇到 io.EOF
// 2、遇到读写错误
// 3、“指定部分”的长度超过了缓存的长度
// 如果要对数据进行更多的控制，比如的错误处理或扫描更大的“指定部分”或顺序扫描
// 则应该使用 bufio.Reader

// type Scanner struct {
//	r            io.Reader // The reader provided by the client.
//	split        SplitFunc // The function to split the tokens.
//	maxTokenSize int       // Maximum size of a token; modified by tests.
//	token        []byte    // Last token returned by split.
//	buf          []byte    // Buffer used as argument to split.
//	start        int       // First non-processed byte in buf.
//	end          int       // End of data in buf.
//	err          error     // Sticky error.
//}

// SplitFunc 用来定义“切分函数”类型
// data 是要扫描的数据
// atEOF 标记底层 io.Reader 中的数据是否已经读完
// advance 返回 data 中已处理的数据长度
// token 返回找到的“指定部分”
// err 返回错误信息
// 如果在 data 中无法找到一个完整的“指定部分”
// 则 SplitFunc 返回 (0, nil) 来告诉 Scanner
// 向缓存中填充更多数据，然后再次扫描
//
// 如果返回的 err 是非 nil 值，扫描将被终止，并返回错误信息
//
// 如果 data 为空，则“切分函数”将不被调用
// 意思是在 SplitFunc 中不必考虑 data 为空的情况
//
// SplitFunc 的作用很简单，从 data 中找出你感兴趣的数据，然后返回
// 并告诉调用者，data 中有多少数据你已经处理过了
//type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

// NewScanner 创建一个 Scanner 来扫描 r
// 默认切分函数为 ScanLines
//func NewScanner(r io.Reader) *Scanner

// Err 返回扫描过程中遇到的非 EOF 错误
// 供用户调用，以便获取错误信息
//func (s *Scanner) Err() error

//------------------------------------------------------------

// Bytes 将最后一次扫描出的“指定部分”作为一个切片返回（引用传递）
// 下一次的 Scan 操作会覆盖本次返回的结果
//func (s *Scanner) Bytes() []byte

// Text 将最后一次扫描出的“指定部分”作为字符串返回（值传递）
//func (s *Scanner) Text() string

//------------------------------------------------------------

// Scan 在 Scanner 的数据中扫描“指定部分”
// 找到后，用户可以通过 Bytes 或 Text 方法来取出“指定部分”
// 如果扫描过程中遇到错误，则终止扫描，并返回 false
//func (s *Scanner) Scan() bool
func test_scan() {
	s := strings.NewReader("ABC\nDEF\r\nGHI\nJKL")
	bs := bufio.NewScanner(s)
	for bs.Scan() {
		fmt.Printf("%s %v\n", bs.Bytes(), bs.Text())
	}
	// ABC ABC
	// DEF DEF
	// GHI GHI
	// JKL JKL
}

//------------------------------------------------------------

// Split 用于设置 Scanner 的“切分函数”
// 这个函数必须在调用 Scan 前执行
//func (s *Scanner) Split(split SplitFunc)

func test_split() {
	s := strings.NewReader("ABC DEF GHI JKL")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords)
	for bs.Scan() {
		fmt.Println(bs.Text())
	}
	// ABC
	// DEF
	// GHI
	// JKL
}

//------------------------------------------------------------

// ScanBytes 是一个“切分函数”
// 用来找出 data 中的单个字节并返回
//func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
func test_ScanBytes() {
	s := strings.NewReader("Hello 世界！")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanBytes)
	for bs.Scan() {
		fmt.Printf("%s ", bs.Text())
	}
}

//------------------------------------------------------------

// ScanRunes 是一个“切分函数”
// 用来找出 data 中的单个 UTF8 字符的编码并返回
// 如果 UTF8 解码出错，则返回的 U+FFFD 会被做为 "\xef\xbf\xbd" 返回
// 这使得用户无法区分“真正的U+FFFD字符”和“解码错误的返回值”
//func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)
func test_ScanRunes() {
	s := strings.NewReader("Hello 世界！")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanRunes)
	for bs.Scan() {
		fmt.Printf("%s ", bs.Text())
	} // H e l l o 世 界 ！
}

//------------------------------------------------------------

// ScanLines 是一个“切分函数”
// 用来找出 data 中的单行数据并返回（包括空行）
// 行尾标记可能是 \n 或 \r\n（返回值不包括行尾标记）
//func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)

///------------------------------------------------------------

// ScanWords 是一个“切分函数”
// 用来找出 data 中的单词
// 单词以空白字符分隔，空白字符由 unicode.IsSpace 定义
//func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
func main() {
	test_scan()
	test_ScanBytes()
	test_ScanRunes()
	test_split()

}
