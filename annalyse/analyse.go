package annalyse

import (
	coldfire "github.com/redcode-labs/Coldfire"
)

func Analyser() map[string]interface{} {

	user, _ := coldfire.GetUser()

	return map[string]interface{}{
		"user": user,
	}
}
