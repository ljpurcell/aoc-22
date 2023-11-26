package utils

import "log"

func CheckAndLog(err error, msg string) {
    if err != nil {
        log.Printf(msg)
        panic(err)
    }
}
