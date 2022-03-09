package devices

import (
	"fmt"
	"github.com/xxxserxxx/gotop/v4/metricapi"
	"os"
)

func init() {
	f := func(cpus map[string]int, l bool) map[string]error {
		//cpuCount, err := CpuCount()
		//if err != nil {
		//	return nil
		//}
		cpuCount := metricapi.CpuCount(os.Getenv("REMOTE_SERVER"))
		formatString := "CPU%1d"
		if cpuCount > 10 {
			formatString = "CPU%02d"
		}
		if cpuCount > 100 {
			formatString = "CPU%03d"
		}
		//vals, err := psCpu.Percent(0, l)
		//if err != nil {
		//	return map[string]error{"gopsutil": err}
		//}
		vals := metricapi.CpuPercent(os.Getenv("REMOTE_SERVER"),l)
		for i := 0; i < len(vals); i++ {
			key := fmt.Sprintf(formatString, i)
			v := vals[i]
			if v > 100 {
				v = 100
			}
			cpus[key] = int(v)
		}
		return nil
	}
	RegisterCPU(f)
}
