package cacheconfig

import (
	"github.com/herb-go/datamodules/herbcache/cachepreset"
	"github.com/herb-go/herbdata/dataencoding/msgpackencoding"
)

var DefaultEncoding = msgpackencoding.Encoding

type Preset struct {
	Namespace string
	Prefix    string
	TTL       int64
}

var DefaultTTL = int64(1800)

func (p *Preset) Exec(preset *cachepreset.Preset) (*cachepreset.Preset, error) {
	commands := cachepreset.NewCommands()
	if p.Namespace != "" {
		commands = commands.Concat(cachepreset.Allocate(p.Namespace))
	}
	if p.Prefix != "" {
		commands = commands.Concat(cachepreset.PrefixCache(p.Prefix))
	}
	ttl := p.TTL
	if ttl <= 0 {
		ttl = DefaultTTL
	}
	encoding := DefaultEncoding
	commands = commands.Concat(cachepreset.TTL(ttl), cachepreset.Encoding(encoding))
	return commands.Exec(preset)
}
