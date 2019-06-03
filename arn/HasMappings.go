package arn

// HasMappings implements common mapping methods.
type hasMappings struct {
	Mappings []*Mapping `json:"mappings" editable:"true"`
}

// SetMapping sets the ID of an external site to the obj.
func (obj *hasMappings) SetMapping(serviceName string, serviceID string) {
	// Is the ID valid?
	if serviceID == "" {
		return
	}

	// If it already exists we don't need to add it
	for _, external := range obj.Mappings {
		if external.Service == serviceName {
			external.ServiceID = serviceID
			return
		}
	}

	// Add the mapping
	obj.Mappings = append(obj.Mappings, &Mapping{
		Service:   serviceName,
		ServiceID: serviceID,
	})
}

// GetMapping returns the external ID for the given service.
func (obj *hasMappings) GetMapping(name string) string {
	for _, external := range obj.Mappings {
		if external.Service == name {
			return external.ServiceID
		}
	}

	return ""
}

// RemoveMapping removes all mappings with the given service name and ID.
func (obj *hasMappings) RemoveMapping(name string) bool {
	for index, external := range obj.Mappings {
		if external.Service == name {
			obj.Mappings = append(obj.Mappings[:index], obj.Mappings[index+1:]...)
			return true
		}
	}

	return false
}
