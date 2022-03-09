package metricapi

import (
	"encoding/json"
	"fmt"
	psNet "github.com/shirou/gopsutil/net"
	"io/ioutil"
	"net/http"
)

func NetIOCounters(server string)[]psNet.IOCountersStat{

	url := "http://"+server+"/metric/netiocounter"
	method := "GET"
	result := make([]psNet.IOCountersStat,0)
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return result
	}
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