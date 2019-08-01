package relay

import (
	"math/big"

	"github.com/ipfs/go-log"

	relayChain "github.com/keep-network/keep-core/pkg/beacon/relay/chain"
	"github.com/keep-network/keep-core/pkg/beacon/relay/entry"

	"github.com/keep-network/keep-core/pkg/beacon/relay/config"
	"github.com/keep-network/keep-core/pkg/beacon/relay/registry"
	"github.com/keep-network/keep-core/pkg/chain"
	"github.com/keep-network/keep-core/pkg/net"
)

var logger = log.Logger("keep-relay")

// NewNode returns an empty Node with no group, zero group count, and a nil last
// seen entry, tied to the given net.Provider.
func NewNode(
	staker chain.Staker,
	netProvider net.Provider,
	blockCounter chain.BlockCounter,
	chainConfig *config.Chain,
	groupRegistry *registry.Groups,
) Node {
	return Node{
		Staker:        staker,
		netProvider:   netProvider,
		blockCounter:  blockCounter,
		chainConfig:   chainConfig,
		groupRegistry: groupRegistry,
	}
}

// GenerateRelayEntryIfEligible takes a relay request and checks if this client
// is one of the nodes elected by that request to create a new relay entry.
// If it is, this client enters the threshold signature creation process and,
// upon successfully completing it, submits the signature as a new relay entry
// to the passed in relayChain. Note that this function returns immediately after
// determining whether the node is or is not is a member of the requested group, and
// signature creation and submission is performed in a background goroutine.
func (n *Node) GenerateRelayEntryIfEligible(
	previousEntry *big.Int,
	seed *big.Int,
	relayChain relayChain.Interface,
	groupPublicKey []byte,
	startBlockHeight uint64,
) {
	memberships := n.groupRegistry.GetGroup(groupPublicKey)

	if len(memberships) < 1 {
		return
	}

	for _, signer := range memberships {
		go func(signer *registry.Membership) {
			channel, err := n.netProvider.ChannelFor(signer.ChannelName)
			if err != nil {
				logger.Errorf(
					"could not create broadcast channel with name [%v]: [%v]",
					signer.ChannelName,
					err,
				)
				return
			}

			err = entry.SignAndSubmit(
				n.blockCounter,
				channel,
				relayChain,
				previousEntry,
				seed,
				n.chainConfig.HonestThreshold(),
				signer.Signer,
				startBlockHeight,
			)
			if err != nil {
				logger.Errorf(
					"error creating threshold signature: [%v]",
					err,
				)
				return
			}
		}(signer)
	}
}
