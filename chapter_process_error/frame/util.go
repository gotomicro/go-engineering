package frame

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gotomicro/ego/core/util/xcolor"
)

// MakeReqAndResError 以error级别打印行号、配置名、目标地址、耗时、请求数据、响应数据、错误码和信息
func MakeReqAndResError(line string, compName string, addr string, cost time.Duration, method string, req string, res string, grpcStatus string, stack string) string {
	return fmt.Sprintf("%s %s %s %s & %s => %s & %s \n%s", xcolor.Green(line), xcolor.Red(compName+"@"+addr), xcolor.Yellow(fmt.Sprintf("[%vms]", float64(cost.Microseconds())/1000)), xcolor.Blue(fmt.Sprintf("%v", method)), xcolor.Blue(fmt.Sprintf("%v", req)), xcolor.Blue(fmt.Sprintf("%v", res)), xcolor.Red(grpcStatus), xcolor.Red(stack))
}

// MakeReqAndResInfo 以info级别打印行号、配置名、目标地址、耗时、请求数据、响应数据、错误码和信息
func MakeReqAndResInfo(line string, compName string, addr string, cost time.Duration, method string, req interface{}, reply interface{}, grpcStatus string) string {
	return fmt.Sprintf("%s %s %s %s & %s => %s & %s", xcolor.Green(line), xcolor.Yellow(fmt.Sprintf("[%vms]", float64(cost.Microseconds())/1000)), xcolor.Green(compName+"@"+addr), xcolor.Blue(method), xcolor.Blue(fmt.Sprintf("%v", req)), xcolor.Blue(fmt.Sprintf("%v", reply)), xcolor.Green(grpcStatus))
}

func fileWithLineNum() string {
	for i := 2; i < 20; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if (!strings.HasSuffix(file, ".pb.go") && !strings.Contains(file, "google.golang.org")) || strings.HasSuffix(file, "_test.go") {
			return file + ":" + strconv.FormatInt(int64(line), 10)
		}
	}
	return ""
}
