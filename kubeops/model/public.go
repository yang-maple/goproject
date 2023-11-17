package model

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"strconv"
	"time"
)

// GetAge 获取node 时间戳
func GetAge(CTT int64) string {
	age := time.Now().Unix() - CTT
	switch {
	case age < 60:
		//秒
		return strconv.FormatInt(age, 10) + "s"
	case age < 3600:
		//分
		return strconv.FormatInt(age/60, 10) + "m"
	case age < 86400:
		//时
		return strconv.FormatInt(age/3600, 10) + "h"
	case age < 31536000:
		//天
		return strconv.FormatInt(age/86400, 10) + "d"
	default:
		//年
		return strconv.FormatInt(age/31536000, 10) + "y"
	}

}

// GetImage 获取image 列表
func GetImage(pod corev1.PodSpec) []string {
	image := make([]string, 0)
	for _, im := range pod.Containers {
		image = append(image, im.Image)
	}
	return image
}

// GetRestart  获取实例 重启次数
func GetRestart(status corev1.PodStatus) int32 {
	var count int32
	for _, v := range status.ContainerStatuses {
		count = count + v.RestartCount
	}
	return count
}

// GetResourcesRequests 获取实例的 Cpu 和 Memory 资源request 总量
func GetResourcesRequests(name string, spec corev1.PodSpec) string {
	if name == "Cpu" {
		var cpuTotal int64
		for _, v := range spec.Containers {
			cpuTotal = cpuTotal + v.Resources.Requests.Cpu().MilliValue()
		}
		if cpuTotal == 0 {
			return strconv.Itoa(int(cpuTotal))
		}
		return strconv.Itoa(int(cpuTotal)) + "m"
	} else if name == "Memory" {
		var memoryTotal int64
		for _, v := range spec.Containers {
			memoryTotal = memoryTotal + v.Resources.Requests.Memory().Value()
		}
		// 判断单位是 M 还是Mi 整除1000000的是M 否则是Mi
		if memoryTotal%1000000 == 0 {
			fmt.Println(spec.Containers[0].Name, ":", memoryTotal)
			switch {
			case memoryTotal == 0:
				return strconv.Itoa(int(memoryTotal))
			case memoryTotal < 1000000000:
				return strconv.Itoa(int(memoryTotal)/1000000) + "M"
			case memoryTotal < 1000000000000:
				return strconv.Itoa(int(memoryTotal)/1000000000) + "G"

			}
		} else {
			switch {
			case memoryTotal == 0:
				return strconv.Itoa(int(memoryTotal))
			case memoryTotal < 1073741824:
				return strconv.Itoa(int(memoryTotal)/1048576) + "Mi"
			case memoryTotal < 1099511627776:
				return strconv.Itoa(int(memoryTotal)/1073741824) + "Gi"
			}
		}
	}
	return ""
}
