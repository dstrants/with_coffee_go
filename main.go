package main

import (
	"with_coffee/lib/covid"
)

func main() {
	covid.ImportCovidCases("GR")
	//msg := covid.LoadCovidCases()

	//slack.Send(msg)
}
