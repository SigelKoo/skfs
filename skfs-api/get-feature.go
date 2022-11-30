package api

import (
	"bufio"
	"math/big"
)

// 将分布式文件存储系统中的文件拉取到本地
func Get(remote_file_path, local_file_path string) bool {
	return false
}

// 打开分布式存储系统中的一个文件，返回该文件操作对象
func Open(remote_file_path string) *fileOperation {
	return nil
}

// 基于已打开的文件操作对象进行字符流读取
func Read(operation *fileOperation) *bufio.Reader {
	return nil
}

// 将已打开的文件操作对象的读取 offset 设定到特定值
func Seek(operation *fileOperation, offset big.Int) {

}

// 基于已打开的文件操作对象进行特定 offset 的字符流读取
func PRead(operation *fileOperation, offset big.Int) *bufio.Reader {
	return nil
}

// 关闭文件操作对象
func Close() {

}
