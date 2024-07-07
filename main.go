package main

import "github.com/pytgcalls/gotgcalls/ntgcalls"

func main() {
	var gotgcalls *ntgcalls.Context
	gotgcalls = ntgcalls.Client(6, "fhwaoifhawiufhalhafi", "sas", false)
	gotgcalls.Run()
}
