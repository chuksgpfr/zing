package main

import "github.com/chuksgpfr/zing/zing"

func main() {
	zs := Zing()
	zingCmd := zing.ZingCommand(zs)

	if err := zingCmd.Execute(); err != nil {
		panic(err)
	}

}

func Zing() *zing.Services {
	store, err := zing.NewStore("./zing-store")
	if err != nil {
		panic(err)
	}

	// defer store.Close()

	service := zing.NewServices(store)

	return service
}
