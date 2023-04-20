package bot

import (
	"github.com/bwmarrin/discordgo"
)

type GuildsChannels map[string][]*discordgo.Channel

func GetGuildWithChannels(s *discordgo.Session) GuildsChannels {
	var guildsCh GuildsChannels
	for _, guild := range s.State.Guilds {

		// Get channel for guilds
		channels, _ := s.GuildChannels(guild.ID)
		guildsCh[guild.ID] = append(guildsCh[guild.ID], channels...)
	}

	return guildsCh
}
