package flavor

import (
	"github.com/kadende/kadende-interfaces/spi"
	"github.com/kadende/kadende-interfaces/spi/instance"
	"github.com/kadende/kadende-interfaces/pkg/types"
	"github.com/kadende/kadende-interfaces/spi/group"
	"github.com/kadende/kadende-interfaces/spi/controller"
)

const (
	// ClusterIDTag is the name of the tag that contains unique ID for the cluster
	ClusterIDTag = "kadende.cluster.id"
)

// InterfaceSpec is the current name and version of the Flavor API.
var InterfaceSpec = spi.InterfaceSpec{
	Name:    "Flavor",
	Version: "0.0.1",
}

// Health is an indication of whether the Flavor is functioning properly.
type Health int

const (
	// Unknown indicates that the Health cannot currently be confirmed.
	Unknown Health = iota

	// Healthy indicates that the Flavor is confirmed to be functioning.
	Healthy

	// Unhealthy indicates that the Flavor is confirmed to not be functioning properly.
	Unhealthy
)

// Plugin defines custom behavior for what runs on instances.
type Plugin interface {

	// Validate checks whether the helper can support a configuration.
	Validate(flavorProperties *types.Any, clusterContext controller.Index) error

	//Commands to be executed on the created instance
	Commands(flavorProperties *types.Any, spec instance.Spec, clusterContext controller.Index) (instance.Spec, error)

	// Healthy determines the Health of this Flavor on an instance.
	Healthy(flavorProperties *types.Any, inst instance.Description, clusterContext controller.Index) (Health, error)

	// Drain allows the flavor to perform a best-effort cleanup operation before the instance is destroyed.
	Drain(flavorProperties *types.Any, inst instance.Description, clusterContext controller.Index) error
}
