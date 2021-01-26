package failure

import "log"

func Fail(message string) {
	log.Fatal(message)
}

func FailOnError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}
}
