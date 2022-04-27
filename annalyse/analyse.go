package annalyse

import (
	coldfire "github.com/redcode-labs/Coldfire"
)

/*
* Get information about the device
 */
func Infos() map[string]string {

	info := coldfire.Info()
	user, err := coldfire.GetUser()

	if err != nil {
		coldfire.PrintError("Info : GetUser command error")
		user = info["username"]
	} else {
		coldfire.PrintGood("Infos : OK")
	}

	return map[string]string{
		// host
		"user":     user,
		"hostname": info["hostname"],
		"os":       info["os"],
		"kernel":   info["kernel"],
		// network
		"ip":     info["local_ip"],
		"mac":    info["mac"],
		"globIp": info["global_ip"],
	}
}

/*
* Get information about a user on the device
 */
func UserInfo(us string) map[string]interface{} {

	coldfire.PrintInfo(us)

	return nil
}
