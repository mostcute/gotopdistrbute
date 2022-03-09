package metricapi

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/host"
	"io/ioutil"
	"net/http"
)

func SensorsTemperatures(server string)[]host.TemperatureStat{

	url := "http://"+server+"/metric/sensorstemperatures"
	method := "GET"
	result := make([]host.TemperatureStat,0)
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