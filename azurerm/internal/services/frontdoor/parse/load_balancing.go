package parse

import "fmt"

type LoadBalancingId struct {
	ResourceGroup string
	FrontDoorName string
	Name          string
}

func NewLoadBalancingID(id FrontDoorId, name string) LoadBalancingId {
	return LoadBalancingId{
		ResourceGroup: id.ResourceGroup,
		FrontDoorName: id.Name,
		Name:          name,
	}
}

func (id LoadBalancingId) ID(subscriptionId string) string {
	base := NewFrontDoorID(id.ResourceGroup, id.Name).ID(subscriptionId)
	return fmt.Sprintf("%s/loadBalancingSettings/%s", base, id.Name)
}

func LoadBalancingID(input string) (*LoadBalancingId, error) {
	frontDoorId, id, err := parseFrontDoorChildResourceId(input)
	if err != nil {
		return nil, fmt.Errorf("parsing FrontDoor Load Balancing ID %q: %+v", input, err)
	}

	loadBalancingId := LoadBalancingId{
		ResourceGroup: frontDoorId.ResourceGroup,
		FrontDoorName: frontDoorId.Name,
	}

	// TODO: handle this being case-insensitive
	// https://github.com/Azure/azure-sdk-for-go/issues/6762
	if loadBalancingId.Name, err = id.PopSegment("loadBalancingSettings"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &loadBalancingId, nil
}
