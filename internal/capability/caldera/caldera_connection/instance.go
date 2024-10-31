package caldera_connection

import "sync"

type calderaInstance struct {
	Url  string
	Port uint16
}

var cInstance *calderaInstance
var instanceLock = &sync.Mutex{}

func newCalderaInstance() (*calderaInstance, error) {
	//TODO Make the connection and do a cleanup
	return &calderaInstance{}, nil
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
