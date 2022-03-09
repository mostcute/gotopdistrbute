// +build !freebsd

package devices

import (
	"github.com/xxxserxxx/gotop/v4/metricapi"
	"os"
)

func init() {
	mf := func(mems map[string]MemoryInfo) map[string]error {
		//memory, err := psMem.SwapMemory()
		//if err != nil {
		//	return map[string]error{"Swap": err}
		//}
		memory := metricapi.VirtualMemorySwap(os.Getenv("REMOTE_SERVER"))
		mems["Swap"] = MemoryInfo{
			Total:       memory.Total,
			Used:        memory.Used,
			UsedPercent: memory.UsedPercent,
		}
		return nil
	}
	RegisterMem(mf)
}
