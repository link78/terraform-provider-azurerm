package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type GlobalScheduleId struct {
	ResourceGroup string
	Name          string
}

func GlobalScheduleID(input string) (*GlobalScheduleId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Global Schedule ID %q: %+v", input, err)
	}

	service := GlobalScheduleId{
		ResourceGroup: id.ResourceGroup,
	}

	if service.Name, err = id.PopSegment("schedules"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &service, nil
}
