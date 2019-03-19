package handlers

import (
	"encoding/json"
	"testing"

	"github.com/tokenized/smart-contract/internal/contract"
	"github.com/tokenized/smart-contract/internal/platform"
	"github.com/tokenized/smart-contract/internal/platform/state"
	"github.com/tokenized/smart-contract/pkg/protocol"
)

func TestContract(t *testing.T) {
	// Offer message
	offerData := protocol.ContractOffer{
		ContractName: "Test Name",
		KeyRoles: []protocol.KeyRole{
			protocol.KeyRole{Type: 1, Name: "John Smith"},
		},
	}
	offer := &offerData

	// Formation data
	var formationData protocol.ContractFormation
	formation := &formationData

	// State data
	var stateData state.Contract
	st := &stateData

	// Create data
	var createData contract.NewContract
	create := &createData

	// Update data
	var updateData contract.UpdateContract
	update := &updateData

	// Conversions
	var err error
	var text []byte
	text, _ = json.MarshalIndent(offer, "", "  ")
	t.Logf("Offer : \n%s\n", string(text))

	// Formation
	err = platform.Convert(offer, formation)
	if err != nil {
		t.Errorf("Failed to convert offer to formation : %s", err)
	}

	text, _ = json.MarshalIndent(formation, "", "  ")
	t.Logf("Formation : \n%s\n", string(text))

	// State
	err = platform.Convert(formation, st)
	if err != nil {
		t.Errorf("Failed to convert formation to state : %s", err)
	}

	text, _ = json.MarshalIndent(st, "", "  ")
	t.Logf("State : \n%s\n", string(text))

	// Create
	err = platform.Convert(formation, create)
	if err != nil {
		t.Errorf("Failed to convert formation to create : %s", err)
	}

	text, _ = json.MarshalIndent(create, "", "  ")
	t.Logf("Create : \n%s\n", string(text))

	// Update
	err = platform.Convert(formation, update)
	if err != nil {
		t.Errorf("Failed to convert formation to update : %s", err)
	}

	text, _ = json.MarshalIndent(update, "", "  ")
	t.Logf("Update : \n%s\n", string(text))

	// TODO Add test to ensure all important fields are being passed through in convert.
}