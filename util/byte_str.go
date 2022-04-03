/**
 * @Author: quqiang
 * @Email: 77347042@qq.com
 * @Version: 1.0.0
 * @Date: 2021/3/3 1:14 下午
 */

package util

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
	EB = 1 << (10 * iota)
)

func FormatFileSize(fileSize float64) (size string) {
	if fileSize < 1024 {
		return fmt.Sprintf("%.2fB", fileSize/float64(1))
	} else if fileSize < (MB) {
		return fmt.Sprintf("%.2fKB", fileSize/float64(KB))
	} else if fileSize < (GB) {
		return fmt.Sprintf("%.2fMB", fileSize/float64(MB))
	} else if fileSize < (TB) {
		return fmt.Sprintf("%.2fGB", fileSize/float64(GB))
	} else if fileSize < (PB) {
		return fmt.Sprintf("%.2fTB", fileSize/float64(TB))
	} else if fileSize < (EB){
		return fmt.Sprintf("%.2fPB", fileSize/float64(PB))
	} else {
		return fmt.Sprintf("%.2fEB", fileSize/float64(EB))
	}
}
