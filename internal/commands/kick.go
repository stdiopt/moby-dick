package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/syncship/moby-dick/pkg/router"
)

func init() {
	Router.AddCommand(router.Command{
		Name: "kick",
		Run: func(s *discordgo.Session, m *discordgo.MessageCreate, a router.Arguments) {
			//
		},
		Permissions: []int{discordgo.PermissionKickMembers},
	})
}
