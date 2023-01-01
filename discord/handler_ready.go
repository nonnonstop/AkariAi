package discord

import (
	"github.com/bwmarrin/discordgo"
)

var buffer = make([][]byte, 0)

func (d *Discord) onReady(s *discordgo.Session, r *discordgo.Ready) {
	d.logger.Infoln("Dicord client is ready")
}
