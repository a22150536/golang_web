package alert

import (
	"errors"
	"fmt"
)

func GetServiceInfo(message AlertMessage) []DeploymentInfo {
	var deploymentInfo = []DeploymentInfo{}
	var hpainfo = HpaInfo{}
	var service, namespace string
	for _, alert := range message.Alerts {
		for k, v := range alert.Label {
			if k == "deployment" {
				service = v
			} else if k == "namespace" {
				namespace = v
			} else if k == "hpa" {
				hpainfo.Name = v
			}
		}
		deploymentInfo = append(deploymentInfo, DeploymentInfo{Name: service, Namespace: namespace, Hpa: hpainfo})
	}
	return deploymentInfo
}

func ScaleService(info []DeploymentInfo) (err error) {
	errmessage := ""
	for _, v := range info {
		err = v.GetDeployment()
		fmt.Println("errrr11")
		if err != nil {
			errmessage += v.Name + "scale失敗"
			err = errors.New(errmessage)
			errmessage += ","
			continue
		}
		fmt.Println("errrr12")
		err = v.GetHPAResource()
		fmt.Println("errrr13")
		if err != nil {
			errmessage += v.Name + "scale失敗"
			err = errors.New(errmessage)
			errmessage += ","
			continue
		}
		fmt.Println("errrr14")
		err = v.SetHPAResource()
		if err != nil {
			errmessage += v.Name + "scale失敗"
			err = errors.New(errmessage)
			errmessage += ","
			continue
		}
	}
	fmt.Println("errrr7")
	return
}
