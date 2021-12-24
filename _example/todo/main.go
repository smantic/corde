package main

import (
	"log"
	"os"
	"sync"

	"github.com/Karitham/corde"
)

var commands = []corde.Command{
	corde.Command{
		Name:        "todo",
		Description: "view edit and remove todos",
		Type:        corde.COMMAND_CHAT_INPUT,
		Options: []corde.Option{
			{
				Name:        "list",
				Description: "list todos",
				Type:        corde.OPTION_SUB_COMMAND,
			},
			{
				Name:        "add",
				Description: "add a todo",
				Type:        corde.OPTION_SUB_COMMAND,
				Options: []corde.Option{
					{
						Name:        "name",
						Type:        corde.OPTION_STRING,
						Description: "todo name",
						Required:    true,
					},
					{
						Name:        "value",
						Type:        corde.OPTION_STRING,
						Description: "todo value",
						Required:    true,
					},
				},
			},
			{
				Name:        "rm",
				Description: "remove a todo",
				Type:        corde.OPTION_SUB_COMMAND,
				Options: []corde.Option{
					{
						Name:        "name",
						Type:        corde.OPTION_STRING,
						Description: "todo name",
						Required:    true,
					},
				},
			},
		},
	},
}

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatalln("DISCORD_BOT_TOKEN not set")
	}
	appID := corde.SnowflakeFromString(os.Getenv("DISCORD_APP_ID"))
	if appID == 0 {
		log.Fatalln("DISCORD_APP_ID not set")
	}
	pk := os.Getenv("DISCORD_PUBLIC_KEY")
	if pk == "" {
		log.Fatalln("DISCORD_PUBLIC_KEY not set")
	}

	t := todo{
		mu:   sync.Mutex{},
		list: make(map[string]string),
	}

	m := corde.NewMux(pk, appID, token)
	m.SetRoute(corde.InteractionCommand{Type: corde.APPLICATION_COMMAND, Route: "todo/add"}, t.addHandler)
	m.SetRoute(corde.InteractionCommand{Type: corde.APPLICATION_COMMAND, Route: "todo/rm"}, t.removeHandler)
	m.SetRoute(corde.InteractionCommand{Type: corde.APPLICATION_COMMAND, Route: "todo/list"}, t.listHandler)

	g := corde.Guild(corde.SnowflakeFromString(os.Getenv("DISCORD_GUILD_ID")))
	for _, c := range commands {
		if err := m.RegisterCommand(c, g); err != nil {
			log.Fatalln("error registering command: ", err)
		}
	}

	if err := m.ListenAndServe(":8070"); err != nil {
		log.Fatalln(err)
	}
}
