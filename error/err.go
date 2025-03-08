package erration

import "log"

func LogError(err error, msg string) {
	if err != nil {
		log.Println(msg)
	}
}