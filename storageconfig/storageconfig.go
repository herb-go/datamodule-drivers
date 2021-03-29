package storageconfig

import (
	"github.com/herb-go/datamodule-drivers/overseers/cachestoragebuilderoverseer"
	"github.com/herb-go/datamodules/herbcache"
	"github.com/herb-go/worker"
)

type Directive struct {
	ID     string
	Engine func(v interface{}) error `config:", lazyload"`
}

func (d *Directive) ApplyTo(s *herbcache.Storage) error {
	f := cachestoragebuilderoverseer.GetCacheStorageFactoryByID(d.ID)
	if f == nil {
		return worker.NewWorkerNotFounderError(d.ID)
	}
	return f(s, d.Engine)
}
