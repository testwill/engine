package api

import (
	"context"

	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/instance"
	protobuf_api "github.com/mesg-foundation/engine/protobuf/api"
	"github.com/mesg-foundation/engine/protobuf/types"
	"github.com/mesg-foundation/engine/sdk"
	instancesdk "github.com/mesg-foundation/engine/sdk/instance"
)

// InstanceServer is the type to aggregate all Instance APIs.
type InstanceServer struct {
	sdk *sdk.SDK
}

// NewInstanceServer creates a new ServiceServer.
func NewInstanceServer(sdk *sdk.SDK) *InstanceServer {
	return &InstanceServer{sdk: sdk}
}

// List instances.
func (s *InstanceServer) List(ctx context.Context, request *protobuf_api.ListInstancesRequest) (*protobuf_api.ListInstancesResponse, error) {
	var (
		h   hash.Hash
		err error
	)
	if request.ServiceHash != "" {
		h, err = hash.Decode(request.ServiceHash)
		if err != nil {
			return nil, err
		}
	}
	instances, err := s.sdk.Instance.List(&instancesdk.Filter{ServiceHash: h})
	if err != nil {
		return nil, err
	}
	return &protobuf_api.ListInstancesResponse{Instances: toProtoInstances(instances)}, nil
}

// Create creates a new instance from service.
func (s *InstanceServer) Create(ctx context.Context, request *protobuf_api.CreateInstanceRequest) (*protobuf_api.CreateInstanceResponse, error) {
	hash, err := hash.Decode(request.ServiceHash)
	if err != nil {
		return nil, err
	}

	i, err := s.sdk.Instance.Create(hash, request.Env)
	if err != nil {
		return nil, err
	}
	return &protobuf_api.CreateInstanceResponse{Hash: i.Hash.String()}, nil
}

// Get retrives instance.
func (s *InstanceServer) Get(ctx context.Context, request *protobuf_api.GetInstanceRequest) (*types.Instance, error) {
	hash, err := hash.Decode(request.Hash)
	if err != nil {
		return nil, err
	}
	i, err := s.sdk.Instance.Get(hash)
	if err != nil {
		return nil, err
	}
	return toProtoInstance(i), nil
}

// Delete an instance
func (s *InstanceServer) Delete(ctx context.Context, request *protobuf_api.DeleteInstanceRequest) (*protobuf_api.DeleteInstanceResponse, error) {
	hash, err := hash.Decode(request.Hash)
	if err != nil {
		return nil, err
	}
	if err := s.sdk.Instance.Delete(hash, request.DeleteData); err != nil {
		return nil, err
	}
	return &protobuf_api.DeleteInstanceResponse{}, nil
}

func toProtoInstances(instances []*instance.Instance) []*types.Instance {
	inst := make([]*types.Instance, len(instances))
	for i, instance := range instances {
		inst[i] = toProtoInstance(instance)
	}
	return inst
}

func toProtoInstance(i *instance.Instance) *types.Instance {
	return &types.Instance{
		Hash:        i.Hash.String(),
		ServiceHash: i.ServiceHash.String(),
	}
}