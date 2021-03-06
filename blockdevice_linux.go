// +build cgo

package main

import (
	"os"
)

// #include <linux/fs.h>
// #include <sys/ioctl.h>
//
// unsigned long long get_size(int fd) {
//	unsigned long long file_size_in_bytes = 0;
// 	ioctl(fd, BLKGETSIZE64, &file_size_in_bytes);
//	return file_size_in_bytes;
// }
import "C"

func getBlockDeviceSize(f *os.File) int64 {
	size := int64(C.get_size(C.int(f.Fd())))
	return size
}
