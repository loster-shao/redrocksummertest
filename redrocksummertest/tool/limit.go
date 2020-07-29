package tool

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/juju/ratelimit"
)

//TODO 还没加进去，测试过了，应该可以，就是看用户等级多少，等级决定最高上限
func Limit(n int64)  {
	//源容量
	src := bytes.NewReader(make([]byte, 1024*1024))
	//创建
	dst := &bytes.Buffer{}

	//每秒增加100KB的存储桶，最大容量为100KB
	bucket := ratelimit.NewBucketWithRate(500*1024, n*1024)

	start := time.Now()

	//将源复制到目标，但用速率有限的一个包装我们的阅读器
	io.Copy(dst, ratelimit.Reader(src, bucket))

	fmt.Printf("Copied %d bytes in %s\n", dst.Len(), time.Since(start))
}
