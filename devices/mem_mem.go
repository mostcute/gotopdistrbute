package devices

import (
	"github.com/xxxserxxx/gotop/v4/metricapi"
	"os"
)

func init() {
	mf := func(mems map[string]MemoryInfo) map[string]error {
		//mainMemory, err := psMem.VirtualMemory()
		//if err != nil {
		//	return map[string]error{"Main": err}
		//}
		mainMemory := metricapi.VirtualMemory(os.Getenv("REMOTE_SERVER"))
		mems["Main"] = MemoryInfo{
			Total:       mainMemory.Total,
			Used:        mainMemory.Used,
			UsedPercent: mainMemory.UsedPercent,
		}
		return nil
	}
	RegisterMem(mf)
}
