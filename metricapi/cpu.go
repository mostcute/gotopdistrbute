package metricapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


var HOST_CHANGED bool

func CpuCount(server string)int{

	url := "http://"+server+"/metric/cpucount"
	method := "GET"
	result := int(0)
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

func CpuPercent(server string,percpu bool)[]float64{

	url := "http://"+server+"/metric/cpupercent"
	method := "GET"
	result := make([]float64,0)
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return result
	}

	if percpu {
		req.Header.Add("percpu","true")
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