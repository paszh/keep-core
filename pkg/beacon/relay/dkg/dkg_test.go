package dkg

import (
	"testing"

	"github.com/keep-network/keep-core/pkg/beacon/relay/gjkr"
	"github.com/keep-network/keep-core/pkg/beacon/relay/group"
	"github.com/keep-network/keep-core/pkg/internal/testutils"
	"github.com/keep-network/keep-core/pkg/net"
)

func TestExecute_HappyPath(t *testing.T) {
	groupSize := 5
	threshold := 3

	interceptor := func(msg net.TaggedMarshaler) net.TaggedMarshaler {
		return msg
	}
	network := testutils.NewInterceptingNetwork(interceptor)

	result, signers, err := executeDKG(groupSize, threshold, network)
	if err != nil {
		t.Fatal(err)
	}

	assertSignersCount(t, signers, groupSize)
	assertSamePublicKey(t, result, signers)
	// TODO: assert no DQ
	// TODO: assert no IA
	// TODO: assert key is valid
}

func TestExecute_IA_member1_commitmentPhase(t *testing.T) {
	groupSize := 5
	threshold := 3

	interceptor := func(msg net.TaggedMarshaler) net.TaggedMarshaler {
		// drop commitment message from member 1
		commitmentMsg, ok := msg.(*gjkr.MemberCommitmentsMessage)
		if ok && commitmentMsg.SenderID() == group.MemberIndex(1) {
			return nil
		}

		return msg
	}
	network := testutils.NewInterceptingNetwork(interceptor)

	result, signers, err := executeDKG(groupSize, threshold, network)
	if err != nil {
		t.Fatal(err)
	}

	assertSignersCount(t, signers, groupSize)
	honestSigners := filterOutMisbehavingSigners(signers, group.MemberIndex(1))
	assertSamePublicKey(t, result, honestSigners)
	// TODO: assert no DQ
	// TODO: assert member 1 is IA
	// TODO: assert key is valid
}
