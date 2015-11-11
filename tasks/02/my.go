package main

import "fmt"

func bufferLog(log chan string, logIndex int) chan string {
	bufferedLog := make(chan string, 100)
	go func(bufferedLog chan string, log chan string) {
		for message := range log {
			bufferedLog <- fmt.Sprintf("%d\t%s", logIndex, message)
		}
		close(bufferedLog)
	}(bufferedLog, log)
	return bufferedLog
}

func bufferLogs(logs chan (chan string)) chan (chan string) {
	bufferedLogs := make(chan (chan string), 100)
	go func(bufferedLogs chan (chan string), logs chan (chan string)) {
		var logIndex int
		for log := range logs {
			logIndex++
			bufferedLogs <- bufferLog(log, logIndex)
		}
		close(bufferedLogs)
	}(bufferedLogs, logs)
	return bufferedLogs
}

func actualLogDrainer(logs chan (chan string), result chan string) {
	for log := range bufferLogs(logs) {
		for message := range log {
			result <- message
		}
	}
	close(result)
}

func OrderedLogDrainer(logs chan (chan string)) chan string {
	out := make(chan string)
	go actualLogDrainer(logs, out)
	return out
}
