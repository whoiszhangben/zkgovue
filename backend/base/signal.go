// signal.go
package base

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type signalHandler func(s os.Signal, arg interface{})

type signalSet struct {
	m map[os.Signal]signalHandler
}

func signalSetNew() *signalSet {
	ss := new(signalSet)
	ss.m = make(map[os.Signal]signalHandler)
	return ss
}

func (set *signalSet) register(s os.Signal, handler signalHandler) {
	if _, found := set.m[s]; !found {
		set.m[s] = handler
	}
}

func (set *signalSet) handle(sig os.Signal, arg interface{}) (err error) {
	if _, found := set.m[sig]; found {
		set.m[sig](sig, arg)
		return nil
	} else {
		return fmt.Errorf("No handler available for signal %v", sig)
	}

	panic("won't reach here")
}

func SysSignalHandleDeamon() {
	ss := signalSetNew()
	handler := func(s os.Signal, arg interface{}) {
		fmt.Println("recv handle signal: %v\n", s)
		if s == syscall.SIGINT || s == syscall.SIGTERM {
			logs := Logs
			logs.Info("recv exit signal to exit !\n")
			logs.CloseLog()
			os.Exit(1)
		}
	}

	ss.register(syscall.SIGINT, handler)
	ss.register(syscall.SIGTERM, handler)

	for {
		c := make(chan os.Signal)
		var sigs []os.Signal
		for sig := range ss.m {
			sigs = append(sigs, sig)
		}
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		sig := <-c

		err := ss.handle(sig, nil)
		if err != nil {
			fmt.Printf("unknown signal received: %v\n", sig)
			os.Exit(1)
		}
	}
}
