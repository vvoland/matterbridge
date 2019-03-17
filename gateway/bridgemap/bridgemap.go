package bridgemap

import (
	"github.com/vvoland/matterbridge/bridge"
	"github.com/vvoland/matterbridge/bridge/api"
	"github.com/vvoland/matterbridge/bridge/discord"
	"github.com/vvoland/matterbridge/bridge/gitter"
	"github.com/vvoland/matterbridge/bridge/irc"
	"github.com/vvoland/matterbridge/bridge/matrix"
	"github.com/vvoland/matterbridge/bridge/mattermost"
	"github.com/vvoland/matterbridge/bridge/rocketchat"
	"github.com/vvoland/matterbridge/bridge/slack"
	"github.com/vvoland/matterbridge/bridge/sshchat"
	"github.com/vvoland/matterbridge/bridge/steam"
	"github.com/vvoland/matterbridge/bridge/telegram"
	"github.com/vvoland/matterbridge/bridge/whatsapp"
	"github.com/vvoland/matterbridge/bridge/xmpp"
	"github.com/vvoland/matterbridge/bridge/zulip"
)

var FullMap = map[string]bridge.Factory{
	"api":          api.New,
	"discord":      bdiscord.New,
	"gitter":       bgitter.New,
	"irc":          birc.New,
	"mattermost":   bmattermost.New,
	"matrix":       bmatrix.New,
	"rocketchat":   brocketchat.New,
	"slack-legacy": bslack.NewLegacy,
	"slack":        bslack.New,
	"sshchat":      bsshchat.New,
	"steam":        bsteam.New,
	"telegram":     btelegram.New,
	"whatsapp":     bwhatsapp.New,
	"xmpp":         bxmpp.New,
	"zulip":        bzulip.New,
}
