package caldera_connection

import (
	"soarca/internal/capability/caldera/api/client"
	"sync"
)

type calderaInstance struct {
	send client.Caldera
}

var cInstance *calderaInstance
var instanceLock = &sync.Mutex{}

func newCalderaInstance() (*calderaInstance, error) {
	var config = client.DefaultTransportConfig()
	//TODO set the correct host by ENV or default docker instance
	var calderaClient = client.NewHTTPClientWithConfig(nil, config)

	//TODO Do an initial cleanup
	return &calderaInstance{*calderaClient}, nil
}

func getCalderaInstance() (*calderaInstance, error) {
	if cInstance == nil {
		instanceLock.Lock()
		defer instanceLock.Unlock()
		if cInstance == nil {
			var instance, err = newCalderaInstance()
			if err != nil {
				return nil, err
			}
			cInstance = instance
		}
	}
	return cInstance, nil
}
