package caldera

import (
	httptransport "github.com/go-openapi/runtime/client"
	"os"
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
	config.Host = os.Getenv("CALDERA_BASE_URL")

	var calderaClient = client.NewHTTPClientWithConfig(nil, config)

	r := httptransport.New(config.Host, config.BasePath, config.Schemes)
	r.DefaultAuthentication = httptransport.APIKeyAuth("Key", "header", os.Getenv("API_KEY"))
	calderaClient.SetTransport(r)

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
