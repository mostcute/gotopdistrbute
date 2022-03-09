package metricapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetNvidiaInfo(server string)[]byte{

	url := "http://"+server+"/metric/nvidiainfo"
	method := "GET"
	result := []byte{}
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