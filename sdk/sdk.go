package sdk

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cskr/pubsub"
	"github.com/mesg-foundation/engine/container"
	"github.com/mesg-foundation/engine/cosmos"
	"github.com/mesg-foundation/engine/database"
	accountsdk "github.com/mesg-foundation/engine/sdk/account"
	eventsdk "github.com/mesg-foundation/engine/sdk/event"
	executionsdk "github.com/mesg-foundation/engine/sdk/execution"
	instancesdk "github.com/mesg-foundation/engine/sdk/instance"
	processesdk "github.com/mesg-foundation/engine/sdk/process"
	servicesdk "github.com/mesg-foundation/engine/sdk/service"
)

// SDK exposes all functionalities of MESG Engine.
type SDK struct {
	Service   servicesdk.Service
	Instance  *instancesdk.Instance
	Execution *executionsdk.Execution
	Event     *eventsdk.Event
	Process   *processesdk.Process
	Account   *accountsdk.SDK
}

// New creates a new SDK with given options.
func New(client *cosmos.Client, cdc *codec.Codec, kb *cosmos.Keybase, c container.Container, instanceDB database.InstanceDB, execDB database.ExecutionDB, processDB database.ProcessDB, engineName, port string) *SDK {
	ps := pubsub.New(0)
	serviceSDK := servicesdk.New(cdc, client)
	instanceSDK := instancesdk.New(c, serviceSDK, instanceDB, engineName, port)
	processSDK := processesdk.New(instanceSDK, processDB)
	executionSDK := executionsdk.New(ps, serviceSDK, instanceSDK, processSDK, execDB)
	eventSDK := eventsdk.New(ps, serviceSDK, instanceSDK)
	accountSDK := accountsdk.NewSDK(kb)
	return &SDK{
		Service:   serviceSDK,
		Instance:  instanceSDK,
		Execution: executionSDK,
		Event:     eventSDK,
		Process:   processSDK,
		Account:   accountSDK,
	}
}
