package caldera_connection

type connectionState int

const (
	uninitiated connectionState = iota
	building
	built
	runningOperation
	finishedOperation
	receivedFacts
	cleared
	done
)

type calderaConnection struct {
	Instance *calderaInstance
	state    connectionState
}

func New() (*calderaConnection, error) {
	var instance, err = getCalderaInstance()
	if err != nil {
		return nil, err
	}
	return &calderaConnection{instance, uninitiated}, nil
}

func (cc calderaConnection) CreateAbility(body []byte) error {
	//TODO add api call
	return nil
}
func (cc calderaConnection) DeleteAbility(abilityId string) error {
	//TODO add api call
	return nil
}
func (cc calderaConnection) CreateAgentGroup(agentIds []string) error {
	//TODO add api call
	return nil
}
func (cc calderaConnection) DeleteAgentGroup(agentGroupId string) error {
	//TODO add api call
	return nil
}
func (cc calderaConnection) CreateAdversary(ablilityId string) error {
	//TODO add api call
	return nil
}
func (cc calderaConnection) DeleteAdversary(adversaryId string) error {
	//TODO add api call
	return nil
}
func (cc calderaConnection) CreateOperation(agentGroupId string, adversaryId string) error {
	//TODO add api call
	return nil
}
func (cc calderaConnection) PollOperationStatus(operationId string) error {
	//TODO add api call
	return nil
}
func (cc calderaConnection) RequestFacts(operationId string) error {
	//TODO add api call
	return nil
}
