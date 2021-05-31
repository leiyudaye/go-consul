package napodat

/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-03-07 21:56:29
 * @LastEditors: lly
 * @LastEditTime: 2021-04-15 00:59:09
 */
import (
	"log"
	"strconv"

	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
)

type KitDiscoverClient struct {
	Host   string
	Port   int
	client consul.Client
}

func NewKitDiscoverClient(consulHost string, consulPort int) (DiscoverClient, error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = consulHost + ":" + strconv.Itoa(consulPort)
	apiClient, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, err
	}
	client := consul.NewClient(apiClient)
	return &KitDiscoverClient{
		Host:   consulHost,
		Port:   consulPort,
		client: client,
	}, err
}

// 服务注册
func (consulClient *KitDiscoverClient) Register(serviceName, instanceId, healthCheckUrl, instanceHost string,
	instancePort int, meta map[string]string, loggeer *log.Logger) bool {
	// 构建服务实例元数据
	registration := &api.AgentServiceRegistration{
		ID:      instanceId,
		Name:    serviceName,
		Address: instanceHost,
		Port:    instancePort,
		Meta:    meta,
		Check: &api.AgentServiceCheck{
			DeregisterCriticalServiceAfter: "30s",
			HTTP:                           "http://" + instanceHost + ":" + strconv.Itoa(instancePort) + healthCheckUrl,
			Interval:                       "15s",
		},
	}

	// 注册服务
	err := consulClient.client.Register(registration)
	if err != nil {
		log.Println("Register Service Error", err)
		return false
	}
	log.Println("Register Service Success")
	return true
}

// 服务注销
func (consulClient *KitDiscoverClient) DeRegister(instanceID string, logger *log.Logger) bool {
	return true
}

// 服务发现
func (consulClient *KitDiscoverClient) DiscoverServices(serviceName string, logger *log.Logger) []interface{} {
	return nil
}
