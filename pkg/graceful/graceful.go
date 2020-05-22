package graceful

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cloudflare/tableflip"
)

func StartGracefulServer(server *http.Server, pidFile string) {
	upg, err := tableflip.New(tableflip.Options{
		PIDFile: pidFile,
	})
	if err != nil {
		panic(err)
	}
	defer upg.Stop()

	// Do an upgrade on SIGHUP
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP)
		for range sig {
			err := upg.Upgrade()
			if err != nil {
				log.Println("Upgrade failed:", err)
			}
		}
	}()

	// Listen must be called before Ready
	ln, err := upg.Listen("tcp", server.Addr)
	if err != nil {
		log.Fatalln("Can't listen:", err)
	}

	// http Serve
	go func() {
		err := server.Serve(ln)
		if err != http.ErrServerClosed {
			log.Println("HTTP server:", err)
		}
	}()

	log.Printf("ready")
	if err := upg.Ready(); err != nil {
		panic(err)
	}
	<-upg.Exit()

	// Make sure to set a deadline on exiting the process
	// after upg.Exit() is closed. No new upgrades can be
	// performed if the parent doesn't exit.
	time.AfterFunc(30*time.Second, func() {
		log.Println("Graceful shutdown timed out")
		os.Exit(1)
	})

	// Wait for connections to drain.
	_ = server.Shutdown(context.Background())
}
