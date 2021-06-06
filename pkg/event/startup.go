package event

import (
	"log"

	"github.com/junglemc/JungleTree/pkg"
)

const (
	thickLine = "===================================="
	thinLine  = "------------------------------------"
)

type (
	ServerStartupEvent    struct{}
	ServerStartupListener struct{}
)

func (e ServerStartupEvent) IsAsync() bool {
	return false
}

type (
	ServerLoadedEvent    struct{}
	ServerLoadedListener struct{}
)

func (e ServerLoadedEvent) IsAsync() bool {
	return false
}

func (l ServerStartupListener) OnEvent(event Event) error {
	log.Println(thickLine)
	log.Println("Starting JungleTree Server " + pkg.Version)
	log.Println(thickLine)
	return nil
}

func (l ServerLoadedListener) OnEvent(event Event) error {
	log.Println(thinLine)
	log.Println("Done!")
	return nil
}
