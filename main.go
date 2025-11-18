package main

import "github.com/chuksgpfr/zing/zing"

func main() {
	zs := Zing()
	zingCmd := zing.ZingCommand(zs)

	if err := zingCmd.Execute(); err != nil {
		panic(err)
	}

}
