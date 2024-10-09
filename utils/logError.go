package utils

import "log"

func LogError(err error, msg string) {

	if err!= nil {
		log.Fatalf("%v : %s",err,msg);
	}

}
