package api

import (
	"math/big"
)

type remoteFileStat struct {
	Name string
	Len  big.Int
}

type remoteDirStat struct {
}

// 删除分布式文件存储系统中的文件
func Delete(remote_file_path string) bool {
	return false
}

// 获取分布式文件存储系统中的文件元数据信息，参考linux
func Stat(remote_file_path string) *remoteFileStat {
	return nil
}

// 在分布式文件存储系统中创建给定的目录路径
func Mkdir(remote_file_path string) bool {
	return false
}

// 将分布式文件系统中的原始路径重命名为新的目标路径
func Rename(rename_src_path, rename_dest_path string) bool {
	return false
}

// 获取分布式文件存储系统中的目录元数据信息
func List(remote_dir_path string) *remoteDirStat {
	return nil
}
