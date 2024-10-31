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

type CalderaConnection struct {
	instance *calderaInstance
	state    connectionState
}
type CalderaFacts map[string]string

func New() (*CalderaConnection, error) {
	var instance, err = getCalderaInstance()
	if err != nil {
		return nil, err
	}
	return &CalderaConnection{instance, uninitiated}, nil
}

func (cc CalderaConnection) CreateAbility(body []byte) (string, error) {
	//TODO add api call
	return "", nil
}
func (cc CalderaConnection) DeleteAbility(abilityId string) error {
	//TODO add api call
	return nil
}
func (cc CalderaConnection) CreateAgentGroup(agentIds []string) error {
	//TODO add api call
	return nil
}
func (cc CalderaConnection) DeleteAgentGroup(agentGroupId string) error {
	//TODO add api call
	return nil
}
func (cc CalderaConnection) CreateAdversary(ablilityId string) (string, error) {
	//TODO add api call
	return "", nil
}
func (cc CalderaConnection) DeleteAdversary(adversaryId string) error {
	//TODO add api call
	return nil
}

func (cc CalderaConnection) CreateOperation(
	agentGroupId string,
	adversaryId string,
) (string, error) {
	//TODO add api call
	return "", nil
}
func (cc CalderaConnection) IsOperationFinished(operationId string) (bool, error) {
	//TODO add api call
	return false, nil
}
func (cc CalderaConnection) RequestFacts(operationId string) (CalderaFacts, error) {
	//TODO add api call
	return nil, nil
}
