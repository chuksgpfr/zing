package main

import "github.com/chuksgpfr/zing/zing"

func Zing() *zing.Services {
	store, err := zing.NewStore("./zing-store")
	if err != nil {
		panic(err)
	}

	// defer store.Close()

	service := zing.NewServices(store)

	return service
}
