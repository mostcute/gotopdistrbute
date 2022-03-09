package metricapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
type Proc struct {
	Pid         int
	CommandName string
	FullCommand string
	CPU         float64
	Mem         float64
}
func GetProcs(server string)[]byte{

	url := "http://"+server+"/metric/proc"
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


	//json.Unmarshal(body,&result)
	return body

}