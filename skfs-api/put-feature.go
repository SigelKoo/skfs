package api

// 将本地文件上传分布式文件存储系统
func Put(local_file_path, remote_file_path string) bool {
	return false
}

// 创建一个空文件并将其打开，返回该文件操作对象
func Create(local_file_path, remote_file_path string) *fileOperation {
	return nil
}

// 基于已打开的文件操作对象进行字符流写入
func Write(operation *fileOperation, local_file_path string) *fileOperation {
	return nil
}

// 关闭文件操作对象
func Close1(*fileOperation) {

}
