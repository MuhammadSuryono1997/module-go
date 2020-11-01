package utils

import (
	"time"
)

const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorYellow = "\033[33m"
const colorBlue = "\033[34m"
const colorPurple = "\033[35m"
const colorCyan = "\033[36m"
const colorWhite = "\033[37m"
const formattime = time.RFC3339

func ColorYellow() string {
	return colorYellow
}

func ColorCyan() string {
	return colorCyan
}

func ColorRed() string {
	return colorRed
}
