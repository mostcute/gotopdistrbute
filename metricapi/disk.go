package metricapi

import (
	"encoding/json"
	"fmt"
	psDisk "github.com/shirou/gopsutil/disk"
	"io/ioutil"
	"net/http"
)

func Partitions(server string)[]psDisk.PartitionStat{

	url := "http://"+server+"/metric/disk"
	method := "GET"

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return []psDisk.PartitionStat{}
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []psDisk.PartitionStat{}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return []psDisk.PartitionStat{}
	}

	var result []psDisk.PartitionStat
	 json.Unmarshal(body,&result)
	return result

}

func DiskUsage(server ,path string)psDisk.UsageStat{

	url := "http://"+server+"/metric/diskusage"
	method := "GET"

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return psDisk.UsageStat{}
	}
	req.Header.Add("path",path)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return psDisk.UsageStat{}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return psDisk.UsageStat{}
	}

	var result psDisk.UsageStat
	json.Unmarshal(body,&result)
	return result

}


func DiskIOCounters(server ,path string)map[string]psDisk.IOCountersStat{

	url := "http://"+server+"/metric/diskiocounter"
	method := "GET"
	result := make(map[string]psDisk.IOCountersStat)
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return result
	}
	req.Header.Add("path",path)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return result
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return result
	}


	json.Unmarshal(body,&result)
	return result

}