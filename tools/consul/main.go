package main

import (
	"fmt"
	capi "github.com/hashicorp/consul/api"
)

func main() {
	// 定义consul 配置信息

	config := capi.Config{
		Address:  "10.1.131.31:8500",
		WaitTime: 0,
		Token:    "74564183-951d-27e3-493e-bb57bfbb3892",
		//Token:     "9cfa076f-a143-5293-f96f-e8b772b1a2db",
		TLSConfig: capi.TLSConfig{},
	}
	client, err := capi.NewClient(&config)
	if err != nil {
		panic(err)
	}

	//agent := client.Agent()
	//agent.Reload()
	//getagent(client)
	//Register(client, "10.1.131.31", 19991)
	DelRegister(client, "node-exporter-10.1.131.35")
	// Get a handle to the KV API
	//kv := client.KV()
	//
	//// PUT a new KV pair
	//p := &capi.KVPair{Key: "REDIS_MAXCLIENTS", Value: []byte("1000")}
	//_, err = kv.Put(p, nil)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Lookup the pair
	//pair, _, err := kv.Get("upstreams/sso-proxy/1008611", nil)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
	//// create a services
	//
	//// delete agent
	//agent.Services()

	//// 解析consul 中的值
	//aa := new(struct {
	//	Weight      int `json:"weight"`
	//	MaxFails    int `json:"max_fails"`
	//	FailTimeout int `json:"fail_timeout"`
	//})
	//err = json.Unmarshal(pair.Value, aa)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(aa.FailTimeout)

}

func getagent(client *capi.Client) {
	agent := client.Agent()
	servces, _ := agent.Services()
	for _, v := range servces {
		fmt.Println(v.Address)
		//fmt.Sprintf("k: %s  kind: %s", k, v.Kind)
	}

}

func Register(client *capi.Client, service_ip string, service_port int) {
	// 配置service 信息
	id := fmt.Sprintf("node-exporter-%s", service_ip)
	newservice := capi.AgentServiceRegistration{
		Kind:              capi.ServiceKindTypical,
		ID:                id,
		Name:              id,
		Tags:              nil,
		Port:              service_port,
		Address:           service_ip,
		SocketPath:        "",
		TaggedAddresses:   nil,
		EnableTagOverride: false,
		Meta: map[string]string{
			"a": "b",
		},
		Weights:   nil,
		Check:     nil,
		Checks:    nil,
		Proxy:     nil,
		Connect:   nil,
		Namespace: "",
		Partition: "",
		Locality:  nil,
	}
	// create agent
	agent := client.Agent()
	err := agent.ServiceRegister(&newservice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("register success")
}

func DelRegister(client *capi.Client, service_ip string) {
	agent := client.Agent()
	err := agent.ServiceDeregister(service_ip) //删除服务
	if err != nil {
		fmt.Println("del failed" + err.Error())
	} else {
		fmt.Println("del success")
	}
}
