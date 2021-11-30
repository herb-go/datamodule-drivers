package cacheconfig

import (
	"github.com/herb-go/datamodule-drivers/overseers/cachestoragebuilderoverseer"
	"github.com/herb-go/datamodules/herbcache"
	"github.com/herb-go/worker"
)

type Storage struct {
	ID     string
	Engine func(v interface{}) error `config:", lazyload"`
}

func (s *Storage) ApplyTo(storage *herbcache.Storage) error {
	f := cachestoragebuilderoverseer.GetCacheStorageFactoryByID(s.ID)
	if f == nil {
		return worker.NewWorkerNotFounderError(s.ID)
	}
	return f(storage, s.Engine)
}
