package rpchandlers

import (
	"github.com/karlsen-network/karlsend/v2/app/appmessage"
	"github.com/karlsen-network/karlsend/v2/app/rpc/rpccontext"
	"github.com/karlsen-network/karlsend/v2/infrastructure/network/netadapter/router"
)

// HandleGetSubnetwork handles the respectively named RPC command
func HandleGetSubnetwork(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetSubnetworkResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
