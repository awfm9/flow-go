package backend

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/onflow/flow/protobuf/go/flow/access"
	"github.com/onflow/flow/protobuf/go/flow/execution"

	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/module"
	"github.com/dapperlabs/flow-go/state/protocol"
	"github.com/dapperlabs/flow-go/storage"
)

// Backends implements the Access API.
//
// It is composed of several sub-backends that implement part of the Access API.
//
// Script related calls are handled by backendScripts.
// Transaction related calls are handled by backendTransactions.
// Block Header related calls are handled by backendBlockHeaders.
// Block details related calls are handled by backendBlockDetails.
// Event related calls are handled by backendEvents.
// Account related calls are handled by backendAccounts.
//
// All remaining calls are handled by the base Backend in this file.
type Backend struct {
	backendScripts
	backendTransactions
	backendEvents
	backendBlockHeaders
	backendBlockDetails
	backendAccounts

	executionRPC execution.ExecutionAPIClient
	state        protocol.State
	chainID      flow.ChainID
	collections  storage.Collections
}

type NetworkParameters struct {
	ChainID flow.ChainID
}

func New(
	state protocol.State,
	executionRPC execution.ExecutionAPIClient,
	collectionRPC access.AccessAPIClient,
	blocks storage.Blocks,
	headers storage.Headers,
	collections storage.Collections,
	transactions storage.Transactions,
	chainID flow.ChainID,
	transactionMetrics module.TransactionMetrics,
) *Backend {
	retry := newRetry()

	b := &Backend{
		executionRPC: executionRPC,
		state:        state,
		// create the sub-backends
		backendScripts: backendScripts{
			headers:      headers,
			executionRPC: executionRPC,
			state:        state,
		},
		backendTransactions: backendTransactions{
			collectionRPC:      collectionRPC,
			executionRPC:       executionRPC,
			state:              state,
			chainID:            chainID,
			collections:        collections,
			blocks:             blocks,
			transactions:       transactions,
			transactionMetrics: transactionMetrics,
			retry:              retry,
		},
		backendEvents: backendEvents{
			executionRPC: executionRPC,
			state:        state,
			blocks:       blocks,
		},
		backendBlockHeaders: backendBlockHeaders{
			headers: headers,
			state:   state,
		},
		backendBlockDetails: backendBlockDetails{
			blocks: blocks,
			state:  state,
		},
		backendAccounts: backendAccounts{
			executionRPC: executionRPC,
			state:        state,
			headers:      headers,
		},
		collections: collections,
		chainID:     chainID,
	}

	retry.SetBackend(b)

	return b
}

// Ping responds to requests when the server is up.
func (b *Backend) Ping(ctx context.Context) error {
	_, err := b.executionRPC.Ping(ctx, &execution.PingRequest{})
	if err != nil {
		return fmt.Errorf("could not ping execution node: %w", err)
	}

	_, err = b.collectionRPC.Ping(ctx, &access.PingRequest{})
	if err != nil {
		return fmt.Errorf("could not ping collection node: %w", err)
	}

	return nil
}

func (b *Backend) GetCollectionByID(_ context.Context, colID flow.Identifier) (*flow.LightCollection, error) {
	// retrieve the collection from the collection storage
	col, err := b.collections.LightByID(colID)
	if err != nil {
		err = convertStorageError(err)
		return nil, err
	}

	return col, nil
}

func (b *Backend) GetNetworkParameters(_ context.Context) NetworkParameters {
	return NetworkParameters{
		ChainID: b.chainID,
	}
}

func convertStorageError(err error) error {
	if errors.Is(err, storage.ErrNotFound) {
		return status.Errorf(codes.NotFound, "not found: %v", err)
	}

	return status.Errorf(codes.Internal, "failed to find: %v", err)
}