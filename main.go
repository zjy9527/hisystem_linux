package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"

	"github.com/zserge/lorca"
)

func main() {
		var args []string
		if runtime.GOOS == "linux"{
			args = append(args,"--class=Lorca")
		}
		ui, err := lorca.New("","", 1400,800, args...)
		if err != nil{
			log.Fatal(err)
		}
		defer ui.Close()
		ln, err := net.Listen("tcp","127.0.0.1:7000")
		if err != nil{
			log.Fatal(err)
		}
		defer ln.Close()
		go http.Serve(ln, http.FileServer(http.Dir("./www")))
		ui.Load(fmt.Sprintf("http://%s/%s",ln.Addr(),"index.html"))


		sig := make(chan os.Signal)
		signal.Notify(sig, os.Interrupt)
		select {
		case <- sig:
		case <- ui.Done():
		}

		log.Println("exiting...")
}
