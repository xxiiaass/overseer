package main

import (
	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
	"net"
)

type TestScript struct {
}

func (TestScript) OnOver() {

}

func (TestScript) SafeHandler(state *overseer.State) {

}

func (TestScript) Init(resource overseer.ForkResource, config overseer.Config) ([]net.Listener, error) {
	return []net.Listener{}, nil
}

//then create another 'main' which runs the upgrades
//'main()' is run in the initial process
func main() {
	overseer.Run(overseer.Config{
		Fetcher: &fetcher.File{Path: "script"},
		Debug:   false, //display log of overseer actions
		Grace:   new(TestScript),
	})
}
