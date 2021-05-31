package napodat

import "log"

/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-03-07 11:35:59
 * @LastEditors: lly
 * @LastEditTime: 2021-04-15 00:19:47
 */

type DiscoverClient interface {
	Register(serviceName, intstanceID, healthCheckUrl string, instanceHost string,
		instancePort int, meta map[string]string, logger *log.Logger) bool

	DeRegister(instanceID string, logger *log.Logger) bool

	DiscoverServices(serviceName string, logger *log.Logger) []interface{}
}
