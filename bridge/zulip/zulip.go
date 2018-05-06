package bzulip

import (
	"strconv"
	"time"

	"github.com/42wim/matterbridge/bridge"
	"github.com/42wim/matterbridge/bridge/config"
	gzb "github.com/ifo/gozulipbot"
)

type Bzulip struct {
	q       *gzb.Queue
	bot     *gzb.Bot
	streams map[int]string
	*bridge.Config
}

func New(cfg *bridge.Config) bridge.Bridger {
	return &Bzulip{Config: cfg, streams: make(map[int]string)}
}

func (b *Bzulip) Connect() error {
	bot := gzb.Bot{APIKey: b.GetString("token"), APIURL: b.GetString("server") + "/api/v1/", Email: b.GetString("login")}
	bot.Init()
	q, err := bot.RegisterAll()
	b.q = q
	b.bot = &bot
	if err != nil {
		b.Log.Debugf("%#v", err)
		return err
	}
	// init stream
	b.getChannel(0)
	b.Log.Info("Connection succeeded")
	go b.handleQueue()
	return nil
}

func (b *Bzulip) Disconnect() error {
	return nil
}

func (b *Bzulip) JoinChannel(channel config.ChannelInfo) error {
	return nil
}

func (b *Bzulip) Send(msg config.Message) (string, error) {
	m := gzb.Message{
		Stream:  msg.Channel,
		Topic:   "test",
		Content: msg.Username + msg.Text,
	}
	resp, err := b.bot.Message(m)
	if err != nil {
		b.Log.Debugf("resp err %#v %#v", resp, err)
	}
	b.Log.Debugf("resp err %#v %#v", resp, err)
	return "", nil
}

func (b *Bzulip) getChannel(id int) string {
	if name, ok := b.streams[id]; ok {
		return name
	}
	streams, err := b.bot.GetRawStreams()
	if err != nil {
		b.Log.Debugf("%#v", err)
		return ""
	}
	for _, stream := range streams.Streams {
		b.streams[stream.StreamID] = stream.Name
	}
	if name, ok := b.streams[id]; ok {
		return name
	}
	return ""
}

func (b *Bzulip) handleQueue() error {
	for {
		messages, _ := b.q.GetEvents()
		for _, m := range messages {
			if m.SenderEmail == b.GetString("login") {
				continue
			}
			b.Log.Debugf("%#v", m)
			rmsg := config.Message{Username: m.SenderFullName, Text: m.Content, Channel: b.getChannel(m.StreamID), Account: b.Account, UserID: strconv.Itoa(m.SenderID)}
			b.Log.Debugf("<= Message %#v", rmsg)
			b.Remote <- rmsg
			b.q.LastEventID = m.ID
		}
		time.Sleep(time.Second * 3)
	}
}
