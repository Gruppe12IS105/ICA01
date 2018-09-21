package counzl_chat

import (
	"../users"
	"fmt"
	"gopkg.in/abiosoft/ishell.v2"
)

// create new shell.
// by default, new shell includes 'exit', 'help' and 'clear' commands.
// the main file will work as a frontpage/menu -> make new shell for each program module
// added to the project...

//blockchain.RunBlockChain()

func Chatroom() {
	shell := ishell.New()

	// display welcome info.
	shell.Println("Velkommen til Chat room!                                                                                   ")
	shell.Println("Denne modulen er fortsatt under utvikling :( du må nok smøre deg med litt tålmodighet")

	// register a function for "greet" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "whoami",
		Help: "Hent informasjon om din egen bruker",
		Func: func(c *ishell.Context) {
			currentUser := shell.Get("brukernavn")
			_, id, online := users.FetchUser(currentUser.(string))
			fmt.Println("Info om din nåværende bruker:\nBrukernavn: " + currentUser.(string) + "\nID : " + id + "\nRegistrert på server: " + online)
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "send",
		Help: "send melding til server",
		Func: func(c *ishell.Context) {
			//client.DialServer()
		},
	})
	shell.Run()
}
