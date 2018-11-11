package request

import (
	"context"
	"errors"

	"github.com/tokenized/smart-contract/internal/app/state/contract"
	"github.com/tokenized/smart-contract/pkg/txbuilder"
	"github.com/tokenized/smart-contract/pkg/protocol"
)

type referendumHandler struct{}

func newReferendumHandler() referendumHandler {
	return referendumHandler{}
}

func (h referendumHandler) handle(ctx context.Context,
	r contractRequest) (*contractResponse, error) {

	referendum, ok := r.m.(*protocol.Referendum)
	if !ok {
		return nil, errors.New("Not *protocol.Referendum")
	}

	// Contract
	c := r.contract

	// create the protocol.Vote
	vote := buildVoteFromReferendum(r.hash, referendum)

	// create the Vote
	v := contract.NewVoteFromProtocolVote(r.senders[0].EncodeAddress(), &vote)

	// record the UTXO for later when we need to send the Result when the
	// Vote cutoff time passes.
	v.UTXO = txbuilder.NewUTXOFromTX(*r.tx, 1)

	// add the Vote to the Contract
	c.Votes[v.RefTxnIDHash] = v

	resp := contractResponse{
		Contract: c,
		Message:  &vote,
	}

	return &resp, nil
}
