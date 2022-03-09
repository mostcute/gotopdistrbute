package metricapi

import (
	"encoding/json"
	"fmt"
	psMem "github.com/shirou/gopsutil/mem"
	"io/ioutil"
	"net/http"
)

func VirtualMemory(server string)*psMem.VirtualMemoryStat{

	url := "http://"+server+"/metric/mem"
	method := "GET"
	result := psMem.VirtualMemoryStat{}
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return &result
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return &result
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return &result
	}


	json.Unmarshal(body,&result)
	return &result

}

func VirtualMemorySwap(server string)*psMem.SwapMemoryStat{

	url := "http://"+server+"/metric/memswap"
	method := "GET"
	result := psMem.SwapMemoryStat{}
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return &result
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return &result
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return &result
	}


	json.Unmarshal(body,&result)
	return &result

}