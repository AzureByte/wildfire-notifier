package main

import (
  "github.com/bwmarrin/discordgo"
  "flag"
  "fmt"
  "os/signal"
  "syscall"
  "os"
)

var (
  Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()

	if Token == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main()  {
  wildfire, err := discordgo.New("Bot " + Token)
  if err != nil {
    fmt.Println("error creating Discord session,", err)
    return
  }

  err = wildfire.Open()
  if err != nil {
    fmt.Println("Error opening Discord session: ", err)
  }

  fmt.Println("Wildfire Notifier is now running.  Press CTRL-C to exit.")
  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <-sc

  // Cleanly close down the Discord session.
  wildfire.Close()
  fmt.Println("The fire has been extinguished and the watchers are now asleep.")

}
