package cachestoragebuilderoverseer

import (
	"github.com/herb-go/datamodules/herbcache"
	"github.com/herb-go/datamodules/herbcache/kvengine"
	"github.com/herb-go/worker"
)

var builderworker func(*herbcache.Storage, func(v interface{}) error) error

var Team = worker.GetWorkerTeam(&builderworker)

func GetCacheStorageFactoryByID(id string) func(*herbcache.Storage, func(v interface{}) error) error {
	if id == "" {
		return kvengine.Builder
	}
	w := worker.FindWorker(id)
	if w == nil {
		return nil
	}
	c, ok := w.Interface.(*func(*herbcache.Storage, func(v interface{}) error) error)
	if ok == false || c == nil {
		return nil
	}
	return *c
}
