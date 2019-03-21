package seo

import (
	"github.com/ecletus/plug"
)

var E_SEO_COLLECTION = PREFIX + ".collection"

func ECollection(name string) string {
	if name == "" {
		panic("Name is empty")
	}
	return E_SEO_COLLECTION + ":" + name
}

type CollectionEvent struct {
	plug.PluginEventInterface
	Collection     *Collection
	CollectionName string
	PluginEvent    plug.PluginEventInterface
}

type Plugin struct {
	plug.EventDispatcher
	CollectionNames []string
}

func (p *Plugin) RequireOptions() []string {
	return p.CollectionNames
}

func (p *Plugin) OnRegister() {
	p.On(plug.E_POST_INIT, func(e plug.PluginEventInterface) (err error) {
		for _, name := range p.CollectionNames {
			coll := e.Options().GetInterface(name).(*Collection)
			err = e.PluginDispatcher().TriggerPlugins(&CollectionEvent{plug.NewPluginEvent(ECollection(name)), coll, name, e})
			if err != nil {
				return
			}
		}
		return
	})
}
