package adapter

import (
	"code.coolops.cn/prometheus-alert-sms/alertMessage"
	"code.coolops.cn/prometheus-alert-sms/utils"
	"encoding/json"
	"fmt"
//	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"net/http"
	"strings"
	"io/ioutil"
	"log"
)

type xxh struct {
	baseUrl string
	userId    string
	receiverIds string
	sendData string
}

func InitXxh(baseUrl,userId,receiverIds string)*xxh{
	return &xxh{
		baseUrl:	baseUrl,
		userId:		userId,
		receiverIds:	receiverIds,
	}
}

func (a xxh)Cmd(sendData alertMessage.AlertMessage) {
        
        for _, alert := range sendData.Alerts{
                a.sendData = utils.FormatData(alert)
                //a.sendData = a.formatData(alert)
		fmt.Println("xxh.sendData:%s",a.sendData)
		fmt.Println("xxh.baseUrl=%s",a.baseUrl)
                SMS := "content=" + a.sendData + "&userId=" + a.userId + "&receiverIds=" + a.receiverIds
		fmt.Println("SMS=%s",SMS)
	        resp, err := http.Post(a.baseUrl,"application/x-www-form-urlencoded",strings.NewReader(SMS))
		fmt.Println(err)
       		if resp.StatusCode == http.StatusOK{
                	log.Println("发送报警信息到XXH成功！！！")
        	}
	        if err != nil {
	            log.Println(err)
	        }
       		defer resp.Body.Close()
        	body, err := ioutil.ReadAll(resp.Body)
        	if err != nil {
	            // handle error
       		}

        	log.Println(string(body))
	}
}

func (a xxh)formatData(sendData alertMessage.AlertMessage)string{
	//alterType := sendData["告警类型"].(string)
	//alterHost := sendData["实例名称"].(string)
	//alterTime := sendData["故障时间"].(string)
	//alterDetails := sendData["告警详情"].(string)
	marshal, err := json.Marshal(sendData)
	if err != nil {
		log.Println("待发送数据转换失败")
		panic(err)
	}
	return string(marshal)
}
