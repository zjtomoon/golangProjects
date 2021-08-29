//io.Reader接口示例
package Reader_interface

import (
	"io"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

// 从标准输入读取
//data, err = ReadFrom(os.Stdin, 11)

// 从普通文件读取，其中 file 是 os.File 的实例
//data, err = ReadFrom(file, 9)

// 从字符串读取
//data, err = ReadFrom(strings.NewReader("from string"), 12)
