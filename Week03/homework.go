package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		c = make(chan os.Signal)
		//监听指定信号 ctrl+c kill
	)
	// Launch a goroutine to fetch the URL.
	group, _ := errgroup.WithContext(context.Background())
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	group.Go(func() error {
		// Fetch the URL.
		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			return err
		}
		fmt.Println(":8081 exit")
		return nil
	})
	group.Go(func() error {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("Program Exit...", s)
				defer GracefullExit()
				return errors.New("Program Exit...")
			case syscall.SIGUSR1:
				fmt.Println("usr1 signal", s)
			case syscall.SIGUSR2:
				fmt.Println("usr2 signal", s)
			default:
				fmt.Println("other signal", s)
			}
		}
		return nil
	})
	
	if err := group.Wait(); err == nil {
		fmt.Println("Successfully ")
	}else{
		fmt.Println("failed "+err.Error())
	}
}

func GracefullExit() {
	fmt.Println("Start Exit...")
	fmt.Println("Execute Clean...")
	fmt.Println("End Exit...")
	os.Exit(0)
}

