package protocol

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	// CodeAssetDefinition identifies data as a AssetDefinition message.
	CodeAssetDefinition = "A1"

	// CodeAssetCreation identifies data as a AssetCreation message.
	CodeAssetCreation = "A2"

	// CodeAssetModification identifies data as a AssetModification message.
	CodeAssetModification = "A3"

	// CodeContractOffer identifies data as a ContractOffer message.
	CodeContractOffer = "C1"

	// CodeContractFormation identifies data as a ContractFormation message.
	CodeContractFormation = "C2"

	// CodeContractAmendment identifies data as a ContractAmendment message.
	CodeContractAmendment = "C3"

	// CodeStaticContractFormation identifies data as a
	// StaticContractFormation message.
	CodeStaticContractFormation = "C4"

	// CodeOrder identifies data as a Order message.
	CodeOrder = "E1"

	// CodeFreeze identifies data as a Freeze message.
	CodeFreeze = "E2"

	// CodeThaw identifies data as a Thaw message.
	CodeThaw = "E3"

	// CodeConfiscation identifies data as a Confiscation message.
	CodeConfiscation = "E4"

	// CodeReconciliation identifies data as a Reconciliation message.
	CodeReconciliation = "E5"

	// CodeInitiative identifies data as a Initiative message.
	CodeInitiative = "G1"

	// CodeReferendum identifies data as a Referendum message.
	CodeReferendum = "G2"

	// CodeVote identifies data as a Vote message.
	CodeVote = "G3"

	// CodeBallotCast identifies data as a BallotCast message.
	CodeBallotCast = "G4"

	// CodeBallotCounted identifies data as a BallotCounted message.
	CodeBallotCounted = "G5"

	// CodeResult identifies data as a Result message.
	CodeResult = "G6"

	// CodeMessage identifies data as a Message message.
	CodeMessage = "M1"

	// CodeRejection identifies data as a Rejection message.
	CodeRejection = "M2"

	// CodeEstablishment identifies data as a Establishment message.
	CodeEstablishment = "R1"

	// CodeAddition identifies data as a Addition message.
	CodeAddition = "R2"

	// CodeAlteration identifies data as a Alteration message.
	CodeAlteration = "R3"

	// CodeRemoval identifies data as a Removal message.
	CodeRemoval = "R4"

	// CodeTransfer identifies data as a Transfer message.
	CodeTransfer = "T1"

	// CodeSettlement identifies data as a Settlement message.
	CodeSettlement = "T4"

	// ComplianceActionFreeze identifies a freeze type
	ComplianceActionFreeze = byte('F')

	// ComplianceActionThaw identifies a thaw type
	ComplianceActionThaw = byte('T')

	// ComplianceActionConfiscation identifies a confiscation type
	ComplianceActionConfiscation = byte('C')

	// ComplianceActionReconciliation identifies a reconcilation type
	ComplianceActionReconciliation = byte('R')
)

// TypeMapping holds a mapping of action codes to action types.
func TypeMapping(code string) OpReturnMessage {
	switch code {
	case CodeAssetDefinition:
		result := AssetDefinition{}
		return &result
	case CodeAssetCreation:
		result := AssetCreation{}
		return &result
	case CodeAssetModification:
		result := AssetModification{}
		return &result
	case CodeContractOffer:
		result := ContractOffer{}
		return &result
	case CodeContractFormation:
		result := ContractFormation{}
		return &result
	case CodeContractAmendment:
		result := ContractAmendment{}
		return &result
	case CodeStaticContractFormation:
		result := StaticContractFormation{}
		return &result
	case CodeOrder:
		result := Order{}
		return &result
	case CodeFreeze:
		result := Freeze{}
		return &result
	case CodeThaw:
		result := Thaw{}
		return &result
	case CodeConfiscation:
		result := Confiscation{}
		return &result
	case CodeReconciliation:
		result := Reconciliation{}
		return &result
	case CodeInitiative:
		result := Initiative{}
		return &result
	case CodeReferendum:
		result := Referendum{}
		return &result
	case CodeVote:
		result := Vote{}
		return &result
	case CodeBallotCast:
		result := BallotCast{}
		return &result
	case CodeBallotCounted:
		result := BallotCounted{}
		return &result
	case CodeResult:
		result := Result{}
		return &result
	case CodeMessage:
		result := Message{}
		return &result
	case CodeRejection:
		result := Rejection{}
		return &result
	case CodeEstablishment:
		result := Establishment{}
		return &result
	case CodeAddition:
		result := Addition{}
		return &result
	case CodeAlteration:
		result := Alteration{}
		return &result
	case CodeRemoval:
		result := Removal{}
		return &result
	case CodeTransfer:
		result := Transfer{}
		return &result
	case CodeSettlement:
		result := Settlement{}
		return &result
	default:
		return nil
	}
}

// AssetDefinition Asset Definition Action - This action is used by the
// issuer to define the properties/characteristics of the Asset (token)
// that it wants to create.
type AssetDefinition struct {
	Header                      Header    `json:"header,omitempty"`                        // Common header data for all actions
	AssetType                   string    `json:"asset_type,omitempty"`                    // eg. Share - Common
	AssetCode                   AssetCode `json:"asset_code,omitempty"`                    // 32 randomly generated bytes.  Each Asset Code should be unique.  However, an Asset Code is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type + Asset Code = Asset Code.  An Asset Code is a human readable identifier that can be used in a similar way to a Bitcoin (BSV) address.
	AssetAuthFlags              [8]byte   `json:"asset_auth_flags,omitempty"`              // Authorization Flags,  bitwise operation
	TransfersPermitted          bool      `json:"transfers_permitted,omitempty"`           // 1 = Transfers are permitted.  0 = Transfers are not permitted.
	TradeRestrictions           string    `json:"trade_restrictions,omitempty"`            // Asset can only be traded within the trade restrictions.  Eg. AUS - Australian residents only.  EU - European Union residents only.
	EnforcementOrdersPermitted  bool      `json:"enforcement_orders_permitted,omitempty"`  // 1 = Enforcement Orders are permitted. 0 = Enforcement Orders are not permitted.
	VoteMultiplier              uint8     `json:"vote_multiplier,omitempty"`               // Multiplies the vote by the integer. 1 token = 1 vote with a 1 for vote multipler (normal).  1 token = 3 votes with a multiplier of 3, for example.
	ReferendumProposal          bool      `json:"referendum_proposal,omitempty"`           // A Referendum is permitted for Asset-Wide Proposals (outside of smart contract scope) if also permitted by the contract. If the contract has proposals by referendum restricted, then this flag is meaningless.
	InitiativeProposal          bool      `json:"initiative_proposal,omitempty"`           // An initiative is permitted for Asset-Wide Proposals (outside of smart contract scope) if also permitted by the contract. If the contract has proposals by initiative restricted, then this flag is meaningless.
	AssetModificationGovernance bool      `json:"asset_modification_governance,omitempty"` // 1 - Contract-wide Asset Governance.  0 - Asset-wide Asset Governance.  If a referendum or initiative is used to propose a modification to a subfield controlled by the asset auth flags, then the vote will either be a contract-wide vote (all assets vote on the referendum/initiative) or an asset-wide vote (all assets vote on the referendum/initiative) depending on the value in this subfield.  The voting system specifies the voting rules.
	TokenQty                    uint64    `json:"token_qty,omitempty"`                     // Quantity of token - 0 is valid. Fungible 'shares' of the Asset. 1 is used for non-fungible tokens.  Asset Codes become the non-fungible Asset Code and many Asset Codes can be associated with a particular Contract.
	ContractFeeCurrency         string    `json:"contract_fee_currency,omitempty"`         // BSV, USD, AUD, EUR, etc.
	ContractFeeVar              float32   `json:"contract_fee_var,omitempty"`              // Percent of the value of the transaction
	ContractFeeFixed            float32   `json:"contract_fee_fixed,omitempty"`            // Fixed fee (payment made in BSV)
	AssetPayload                []byte    `json:"asset_payload,omitempty"`                 // Payload length is dependent on the asset type. Each asset is made up of a defined set of information pertaining to the specific asset type, and may contain fields of variable length type (nvarchar8, 16, 32)
}

// Type returns the type identifer for this message.
func (action AssetDefinition) Type() string {
	return CodeAssetDefinition
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *AssetDefinition) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *AssetDefinition) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// AssetType (string)
	// fmt.Printf("Serializing AssetType\n")
	if err := WriteFixedChar(buf, action.AssetType, 3); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AssetType : buf len %d\n", buf.Len())

	// AssetCode (AssetCode)
	// fmt.Printf("Serializing AssetCode\n")
	{
		b, err := action.AssetCode.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized AssetCode : buf len %d\n", buf.Len())

	// AssetAuthFlags ([8]byte)
	// fmt.Printf("Serializing AssetAuthFlags\n")
	if err := write(buf, action.AssetAuthFlags); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AssetAuthFlags : buf len %d\n", buf.Len())

	// TransfersPermitted (bool)
	// fmt.Printf("Serializing TransfersPermitted\n")
	if err := write(buf, action.TransfersPermitted); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized TransfersPermitted : buf len %d\n", buf.Len())

	// TradeRestrictions (string)
	// fmt.Printf("Serializing TradeRestrictions\n")
	if err := WriteFixedChar(buf, action.TradeRestrictions, 3); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized TradeRestrictions : buf len %d\n", buf.Len())

	// EnforcementOrdersPermitted (bool)
	// fmt.Printf("Serializing EnforcementOrdersPermitted\n")
	if err := write(buf, action.EnforcementOrdersPermitted); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized EnforcementOrdersPermitted : buf len %d\n", buf.Len())

	// VoteMultiplier (uint8)
	// fmt.Printf("Serializing VoteMultiplier\n")
	if err := write(buf, action.VoteMultiplier); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized VoteMultiplier : buf len %d\n", buf.Len())

	// ReferendumProposal (bool)
	// fmt.Printf("Serializing ReferendumProposal\n")
	if err := write(buf, action.ReferendumProposal); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ReferendumProposal : buf len %d\n", buf.Len())

	// InitiativeProposal (bool)
	// fmt.Printf("Serializing InitiativeProposal\n")
	if err := write(buf, action.InitiativeProposal); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized InitiativeProposal : buf len %d\n", buf.Len())

	// AssetModificationGovernance (bool)
	// fmt.Printf("Serializing AssetModificationGovernance\n")
	if err := write(buf, action.AssetModificationGovernance); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AssetModificationGovernance : buf len %d\n", buf.Len())

	// TokenQty (uint64)
	// fmt.Printf("Serializing TokenQty\n")
	if err := write(buf, action.TokenQty); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized TokenQty : buf len %d\n", buf.Len())

	// ContractFeeCurrency (string)
	// fmt.Printf("Serializing ContractFeeCurrency\n")
	if err := WriteFixedChar(buf, action.ContractFeeCurrency, 3); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractFeeCurrency : buf len %d\n", buf.Len())

	// ContractFeeVar (float32)
	// fmt.Printf("Serializing ContractFeeVar\n")
	if err := write(buf, action.ContractFeeVar); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractFeeVar : buf len %d\n", buf.Len())

	// ContractFeeFixed (float32)
	// fmt.Printf("Serializing ContractFeeFixed\n")
	if err := write(buf, action.ContractFeeFixed); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractFeeFixed : buf len %d\n", buf.Len())

	// AssetPayload ([]byte)
	// fmt.Printf("Serializing AssetPayload\n")
	if err := WriteVarBin(buf, action.AssetPayload, 16); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AssetPayload : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in AssetDefinition from the byte slice
func (action *AssetDefinition) write(b []byte) (int, error) {
	// fmt.Printf("Reading AssetDefinition : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// AssetType (string)
	// fmt.Printf("Reading AssetType : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read AssetType : %d bytes remaining\n%+v\n", buf.Len(), action.AssetType)

	// AssetCode (AssetCode)
	// fmt.Printf("Reading AssetCode : %d bytes remaining\n", buf.Len())
	if err := action.AssetCode.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read AssetCode : %d bytes remaining\n%+v\n", buf.Len(), action.AssetCode)

	// AssetAuthFlags ([8]byte)
	// fmt.Printf("Reading AssetAuthFlags : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.AssetAuthFlags); err != nil {
		return 0, err
	}

	// fmt.Printf("Read AssetAuthFlags : %d bytes remaining\n%+v\n", buf.Len(), action.AssetAuthFlags)

	// TransfersPermitted (bool)
	// fmt.Printf("Reading TransfersPermitted : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.TransfersPermitted); err != nil {
		return 0, err
	}

	// fmt.Printf("Read TransfersPermitted : %d bytes remaining\n%+v\n", buf.Len(), action.TransfersPermitted)

	// TradeRestrictions (string)
	// fmt.Printf("Reading TradeRestrictions : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.TradeRestrictions, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read TradeRestrictions : %d bytes remaining\n%+v\n", buf.Len(), action.TradeRestrictions)

	// EnforcementOrdersPermitted (bool)
	// fmt.Printf("Reading EnforcementOrdersPermitted : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.EnforcementOrdersPermitted); err != nil {
		return 0, err
	}

	// fmt.Printf("Read EnforcementOrdersPermitted : %d bytes remaining\n%+v\n", buf.Len(), action.EnforcementOrdersPermitted)

	// VoteMultiplier (uint8)
	// fmt.Printf("Reading VoteMultiplier : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.VoteMultiplier); err != nil {
		return 0, err
	}

	// fmt.Printf("Read VoteMultiplier : %d bytes remaining\n%+v\n", buf.Len(), action.VoteMultiplier)

	// ReferendumProposal (bool)
	// fmt.Printf("Reading ReferendumProposal : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ReferendumProposal); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ReferendumProposal : %d bytes remaining\n%+v\n", buf.Len(), action.ReferendumProposal)

	// InitiativeProposal (bool)
	// fmt.Printf("Reading InitiativeProposal : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.InitiativeProposal); err != nil {
		return 0, err
	}

	// fmt.Printf("Read InitiativeProposal : %d bytes remaining\n%+v\n", buf.Len(), action.InitiativeProposal)

	// AssetModificationGovernance (bool)
	// fmt.Printf("Reading AssetModificationGovernance : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.AssetModificationGovernance); err != nil {
		return 0, err
	}

	// fmt.Printf("Read AssetModificationGovernance : %d bytes remaining\n%+v\n", buf.Len(), action.AssetModificationGovernance)

	// TokenQty (uint64)
	// fmt.Printf("Reading TokenQty : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.TokenQty); err != nil {
		return 0, err
	}

	// fmt.Printf("Read TokenQty : %d bytes remaining\n%+v\n", buf.Len(), action.TokenQty)

	// ContractFeeCurrency (string)
	// fmt.Printf("Reading ContractFeeCurrency : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.ContractFeeCurrency, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read ContractFeeCurrency : %d bytes remaining\n%+v\n", buf.Len(), action.ContractFeeCurrency)

	// ContractFeeVar (float32)
	// fmt.Printf("Reading ContractFeeVar : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ContractFeeVar); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ContractFeeVar : %d bytes remaining\n%+v\n", buf.Len(), action.ContractFeeVar)

	// ContractFeeFixed (float32)
	// fmt.Printf("Reading ContractFeeFixed : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ContractFeeFixed); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ContractFeeFixed : %d bytes remaining\n%+v\n", buf.Len(), action.ContractFeeFixed)

	// AssetPayload ([]byte)
	// fmt.Printf("Reading AssetPayload : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.AssetPayload, err = ReadVarBin(buf, 16)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read AssetPayload : %d bytes remaining\n%+v\n", buf.Len(), action.AssetPayload)

	// fmt.Printf("Read AssetDefinition : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action AssetDefinition) PayloadMessage() (PayloadMessage, error) {
	p := AssetTypeMapping(action.AssetType)
	if p == nil {
		return nil, fmt.Errorf("Undefined asset type : %s", action.AssetType)
	}

	if _, err := p.Write(action.AssetPayload); err != nil {
		return nil, err
	}

	return p, nil
}

func (action AssetDefinition) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("AssetType:%#+v", action.AssetType))
	vals = append(vals, fmt.Sprintf("AssetCode:%#+v", action.AssetCode))
	vals = append(vals, fmt.Sprintf("AssetAuthFlags:%#+v", action.AssetAuthFlags))
	vals = append(vals, fmt.Sprintf("TransfersPermitted:%#+v", action.TransfersPermitted))
	vals = append(vals, fmt.Sprintf("TradeRestrictions:%#+v", action.TradeRestrictions))
	vals = append(vals, fmt.Sprintf("EnforcementOrdersPermitted:%#+v", action.EnforcementOrdersPermitted))
	vals = append(vals, fmt.Sprintf("VoteMultiplier:%v", action.VoteMultiplier))
	vals = append(vals, fmt.Sprintf("ReferendumProposal:%#+v", action.ReferendumProposal))
	vals = append(vals, fmt.Sprintf("InitiativeProposal:%#+v", action.InitiativeProposal))
	vals = append(vals, fmt.Sprintf("AssetModificationGovernance:%#+v", action.AssetModificationGovernance))
	vals = append(vals, fmt.Sprintf("TokenQty:%v", action.TokenQty))
	vals = append(vals, fmt.Sprintf("ContractFeeCurrency:%#+v", action.ContractFeeCurrency))
	vals = append(vals, fmt.Sprintf("ContractFeeVar:%v", action.ContractFeeVar))
	vals = append(vals, fmt.Sprintf("ContractFeeFixed:%v", action.ContractFeeFixed))
	vals = append(vals, fmt.Sprintf("AssetPayload:%#x", action.AssetPayload))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// AssetCreation Asset Creation Action - This action creates an Asset in
// response to the Issuer's instructions in the Definition Action.
type AssetCreation struct {
	Header                      Header    `json:"header,omitempty"`                        // Common header data for all actions
	AssetType                   string    `json:"asset_type,omitempty"`                    // eg. Share - Common
	AssetCode                   AssetCode `json:"asset_code,omitempty"`                    // 32 randomly generated bytes.  Each Asset Code should be unique.  However, an Asset Code is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type + Asset Code = Asset Code.  An Asset Code is a human readable identifier that can be used in a similar way to a Bitcoin (BSV) address.
	AssetAuthFlags              [8]byte   `json:"asset_auth_flags,omitempty"`              // Authorization Flags,  bitwise operation
	TransfersPermitted          bool      `json:"transfers_permitted,omitempty"`           // 1 = Transfers are permitted.  0 = Transfers are not permitted.
	TradeRestrictions           string    `json:"trade_restrictions,omitempty"`            // Asset can only be traded within the trade restrictions.  Eg. AUS - Australian residents only.  EU - European Union residents only.
	EnforcementOrdersPermitted  bool      `json:"enforcement_orders_permitted,omitempty"`  // 1 = Enforcement Orders are permitted. 0 = Enforcement Orders are not permitted.
	VoteMultiplier              uint8     `json:"vote_multiplier,omitempty"`               // Multiplies the vote by the integer. 1 token = 1 vote with a 1 for vote multipler (normal).  1 token = 3 votes with a multiplier of 3, for example.
	ReferendumProposal          bool      `json:"referendum_proposal,omitempty"`           // A Referendum is permitted for Asset-Wide Proposals (outside of smart contract scope) if also permitted by the contract. If the contract has proposals by referendum restricted, then this flag is meaningless.
	InitiativeProposal          bool      `json:"initiative_proposal,omitempty"`           // An initiative is permitted for Asset-Wide Proposals (outside of smart contract scope) if also permitted by the contract. If the contract has proposals by initiative restricted, then this flag is meaningless.
	AssetModificationGovernance bool      `json:"asset_modification_governance,omitempty"` // 1 - Contract-wide Asset Governance.  0 - Asset-wide Asset Governance.  If a referendum or initiative is used to propose a modification to a subfield controlled by the asset auth flags, then the vote will either be a contract-wide vote (all assets vote on the referendum/initiative) or an asset-wide vote (all assets vote on the referendum/initiative).  The voting system specifies the voting rules.
	TokenQty                    uint64    `json:"token_qty,omitempty"`                     // Quantity of token - 0 is valid. Fungible 'shares' of the Asset. 1 is used for non-fungible tokens.  Asset Codes become the non-fungible Asset Code and many Asset Codes can be associated with a particular Contract.
	ContractFeeCurrency         string    `json:"contract_fee_currency,omitempty"`         // BSV, USD, AUD, EUR, etc.
	ContractFeeVar              float32   `json:"contract_fee_var,omitempty"`              // Percent of the value of the transaction
	ContractFeeFixed            float32   `json:"contract_fee_fixed,omitempty"`            // Fixed fee (payment made in BSV)
	AssetPayload                []byte    `json:"asset_payload,omitempty"`                 // Payload length is dependent on the asset type. Each asset is made up of a defined set of information pertaining to the specific asset type, and may contain fields of variable length type (nvarchar8, 16, 32)
	AssetRevision               uint32    `json:"asset _revision,omitempty"`               // Counter 0 to (2^32)-1
	Timestamp                   Timestamp `json:"timestamp,omitempty"`                     // Timestamp in nanoseconds of when the smart contract created the action.
}

// Type returns the type identifer for this message.
func (action AssetCreation) Type() string {
	return CodeAssetCreation
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *AssetCreation) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *AssetCreation) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// AssetType (string)
	// fmt.Printf("Serializing AssetType\n")
	if err := WriteFixedChar(buf, action.AssetType, 3); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AssetType : buf len %d\n", buf.Len())

	// AssetCode (AssetCode)
	// fmt.Printf("Serializing AssetCode\n")
	{
		b, err := action.AssetCode.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized AssetCode : buf len %d\n", buf.Len())

	// AssetAuthFlags ([8]byte)
	// fmt.Printf("Serializing AssetAuthFlags\n")
	if err := write(buf, action.AssetAuthFlags); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AssetAuthFlags : buf len %d\n", buf.Len())

	// TransfersPermitted (bool)
	// fmt.Printf("Serializing TransfersPermitted\n")
	if err := write(buf, action.TransfersPermitted); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized TransfersPermitted : buf len %d\n", buf.Len())

	// TradeRestrictions (string)
	// fmt.Printf("Serializing TradeRestrictions\n")
	if err := WriteFixedChar(buf, action.TradeRestrictions, 3); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized TradeRestrictions : buf len %d\n", buf.Len())

	// EnforcementOrdersPermitted (bool)
	// fmt.Printf("Serializing EnforcementOrdersPermitted\n")
	if err := write(buf, action.EnforcementOrdersPermitted); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized EnforcementOrdersPermitted : buf len %d\n", buf.Len())

	// VoteMultiplier (uint8)
	// fmt.Printf("Serializing VoteMultiplier\n")
	if err := write(buf, action.VoteMultiplier); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized VoteMultiplier : buf len %d\n", buf.Len())

	// ReferendumProposal (bool)
	// fmt.Printf("Serializing ReferendumProposal\n")
	if err := write(buf, action.ReferendumProposal); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ReferendumProposal : buf len %d\n", buf.Len())

	// InitiativeProposal (bool)
	// fmt.Printf("Serializing InitiativeProposal\n")
	if err := write(buf, action.InitiativeProposal); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized InitiativeProposal : buf len %d\n", buf.Len())

	// AssetModificationGovernance (bool)
	// fmt.Printf("Serializing AssetModificationGovernance\n")
	if err := write(buf, action.AssetModificationGovernance); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AssetModificationGovernance : buf len %d\n", buf.Len())

	// TokenQty (uint64)
	// fmt.Printf("Serializing TokenQty\n")
	if err := write(buf, action.TokenQty); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized TokenQty : buf len %d\n", buf.Len())

	// ContractFeeCurrency (string)
	// fmt.Printf("Serializing ContractFeeCurrency\n")
	if err := WriteFixedChar(buf, action.ContractFeeCurrency, 3); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractFeeCurrency : buf len %d\n", buf.Len())

	// ContractFeeVar (float32)
	// fmt.Printf("Serializing ContractFeeVar\n")
	if err := write(buf, action.ContractFeeVar); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractFeeVar : buf len %d\n", buf.Len())

	// ContractFeeFixed (float32)
	// fmt.Printf("Serializing ContractFeeFixed\n")
	if err := write(buf, action.ContractFeeFixed); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractFeeFixed : buf len %d\n", buf.Len())

	// AssetPayload ([]byte)
	// fmt.Printf("Serializing AssetPayload\n")
	if err := WriteVarBin(buf, action.AssetPayload, 16); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AssetPayload : buf len %d\n", buf.Len())

	// AssetRevision (uint32)
	// fmt.Printf("Serializing AssetRevision\n")
	if err := write(buf, action.AssetRevision); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AssetRevision : buf len %d\n", buf.Len())

	// Timestamp (Timestamp)
	// fmt.Printf("Serializing Timestamp\n")
	{
		b, err := action.Timestamp.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Timestamp : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in AssetCreation from the byte slice
func (action *AssetCreation) write(b []byte) (int, error) {
	// fmt.Printf("Reading AssetCreation : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// AssetType (string)
	// fmt.Printf("Reading AssetType : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read AssetType : %d bytes remaining\n%+v\n", buf.Len(), action.AssetType)

	// AssetCode (AssetCode)
	// fmt.Printf("Reading AssetCode : %d bytes remaining\n", buf.Len())
	if err := action.AssetCode.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read AssetCode : %d bytes remaining\n%+v\n", buf.Len(), action.AssetCode)

	// AssetAuthFlags ([8]byte)
	// fmt.Printf("Reading AssetAuthFlags : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.AssetAuthFlags); err != nil {
		return 0, err
	}

	// fmt.Printf("Read AssetAuthFlags : %d bytes remaining\n%+v\n", buf.Len(), action.AssetAuthFlags)

	// TransfersPermitted (bool)
	// fmt.Printf("Reading TransfersPermitted : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.TransfersPermitted); err != nil {
		return 0, err
	}

	// fmt.Printf("Read TransfersPermitted : %d bytes remaining\n%+v\n", buf.Len(), action.TransfersPermitted)

	// TradeRestrictions (string)
	// fmt.Printf("Reading TradeRestrictions : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.TradeRestrictions, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read TradeRestrictions : %d bytes remaining\n%+v\n", buf.Len(), action.TradeRestrictions)

	// EnforcementOrdersPermitted (bool)
	// fmt.Printf("Reading EnforcementOrdersPermitted : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.EnforcementOrdersPermitted); err != nil {
		return 0, err
	}

	// fmt.Printf("Read EnforcementOrdersPermitted : %d bytes remaining\n%+v\n", buf.Len(), action.EnforcementOrdersPermitted)

	// VoteMultiplier (uint8)
	// fmt.Printf("Reading VoteMultiplier : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.VoteMultiplier); err != nil {
		return 0, err
	}

	// fmt.Printf("Read VoteMultiplier : %d bytes remaining\n%+v\n", buf.Len(), action.VoteMultiplier)

	// ReferendumProposal (bool)
	// fmt.Printf("Reading ReferendumProposal : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ReferendumProposal); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ReferendumProposal : %d bytes remaining\n%+v\n", buf.Len(), action.ReferendumProposal)

	// InitiativeProposal (bool)
	// fmt.Printf("Reading InitiativeProposal : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.InitiativeProposal); err != nil {
		return 0, err
	}

	// fmt.Printf("Read InitiativeProposal : %d bytes remaining\n%+v\n", buf.Len(), action.InitiativeProposal)

	// AssetModificationGovernance (bool)
	// fmt.Printf("Reading AssetModificationGovernance : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.AssetModificationGovernance); err != nil {
		return 0, err
	}

	// fmt.Printf("Read AssetModificationGovernance : %d bytes remaining\n%+v\n", buf.Len(), action.AssetModificationGovernance)

	// TokenQty (uint64)
	// fmt.Printf("Reading TokenQty : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.TokenQty); err != nil {
		return 0, err
	}

	// fmt.Printf("Read TokenQty : %d bytes remaining\n%+v\n", buf.Len(), action.TokenQty)

	// ContractFeeCurrency (string)
	// fmt.Printf("Reading ContractFeeCurrency : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.ContractFeeCurrency, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read ContractFeeCurrency : %d bytes remaining\n%+v\n", buf.Len(), action.ContractFeeCurrency)

	// ContractFeeVar (float32)
	// fmt.Printf("Reading ContractFeeVar : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ContractFeeVar); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ContractFeeVar : %d bytes remaining\n%+v\n", buf.Len(), action.ContractFeeVar)

	// ContractFeeFixed (float32)
	// fmt.Printf("Reading ContractFeeFixed : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ContractFeeFixed); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ContractFeeFixed : %d bytes remaining\n%+v\n", buf.Len(), action.ContractFeeFixed)

	// AssetPayload ([]byte)
	// fmt.Printf("Reading AssetPayload : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.AssetPayload, err = ReadVarBin(buf, 16)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read AssetPayload : %d bytes remaining\n%+v\n", buf.Len(), action.AssetPayload)

	// AssetRevision (uint32)
	// fmt.Printf("Reading AssetRevision : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.AssetRevision); err != nil {
		return 0, err
	}

	// fmt.Printf("Read AssetRevision : %d bytes remaining\n%+v\n", buf.Len(), action.AssetRevision)

	// Timestamp (Timestamp)
	// fmt.Printf("Reading Timestamp : %d bytes remaining\n", buf.Len())
	if err := action.Timestamp.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Timestamp : %d bytes remaining\n%+v\n", buf.Len(), action.Timestamp)

	// fmt.Printf("Read AssetCreation : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action AssetCreation) PayloadMessage() (PayloadMessage, error) {
	p := AssetTypeMapping(action.AssetType)
	if p == nil {
		return nil, fmt.Errorf("Undefined asset type : %s", action.AssetType)
	}

	if _, err := p.Write(action.AssetPayload); err != nil {
		return nil, err
	}

	return p, nil
}

func (action AssetCreation) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("AssetType:%#+v", action.AssetType))
	vals = append(vals, fmt.Sprintf("AssetCode:%#+v", action.AssetCode))
	vals = append(vals, fmt.Sprintf("AssetAuthFlags:%#+v", action.AssetAuthFlags))
	vals = append(vals, fmt.Sprintf("TransfersPermitted:%#+v", action.TransfersPermitted))
	vals = append(vals, fmt.Sprintf("TradeRestrictions:%#+v", action.TradeRestrictions))
	vals = append(vals, fmt.Sprintf("EnforcementOrdersPermitted:%#+v", action.EnforcementOrdersPermitted))
	vals = append(vals, fmt.Sprintf("VoteMultiplier:%v", action.VoteMultiplier))
	vals = append(vals, fmt.Sprintf("ReferendumProposal:%#+v", action.ReferendumProposal))
	vals = append(vals, fmt.Sprintf("InitiativeProposal:%#+v", action.InitiativeProposal))
	vals = append(vals, fmt.Sprintf("AssetModificationGovernance:%#+v", action.AssetModificationGovernance))
	vals = append(vals, fmt.Sprintf("TokenQty:%v", action.TokenQty))
	vals = append(vals, fmt.Sprintf("ContractFeeCurrency:%#+v", action.ContractFeeCurrency))
	vals = append(vals, fmt.Sprintf("ContractFeeVar:%v", action.ContractFeeVar))
	vals = append(vals, fmt.Sprintf("ContractFeeFixed:%v", action.ContractFeeFixed))
	vals = append(vals, fmt.Sprintf("AssetPayload:%#x", action.AssetPayload))
	vals = append(vals, fmt.Sprintf("AssetRevision:%v", action.AssetRevision))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", action.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// AssetModification Asset Modification Action - Token Dilutions, Call
// Backs/Revocations, burning etc.
type AssetModification struct {
	Header        Header      `json:"header,omitempty"`         // Common header data for all actions
	AssetType     string      `json:"asset_type,omitempty"`     // eg. Share - Common
	AssetCode     AssetCode   `json:"asset_code,omitempty"`     // 32 randomly generated bytes.  Each Asset Code should be unique.  However, an Asset Code is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type + Asset Code = Asset Code.  An Asset Code is a human readable identifier that can be used in a similar way to a Bitcoin (BSV) address.
	AssetRevision uint32      `json:"asset_revision,omitempty"` // Counter. (Subfield cannot be manually changed by Asset Modification Action.  Only SC can increment by 1 with each AC action. SC will reject AM actions where the wrong asset revision has been selected.
	Modifications []Amendment `json:"modifications,omitempty"`  //
	RefTxID       TxId        `json:"ref_tx_id,omitempty"`      // Tx-ID of the associated Result action (governance) that permitted the modifications.
}

// Type returns the type identifer for this message.
func (action AssetModification) Type() string {
	return CodeAssetModification
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *AssetModification) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *AssetModification) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// AssetType (string)
	// fmt.Printf("Serializing AssetType\n")
	if err := WriteFixedChar(buf, action.AssetType, 3); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AssetType : buf len %d\n", buf.Len())

	// AssetCode (AssetCode)
	// fmt.Printf("Serializing AssetCode\n")
	{
		b, err := action.AssetCode.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized AssetCode : buf len %d\n", buf.Len())

	// AssetRevision (uint32)
	// fmt.Printf("Serializing AssetRevision\n")
	if err := write(buf, action.AssetRevision); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AssetRevision : buf len %d\n", buf.Len())

	// Modifications ([]Amendment)
	// fmt.Printf("Serializing Modifications\n")
	if err := WriteVariableSize(buf, uint64(len(action.Modifications)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.Modifications {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Modifications : buf len %d\n", buf.Len())

	// RefTxID (TxId)
	// fmt.Printf("Serializing RefTxID\n")
	{
		b, err := action.RefTxID.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized RefTxID : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in AssetModification from the byte slice
func (action *AssetModification) write(b []byte) (int, error) {
	// fmt.Printf("Reading AssetModification : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// AssetType (string)
	// fmt.Printf("Reading AssetType : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read AssetType : %d bytes remaining\n%+v\n", buf.Len(), action.AssetType)

	// AssetCode (AssetCode)
	// fmt.Printf("Reading AssetCode : %d bytes remaining\n", buf.Len())
	if err := action.AssetCode.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read AssetCode : %d bytes remaining\n%+v\n", buf.Len(), action.AssetCode)

	// AssetRevision (uint32)
	// fmt.Printf("Reading AssetRevision : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.AssetRevision); err != nil {
		return 0, err
	}

	// fmt.Printf("Read AssetRevision : %d bytes remaining\n%+v\n", buf.Len(), action.AssetRevision)

	// Modifications ([]Amendment)
	// fmt.Printf("Reading Modifications : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.Modifications = make([]Amendment, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Amendment
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.Modifications = append(action.Modifications, newValue)
		}
	}

	// fmt.Printf("Read Modifications : %d bytes remaining\n%+v\n", buf.Len(), action.Modifications)

	// RefTxID (TxId)
	// fmt.Printf("Reading RefTxID : %d bytes remaining\n", buf.Len())
	if err := action.RefTxID.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read RefTxID : %d bytes remaining\n%+v\n", buf.Len(), action.RefTxID)

	// fmt.Printf("Read AssetModification : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action AssetModification) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action AssetModification) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("AssetType:%#+v", action.AssetType))
	vals = append(vals, fmt.Sprintf("AssetCode:%#+v", action.AssetCode))
	vals = append(vals, fmt.Sprintf("AssetRevision:%v", action.AssetRevision))
	vals = append(vals, fmt.Sprintf("Modifications:%#+v", action.Modifications))
	vals = append(vals, fmt.Sprintf("RefTxID:%#+v", action.RefTxID))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// ContractOffer The Contract Offer action allows the Issuer to tell the
// smart contract what they want the details (labels, data, T&C's, etc.) of
// the Contract to be on-chain in a public and immutable way. The Contract
// Offer action 'initializes' a generic smart contract that has been spun
// up by either the Smart Contract Operator or the Issuer. This on-chain
// action allows for the positive response from the smart contract with
// either a Contract Formation Action or a Rejection Action.
type ContractOffer struct {
	Header                     Header         `json:"header,omitempty"`                        // Common header data for all actions
	ContractName               string         `json:"contract_name,omitempty"`                 // Can be any unique identifying string, including human readable names for branding/vanity purposes.   [Contract identifier (instance) is the bitcoin public key hash address. If the Public Address is lost, then the issuer will have to reissue the entire contract, Asset definition and tokens with the new public address.]. Smart contracts can be branded and specialized to suit any terms and conditions.
	ContractFileType           uint8          `json:"contract_file_type,omitempty"`            // 1 - SHA-256 Hash, 2 - Markdown
	ContractFile               []byte         `json:"contract_file,omitempty"`                 // SHA-256 hash of the contract file or markdown data for contract file specific to the smart contract and relevant Assets.  Legal and technical information. (eg. pdf)
	SupportingDocsFileType     uint8          `json:"supporting_docs_file_type,omitempty"`     // 1 - 7z
	SupportingDocs             string         `json:"supporting_docs,omitempty"`               //
	GoverningLaw               string         `json:"governing_law,omitempty"`                 // 5 Letter Code to Identify which governing law the contract will adhere to.  Disputes are to be settled by this law in the jurisdiction specified below. Private dispute resolution organizations can be used as well.  A custom code just needs to be defined.
	Jurisdiction               string         `json:"jurisdiction,omitempty"`                  // Legal proceedings/arbitration will take place using the specified Governing Law in this location.
	ContractExpiration         Timestamp      `json:"contract_expiration,omitempty"`           // All actions related to the contract will cease to work after this timestamp. The smart contract will stop running.  This will allow many token use cases to be able to calculate total smart contract running costs for the entire life of the contract. Eg. an issuer is creating tickets for an event on the 5th of June 2018.  The smart contract will facilitate exchange and send transactions up until the 6th of June.  Wallets can use this to forget tokens that are no longer valid - or at least store them in an 'Expired' folder.
	ContractURI                string         `json:"contract_uri,omitempty"`                  // Points to an information page that also has a copy of the Contract.  Anyone can go to the website to have a look at the price/token, information about the Issuer (company), information about the Asset, legal information, etc.  There will also be a way for Token Owners to vote on this page and contact details with the Issuer/tokenized companies. Could be a IPv6/IPv4, an IPFS address (hash) or txn-id for on-chain information or even a public address (DNS).
	IssuerName                 string         `json:"issuer_name,omitempty"`                   // Length 0-255 bytes. 0 is not valid.Issuing entity (company, organization, individual).  Can be any unique identifying string, including human readable names for branding/vanity purposes.
	IssuerType                 byte           `json:"issuer_type,omitempty"`                   // P - Public Company Limited by Shares, C - Private Company Limited by Shares, I - Individual, L - Limited Partnership, U -Unlimited Partnership, T - Sole Proprietorship, S - Statutory Company, O - Non-Profit Organization, N - Nation State, G - Government Agency, U - Unit Trust, D - Discretionary Trust.  Found in 'Entities' (Specification/Resources).
	IssuerLEI                  string         `json:"issuer_lei,omitempty"`                    // Null is valid. A Legal Entity Identifier (or LEI) is an international identifier made up of a 20-character identifier that identifies distinct legal entities that engage in financial transactions. It is defined by ISO 17442.[1] Natural persons are not required to have an LEI; they’re eligible to have one issued, however, but only if they act in an independent business capacity.[2] The LEI is a global standard, designed to be non-proprietary data that is freely accessible to all.[3] As of December 2018, over 1,300,000 legal entities from more than 200 countries have now been issued with LEIs.
	IssuerLogoURL              string         `json:"issuer_logo_url,omitempty"`               // The URL of the Issuers logo.
	ContractOperatorID         string         `json:"contract_operator_id,omitempty"`          // Length 0-255 bytes. 0 is valid. Smart Contract Operator identifier. Can be any unique identifying string, including human readable names for branding/vanity purposes. Can also be null or the Issuer.
	OperatorLEI                string         `json:"operator_lei,omitempty"`                  // Null is valid. A Legal Entity Identifier (or LEI) is an international identifier made up of a 20-character identifier that identifies distinct legal entities that engage in financial transactions. It is defined by ISO 17442.[1] Natural persons are not required to have an LEI; they’re eligible to have one issued, however, but only if they act in an independent business capacity.[2] The LEI is a global standard, designed to be non-proprietary data that is freely accessible to all.[3] As of December 2018, over 1,300,000 legal entities from more than 200 countries have now been issued with LEIs.
	ContractAuthFlags          [16]byte       `json:"contract_auth_flags,omitempty"`           // Authorization Flags aka Terms and Conditions that the smart contract can enforce.  Other terms and conditions that are out of the smart contract's control are listed in the actual Contract File.
	VotingSystems              []VotingSystem `json:"voting_systems,omitempty"`                // A list of voting systems.
	RestrictedQtyAssets        uint64         `json:"restricted_qty_assets,omitempty"`         // Number of Assets (non-fungible) permitted on this contract. 0 if unlimited which will display an infinity symbol in UI
	ReferendumProposal         bool           `json:"referendum_proposal,omitempty"`           // A Referendum is permitted for Proposals (outside of smart contract scope).
	InitiativeProposal         bool           `json:"initiative_proposal,omitempty"`           // An initiative is permitted for Proposals (outside of smart contract scope).
	Registries                 []Registry     `json:"registries,omitempty"`                    // A list Registries
	IssuerAddress              bool           `json:"issuer_address,omitempty"`                // Physical/mailing address. Y/N, N means there is no issuer address.
	UnitNumber                 string         `json:"unit_number,omitempty"`                   // Issuer Address Details (eg. HQ)
	BuildingNumber             string         `json:"building_number,omitempty"`               //
	Street                     string         `json:"street,omitempty"`                        //
	SuburbCity                 string         `json:"suburb_city,omitempty"`                   //
	TerritoryStateProvinceCode string         `json:"territory_state_province_code,omitempty"` //
	CountryCode                string         `json:"country_code,omitempty"`                  //
	PostalZIPCode              string         `json:"postal_zip_code,omitempty"`               //
	EmailAddress               string         `json:"email_address,omitempty"`                 // Length 0-255 bytes. 0 is valid (no ContactAddress). Address for text-based communication: eg. email address, Bitcoin address
	PhoneNumber                string         `json:"phone_number,omitempty"`                  // Length 0-50 bytes. 0 is valid (no Phone subfield).Phone Number for Entity.
	KeyRoles                   []KeyRole      `json:"key_roles,omitempty"`                     // A list of Key Roles.
	NotableRoles               []NotableRole  `json:"notable_roles,omitempty"`                 // A list of Notable Roles.
}

// SetIssuerAddress Sets the issuer's mailing address on the contract..
func (action *ContractOffer) SetIssuerAddress(unit string, building string, street string, city string, state string, countryCode string, postalCode string) {
	action.UnitNumber = unit
	action.BuildingNumber = building
	action.Street = street
	action.SuburbCity = city
	action.TerritoryStateProvinceCode = state
	action.CountryCode = countryCode
	action.PostalZIPCode = postalCode
	action.IssuerAddress = true
}

// AddKeyRole Adds a key role to the contract.
func (action *ContractOffer) AddKeyRole(roleType uint8, name string) {
	action.KeyRoles = append(action.KeyRoles, *NewKeyRole(roleType, name))
}

// AddNotableRole Adds a notable role to the contract.
func (action *ContractOffer) AddNotableRole(roleType uint8, name string) {
	action.NotableRoles = append(action.NotableRoles, *NewNotableRole(roleType, name))
}

// Type returns the type identifer for this message.
func (action ContractOffer) Type() string {
	return CodeContractOffer
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *ContractOffer) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *ContractOffer) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// ContractName (string)
	// fmt.Printf("Serializing ContractName\n")
	if err := WriteVarChar(buf, action.ContractName, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractName : buf len %d\n", buf.Len())

	// ContractFileType (uint8)
	// fmt.Printf("Serializing ContractFileType\n")
	if err := write(buf, action.ContractFileType); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractFileType : buf len %d\n", buf.Len())

	// ContractFile ([]byte)
	// fmt.Printf("Serializing ContractFile\n")
	if err := WriteVarBin(buf, action.ContractFile, 32); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractFile : buf len %d\n", buf.Len())

	// SupportingDocsFileType (uint8)
	// fmt.Printf("Serializing SupportingDocsFileType\n")
	if err := write(buf, action.SupportingDocsFileType); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized SupportingDocsFileType : buf len %d\n", buf.Len())

	// SupportingDocs (string)
	// fmt.Printf("Serializing SupportingDocs\n")
	if err := WriteVarChar(buf, action.SupportingDocs, 32); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized SupportingDocs : buf len %d\n", buf.Len())

	// GoverningLaw (string)
	// fmt.Printf("Serializing GoverningLaw\n")
	if err := WriteFixedChar(buf, action.GoverningLaw, 5); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized GoverningLaw : buf len %d\n", buf.Len())

	// Jurisdiction (string)
	// fmt.Printf("Serializing Jurisdiction\n")
	if err := WriteFixedChar(buf, action.Jurisdiction, 5); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized Jurisdiction : buf len %d\n", buf.Len())

	// ContractExpiration (Timestamp)
	// fmt.Printf("Serializing ContractExpiration\n")
	{
		b, err := action.ContractExpiration.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized ContractExpiration : buf len %d\n", buf.Len())

	// ContractURI (string)
	// fmt.Printf("Serializing ContractURI\n")
	if err := WriteVarChar(buf, action.ContractURI, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractURI : buf len %d\n", buf.Len())

	// IssuerName (string)
	// fmt.Printf("Serializing IssuerName\n")
	if err := WriteVarChar(buf, action.IssuerName, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized IssuerName : buf len %d\n", buf.Len())

	// IssuerType (byte)
	// fmt.Printf("Serializing IssuerType\n")
	if err := write(buf, action.IssuerType); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized IssuerType : buf len %d\n", buf.Len())

	// IssuerLEI (string)
	// fmt.Printf("Serializing IssuerLEI\n")
	if err := WriteFixedChar(buf, action.IssuerLEI, 20); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized IssuerLEI : buf len %d\n", buf.Len())

	// IssuerLogoURL (string)
	// fmt.Printf("Serializing IssuerLogoURL\n")
	if err := WriteVarChar(buf, action.IssuerLogoURL, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized IssuerLogoURL : buf len %d\n", buf.Len())

	// ContractOperatorID (string)
	// fmt.Printf("Serializing ContractOperatorID\n")
	if err := WriteVarChar(buf, action.ContractOperatorID, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractOperatorID : buf len %d\n", buf.Len())

	// OperatorLEI (string)
	// fmt.Printf("Serializing OperatorLEI\n")
	if err := WriteFixedChar(buf, action.OperatorLEI, 20); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized OperatorLEI : buf len %d\n", buf.Len())

	// ContractAuthFlags ([16]byte)
	// fmt.Printf("Serializing ContractAuthFlags\n")
	if err := write(buf, action.ContractAuthFlags); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractAuthFlags : buf len %d\n", buf.Len())

	// VotingSystems ([]VotingSystem)
	// fmt.Printf("Serializing VotingSystems\n")
	if err := WriteVariableSize(buf, uint64(len(action.VotingSystems)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.VotingSystems {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized VotingSystems : buf len %d\n", buf.Len())

	// RestrictedQtyAssets (uint64)
	// fmt.Printf("Serializing RestrictedQtyAssets\n")
	if err := write(buf, action.RestrictedQtyAssets); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized RestrictedQtyAssets : buf len %d\n", buf.Len())

	// ReferendumProposal (bool)
	// fmt.Printf("Serializing ReferendumProposal\n")
	if err := write(buf, action.ReferendumProposal); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ReferendumProposal : buf len %d\n", buf.Len())

	// InitiativeProposal (bool)
	// fmt.Printf("Serializing InitiativeProposal\n")
	if err := write(buf, action.InitiativeProposal); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized InitiativeProposal : buf len %d\n", buf.Len())

	// Registries ([]Registry)
	// fmt.Printf("Serializing Registries\n")
	if err := WriteVariableSize(buf, uint64(len(action.Registries)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.Registries {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Registries : buf len %d\n", buf.Len())

	// IssuerAddress (bool)
	// fmt.Printf("Serializing IssuerAddress\n")
	if err := write(buf, action.IssuerAddress); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized IssuerAddress : buf len %d\n", buf.Len())

	// UnitNumber (string)
	// fmt.Printf("Serializing UnitNumber\n")
	if err := WriteVarChar(buf, action.UnitNumber, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized UnitNumber : buf len %d\n", buf.Len())

	// BuildingNumber (string)
	// fmt.Printf("Serializing BuildingNumber\n")
	if err := WriteVarChar(buf, action.BuildingNumber, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized BuildingNumber : buf len %d\n", buf.Len())

	// Street (string)
	// fmt.Printf("Serializing Street\n")
	if err := WriteVarChar(buf, action.Street, 16); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized Street : buf len %d\n", buf.Len())

	// SuburbCity (string)
	// fmt.Printf("Serializing SuburbCity\n")
	if err := WriteVarChar(buf, action.SuburbCity, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized SuburbCity : buf len %d\n", buf.Len())

	// TerritoryStateProvinceCode (string)
	// fmt.Printf("Serializing TerritoryStateProvinceCode\n")
	if err := WriteFixedChar(buf, action.TerritoryStateProvinceCode, 5); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized TerritoryStateProvinceCode : buf len %d\n", buf.Len())

	// CountryCode (string)
	// fmt.Printf("Serializing CountryCode\n")
	if err := WriteFixedChar(buf, action.CountryCode, 3); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized CountryCode : buf len %d\n", buf.Len())

	// PostalZIPCode (string)
	// fmt.Printf("Serializing PostalZIPCode\n")
	if err := WriteVarChar(buf, action.PostalZIPCode, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized PostalZIPCode : buf len %d\n", buf.Len())

	// EmailAddress (string)
	// fmt.Printf("Serializing EmailAddress\n")
	if err := WriteVarChar(buf, action.EmailAddress, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized EmailAddress : buf len %d\n", buf.Len())

	// PhoneNumber (string)
	// fmt.Printf("Serializing PhoneNumber\n")
	if err := WriteVarChar(buf, action.PhoneNumber, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized PhoneNumber : buf len %d\n", buf.Len())

	// KeyRoles ([]KeyRole)
	// fmt.Printf("Serializing KeyRoles\n")
	if err := WriteVariableSize(buf, uint64(len(action.KeyRoles)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.KeyRoles {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized KeyRoles : buf len %d\n", buf.Len())

	// NotableRoles ([]NotableRole)
	// fmt.Printf("Serializing NotableRoles\n")
	if err := WriteVariableSize(buf, uint64(len(action.NotableRoles)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.NotableRoles {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized NotableRoles : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in ContractOffer from the byte slice
func (action *ContractOffer) write(b []byte) (int, error) {
	// fmt.Printf("Reading ContractOffer : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// ContractName (string)
	// fmt.Printf("Reading ContractName : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.ContractName, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read ContractName : %d bytes remaining\n%+v\n", buf.Len(), action.ContractName)

	// ContractFileType (uint8)
	// fmt.Printf("Reading ContractFileType : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ContractFileType); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ContractFileType : %d bytes remaining\n%+v\n", buf.Len(), action.ContractFileType)

	// ContractFile ([]byte)
	// fmt.Printf("Reading ContractFile : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.ContractFile, err = ReadVarBin(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read ContractFile : %d bytes remaining\n%+v\n", buf.Len(), action.ContractFile)

	// SupportingDocsFileType (uint8)
	// fmt.Printf("Reading SupportingDocsFileType : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.SupportingDocsFileType); err != nil {
		return 0, err
	}

	// fmt.Printf("Read SupportingDocsFileType : %d bytes remaining\n%+v\n", buf.Len(), action.SupportingDocsFileType)

	// SupportingDocs (string)
	// fmt.Printf("Reading SupportingDocs : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.SupportingDocs, err = ReadVarChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read SupportingDocs : %d bytes remaining\n%+v\n", buf.Len(), action.SupportingDocs)

	// GoverningLaw (string)
	// fmt.Printf("Reading GoverningLaw : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.GoverningLaw, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read GoverningLaw : %d bytes remaining\n%+v\n", buf.Len(), action.GoverningLaw)

	// Jurisdiction (string)
	// fmt.Printf("Reading Jurisdiction : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.Jurisdiction, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read Jurisdiction : %d bytes remaining\n%+v\n", buf.Len(), action.Jurisdiction)

	// ContractExpiration (Timestamp)
	// fmt.Printf("Reading ContractExpiration : %d bytes remaining\n", buf.Len())
	if err := action.ContractExpiration.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ContractExpiration : %d bytes remaining\n%+v\n", buf.Len(), action.ContractExpiration)

	// ContractURI (string)
	// fmt.Printf("Reading ContractURI : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.ContractURI, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read ContractURI : %d bytes remaining\n%+v\n", buf.Len(), action.ContractURI)

	// IssuerName (string)
	// fmt.Printf("Reading IssuerName : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.IssuerName, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read IssuerName : %d bytes remaining\n%+v\n", buf.Len(), action.IssuerName)

	// IssuerType (byte)
	// fmt.Printf("Reading IssuerType : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.IssuerType); err != nil {
		return 0, err
	}

	// fmt.Printf("Read IssuerType : %d bytes remaining\n%+v\n", buf.Len(), action.IssuerType)

	// IssuerLEI (string)
	// fmt.Printf("Reading IssuerLEI : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.IssuerLEI, err = ReadFixedChar(buf, 20)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read IssuerLEI : %d bytes remaining\n%+v\n", buf.Len(), action.IssuerLEI)

	// IssuerLogoURL (string)
	// fmt.Printf("Reading IssuerLogoURL : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.IssuerLogoURL, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read IssuerLogoURL : %d bytes remaining\n%+v\n", buf.Len(), action.IssuerLogoURL)

	// ContractOperatorID (string)
	// fmt.Printf("Reading ContractOperatorID : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.ContractOperatorID, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read ContractOperatorID : %d bytes remaining\n%+v\n", buf.Len(), action.ContractOperatorID)

	// OperatorLEI (string)
	// fmt.Printf("Reading OperatorLEI : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.OperatorLEI, err = ReadFixedChar(buf, 20)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read OperatorLEI : %d bytes remaining\n%+v\n", buf.Len(), action.OperatorLEI)

	// ContractAuthFlags ([16]byte)
	// fmt.Printf("Reading ContractAuthFlags : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ContractAuthFlags); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ContractAuthFlags : %d bytes remaining\n%+v\n", buf.Len(), action.ContractAuthFlags)

	// VotingSystems ([]VotingSystem)
	// fmt.Printf("Reading VotingSystems : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.VotingSystems = make([]VotingSystem, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue VotingSystem
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.VotingSystems = append(action.VotingSystems, newValue)
		}
	}

	// fmt.Printf("Read VotingSystems : %d bytes remaining\n%+v\n", buf.Len(), action.VotingSystems)

	// RestrictedQtyAssets (uint64)
	// fmt.Printf("Reading RestrictedQtyAssets : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.RestrictedQtyAssets); err != nil {
		return 0, err
	}

	// fmt.Printf("Read RestrictedQtyAssets : %d bytes remaining\n%+v\n", buf.Len(), action.RestrictedQtyAssets)

	// ReferendumProposal (bool)
	// fmt.Printf("Reading ReferendumProposal : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ReferendumProposal); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ReferendumProposal : %d bytes remaining\n%+v\n", buf.Len(), action.ReferendumProposal)

	// InitiativeProposal (bool)
	// fmt.Printf("Reading InitiativeProposal : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.InitiativeProposal); err != nil {
		return 0, err
	}

	// fmt.Printf("Read InitiativeProposal : %d bytes remaining\n%+v\n", buf.Len(), action.InitiativeProposal)

	// Registries ([]Registry)
	// fmt.Printf("Reading Registries : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.Registries = make([]Registry, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Registry
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.Registries = append(action.Registries, newValue)
		}
	}

	// fmt.Printf("Read Registries : %d bytes remaining\n%+v\n", buf.Len(), action.Registries)

	// IssuerAddress (bool)
	// fmt.Printf("Reading IssuerAddress : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.IssuerAddress); err != nil {
		return 0, err
	}

	// fmt.Printf("Read IssuerAddress : %d bytes remaining\n%+v\n", buf.Len(), action.IssuerAddress)

	// UnitNumber (string)
	// fmt.Printf("Reading UnitNumber : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.UnitNumber, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read UnitNumber : %d bytes remaining\n%+v\n", buf.Len(), action.UnitNumber)

	// BuildingNumber (string)
	// fmt.Printf("Reading BuildingNumber : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.BuildingNumber, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read BuildingNumber : %d bytes remaining\n%+v\n", buf.Len(), action.BuildingNumber)

	// Street (string)
	// fmt.Printf("Reading Street : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.Street, err = ReadVarChar(buf, 16)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read Street : %d bytes remaining\n%+v\n", buf.Len(), action.Street)

	// SuburbCity (string)
	// fmt.Printf("Reading SuburbCity : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.SuburbCity, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read SuburbCity : %d bytes remaining\n%+v\n", buf.Len(), action.SuburbCity)

	// TerritoryStateProvinceCode (string)
	// fmt.Printf("Reading TerritoryStateProvinceCode : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.TerritoryStateProvinceCode, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read TerritoryStateProvinceCode : %d bytes remaining\n%+v\n", buf.Len(), action.TerritoryStateProvinceCode)

	// CountryCode (string)
	// fmt.Printf("Reading CountryCode : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.CountryCode, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read CountryCode : %d bytes remaining\n%+v\n", buf.Len(), action.CountryCode)

	// PostalZIPCode (string)
	// fmt.Printf("Reading PostalZIPCode : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.PostalZIPCode, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read PostalZIPCode : %d bytes remaining\n%+v\n", buf.Len(), action.PostalZIPCode)

	// EmailAddress (string)
	// fmt.Printf("Reading EmailAddress : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.EmailAddress, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read EmailAddress : %d bytes remaining\n%+v\n", buf.Len(), action.EmailAddress)

	// PhoneNumber (string)
	// fmt.Printf("Reading PhoneNumber : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.PhoneNumber, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read PhoneNumber : %d bytes remaining\n%+v\n", buf.Len(), action.PhoneNumber)

	// KeyRoles ([]KeyRole)
	// fmt.Printf("Reading KeyRoles : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.KeyRoles = make([]KeyRole, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue KeyRole
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.KeyRoles = append(action.KeyRoles, newValue)
		}
	}

	// fmt.Printf("Read KeyRoles : %d bytes remaining\n%+v\n", buf.Len(), action.KeyRoles)

	// NotableRoles ([]NotableRole)
	// fmt.Printf("Reading NotableRoles : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.NotableRoles = make([]NotableRole, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue NotableRole
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.NotableRoles = append(action.NotableRoles, newValue)
		}
	}

	// fmt.Printf("Read NotableRoles : %d bytes remaining\n%+v\n", buf.Len(), action.NotableRoles)

	// fmt.Printf("Read ContractOffer : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action ContractOffer) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action ContractOffer) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("ContractName:%#+v", action.ContractName))
	vals = append(vals, fmt.Sprintf("ContractFileType:%v", action.ContractFileType))
	vals = append(vals, fmt.Sprintf("ContractFile:%#x", action.ContractFile))
	vals = append(vals, fmt.Sprintf("SupportingDocsFileType:%v", action.SupportingDocsFileType))
	vals = append(vals, fmt.Sprintf("SupportingDocs:%#+v", action.SupportingDocs))
	vals = append(vals, fmt.Sprintf("GoverningLaw:%#+v", action.GoverningLaw))
	vals = append(vals, fmt.Sprintf("Jurisdiction:%#+v", action.Jurisdiction))
	vals = append(vals, fmt.Sprintf("ContractExpiration:%#+v", action.ContractExpiration))
	vals = append(vals, fmt.Sprintf("ContractURI:%#+v", action.ContractURI))
	vals = append(vals, fmt.Sprintf("IssuerName:%#+v", action.IssuerName))
	vals = append(vals, fmt.Sprintf("IssuerType:%#+v", action.IssuerType))
	vals = append(vals, fmt.Sprintf("IssuerLEI:%#+v", action.IssuerLEI))
	vals = append(vals, fmt.Sprintf("IssuerLogoURL:%#+v", action.IssuerLogoURL))
	vals = append(vals, fmt.Sprintf("ContractOperatorID:%#+v", action.ContractOperatorID))
	vals = append(vals, fmt.Sprintf("OperatorLEI:%#+v", action.OperatorLEI))
	vals = append(vals, fmt.Sprintf("ContractAuthFlags:%#+v", action.ContractAuthFlags))
	vals = append(vals, fmt.Sprintf("VotingSystems:%#+v", action.VotingSystems))
	vals = append(vals, fmt.Sprintf("RestrictedQtyAssets:%v", action.RestrictedQtyAssets))
	vals = append(vals, fmt.Sprintf("ReferendumProposal:%#+v", action.ReferendumProposal))
	vals = append(vals, fmt.Sprintf("InitiativeProposal:%#+v", action.InitiativeProposal))
	vals = append(vals, fmt.Sprintf("Registries:%#+v", action.Registries))
	vals = append(vals, fmt.Sprintf("IssuerAddress:%#+v", action.IssuerAddress))
	vals = append(vals, fmt.Sprintf("UnitNumber:%#+v", action.UnitNumber))
	vals = append(vals, fmt.Sprintf("BuildingNumber:%#+v", action.BuildingNumber))
	vals = append(vals, fmt.Sprintf("Street:%#+v", action.Street))
	vals = append(vals, fmt.Sprintf("SuburbCity:%#+v", action.SuburbCity))
	vals = append(vals, fmt.Sprintf("TerritoryStateProvinceCode:%#+v", action.TerritoryStateProvinceCode))
	vals = append(vals, fmt.Sprintf("CountryCode:%#+v", action.CountryCode))
	vals = append(vals, fmt.Sprintf("PostalZIPCode:%#+v", action.PostalZIPCode))
	vals = append(vals, fmt.Sprintf("EmailAddress:%#+v", action.EmailAddress))
	vals = append(vals, fmt.Sprintf("PhoneNumber:%#+v", action.PhoneNumber))
	vals = append(vals, fmt.Sprintf("KeyRoles:%#+v", action.KeyRoles))
	vals = append(vals, fmt.Sprintf("NotableRoles:%#+v", action.NotableRoles))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// ContractFormation This txn is created by the Contract (smart
// contract/off-chain agent/token contract) upon receipt of a valid
// Contract Offer Action from the issuer. The Smart Contract will execute
// on a server controlled by the Issuer. or a Smart Contract Operator on
// their behalf.
type ContractFormation struct {
	Header                     Header         `json:"header,omitempty"`                        // Common header data for all actions
	ContractName               string         `json:"contract_name,omitempty"`                 // Can be any unique identifying string, including human readable names for branding/vanity purposes.   [Contract identifier (instance) is the bitcoin public key hash address. If the Public Address is lost, then the issuer will have to reissue the entire contract, Asset definition and tokens with the new public address.]. Smart contracts can be branded and specialized to suit any terms and conditions.
	ContractFileType           uint8          `json:"contract_file_type,omitempty"`            // 1 - SHA-256 Hash, 2 - Markdown file
	ContractFile               []byte         `json:"contract_file,omitempty"`                 // SHA-256 hash of the contract file or markdown data for contract file specific to the smart contract and relevant Assets.  Legal and technical information. (eg. pdf)
	SupportingDocsFileType     uint8          `json:"supporting_docs_file_type,omitempty"`     // 1 - 7z
	SupportingDocs             string         `json:"supporting_docs,omitempty"`               //
	GoverningLaw               string         `json:"governing_law,omitempty"`                 // 5 Letter Code to Identify which governing law the contract will adhere to.  Disputes are to be settled by this law in the jurisdiction specified below. Private dispute resolution organizations can be used as well.  A custom code just needs to be defined.
	Jurisdiction               string         `json:"jurisdiction,omitempty"`                  // Legal proceedings/arbitration will take place using the specified Governing Law in this location.
	ContractExpiration         Timestamp      `json:"contract_expiration,omitempty"`           // All actions related to the contract will cease to work after this timestamp. The smart contract will stop running.  This will allow many token use cases to be able to calculate smart contract running costs. Eg. an issuer is creating tickets for an event on the 5th of June 2018.  The smart contract will facilitate exchange and send transactions up until the 6th of June.  Wallets can use this to forget tokens that are no longer valid - or at least store them in an 'Expired' folder.
	ContractURI                string         `json:"contract_uri,omitempty"`                  // Length 0-255 bytes.  0 is valid. Points to an information page that also has a copy of the Contract.  Anyone can go to the website to have a look at the price/token, information about the Issuer (company), information about the Asset, legal information, etc.  There will also be a way for Token Owners to vote on this page and contact details with the Issuer/tokenized companies. Could be a IPv6/IPv4, an IPFS address (hash) or txn-id for on chain information or even a public address (DNS).
	IssuerName                 string         `json:"issuer_name,omitempty"`                   // Length 0-255 bytes. 0 is not valid. Issuing entity (company, organization, individual).  Can be any unique identifying string, including human readable names for branding/vanity purposes.
	IssuerType                 byte           `json:"issuer_type,omitempty"`                   // P - Public Company Limited by Shares, C - Private Company Limited by Shares, I - Individual, L - Limited Partnership, U -Unlimited Partnership, T - Sole Proprietorship, S - Statutory Company, O - Non-Profit Organization, N - Nation State, G - Government Agency, U - Unit Trust, D - Discretionary Trust.  Found in 'Entities' (Specification/Resources).
	IssuerLEI                  string         `json:"issuer_lei,omitempty"`                    // Null is valid. A Legal Entity Identifier (or LEI) is an international identifier made up of a 20-character identifier that identifies distinct legal entities that engage in financial transactions. It is defined by ISO 17442.[1] Natural persons are not required to have an LEI; they’re eligible to have one issued, however, but only if they act in an independent business capacity.[2] The LEI is a global standard, designed to be non-proprietary data that is freely accessible to all.[3] As of December 2018, over 1,300,000 legal entities from more than 200 countries have now been issued with LEIs.
	IssuerLogoURL              string         `json:"issuer_logo_url,omitempty"`               // The URL of the Issuers logo.
	ContractOperatorID         string         `json:"contract_operator_id,omitempty"`          // Length 0-255 bytes. 0 is valid. Smart Contract Operator identifier. Can be any unique identifying string, including human readable names for branding/vanity purposes. Can also be null or the Issuer.
	OperatorLEI                string         `json:"operator_lei,omitempty"`                  // Null is valid. A Legal Entity Identifier (or LEI) is an international identifier made up of a 20-character identifier that identifies distinct legal entities that engage in financial transactions. It is defined by ISO 17442.[1] Natural persons are not required to have an LEI; they’re eligible to have one issued, however, but only if they act in an independent business capacity.[2] The LEI is a global standard, designed to be non-proprietary data that is freely accessible to all.[3] As of December 2018, over 1,300,000 legal entities from more than 200 countries have now been issued with LEIs.
	ContractAuthFlags          [16]byte       `json:"contract_auth_flags,omitempty"`           // Authorization Flags aka Terms and Conditions that the smart contract can enforce.  Other terms and conditions that are out of the smart contract's control are listed in the actual Contract File.
	VotingSystems              []VotingSystem `json:"voting_systems,omitempty"`                // A list voting systems.
	RestrictedQtyAssets        uint64         `json:"restricted_qty_assets,omitempty"`         // Number of Assets (non-fungible) permitted on this contract. 0 if unlimited which will display an infinity symbol in UI
	ReferendumProposal         bool           `json:"referendum_proposal,omitempty"`           // A Referendum is permitted for Contract-Wide Proposals (outside of smart contract scope).
	InitiativeProposal         bool           `json:"initiative_proposal,omitempty"`           // An initiative is permitted for Contract-Wide Proposals (outside of smart contract scope).
	Registries                 []Registry     `json:"registries,omitempty"`                    // A list Registries
	IssuerAddress              bool           `json:"issuer_address,omitempty"`                // Physical/mailing address. Y/N, N means there is no issuer address.
	UnitNumber                 string         `json:"unit_number,omitempty"`                   // Issuer Address Details (eg. HQ)
	BuildingNumber             string         `json:"building_number,omitempty"`               //
	Street                     string         `json:"street,omitempty"`                        //
	SuburbCity                 string         `json:"suburb_city,omitempty"`                   //
	TerritoryStateProvinceCode string         `json:"territory_state_province_code,omitempty"` //
	CountryCode                string         `json:"country_code,omitempty"`                  //
	PostalZIPCode              string         `json:"postal_zip_code,omitempty"`               //
	EmailAddress               string         `json:"email_address,omitempty"`                 // Address for text-based communication: eg. email address, Bitcoin address
	PhoneNumber                string         `json:"phone_number,omitempty"`                  // Phone Number for Entity. Max acceptable size: 50.
	KeyRoles                   []KeyRole      `json:"key_roles,omitempty"`                     // A list of Key Roles.
	NotableRoles               []NotableRole  `json:"notable_roles,omitempty"`                 // A list of Notable Roles.
	ContractRevision           uint32         `json:"contract_revision,omitempty"`             // Counter. Cannot be manually changed by issuer.  Can only be incremented by 1 by SC when CF action is published.
	Timestamp                  Timestamp      `json:"timestamp,omitempty"`                     // Timestamp in nanoseconds of when the smart contract created the action.
}

// Type returns the type identifer for this message.
func (action ContractFormation) Type() string {
	return CodeContractFormation
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *ContractFormation) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *ContractFormation) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// ContractName (string)
	// fmt.Printf("Serializing ContractName\n")
	if err := WriteVarChar(buf, action.ContractName, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractName : buf len %d\n", buf.Len())

	// ContractFileType (uint8)
	// fmt.Printf("Serializing ContractFileType\n")
	if err := write(buf, action.ContractFileType); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractFileType : buf len %d\n", buf.Len())

	// ContractFile ([]byte)
	// fmt.Printf("Serializing ContractFile\n")
	if err := WriteVarBin(buf, action.ContractFile, 32); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractFile : buf len %d\n", buf.Len())

	// SupportingDocsFileType (uint8)
	// fmt.Printf("Serializing SupportingDocsFileType\n")
	if err := write(buf, action.SupportingDocsFileType); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized SupportingDocsFileType : buf len %d\n", buf.Len())

	// SupportingDocs (string)
	// fmt.Printf("Serializing SupportingDocs\n")
	if err := WriteVarChar(buf, action.SupportingDocs, 32); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized SupportingDocs : buf len %d\n", buf.Len())

	// GoverningLaw (string)
	// fmt.Printf("Serializing GoverningLaw\n")
	if err := WriteFixedChar(buf, action.GoverningLaw, 5); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized GoverningLaw : buf len %d\n", buf.Len())

	// Jurisdiction (string)
	// fmt.Printf("Serializing Jurisdiction\n")
	if err := WriteFixedChar(buf, action.Jurisdiction, 5); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized Jurisdiction : buf len %d\n", buf.Len())

	// ContractExpiration (Timestamp)
	// fmt.Printf("Serializing ContractExpiration\n")
	{
		b, err := action.ContractExpiration.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized ContractExpiration : buf len %d\n", buf.Len())

	// ContractURI (string)
	// fmt.Printf("Serializing ContractURI\n")
	if err := WriteVarChar(buf, action.ContractURI, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractURI : buf len %d\n", buf.Len())

	// IssuerName (string)
	// fmt.Printf("Serializing IssuerName\n")
	if err := WriteVarChar(buf, action.IssuerName, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized IssuerName : buf len %d\n", buf.Len())

	// IssuerType (byte)
	// fmt.Printf("Serializing IssuerType\n")
	if err := write(buf, action.IssuerType); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized IssuerType : buf len %d\n", buf.Len())

	// IssuerLEI (string)
	// fmt.Printf("Serializing IssuerLEI\n")
	if err := WriteFixedChar(buf, action.IssuerLEI, 20); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized IssuerLEI : buf len %d\n", buf.Len())

	// IssuerLogoURL (string)
	// fmt.Printf("Serializing IssuerLogoURL\n")
	if err := WriteVarChar(buf, action.IssuerLogoURL, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized IssuerLogoURL : buf len %d\n", buf.Len())

	// ContractOperatorID (string)
	// fmt.Printf("Serializing ContractOperatorID\n")
	if err := WriteVarChar(buf, action.ContractOperatorID, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractOperatorID : buf len %d\n", buf.Len())

	// OperatorLEI (string)
	// fmt.Printf("Serializing OperatorLEI\n")
	if err := WriteFixedChar(buf, action.OperatorLEI, 20); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized OperatorLEI : buf len %d\n", buf.Len())

	// ContractAuthFlags ([16]byte)
	// fmt.Printf("Serializing ContractAuthFlags\n")
	if err := write(buf, action.ContractAuthFlags); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractAuthFlags : buf len %d\n", buf.Len())

	// VotingSystems ([]VotingSystem)
	// fmt.Printf("Serializing VotingSystems\n")
	if err := WriteVariableSize(buf, uint64(len(action.VotingSystems)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.VotingSystems {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized VotingSystems : buf len %d\n", buf.Len())

	// RestrictedQtyAssets (uint64)
	// fmt.Printf("Serializing RestrictedQtyAssets\n")
	if err := write(buf, action.RestrictedQtyAssets); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized RestrictedQtyAssets : buf len %d\n", buf.Len())

	// ReferendumProposal (bool)
	// fmt.Printf("Serializing ReferendumProposal\n")
	if err := write(buf, action.ReferendumProposal); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ReferendumProposal : buf len %d\n", buf.Len())

	// InitiativeProposal (bool)
	// fmt.Printf("Serializing InitiativeProposal\n")
	if err := write(buf, action.InitiativeProposal); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized InitiativeProposal : buf len %d\n", buf.Len())

	// Registries ([]Registry)
	// fmt.Printf("Serializing Registries\n")
	if err := WriteVariableSize(buf, uint64(len(action.Registries)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.Registries {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Registries : buf len %d\n", buf.Len())

	// IssuerAddress (bool)
	// fmt.Printf("Serializing IssuerAddress\n")
	if err := write(buf, action.IssuerAddress); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized IssuerAddress : buf len %d\n", buf.Len())

	// UnitNumber (string)
	// fmt.Printf("Serializing UnitNumber\n")
	if err := WriteVarChar(buf, action.UnitNumber, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized UnitNumber : buf len %d\n", buf.Len())

	// BuildingNumber (string)
	// fmt.Printf("Serializing BuildingNumber\n")
	if err := WriteVarChar(buf, action.BuildingNumber, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized BuildingNumber : buf len %d\n", buf.Len())

	// Street (string)
	// fmt.Printf("Serializing Street\n")
	if err := WriteVarChar(buf, action.Street, 16); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized Street : buf len %d\n", buf.Len())

	// SuburbCity (string)
	// fmt.Printf("Serializing SuburbCity\n")
	if err := WriteVarChar(buf, action.SuburbCity, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized SuburbCity : buf len %d\n", buf.Len())

	// TerritoryStateProvinceCode (string)
	// fmt.Printf("Serializing TerritoryStateProvinceCode\n")
	if err := WriteFixedChar(buf, action.TerritoryStateProvinceCode, 5); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized TerritoryStateProvinceCode : buf len %d\n", buf.Len())

	// CountryCode (string)
	// fmt.Printf("Serializing CountryCode\n")
	if err := WriteFixedChar(buf, action.CountryCode, 3); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized CountryCode : buf len %d\n", buf.Len())

	// PostalZIPCode (string)
	// fmt.Printf("Serializing PostalZIPCode\n")
	if err := WriteVarChar(buf, action.PostalZIPCode, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized PostalZIPCode : buf len %d\n", buf.Len())

	// EmailAddress (string)
	// fmt.Printf("Serializing EmailAddress\n")
	if err := WriteVarChar(buf, action.EmailAddress, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized EmailAddress : buf len %d\n", buf.Len())

	// PhoneNumber (string)
	// fmt.Printf("Serializing PhoneNumber\n")
	if err := WriteVarChar(buf, action.PhoneNumber, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized PhoneNumber : buf len %d\n", buf.Len())

	// KeyRoles ([]KeyRole)
	// fmt.Printf("Serializing KeyRoles\n")
	if err := WriteVariableSize(buf, uint64(len(action.KeyRoles)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.KeyRoles {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized KeyRoles : buf len %d\n", buf.Len())

	// NotableRoles ([]NotableRole)
	// fmt.Printf("Serializing NotableRoles\n")
	if err := WriteVariableSize(buf, uint64(len(action.NotableRoles)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.NotableRoles {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized NotableRoles : buf len %d\n", buf.Len())

	// ContractRevision (uint32)
	// fmt.Printf("Serializing ContractRevision\n")
	if err := write(buf, action.ContractRevision); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractRevision : buf len %d\n", buf.Len())

	// Timestamp (Timestamp)
	// fmt.Printf("Serializing Timestamp\n")
	{
		b, err := action.Timestamp.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Timestamp : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in ContractFormation from the byte slice
func (action *ContractFormation) write(b []byte) (int, error) {
	// fmt.Printf("Reading ContractFormation : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// ContractName (string)
	// fmt.Printf("Reading ContractName : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.ContractName, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read ContractName : %d bytes remaining\n%+v\n", buf.Len(), action.ContractName)

	// ContractFileType (uint8)
	// fmt.Printf("Reading ContractFileType : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ContractFileType); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ContractFileType : %d bytes remaining\n%+v\n", buf.Len(), action.ContractFileType)

	// ContractFile ([]byte)
	// fmt.Printf("Reading ContractFile : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.ContractFile, err = ReadVarBin(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read ContractFile : %d bytes remaining\n%+v\n", buf.Len(), action.ContractFile)

	// SupportingDocsFileType (uint8)
	// fmt.Printf("Reading SupportingDocsFileType : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.SupportingDocsFileType); err != nil {
		return 0, err
	}

	// fmt.Printf("Read SupportingDocsFileType : %d bytes remaining\n%+v\n", buf.Len(), action.SupportingDocsFileType)

	// SupportingDocs (string)
	// fmt.Printf("Reading SupportingDocs : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.SupportingDocs, err = ReadVarChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read SupportingDocs : %d bytes remaining\n%+v\n", buf.Len(), action.SupportingDocs)

	// GoverningLaw (string)
	// fmt.Printf("Reading GoverningLaw : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.GoverningLaw, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read GoverningLaw : %d bytes remaining\n%+v\n", buf.Len(), action.GoverningLaw)

	// Jurisdiction (string)
	// fmt.Printf("Reading Jurisdiction : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.Jurisdiction, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read Jurisdiction : %d bytes remaining\n%+v\n", buf.Len(), action.Jurisdiction)

	// ContractExpiration (Timestamp)
	// fmt.Printf("Reading ContractExpiration : %d bytes remaining\n", buf.Len())
	if err := action.ContractExpiration.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ContractExpiration : %d bytes remaining\n%+v\n", buf.Len(), action.ContractExpiration)

	// ContractURI (string)
	// fmt.Printf("Reading ContractURI : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.ContractURI, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read ContractURI : %d bytes remaining\n%+v\n", buf.Len(), action.ContractURI)

	// IssuerName (string)
	// fmt.Printf("Reading IssuerName : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.IssuerName, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read IssuerName : %d bytes remaining\n%+v\n", buf.Len(), action.IssuerName)

	// IssuerType (byte)
	// fmt.Printf("Reading IssuerType : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.IssuerType); err != nil {
		return 0, err
	}

	// fmt.Printf("Read IssuerType : %d bytes remaining\n%+v\n", buf.Len(), action.IssuerType)

	// IssuerLEI (string)
	// fmt.Printf("Reading IssuerLEI : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.IssuerLEI, err = ReadFixedChar(buf, 20)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read IssuerLEI : %d bytes remaining\n%+v\n", buf.Len(), action.IssuerLEI)

	// IssuerLogoURL (string)
	// fmt.Printf("Reading IssuerLogoURL : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.IssuerLogoURL, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read IssuerLogoURL : %d bytes remaining\n%+v\n", buf.Len(), action.IssuerLogoURL)

	// ContractOperatorID (string)
	// fmt.Printf("Reading ContractOperatorID : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.ContractOperatorID, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read ContractOperatorID : %d bytes remaining\n%+v\n", buf.Len(), action.ContractOperatorID)

	// OperatorLEI (string)
	// fmt.Printf("Reading OperatorLEI : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.OperatorLEI, err = ReadFixedChar(buf, 20)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read OperatorLEI : %d bytes remaining\n%+v\n", buf.Len(), action.OperatorLEI)

	// ContractAuthFlags ([16]byte)
	// fmt.Printf("Reading ContractAuthFlags : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ContractAuthFlags); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ContractAuthFlags : %d bytes remaining\n%+v\n", buf.Len(), action.ContractAuthFlags)

	// VotingSystems ([]VotingSystem)
	// fmt.Printf("Reading VotingSystems : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.VotingSystems = make([]VotingSystem, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue VotingSystem
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.VotingSystems = append(action.VotingSystems, newValue)
		}
	}

	// fmt.Printf("Read VotingSystems : %d bytes remaining\n%+v\n", buf.Len(), action.VotingSystems)

	// RestrictedQtyAssets (uint64)
	// fmt.Printf("Reading RestrictedQtyAssets : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.RestrictedQtyAssets); err != nil {
		return 0, err
	}

	// fmt.Printf("Read RestrictedQtyAssets : %d bytes remaining\n%+v\n", buf.Len(), action.RestrictedQtyAssets)

	// ReferendumProposal (bool)
	// fmt.Printf("Reading ReferendumProposal : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ReferendumProposal); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ReferendumProposal : %d bytes remaining\n%+v\n", buf.Len(), action.ReferendumProposal)

	// InitiativeProposal (bool)
	// fmt.Printf("Reading InitiativeProposal : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.InitiativeProposal); err != nil {
		return 0, err
	}

	// fmt.Printf("Read InitiativeProposal : %d bytes remaining\n%+v\n", buf.Len(), action.InitiativeProposal)

	// Registries ([]Registry)
	// fmt.Printf("Reading Registries : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.Registries = make([]Registry, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Registry
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.Registries = append(action.Registries, newValue)
		}
	}

	// fmt.Printf("Read Registries : %d bytes remaining\n%+v\n", buf.Len(), action.Registries)

	// IssuerAddress (bool)
	// fmt.Printf("Reading IssuerAddress : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.IssuerAddress); err != nil {
		return 0, err
	}

	// fmt.Printf("Read IssuerAddress : %d bytes remaining\n%+v\n", buf.Len(), action.IssuerAddress)

	// UnitNumber (string)
	// fmt.Printf("Reading UnitNumber : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.UnitNumber, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read UnitNumber : %d bytes remaining\n%+v\n", buf.Len(), action.UnitNumber)

	// BuildingNumber (string)
	// fmt.Printf("Reading BuildingNumber : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.BuildingNumber, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read BuildingNumber : %d bytes remaining\n%+v\n", buf.Len(), action.BuildingNumber)

	// Street (string)
	// fmt.Printf("Reading Street : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.Street, err = ReadVarChar(buf, 16)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read Street : %d bytes remaining\n%+v\n", buf.Len(), action.Street)

	// SuburbCity (string)
	// fmt.Printf("Reading SuburbCity : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.SuburbCity, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read SuburbCity : %d bytes remaining\n%+v\n", buf.Len(), action.SuburbCity)

	// TerritoryStateProvinceCode (string)
	// fmt.Printf("Reading TerritoryStateProvinceCode : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.TerritoryStateProvinceCode, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read TerritoryStateProvinceCode : %d bytes remaining\n%+v\n", buf.Len(), action.TerritoryStateProvinceCode)

	// CountryCode (string)
	// fmt.Printf("Reading CountryCode : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.CountryCode, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read CountryCode : %d bytes remaining\n%+v\n", buf.Len(), action.CountryCode)

	// PostalZIPCode (string)
	// fmt.Printf("Reading PostalZIPCode : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.PostalZIPCode, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read PostalZIPCode : %d bytes remaining\n%+v\n", buf.Len(), action.PostalZIPCode)

	// EmailAddress (string)
	// fmt.Printf("Reading EmailAddress : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.EmailAddress, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read EmailAddress : %d bytes remaining\n%+v\n", buf.Len(), action.EmailAddress)

	// PhoneNumber (string)
	// fmt.Printf("Reading PhoneNumber : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.PhoneNumber, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read PhoneNumber : %d bytes remaining\n%+v\n", buf.Len(), action.PhoneNumber)

	// KeyRoles ([]KeyRole)
	// fmt.Printf("Reading KeyRoles : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.KeyRoles = make([]KeyRole, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue KeyRole
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.KeyRoles = append(action.KeyRoles, newValue)
		}
	}

	// fmt.Printf("Read KeyRoles : %d bytes remaining\n%+v\n", buf.Len(), action.KeyRoles)

	// NotableRoles ([]NotableRole)
	// fmt.Printf("Reading NotableRoles : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.NotableRoles = make([]NotableRole, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue NotableRole
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.NotableRoles = append(action.NotableRoles, newValue)
		}
	}

	// fmt.Printf("Read NotableRoles : %d bytes remaining\n%+v\n", buf.Len(), action.NotableRoles)

	// ContractRevision (uint32)
	// fmt.Printf("Reading ContractRevision : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ContractRevision); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ContractRevision : %d bytes remaining\n%+v\n", buf.Len(), action.ContractRevision)

	// Timestamp (Timestamp)
	// fmt.Printf("Reading Timestamp : %d bytes remaining\n", buf.Len())
	if err := action.Timestamp.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Timestamp : %d bytes remaining\n%+v\n", buf.Len(), action.Timestamp)

	// fmt.Printf("Read ContractFormation : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action ContractFormation) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action ContractFormation) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("ContractName:%#+v", action.ContractName))
	vals = append(vals, fmt.Sprintf("ContractFileType:%v", action.ContractFileType))
	vals = append(vals, fmt.Sprintf("ContractFile:%#x", action.ContractFile))
	vals = append(vals, fmt.Sprintf("SupportingDocsFileType:%v", action.SupportingDocsFileType))
	vals = append(vals, fmt.Sprintf("SupportingDocs:%#+v", action.SupportingDocs))
	vals = append(vals, fmt.Sprintf("GoverningLaw:%#+v", action.GoverningLaw))
	vals = append(vals, fmt.Sprintf("Jurisdiction:%#+v", action.Jurisdiction))
	vals = append(vals, fmt.Sprintf("ContractExpiration:%#+v", action.ContractExpiration))
	vals = append(vals, fmt.Sprintf("ContractURI:%#+v", action.ContractURI))
	vals = append(vals, fmt.Sprintf("IssuerName:%#+v", action.IssuerName))
	vals = append(vals, fmt.Sprintf("IssuerType:%#+v", action.IssuerType))
	vals = append(vals, fmt.Sprintf("IssuerLEI:%#+v", action.IssuerLEI))
	vals = append(vals, fmt.Sprintf("IssuerLogoURL:%#+v", action.IssuerLogoURL))
	vals = append(vals, fmt.Sprintf("ContractOperatorID:%#+v", action.ContractOperatorID))
	vals = append(vals, fmt.Sprintf("OperatorLEI:%#+v", action.OperatorLEI))
	vals = append(vals, fmt.Sprintf("ContractAuthFlags:%#+v", action.ContractAuthFlags))
	vals = append(vals, fmt.Sprintf("VotingSystems:%#+v", action.VotingSystems))
	vals = append(vals, fmt.Sprintf("RestrictedQtyAssets:%v", action.RestrictedQtyAssets))
	vals = append(vals, fmt.Sprintf("ReferendumProposal:%#+v", action.ReferendumProposal))
	vals = append(vals, fmt.Sprintf("InitiativeProposal:%#+v", action.InitiativeProposal))
	vals = append(vals, fmt.Sprintf("Registries:%#+v", action.Registries))
	vals = append(vals, fmt.Sprintf("IssuerAddress:%#+v", action.IssuerAddress))
	vals = append(vals, fmt.Sprintf("UnitNumber:%#+v", action.UnitNumber))
	vals = append(vals, fmt.Sprintf("BuildingNumber:%#+v", action.BuildingNumber))
	vals = append(vals, fmt.Sprintf("Street:%#+v", action.Street))
	vals = append(vals, fmt.Sprintf("SuburbCity:%#+v", action.SuburbCity))
	vals = append(vals, fmt.Sprintf("TerritoryStateProvinceCode:%#+v", action.TerritoryStateProvinceCode))
	vals = append(vals, fmt.Sprintf("CountryCode:%#+v", action.CountryCode))
	vals = append(vals, fmt.Sprintf("PostalZIPCode:%#+v", action.PostalZIPCode))
	vals = append(vals, fmt.Sprintf("EmailAddress:%#+v", action.EmailAddress))
	vals = append(vals, fmt.Sprintf("PhoneNumber:%#+v", action.PhoneNumber))
	vals = append(vals, fmt.Sprintf("KeyRoles:%#+v", action.KeyRoles))
	vals = append(vals, fmt.Sprintf("NotableRoles:%#+v", action.NotableRoles))
	vals = append(vals, fmt.Sprintf("ContractRevision:%v", action.ContractRevision))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", action.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// ContractAmendment Contract Amendment Action - the issuer can initiate an
// amendment to the contract establishment metadata. The ability to make an
// amendment to the contract is restricted by the Authorization Flag set on
// the current revision of Contract Formation action.
type ContractAmendment struct {
	Header                Header      `json:"header,omitempty"`                  // Common header data for all actions
	ChangeIssuerAddress   bool        `json:"change_issuer_address,omitempty"`   // 1 - Yes, 0 - No.  Used to change the issuer address.  The new issuer address must be in the input[1] position.
	ChangeOperatorAddress bool        `json:"change_operator_address,omitempty"` // 1 - Yes, 0 - No.  Used to change the smart contract operator address.  The new operator address must be in the input[1] position.
	ContractRevision      uint32      `json:"contract_revision,omitempty"`       // Counter 0 to (2^32)-1
	Amendments            []Amendment `json:"amendments,omitempty"`              //
	RefTxID               TxId        `json:"ref_tx_id,omitempty"`               // Tx-ID of the associated Result action (governance) that permitted the modifications.
}

// Type returns the type identifer for this message.
func (action ContractAmendment) Type() string {
	return CodeContractAmendment
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *ContractAmendment) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *ContractAmendment) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// ChangeIssuerAddress (bool)
	// fmt.Printf("Serializing ChangeIssuerAddress\n")
	if err := write(buf, action.ChangeIssuerAddress); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ChangeIssuerAddress : buf len %d\n", buf.Len())

	// ChangeOperatorAddress (bool)
	// fmt.Printf("Serializing ChangeOperatorAddress\n")
	if err := write(buf, action.ChangeOperatorAddress); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ChangeOperatorAddress : buf len %d\n", buf.Len())

	// ContractRevision (uint32)
	// fmt.Printf("Serializing ContractRevision\n")
	if err := write(buf, action.ContractRevision); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractRevision : buf len %d\n", buf.Len())

	// Amendments ([]Amendment)
	// fmt.Printf("Serializing Amendments\n")
	if err := WriteVariableSize(buf, uint64(len(action.Amendments)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.Amendments {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Amendments : buf len %d\n", buf.Len())

	// RefTxID (TxId)
	// fmt.Printf("Serializing RefTxID\n")
	{
		b, err := action.RefTxID.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized RefTxID : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in ContractAmendment from the byte slice
func (action *ContractAmendment) write(b []byte) (int, error) {
	// fmt.Printf("Reading ContractAmendment : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// ChangeIssuerAddress (bool)
	// fmt.Printf("Reading ChangeIssuerAddress : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ChangeIssuerAddress); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ChangeIssuerAddress : %d bytes remaining\n%+v\n", buf.Len(), action.ChangeIssuerAddress)

	// ChangeOperatorAddress (bool)
	// fmt.Printf("Reading ChangeOperatorAddress : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ChangeOperatorAddress); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ChangeOperatorAddress : %d bytes remaining\n%+v\n", buf.Len(), action.ChangeOperatorAddress)

	// ContractRevision (uint32)
	// fmt.Printf("Reading ContractRevision : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ContractRevision); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ContractRevision : %d bytes remaining\n%+v\n", buf.Len(), action.ContractRevision)

	// Amendments ([]Amendment)
	// fmt.Printf("Reading Amendments : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.Amendments = make([]Amendment, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Amendment
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.Amendments = append(action.Amendments, newValue)
		}
	}

	// fmt.Printf("Read Amendments : %d bytes remaining\n%+v\n", buf.Len(), action.Amendments)

	// RefTxID (TxId)
	// fmt.Printf("Reading RefTxID : %d bytes remaining\n", buf.Len())
	if err := action.RefTxID.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read RefTxID : %d bytes remaining\n%+v\n", buf.Len(), action.RefTxID)

	// fmt.Printf("Read ContractAmendment : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action ContractAmendment) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action ContractAmendment) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("ChangeIssuerAddress:%#+v", action.ChangeIssuerAddress))
	vals = append(vals, fmt.Sprintf("ChangeOperatorAddress:%#+v", action.ChangeOperatorAddress))
	vals = append(vals, fmt.Sprintf("ContractRevision:%v", action.ContractRevision))
	vals = append(vals, fmt.Sprintf("Amendments:%#+v", action.Amendments))
	vals = append(vals, fmt.Sprintf("RefTxID:%#+v", action.RefTxID))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// StaticContractFormation Static Contract Formation Action
type StaticContractFormation struct {
	Header                 Header       `json:"header,omitempty"`                    // Common header data for all actions
	ContractName           string       `json:"contract_name,omitempty"`             // Length 0-255 bytes. Can be any unique identifying string, including human readable names for branding/vanity purposes.   [Contract identifier (instance) is the bitcoin public address. If the Public Address is lost, then the issuer will have to reissue the entire contract, Asset definition and tokens with the new public address.]. Smart contracts can be branded and specialized to suit any terms and conditions.
	ContractType           string       `json:"contract_type,omitempty"`             //
	ContractCode           ContractCode `json:"contract_code,omitempty"`             // 32 randomly generated bytes.  Each Contract Code should be unique.  The Contract ID will be human facing and will be the Contract Code, with a checksum, encoded in base58 and prefixed by 'CON'. Contract ID = CON + base58(ContractCode + checksum).  Eg. Contract ID = 'CON18RDoKK7Ed5zid2FkKVy7q3rULr4tgfjr4'
	ContractFileType       uint8        `json:"contract_file_type,omitempty"`        // 1 - SHA-256 Hash, 2 - Markdown file
	ContractFile           []byte       `json:"contract_file,omitempty"`             // SHA-256 hash of the contract file or markdown data for contract file specific to the smart contract and relevant Assets.  Legal and technical information. (eg. pdf)
	SupportingDocsFileType uint8        `json:"supporting_docs_file_type,omitempty"` // 1 - 7z
	SupportingDocs         string       `json:"supporting_docs,omitempty"`           //
	ContractRevision       uint32       `json:"contract_revision,omitempty"`         // Counter 0 to (2^32)-1
	GoverningLaw           string       `json:"governing_law,omitempty"`             // 5 Letter Code to Identify which governing law the contract will adhere to.  Disputes are to be settled by this law in the jurisdiction specified below. Private dispute resolution organizations can be used as well.  A custom code just needs to be defined.
	Jurisdiction           string       `json:"jurisdiction,omitempty"`              // Legal proceedings/arbitration will take place using the specified Governing Law in this location.
	EffectiveDate          Timestamp    `json:"effective_date,omitempty"`            // Start date of the contract.
	ContractExpiration     Timestamp    `json:"contract_expiration,omitempty"`       // All actions related to the contract will cease to work after this timestamp. The smart contract will stop running.  This will allow many token use cases to be able to calculate smart contract running costs. Eg. an issuer is creating tickets for an event on the 5th of June 2018.  The smart contract will facilitate exchange and send transactions up until the 6th of June.  Wallets can use this to forget tokens that are no longer valid - or at least store them in an 'Expired' folder.
	ContractURI            string       `json:"contract_uri,omitempty"`              // Length 0-255 bytes. Points to an information page that also has a copy of the Contract.  Anyone can go to the website to have a look at the price/token, information about the Issuer (company), information about the Asset, legal information, etc.  There will also be a way for Token Owners to vote on this page and contact details with the Issuer/tokenized companies. Could be a IPv6/IPv4, an IPFS address (hash) or txn-id for on chain information or even a public address (DNS).
	PrevRevTxID            TxId         `json:"prev_rev_tx_id,omitempty"`            // The Tx-ID of the previous contract revision.
	Entities               []Entity     `json:"entities,omitempty"`                  //
}

// Type returns the type identifer for this message.
func (action StaticContractFormation) Type() string {
	return CodeStaticContractFormation
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *StaticContractFormation) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *StaticContractFormation) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// ContractName (string)
	// fmt.Printf("Serializing ContractName\n")
	if err := WriteVarChar(buf, action.ContractName, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractName : buf len %d\n", buf.Len())

	// ContractType (string)
	// fmt.Printf("Serializing ContractType\n")
	if err := WriteVarChar(buf, action.ContractType, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractType : buf len %d\n", buf.Len())

	// ContractCode (ContractCode)
	// fmt.Printf("Serializing ContractCode\n")
	{
		b, err := action.ContractCode.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized ContractCode : buf len %d\n", buf.Len())

	// ContractFileType (uint8)
	// fmt.Printf("Serializing ContractFileType\n")
	if err := write(buf, action.ContractFileType); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractFileType : buf len %d\n", buf.Len())

	// ContractFile ([]byte)
	// fmt.Printf("Serializing ContractFile\n")
	if err := WriteVarBin(buf, action.ContractFile, 32); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractFile : buf len %d\n", buf.Len())

	// SupportingDocsFileType (uint8)
	// fmt.Printf("Serializing SupportingDocsFileType\n")
	if err := write(buf, action.SupportingDocsFileType); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized SupportingDocsFileType : buf len %d\n", buf.Len())

	// SupportingDocs (string)
	// fmt.Printf("Serializing SupportingDocs\n")
	if err := WriteVarChar(buf, action.SupportingDocs, 32); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized SupportingDocs : buf len %d\n", buf.Len())

	// ContractRevision (uint32)
	// fmt.Printf("Serializing ContractRevision\n")
	if err := write(buf, action.ContractRevision); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractRevision : buf len %d\n", buf.Len())

	// GoverningLaw (string)
	// fmt.Printf("Serializing GoverningLaw\n")
	if err := WriteFixedChar(buf, action.GoverningLaw, 5); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized GoverningLaw : buf len %d\n", buf.Len())

	// Jurisdiction (string)
	// fmt.Printf("Serializing Jurisdiction\n")
	if err := WriteFixedChar(buf, action.Jurisdiction, 5); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized Jurisdiction : buf len %d\n", buf.Len())

	// EffectiveDate (Timestamp)
	// fmt.Printf("Serializing EffectiveDate\n")
	{
		b, err := action.EffectiveDate.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized EffectiveDate : buf len %d\n", buf.Len())

	// ContractExpiration (Timestamp)
	// fmt.Printf("Serializing ContractExpiration\n")
	{
		b, err := action.ContractExpiration.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized ContractExpiration : buf len %d\n", buf.Len())

	// ContractURI (string)
	// fmt.Printf("Serializing ContractURI\n")
	if err := WriteVarChar(buf, action.ContractURI, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ContractURI : buf len %d\n", buf.Len())

	// PrevRevTxID (TxId)
	// fmt.Printf("Serializing PrevRevTxID\n")
	{
		b, err := action.PrevRevTxID.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized PrevRevTxID : buf len %d\n", buf.Len())

	// Entities ([]Entity)
	// fmt.Printf("Serializing Entities\n")
	if err := WriteVariableSize(buf, uint64(len(action.Entities)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.Entities {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Entities : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in StaticContractFormation from the byte slice
func (action *StaticContractFormation) write(b []byte) (int, error) {
	// fmt.Printf("Reading StaticContractFormation : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// ContractName (string)
	// fmt.Printf("Reading ContractName : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.ContractName, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read ContractName : %d bytes remaining\n%+v\n", buf.Len(), action.ContractName)

	// ContractType (string)
	// fmt.Printf("Reading ContractType : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.ContractType, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read ContractType : %d bytes remaining\n%+v\n", buf.Len(), action.ContractType)

	// ContractCode (ContractCode)
	// fmt.Printf("Reading ContractCode : %d bytes remaining\n", buf.Len())
	if err := action.ContractCode.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ContractCode : %d bytes remaining\n%+v\n", buf.Len(), action.ContractCode)

	// ContractFileType (uint8)
	// fmt.Printf("Reading ContractFileType : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ContractFileType); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ContractFileType : %d bytes remaining\n%+v\n", buf.Len(), action.ContractFileType)

	// ContractFile ([]byte)
	// fmt.Printf("Reading ContractFile : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.ContractFile, err = ReadVarBin(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read ContractFile : %d bytes remaining\n%+v\n", buf.Len(), action.ContractFile)

	// SupportingDocsFileType (uint8)
	// fmt.Printf("Reading SupportingDocsFileType : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.SupportingDocsFileType); err != nil {
		return 0, err
	}

	// fmt.Printf("Read SupportingDocsFileType : %d bytes remaining\n%+v\n", buf.Len(), action.SupportingDocsFileType)

	// SupportingDocs (string)
	// fmt.Printf("Reading SupportingDocs : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.SupportingDocs, err = ReadVarChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read SupportingDocs : %d bytes remaining\n%+v\n", buf.Len(), action.SupportingDocs)

	// ContractRevision (uint32)
	// fmt.Printf("Reading ContractRevision : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ContractRevision); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ContractRevision : %d bytes remaining\n%+v\n", buf.Len(), action.ContractRevision)

	// GoverningLaw (string)
	// fmt.Printf("Reading GoverningLaw : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.GoverningLaw, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read GoverningLaw : %d bytes remaining\n%+v\n", buf.Len(), action.GoverningLaw)

	// Jurisdiction (string)
	// fmt.Printf("Reading Jurisdiction : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.Jurisdiction, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read Jurisdiction : %d bytes remaining\n%+v\n", buf.Len(), action.Jurisdiction)

	// EffectiveDate (Timestamp)
	// fmt.Printf("Reading EffectiveDate : %d bytes remaining\n", buf.Len())
	if err := action.EffectiveDate.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read EffectiveDate : %d bytes remaining\n%+v\n", buf.Len(), action.EffectiveDate)

	// ContractExpiration (Timestamp)
	// fmt.Printf("Reading ContractExpiration : %d bytes remaining\n", buf.Len())
	if err := action.ContractExpiration.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ContractExpiration : %d bytes remaining\n%+v\n", buf.Len(), action.ContractExpiration)

	// ContractURI (string)
	// fmt.Printf("Reading ContractURI : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.ContractURI, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read ContractURI : %d bytes remaining\n%+v\n", buf.Len(), action.ContractURI)

	// PrevRevTxID (TxId)
	// fmt.Printf("Reading PrevRevTxID : %d bytes remaining\n", buf.Len())
	if err := action.PrevRevTxID.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read PrevRevTxID : %d bytes remaining\n%+v\n", buf.Len(), action.PrevRevTxID)

	// Entities ([]Entity)
	// fmt.Printf("Reading Entities : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.Entities = make([]Entity, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Entity
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.Entities = append(action.Entities, newValue)
		}
	}

	// fmt.Printf("Read Entities : %d bytes remaining\n%+v\n", buf.Len(), action.Entities)

	// fmt.Printf("Read StaticContractFormation : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action StaticContractFormation) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action StaticContractFormation) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("ContractName:%#+v", action.ContractName))
	vals = append(vals, fmt.Sprintf("ContractType:%#+v", action.ContractType))
	vals = append(vals, fmt.Sprintf("ContractCode:%#+v", action.ContractCode))
	vals = append(vals, fmt.Sprintf("ContractFileType:%v", action.ContractFileType))
	vals = append(vals, fmt.Sprintf("ContractFile:%#x", action.ContractFile))
	vals = append(vals, fmt.Sprintf("SupportingDocsFileType:%v", action.SupportingDocsFileType))
	vals = append(vals, fmt.Sprintf("SupportingDocs:%#+v", action.SupportingDocs))
	vals = append(vals, fmt.Sprintf("ContractRevision:%v", action.ContractRevision))
	vals = append(vals, fmt.Sprintf("GoverningLaw:%#+v", action.GoverningLaw))
	vals = append(vals, fmt.Sprintf("Jurisdiction:%#+v", action.Jurisdiction))
	vals = append(vals, fmt.Sprintf("EffectiveDate:%#+v", action.EffectiveDate))
	vals = append(vals, fmt.Sprintf("ContractExpiration:%#+v", action.ContractExpiration))
	vals = append(vals, fmt.Sprintf("ContractURI:%#+v", action.ContractURI))
	vals = append(vals, fmt.Sprintf("PrevRevTxID:%#+v", action.PrevRevTxID))
	vals = append(vals, fmt.Sprintf("Entities:%#+v", action.Entities))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Order Order Action - Issuer to signal to the smart contract that the
// tokens that a particular public address(es) owns are to be confiscated,
// frozen, thawed or reconciled.
type Order struct {
	Header                 Header          `json:"header,omitempty"`                    // Common header data for all actions
	AssetType              string          `json:"asset_type,omitempty"`                // eg. Share, Bond, Ticket
	AssetCode              AssetCode       `json:"asset_code,omitempty"`                // 32 randomly generated bytes.  Each Asset Code should be unique.  However, an Asset Code is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type + Asset Code = Asset Code.  An Asset Code is a human readable identifier that can be used in a similar way to a Bitcoin (BSV) address.
	ComplianceAction       byte            `json:"compliance_action,omitempty"`         // Freeze (F), Thaw (T), Confiscate (C), Reconciliation (R)
	TargetAddresses        []TargetAddress `json:"target_addresses,omitempty"`          //
	DepositAddress         PublicKeyHash   `json:"deposit_address,omitempty"`           // The public address for confiscated tokens to be deposited in.  Null for Freeze, Thaw, actions. For Reconciliation actions the deposit address is who receives bitcoin.
	AuthorityName          string          `json:"authority_name,omitempty"`            // Length 0-255 bytes. Enforcement Authority Name (eg. Issuer, Queensland Police Service, Tokenized, etc.)
	SigAlgoAddressList     uint8           `json:"sig_algo_address_list,omitempty"`     // 0 = No Registry-signed Message, 1 = ECDSA+secp256k1
	AuthorityPublicKey     string          `json:"authority_public_key,omitempty"`      // Length 0-255 bytes. Public Key associated with the Enforcement Authority
	OrderSignature         string          `json:"order_signature,omitempty"`           // Length 0-255 bytes. Signature for a message that lists out the target addresses and deposit address. Signature of (Contract Address, Asset Code, Compliance Action, Supporting Evidence Hash, Time Out Expiration, TargetAddress1, TargetAddress1Qty, TargetAddressX, TargetAddressXQty,...,DepositAddress)
	SupportingEvidenceTxId TxId            `json:"supporting_evidence_tx_id,omitempty"` // SHA-256: warrant, court order, etc.
	RefTxnID               TxId            `json:"ref_txn_id,omitempty"`                // The settlement action that was dropped from the network.  Not applicable for Freeze, Thaw, and Confiscation orders.  Only applicable for reconcilliation actions.  No subfield when F, T, R is selected as the Compliance Action subfield.
	FreezePeriod           Timestamp       `json:"freeze_period,omitempty"`             // Used for a 'time out'.  Tokens are automatically unfrozen after the expiration timestamp without requiring a Thaw Action. Null value for Thaw, Confiscation and Reconciallitaion orders.
	Message                string          `json:"message,omitempty"`                   //
}

// Type returns the type identifer for this message.
func (action Order) Type() string {
	return CodeOrder
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *Order) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *Order) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// AssetType (string)
	// fmt.Printf("Serializing AssetType\n")
	if err := WriteFixedChar(buf, action.AssetType, 3); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AssetType : buf len %d\n", buf.Len())

	// AssetCode (AssetCode)
	// fmt.Printf("Serializing AssetCode\n")
	{
		b, err := action.AssetCode.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized AssetCode : buf len %d\n", buf.Len())

	// ComplianceAction (byte)
	// fmt.Printf("Serializing ComplianceAction\n")
	if err := write(buf, action.ComplianceAction); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ComplianceAction : buf len %d\n", buf.Len())

	// TargetAddresses ([]TargetAddress)
	// fmt.Printf("Serializing TargetAddresses\n")
	if err := WriteVariableSize(buf, uint64(len(action.TargetAddresses)), 16, 8); err != nil {
		return nil, err
	}
	for _, value := range action.TargetAddresses {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized TargetAddresses : buf len %d\n", buf.Len())

	// DepositAddress (PublicKeyHash)
	// fmt.Printf("Serializing DepositAddress\n")
	{
		b, err := action.DepositAddress.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized DepositAddress : buf len %d\n", buf.Len())

	// AuthorityName (string)
	// fmt.Printf("Serializing AuthorityName\n")
	if err := WriteVarChar(buf, action.AuthorityName, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AuthorityName : buf len %d\n", buf.Len())

	// SigAlgoAddressList (uint8)
	// fmt.Printf("Serializing SigAlgoAddressList\n")
	if err := write(buf, action.SigAlgoAddressList); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized SigAlgoAddressList : buf len %d\n", buf.Len())

	// AuthorityPublicKey (string)
	// fmt.Printf("Serializing AuthorityPublicKey\n")
	if err := WriteVarChar(buf, action.AuthorityPublicKey, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AuthorityPublicKey : buf len %d\n", buf.Len())

	// OrderSignature (string)
	// fmt.Printf("Serializing OrderSignature\n")
	if err := WriteVarChar(buf, action.OrderSignature, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized OrderSignature : buf len %d\n", buf.Len())

	// SupportingEvidenceTxId (TxId)
	// fmt.Printf("Serializing SupportingEvidenceTxId\n")
	{
		b, err := action.SupportingEvidenceTxId.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized SupportingEvidenceTxId : buf len %d\n", buf.Len())

	// RefTxnID (TxId)
	// fmt.Printf("Serializing RefTxnID\n")
	{
		b, err := action.RefTxnID.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized RefTxnID : buf len %d\n", buf.Len())

	// FreezePeriod (Timestamp)
	// fmt.Printf("Serializing FreezePeriod\n")
	{
		b, err := action.FreezePeriod.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized FreezePeriod : buf len %d\n", buf.Len())

	// Message (string)
	// fmt.Printf("Serializing Message\n")
	if err := WriteVarChar(buf, action.Message, 32); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized Message : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in Order from the byte slice
func (action *Order) write(b []byte) (int, error) {
	// fmt.Printf("Reading Order : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// AssetType (string)
	// fmt.Printf("Reading AssetType : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read AssetType : %d bytes remaining\n%+v\n", buf.Len(), action.AssetType)

	// AssetCode (AssetCode)
	// fmt.Printf("Reading AssetCode : %d bytes remaining\n", buf.Len())
	if err := action.AssetCode.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read AssetCode : %d bytes remaining\n%+v\n", buf.Len(), action.AssetCode)

	// ComplianceAction (byte)
	// fmt.Printf("Reading ComplianceAction : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ComplianceAction); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ComplianceAction : %d bytes remaining\n%+v\n", buf.Len(), action.ComplianceAction)

	// TargetAddresses ([]TargetAddress)
	// fmt.Printf("Reading TargetAddresses : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 16, 8)
		if err != nil {
			return 0, err
		}
		action.TargetAddresses = make([]TargetAddress, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue TargetAddress
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.TargetAddresses = append(action.TargetAddresses, newValue)
		}
	}

	// fmt.Printf("Read TargetAddresses : %d bytes remaining\n%+v\n", buf.Len(), action.TargetAddresses)

	// DepositAddress (PublicKeyHash)
	// fmt.Printf("Reading DepositAddress : %d bytes remaining\n", buf.Len())
	if err := action.DepositAddress.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read DepositAddress : %d bytes remaining\n%+v\n", buf.Len(), action.DepositAddress)

	// AuthorityName (string)
	// fmt.Printf("Reading AuthorityName : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.AuthorityName, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read AuthorityName : %d bytes remaining\n%+v\n", buf.Len(), action.AuthorityName)

	// SigAlgoAddressList (uint8)
	// fmt.Printf("Reading SigAlgoAddressList : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.SigAlgoAddressList); err != nil {
		return 0, err
	}

	// fmt.Printf("Read SigAlgoAddressList : %d bytes remaining\n%+v\n", buf.Len(), action.SigAlgoAddressList)

	// AuthorityPublicKey (string)
	// fmt.Printf("Reading AuthorityPublicKey : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.AuthorityPublicKey, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read AuthorityPublicKey : %d bytes remaining\n%+v\n", buf.Len(), action.AuthorityPublicKey)

	// OrderSignature (string)
	// fmt.Printf("Reading OrderSignature : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.OrderSignature, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read OrderSignature : %d bytes remaining\n%+v\n", buf.Len(), action.OrderSignature)

	// SupportingEvidenceTxId (TxId)
	// fmt.Printf("Reading SupportingEvidenceTxId : %d bytes remaining\n", buf.Len())
	if err := action.SupportingEvidenceTxId.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read SupportingEvidenceTxId : %d bytes remaining\n%+v\n", buf.Len(), action.SupportingEvidenceTxId)

	// RefTxnID (TxId)
	// fmt.Printf("Reading RefTxnID : %d bytes remaining\n", buf.Len())
	if err := action.RefTxnID.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read RefTxnID : %d bytes remaining\n%+v\n", buf.Len(), action.RefTxnID)

	// FreezePeriod (Timestamp)
	// fmt.Printf("Reading FreezePeriod : %d bytes remaining\n", buf.Len())
	if err := action.FreezePeriod.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read FreezePeriod : %d bytes remaining\n%+v\n", buf.Len(), action.FreezePeriod)

	// Message (string)
	// fmt.Printf("Reading Message : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.Message, err = ReadVarChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read Message : %d bytes remaining\n%+v\n", buf.Len(), action.Message)

	// fmt.Printf("Read Order : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action Order) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action Order) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("AssetType:%#+v", action.AssetType))
	vals = append(vals, fmt.Sprintf("AssetCode:%#+v", action.AssetCode))
	vals = append(vals, fmt.Sprintf("ComplianceAction:%#+v", action.ComplianceAction))
	vals = append(vals, fmt.Sprintf("TargetAddresses:%#+v", action.TargetAddresses))
	vals = append(vals, fmt.Sprintf("DepositAddress:%#+v", action.DepositAddress))
	vals = append(vals, fmt.Sprintf("AuthorityName:%#+v", action.AuthorityName))
	vals = append(vals, fmt.Sprintf("SigAlgoAddressList:%v", action.SigAlgoAddressList))
	vals = append(vals, fmt.Sprintf("AuthorityPublicKey:%#+v", action.AuthorityPublicKey))
	vals = append(vals, fmt.Sprintf("OrderSignature:%#+v", action.OrderSignature))
	vals = append(vals, fmt.Sprintf("SupportingEvidenceTxId:%#+v", action.SupportingEvidenceTxId))
	vals = append(vals, fmt.Sprintf("RefTxnID:%#+v", action.RefTxnID))
	vals = append(vals, fmt.Sprintf("FreezePeriod:%#+v", action.FreezePeriod))
	vals = append(vals, fmt.Sprintf("Message:%#+v", action.Message))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Freeze Freeze Action - To be used to comply with
// contractual/legal/issuer requirements. The target public address(es)
// will be marked as frozen. However the Freeze action publishes this fact
// to the public blockchain for transparency. The Contract will not respond
// to any actions requested by the frozen address.
type Freeze struct {
	Header    Header          `json:"header,omitempty"`    // Common header data for all actions
	Addresses []PublicKeyHash `json:"addresses,omitempty"` // Addresses holding tokens to be frozen.
	Timestamp Timestamp       `json:"timestamp,omitempty"` // Timestamp in nanoseconds of when the smart contract created the action.
}

// Type returns the type identifer for this message.
func (action Freeze) Type() string {
	return CodeFreeze
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *Freeze) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *Freeze) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Addresses ([]PublicKeyHash)
	// fmt.Printf("Serializing Addresses\n")
	if err := WriteVariableSize(buf, uint64(len(action.Addresses)), 16, 8); err != nil {
		return nil, err
	}
	for _, value := range action.Addresses {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Addresses : buf len %d\n", buf.Len())

	// Timestamp (Timestamp)
	// fmt.Printf("Serializing Timestamp\n")
	{
		b, err := action.Timestamp.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Timestamp : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in Freeze from the byte slice
func (action *Freeze) write(b []byte) (int, error) {
	// fmt.Printf("Reading Freeze : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// Addresses ([]PublicKeyHash)
	// fmt.Printf("Reading Addresses : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 16, 8)
		if err != nil {
			return 0, err
		}
		action.Addresses = make([]PublicKeyHash, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue PublicKeyHash
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.Addresses = append(action.Addresses, newValue)
		}
	}

	// fmt.Printf("Read Addresses : %d bytes remaining\n%+v\n", buf.Len(), action.Addresses)

	// Timestamp (Timestamp)
	// fmt.Printf("Reading Timestamp : %d bytes remaining\n", buf.Len())
	if err := action.Timestamp.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Timestamp : %d bytes remaining\n%+v\n", buf.Len(), action.Timestamp)

	// fmt.Printf("Read Freeze : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action Freeze) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action Freeze) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("Addresses:%#+v", action.Addresses))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", action.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Thaw Thaw Action - to be used to comply with contractual obligations or
// legal requirements. The Alleged Offender's tokens will be unfrozen to
// allow them to resume normal exchange and governance activities.
type Thaw struct {
	Header    Header          `json:"header,omitempty"`    // Common header data for all actions
	Addresses []PublicKeyHash `json:"addresses,omitempty"` // Addresses holding tokens to be thawed.
	RefTxID   TxId            `json:"ref_tx_id,omitempty"` // The related freeze action.
	Timestamp Timestamp       `json:"timestamp,omitempty"` // Timestamp in nanoseconds of when the smart contract created the action.
}

// Type returns the type identifer for this message.
func (action Thaw) Type() string {
	return CodeThaw
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *Thaw) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *Thaw) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Addresses ([]PublicKeyHash)
	// fmt.Printf("Serializing Addresses\n")
	if err := WriteVariableSize(buf, uint64(len(action.Addresses)), 16, 8); err != nil {
		return nil, err
	}
	for _, value := range action.Addresses {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Addresses : buf len %d\n", buf.Len())

	// RefTxID (TxId)
	// fmt.Printf("Serializing RefTxID\n")
	{
		b, err := action.RefTxID.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized RefTxID : buf len %d\n", buf.Len())

	// Timestamp (Timestamp)
	// fmt.Printf("Serializing Timestamp\n")
	{
		b, err := action.Timestamp.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Timestamp : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in Thaw from the byte slice
func (action *Thaw) write(b []byte) (int, error) {
	// fmt.Printf("Reading Thaw : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// Addresses ([]PublicKeyHash)
	// fmt.Printf("Reading Addresses : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 16, 8)
		if err != nil {
			return 0, err
		}
		action.Addresses = make([]PublicKeyHash, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue PublicKeyHash
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.Addresses = append(action.Addresses, newValue)
		}
	}

	// fmt.Printf("Read Addresses : %d bytes remaining\n%+v\n", buf.Len(), action.Addresses)

	// RefTxID (TxId)
	// fmt.Printf("Reading RefTxID : %d bytes remaining\n", buf.Len())
	if err := action.RefTxID.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read RefTxID : %d bytes remaining\n%+v\n", buf.Len(), action.RefTxID)

	// Timestamp (Timestamp)
	// fmt.Printf("Reading Timestamp : %d bytes remaining\n", buf.Len())
	if err := action.Timestamp.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Timestamp : %d bytes remaining\n%+v\n", buf.Len(), action.Timestamp)

	// fmt.Printf("Read Thaw : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action Thaw) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action Thaw) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("Addresses:%#+v", action.Addresses))
	vals = append(vals, fmt.Sprintf("RefTxID:%#+v", action.RefTxID))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", action.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Confiscation Confiscation Action - to be used to comply with contractual
// obligations, legal and/or issuer requirements.
type Confiscation struct {
	Header     Header          `json:"header,omitempty"`      // Common header data for all actions
	Addresses  []PublicKeyHash `json:"addresses,omitempty"`   // Addresses holding tokens to be confiscated.
	DepositQty uint64          `json:"deposit_qty,omitempty"` // Custodian's token balance after confiscation.
	Timestamp  Timestamp       `json:"timestamp,omitempty"`   // Timestamp in nanoseconds of when the smart contract created the action.
}

// Type returns the type identifer for this message.
func (action Confiscation) Type() string {
	return CodeConfiscation
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *Confiscation) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *Confiscation) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Addresses ([]PublicKeyHash)
	// fmt.Printf("Serializing Addresses\n")
	if err := WriteVariableSize(buf, uint64(len(action.Addresses)), 16, 8); err != nil {
		return nil, err
	}
	for _, value := range action.Addresses {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Addresses : buf len %d\n", buf.Len())

	// DepositQty (uint64)
	// fmt.Printf("Serializing DepositQty\n")
	if err := write(buf, action.DepositQty); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized DepositQty : buf len %d\n", buf.Len())

	// Timestamp (Timestamp)
	// fmt.Printf("Serializing Timestamp\n")
	{
		b, err := action.Timestamp.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Timestamp : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in Confiscation from the byte slice
func (action *Confiscation) write(b []byte) (int, error) {
	// fmt.Printf("Reading Confiscation : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// Addresses ([]PublicKeyHash)
	// fmt.Printf("Reading Addresses : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 16, 8)
		if err != nil {
			return 0, err
		}
		action.Addresses = make([]PublicKeyHash, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue PublicKeyHash
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.Addresses = append(action.Addresses, newValue)
		}
	}

	// fmt.Printf("Read Addresses : %d bytes remaining\n%+v\n", buf.Len(), action.Addresses)

	// DepositQty (uint64)
	// fmt.Printf("Reading DepositQty : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.DepositQty); err != nil {
		return 0, err
	}

	// fmt.Printf("Read DepositQty : %d bytes remaining\n%+v\n", buf.Len(), action.DepositQty)

	// Timestamp (Timestamp)
	// fmt.Printf("Reading Timestamp : %d bytes remaining\n", buf.Len())
	if err := action.Timestamp.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Timestamp : %d bytes remaining\n%+v\n", buf.Len(), action.Timestamp)

	// fmt.Printf("Read Confiscation : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action Confiscation) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action Confiscation) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("Addresses:%#+v", action.Addresses))
	vals = append(vals, fmt.Sprintf("DepositQty:%v", action.DepositQty))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", action.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Reconciliation Reconciliation Action - to be used at the direction of
// the issuer to fix record keeping errors with bitcoin and token balances.
type Reconciliation struct {
	Header    Header          `json:"header,omitempty"`    // Common header data for all actions
	Addresses []PublicKeyHash `json:"addresses,omitempty"` // Addresses holding tokens to be reconciled.
	Timestamp Timestamp       `json:"timestamp,omitempty"` // Timestamp in nanoseconds of when the smart contract created the action.
}

// Type returns the type identifer for this message.
func (action Reconciliation) Type() string {
	return CodeReconciliation
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *Reconciliation) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *Reconciliation) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Addresses ([]PublicKeyHash)
	// fmt.Printf("Serializing Addresses\n")
	if err := WriteVariableSize(buf, uint64(len(action.Addresses)), 16, 8); err != nil {
		return nil, err
	}
	for _, value := range action.Addresses {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Addresses : buf len %d\n", buf.Len())

	// Timestamp (Timestamp)
	// fmt.Printf("Serializing Timestamp\n")
	{
		b, err := action.Timestamp.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Timestamp : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in Reconciliation from the byte slice
func (action *Reconciliation) write(b []byte) (int, error) {
	// fmt.Printf("Reading Reconciliation : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// Addresses ([]PublicKeyHash)
	// fmt.Printf("Reading Addresses : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 16, 8)
		if err != nil {
			return 0, err
		}
		action.Addresses = make([]PublicKeyHash, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue PublicKeyHash
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.Addresses = append(action.Addresses, newValue)
		}
	}

	// fmt.Printf("Read Addresses : %d bytes remaining\n%+v\n", buf.Len(), action.Addresses)

	// Timestamp (Timestamp)
	// fmt.Printf("Reading Timestamp : %d bytes remaining\n", buf.Len())
	if err := action.Timestamp.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Timestamp : %d bytes remaining\n%+v\n", buf.Len(), action.Timestamp)

	// fmt.Printf("Read Reconciliation : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action Reconciliation) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action Reconciliation) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("Addresses:%#+v", action.Addresses))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", action.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Initiative Initiative Action - Allows Token Owners to propose an
// Initiative (aka Initiative/Shareholder vote). A significant cost -
// specified in the Contract Formation - can be attached to this action to
// reduce spam, as the resulting vote will be put to all token owners.
type Initiative struct {
	Header               Header      `json:"header,omitempty"`                 // Common header data for all actions
	AssetType            string      `json:"asset_type,omitempty"`             // eg. Share, Bond, Ticket
	AssetCode            AssetCode   `json:"asset_code,omitempty"`             // 32 randomly generated bytes.  Each Asset Code should be unique.  However, an Asset Code is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type + Asset Code = Asset Code.  An Asset Code is a human readable identifier that can be used in a similar way to a Bitcoin (BSV) address.
	VoteSystem           uint8       `json:"vote_system,omitempty"`            // X for Vote System X. (1-255, 0 is not valid.)
	Proposal             bool        `json:"proposal,omitempty"`               // 1 for a Proposal, 0 for an initiative that is requesting changes to specific subfields for modification. If this field is true, the subfields should be empty.  The smart contract cannot interpret the results of a vote when Proposal = 1.  All meaning is interpreted by the token owners and smart contract simply facilates the record keeping.  When Proposal = 0, the smart contract always assumes the first choice is a 'yes', or 'pass', if the threshold is met, and will process the proposed changes accordingly.
	ProposedChanges      []Amendment `json:"proposed_changes,omitempty"`       // Each element contains details of which fields to modify, or delete. Because the number of fields in a Contract and Asset is dynamic due to some fields being able to be repeated, the index value of the field needs to be calculated against the Contract or Asset the changes are to apply to. In the event of a Vote being created from this Initiative, the changes will be applied to the version of the Contract or Asset at that time.
	VoteOptions          string      `json:"vote_options,omitempty"`           // Length 1-255 bytes. 0 is not valid. Each byte allows for a different vote option.  Typical votes will likely be multiple choice or Y/N. Vote instances are identified by the Tx-ID. AB000000000 would be chosen for Y/N (binary) type votes.
	VoteMax              uint8       `json:"vote_max,omitempty"`               // Range: 1-X. How many selections can a voter make in a Ballot Cast.  1 is selected for Y/N (binary)
	ProposalDescription  string      `json:"proposal_description,omitempty"`   // Length restricted by the Bitcoin protocol. 0 is valid. Description or details of the vote
	ProposalDocumentHash [32]byte    `json:"proposal_document_hash,omitempty"` // SHA256 Hash of the proposal document to be distributed to voters.
	VoteCutOffTimestamp  Timestamp   `json:"vote_cut_off_timestamp,omitempty"` // Ballot casts after this timestamp will not be included. The vote has finished.
}

// Type returns the type identifer for this message.
func (action Initiative) Type() string {
	return CodeInitiative
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *Initiative) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *Initiative) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// AssetType (string)
	// fmt.Printf("Serializing AssetType\n")
	if err := WriteFixedChar(buf, action.AssetType, 3); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AssetType : buf len %d\n", buf.Len())

	// AssetCode (AssetCode)
	// fmt.Printf("Serializing AssetCode\n")
	{
		b, err := action.AssetCode.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized AssetCode : buf len %d\n", buf.Len())

	// VoteSystem (uint8)
	// fmt.Printf("Serializing VoteSystem\n")
	if err := write(buf, action.VoteSystem); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized VoteSystem : buf len %d\n", buf.Len())

	// Proposal (bool)
	// fmt.Printf("Serializing Proposal\n")
	if err := write(buf, action.Proposal); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized Proposal : buf len %d\n", buf.Len())

	// ProposedChanges ([]Amendment)
	// fmt.Printf("Serializing ProposedChanges\n")
	if err := WriteVariableSize(buf, uint64(len(action.ProposedChanges)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.ProposedChanges {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized ProposedChanges : buf len %d\n", buf.Len())

	// VoteOptions (string)
	// fmt.Printf("Serializing VoteOptions\n")
	if err := WriteVarChar(buf, action.VoteOptions, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized VoteOptions : buf len %d\n", buf.Len())

	// VoteMax (uint8)
	// fmt.Printf("Serializing VoteMax\n")
	if err := write(buf, action.VoteMax); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized VoteMax : buf len %d\n", buf.Len())

	// ProposalDescription (string)
	// fmt.Printf("Serializing ProposalDescription\n")
	if err := WriteVarChar(buf, action.ProposalDescription, 32); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ProposalDescription : buf len %d\n", buf.Len())

	// ProposalDocumentHash ([32]byte)
	// fmt.Printf("Serializing ProposalDocumentHash\n")
	if err := write(buf, action.ProposalDocumentHash); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ProposalDocumentHash : buf len %d\n", buf.Len())

	// VoteCutOffTimestamp (Timestamp)
	// fmt.Printf("Serializing VoteCutOffTimestamp\n")
	{
		b, err := action.VoteCutOffTimestamp.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized VoteCutOffTimestamp : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in Initiative from the byte slice
func (action *Initiative) write(b []byte) (int, error) {
	// fmt.Printf("Reading Initiative : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// AssetType (string)
	// fmt.Printf("Reading AssetType : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read AssetType : %d bytes remaining\n%+v\n", buf.Len(), action.AssetType)

	// AssetCode (AssetCode)
	// fmt.Printf("Reading AssetCode : %d bytes remaining\n", buf.Len())
	if err := action.AssetCode.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read AssetCode : %d bytes remaining\n%+v\n", buf.Len(), action.AssetCode)

	// VoteSystem (uint8)
	// fmt.Printf("Reading VoteSystem : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.VoteSystem); err != nil {
		return 0, err
	}

	// fmt.Printf("Read VoteSystem : %d bytes remaining\n%+v\n", buf.Len(), action.VoteSystem)

	// Proposal (bool)
	// fmt.Printf("Reading Proposal : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.Proposal); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Proposal : %d bytes remaining\n%+v\n", buf.Len(), action.Proposal)

	// ProposedChanges ([]Amendment)
	// fmt.Printf("Reading ProposedChanges : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.ProposedChanges = make([]Amendment, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Amendment
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.ProposedChanges = append(action.ProposedChanges, newValue)
		}
	}

	// fmt.Printf("Read ProposedChanges : %d bytes remaining\n%+v\n", buf.Len(), action.ProposedChanges)

	// VoteOptions (string)
	// fmt.Printf("Reading VoteOptions : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.VoteOptions, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read VoteOptions : %d bytes remaining\n%+v\n", buf.Len(), action.VoteOptions)

	// VoteMax (uint8)
	// fmt.Printf("Reading VoteMax : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.VoteMax); err != nil {
		return 0, err
	}

	// fmt.Printf("Read VoteMax : %d bytes remaining\n%+v\n", buf.Len(), action.VoteMax)

	// ProposalDescription (string)
	// fmt.Printf("Reading ProposalDescription : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.ProposalDescription, err = ReadVarChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read ProposalDescription : %d bytes remaining\n%+v\n", buf.Len(), action.ProposalDescription)

	// ProposalDocumentHash ([32]byte)
	// fmt.Printf("Reading ProposalDocumentHash : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ProposalDocumentHash); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ProposalDocumentHash : %d bytes remaining\n%+v\n", buf.Len(), action.ProposalDocumentHash)

	// VoteCutOffTimestamp (Timestamp)
	// fmt.Printf("Reading VoteCutOffTimestamp : %d bytes remaining\n", buf.Len())
	if err := action.VoteCutOffTimestamp.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read VoteCutOffTimestamp : %d bytes remaining\n%+v\n", buf.Len(), action.VoteCutOffTimestamp)

	// fmt.Printf("Read Initiative : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action Initiative) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action Initiative) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("AssetType:%#+v", action.AssetType))
	vals = append(vals, fmt.Sprintf("AssetCode:%#+v", action.AssetCode))
	vals = append(vals, fmt.Sprintf("VoteSystem:%v", action.VoteSystem))
	vals = append(vals, fmt.Sprintf("Proposal:%#+v", action.Proposal))
	vals = append(vals, fmt.Sprintf("ProposedChanges:%#+v", action.ProposedChanges))
	vals = append(vals, fmt.Sprintf("VoteOptions:%#+v", action.VoteOptions))
	vals = append(vals, fmt.Sprintf("VoteMax:%v", action.VoteMax))
	vals = append(vals, fmt.Sprintf("ProposalDescription:%#+v", action.ProposalDescription))
	vals = append(vals, fmt.Sprintf("ProposalDocumentHash:%#+v", action.ProposalDocumentHash))
	vals = append(vals, fmt.Sprintf("VoteCutOffTimestamp:%#+v", action.VoteCutOffTimestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Referendum Referendum Action - Issuer instructs the Contract to Initiate
// a Token Owner Vote. Usually used for contract amendments, organizational
// governance, etc.
type Referendum struct {
	Header               Header      `json:"header,omitempty"`                 // Common header data for all actions
	AssetSpecificVote    bool        `json:"asset_specific_vote,omitempty"`    // 1 - Yes, 0 - No.  No Asset Type/AssetCode subfields for N - No.
	AssetType            string      `json:"asset_type,omitempty"`             // eg. Share, Bond, Ticket
	AssetCode            AssetCode   `json:"asset_code,omitempty"`             // 32 randomly generated bytes.  Each Asset Code should be unique.  However, an Asset Code is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type + Asset Code = Asset Code.  An Asset Code is a human readable identifier that can be used in a similar way to a Bitcoin (BSV) address.
	VoteSystem           uint8       `json:"vote_system,omitempty"`            // X for Vote System X. (1-255, 0 is not valid.)
	Proposal             bool        `json:"proposal,omitempty"`               // 1 for a Proposal, 0 for an initiative that is requesting changes to specific subfields for modification. If this field is true, the subfields should be empty.  The smart contract cannot interpret the results of a vote when Proposal = 1.  All meaning is interpreted by the token owners and smart contract simply facilates the record keeping.  When Proposal = 0, the smart contract always assumes the first choice is a 'yes', or 'pass', if the threshold is met, and will process the proposed changes accordingly.
	ProposedChanges      []Amendment `json:"proposed_changes,omitempty"`       // Each element contains details of which fields to modify, or delete. Because the number of fields in a Contract and Asset is dynamic due to some fields being able to be repeated, the index value of the field needs to be calculated against the Contract or Asset the changes are to apply to. In the event of a Vote being created from this Initiative, the changes will be applied to the version of the Contract or Asset at that time.
	VoteOptions          string      `json:"vote_options,omitempty"`           // Length 1-255 bytes. 0 is not valid. Each byte allows for a different vote option.  Typical votes will likely be multiple choice or Y/N. Vote instances are identified by the Tx-ID. AB000000000 would be chosen for Y/N (binary) type votes. Only applicable if Proposal Type is set to P for Proposal.  All other Proposal Types will be binary.  Pass/Fail.
	VoteMax              uint8       `json:"vote_max,omitempty"`               // Range: 1-15. How many selections can a voter make in a Ballot Cast.  1 is selected for Y/N (binary)
	ProposalDescription  string      `json:"proposal_description,omitempty"`   // Length restricted by the Bitcoin protocol. 0 is valid. Description of the vote.
	ProposalDocumentHash [32]byte    `json:"proposal_document_hash,omitempty"` // SHA256 Hash of the proposal document to be distributed to voters
	VoteCutOffTimestamp  Timestamp   `json:"vote_cut_off_timestamp,omitempty"` // Ballot casts after this timestamp will not be included. The vote has finished.
}

// Type returns the type identifer for this message.
func (action Referendum) Type() string {
	return CodeReferendum
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *Referendum) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *Referendum) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// AssetSpecificVote (bool)
	// fmt.Printf("Serializing AssetSpecificVote\n")
	if err := write(buf, action.AssetSpecificVote); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AssetSpecificVote : buf len %d\n", buf.Len())

	// AssetType (string)
	// fmt.Printf("Serializing AssetType\n")
	if err := WriteFixedChar(buf, action.AssetType, 3); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AssetType : buf len %d\n", buf.Len())

	// AssetCode (AssetCode)
	// fmt.Printf("Serializing AssetCode\n")
	{
		b, err := action.AssetCode.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized AssetCode : buf len %d\n", buf.Len())

	// VoteSystem (uint8)
	// fmt.Printf("Serializing VoteSystem\n")
	if err := write(buf, action.VoteSystem); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized VoteSystem : buf len %d\n", buf.Len())

	// Proposal (bool)
	// fmt.Printf("Serializing Proposal\n")
	if err := write(buf, action.Proposal); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized Proposal : buf len %d\n", buf.Len())

	// ProposedChanges ([]Amendment)
	// fmt.Printf("Serializing ProposedChanges\n")
	if err := WriteVariableSize(buf, uint64(len(action.ProposedChanges)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.ProposedChanges {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized ProposedChanges : buf len %d\n", buf.Len())

	// VoteOptions (string)
	// fmt.Printf("Serializing VoteOptions\n")
	if err := WriteVarChar(buf, action.VoteOptions, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized VoteOptions : buf len %d\n", buf.Len())

	// VoteMax (uint8)
	// fmt.Printf("Serializing VoteMax\n")
	if err := write(buf, action.VoteMax); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized VoteMax : buf len %d\n", buf.Len())

	// ProposalDescription (string)
	// fmt.Printf("Serializing ProposalDescription\n")
	if err := WriteVarChar(buf, action.ProposalDescription, 32); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ProposalDescription : buf len %d\n", buf.Len())

	// ProposalDocumentHash ([32]byte)
	// fmt.Printf("Serializing ProposalDocumentHash\n")
	if err := write(buf, action.ProposalDocumentHash); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ProposalDocumentHash : buf len %d\n", buf.Len())

	// VoteCutOffTimestamp (Timestamp)
	// fmt.Printf("Serializing VoteCutOffTimestamp\n")
	{
		b, err := action.VoteCutOffTimestamp.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized VoteCutOffTimestamp : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in Referendum from the byte slice
func (action *Referendum) write(b []byte) (int, error) {
	// fmt.Printf("Reading Referendum : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// AssetSpecificVote (bool)
	// fmt.Printf("Reading AssetSpecificVote : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.AssetSpecificVote); err != nil {
		return 0, err
	}

	// fmt.Printf("Read AssetSpecificVote : %d bytes remaining\n%+v\n", buf.Len(), action.AssetSpecificVote)

	// AssetType (string)
	// fmt.Printf("Reading AssetType : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read AssetType : %d bytes remaining\n%+v\n", buf.Len(), action.AssetType)

	// AssetCode (AssetCode)
	// fmt.Printf("Reading AssetCode : %d bytes remaining\n", buf.Len())
	if err := action.AssetCode.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read AssetCode : %d bytes remaining\n%+v\n", buf.Len(), action.AssetCode)

	// VoteSystem (uint8)
	// fmt.Printf("Reading VoteSystem : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.VoteSystem); err != nil {
		return 0, err
	}

	// fmt.Printf("Read VoteSystem : %d bytes remaining\n%+v\n", buf.Len(), action.VoteSystem)

	// Proposal (bool)
	// fmt.Printf("Reading Proposal : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.Proposal); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Proposal : %d bytes remaining\n%+v\n", buf.Len(), action.Proposal)

	// ProposedChanges ([]Amendment)
	// fmt.Printf("Reading ProposedChanges : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.ProposedChanges = make([]Amendment, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Amendment
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.ProposedChanges = append(action.ProposedChanges, newValue)
		}
	}

	// fmt.Printf("Read ProposedChanges : %d bytes remaining\n%+v\n", buf.Len(), action.ProposedChanges)

	// VoteOptions (string)
	// fmt.Printf("Reading VoteOptions : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.VoteOptions, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read VoteOptions : %d bytes remaining\n%+v\n", buf.Len(), action.VoteOptions)

	// VoteMax (uint8)
	// fmt.Printf("Reading VoteMax : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.VoteMax); err != nil {
		return 0, err
	}

	// fmt.Printf("Read VoteMax : %d bytes remaining\n%+v\n", buf.Len(), action.VoteMax)

	// ProposalDescription (string)
	// fmt.Printf("Reading ProposalDescription : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.ProposalDescription, err = ReadVarChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read ProposalDescription : %d bytes remaining\n%+v\n", buf.Len(), action.ProposalDescription)

	// ProposalDocumentHash ([32]byte)
	// fmt.Printf("Reading ProposalDocumentHash : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ProposalDocumentHash); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ProposalDocumentHash : %d bytes remaining\n%+v\n", buf.Len(), action.ProposalDocumentHash)

	// VoteCutOffTimestamp (Timestamp)
	// fmt.Printf("Reading VoteCutOffTimestamp : %d bytes remaining\n", buf.Len())
	if err := action.VoteCutOffTimestamp.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read VoteCutOffTimestamp : %d bytes remaining\n%+v\n", buf.Len(), action.VoteCutOffTimestamp)

	// fmt.Printf("Read Referendum : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action Referendum) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action Referendum) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("AssetSpecificVote:%#+v", action.AssetSpecificVote))
	vals = append(vals, fmt.Sprintf("AssetType:%#+v", action.AssetType))
	vals = append(vals, fmt.Sprintf("AssetCode:%#+v", action.AssetCode))
	vals = append(vals, fmt.Sprintf("VoteSystem:%v", action.VoteSystem))
	vals = append(vals, fmt.Sprintf("Proposal:%#+v", action.Proposal))
	vals = append(vals, fmt.Sprintf("ProposedChanges:%#+v", action.ProposedChanges))
	vals = append(vals, fmt.Sprintf("VoteOptions:%#+v", action.VoteOptions))
	vals = append(vals, fmt.Sprintf("VoteMax:%v", action.VoteMax))
	vals = append(vals, fmt.Sprintf("ProposalDescription:%#+v", action.ProposalDescription))
	vals = append(vals, fmt.Sprintf("ProposalDocumentHash:%#+v", action.ProposalDocumentHash))
	vals = append(vals, fmt.Sprintf("VoteCutOffTimestamp:%#+v", action.VoteCutOffTimestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Vote Vote Action - A vote is created by the Contract in response to a
// valid Referendum (Issuer) or Initiative (User) Action.
type Vote struct {
	Header    Header    `json:"header,omitempty"`    // Common header data for all actions
	Timestamp Timestamp `json:"timestamp,omitempty"` // Timestamp in nanoseconds of when the smart contract created the action.
}

// Type returns the type identifer for this message.
func (action Vote) Type() string {
	return CodeVote
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *Vote) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *Vote) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Timestamp (Timestamp)
	// fmt.Printf("Serializing Timestamp\n")
	{
		b, err := action.Timestamp.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Timestamp : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in Vote from the byte slice
func (action *Vote) write(b []byte) (int, error) {
	// fmt.Printf("Reading Vote : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// Timestamp (Timestamp)
	// fmt.Printf("Reading Timestamp : %d bytes remaining\n", buf.Len())
	if err := action.Timestamp.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Timestamp : %d bytes remaining\n%+v\n", buf.Len(), action.Timestamp)

	// fmt.Printf("Read Vote : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action Vote) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action Vote) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", action.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// BallotCast Ballot Cast Action - Used by Token Owners to cast their
// ballot (vote) on proposals raised by the Issuer (Referendum) or other
// token holders (Initiative). 1 Vote per token unless a vote multiplier is
// specified in the relevant Asset Definition action.
type BallotCast struct {
	Header    Header    `json:"header,omitempty"`     // Common header data for all actions
	AssetType string    `json:"asset_type,omitempty"` // eg. Share, Bond, Ticket
	AssetCode AssetCode `json:"asset_code,omitempty"` // 32 randomly generated bytes.  Each Asset Code should be unique.  However, an Asset Code is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type + Asset Code = Asset Code.  An Asset Code is a human readable identifier that can be used in a similar way to a Bitcoin (BSV) address.
	VoteTxID  TxId      `json:"vote_tx_id,omitempty"` // Tx ID of the Vote the Ballot Cast is being made for.
	Vote      string    `json:"vote,omitempty"`       // Length 1-255 bytes. 0 is not valid. Accept, Reject, Abstain, Spoiled, Multiple Choice, or Preference List. 15 options total. Order of preference.  1st position = 1st choice. 2nd position = 2nd choice, etc.  A is always Accept and B is always reject in a Y/N votes.
}

// Type returns the type identifer for this message.
func (action BallotCast) Type() string {
	return CodeBallotCast
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *BallotCast) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *BallotCast) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// AssetType (string)
	// fmt.Printf("Serializing AssetType\n")
	if err := WriteFixedChar(buf, action.AssetType, 3); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AssetType : buf len %d\n", buf.Len())

	// AssetCode (AssetCode)
	// fmt.Printf("Serializing AssetCode\n")
	{
		b, err := action.AssetCode.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized AssetCode : buf len %d\n", buf.Len())

	// VoteTxID (TxId)
	// fmt.Printf("Serializing VoteTxID\n")
	{
		b, err := action.VoteTxID.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized VoteTxID : buf len %d\n", buf.Len())

	// Vote (string)
	// fmt.Printf("Serializing Vote\n")
	if err := WriteVarChar(buf, action.Vote, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized Vote : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in BallotCast from the byte slice
func (action *BallotCast) write(b []byte) (int, error) {
	// fmt.Printf("Reading BallotCast : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// AssetType (string)
	// fmt.Printf("Reading AssetType : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read AssetType : %d bytes remaining\n%+v\n", buf.Len(), action.AssetType)

	// AssetCode (AssetCode)
	// fmt.Printf("Reading AssetCode : %d bytes remaining\n", buf.Len())
	if err := action.AssetCode.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read AssetCode : %d bytes remaining\n%+v\n", buf.Len(), action.AssetCode)

	// VoteTxID (TxId)
	// fmt.Printf("Reading VoteTxID : %d bytes remaining\n", buf.Len())
	if err := action.VoteTxID.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read VoteTxID : %d bytes remaining\n%+v\n", buf.Len(), action.VoteTxID)

	// Vote (string)
	// fmt.Printf("Reading Vote : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.Vote, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read Vote : %d bytes remaining\n%+v\n", buf.Len(), action.Vote)

	// fmt.Printf("Read BallotCast : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action BallotCast) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action BallotCast) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("AssetType:%#+v", action.AssetType))
	vals = append(vals, fmt.Sprintf("AssetCode:%#+v", action.AssetCode))
	vals = append(vals, fmt.Sprintf("VoteTxID:%#+v", action.VoteTxID))
	vals = append(vals, fmt.Sprintf("Vote:%#+v", action.Vote))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// BallotCounted Ballot Counted Action - The smart contract will respond to
// a Ballot Cast action with a Ballot Counted action if the Ballot Cast is
// valid. If the Ballot Cast is not valid, then the smart contract will
// respond with a Rejection Action.
type BallotCounted struct {
	Header    Header    `json:"header,omitempty"`    // Common header data for all actions
	Timestamp Timestamp `json:"timestamp,omitempty"` // Timestamp in nanoseconds of when the smart contract created the action.
}

// Type returns the type identifer for this message.
func (action BallotCounted) Type() string {
	return CodeBallotCounted
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *BallotCounted) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *BallotCounted) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Timestamp (Timestamp)
	// fmt.Printf("Serializing Timestamp\n")
	{
		b, err := action.Timestamp.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Timestamp : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in BallotCounted from the byte slice
func (action *BallotCounted) write(b []byte) (int, error) {
	// fmt.Printf("Reading BallotCounted : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// Timestamp (Timestamp)
	// fmt.Printf("Reading Timestamp : %d bytes remaining\n", buf.Len())
	if err := action.Timestamp.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Timestamp : %d bytes remaining\n%+v\n", buf.Len(), action.Timestamp)

	// fmt.Printf("Read BallotCounted : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action BallotCounted) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action BallotCounted) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", action.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Result Result Action - Once a vote has been completed the results are
// published.
type Result struct {
	Header           Header      `json:"header,omitempty"`             // Common header data for all actions
	AssetType        string      `json:"asset_type,omitempty"`         // eg. Share, Bond, Ticket
	AssetCode        AssetCode   `json:"asset_code,omitempty"`         // 32 randomly generated bytes.  Each Asset Code should be unique.  However, an Asset Code is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type + Asset Code = Asset Code.  An Asset Code is a human readable identifier that can be used in a similar way to a Bitcoin (BSV) address.
	Proposal         bool        `json:"proposal,omitempty"`           // 1 for a Proposal, 0 for an initiative that is requesting changes to specific subfields for modification. If this field is true, the subfields should be empty.  The smart contract cannot interpret the results of a vote when Proposal = 1.  All meaning is interpreted by the token owners and smart contract simply facilates the record keeping.  When Proposal = 0, the smart contract always assumes the first choice is a 'yes', or 'pass', if the threshold is met, and will process the proposed changes accordingly.
	ProposedChanges  []Amendment `json:"proposed_changes,omitempty"`   // Each element contains details of which fields to modify, or delete. Because the number of fields in a Contract and Asset is dynamic due to some fields being able to be repeated, the index value of the field needs to be calculated against the Contract or Asset the changes are to apply to. In the event of a Vote being created from this Initiative, the changes will be applied to the version of the Contract or Asset at that time.
	VoteTxID         TxId        `json:"vote_tx_id,omitempty"`         // Link to the Vote Action txn.
	VoteOptionsCount uint8       `json:"vote_options_count,omitempty"` // Number of Vote Options to follow.
	OptionXTally     uint64      `json:"option_x_tally,omitempty"`     // Number of valid votes counted for Option X
	Result           string      `json:"result,omitempty"`             // Length 1-255 bytes. 0 is not valid. The Option with the most votes. In the event of a draw for 1st place, all winning options are listed.
	Timestamp        Timestamp   `json:"timestamp,omitempty"`          // Timestamp in nanoseconds of when the smart contract created the action.
}

// Type returns the type identifer for this message.
func (action Result) Type() string {
	return CodeResult
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *Result) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *Result) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// AssetType (string)
	// fmt.Printf("Serializing AssetType\n")
	if err := WriteFixedChar(buf, action.AssetType, 3); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized AssetType : buf len %d\n", buf.Len())

	// AssetCode (AssetCode)
	// fmt.Printf("Serializing AssetCode\n")
	{
		b, err := action.AssetCode.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized AssetCode : buf len %d\n", buf.Len())

	// Proposal (bool)
	// fmt.Printf("Serializing Proposal\n")
	if err := write(buf, action.Proposal); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized Proposal : buf len %d\n", buf.Len())

	// ProposedChanges ([]Amendment)
	// fmt.Printf("Serializing ProposedChanges\n")
	if err := WriteVariableSize(buf, uint64(len(action.ProposedChanges)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.ProposedChanges {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized ProposedChanges : buf len %d\n", buf.Len())

	// VoteTxID (TxId)
	// fmt.Printf("Serializing VoteTxID\n")
	{
		b, err := action.VoteTxID.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized VoteTxID : buf len %d\n", buf.Len())

	// VoteOptionsCount (uint8)
	// fmt.Printf("Serializing VoteOptionsCount\n")
	if err := write(buf, action.VoteOptionsCount); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized VoteOptionsCount : buf len %d\n", buf.Len())

	// OptionXTally (uint64)
	// fmt.Printf("Serializing OptionXTally\n")
	if err := write(buf, action.OptionXTally); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized OptionXTally : buf len %d\n", buf.Len())

	// Result (string)
	// fmt.Printf("Serializing Result\n")
	if err := WriteVarChar(buf, action.Result, 8); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized Result : buf len %d\n", buf.Len())

	// Timestamp (Timestamp)
	// fmt.Printf("Serializing Timestamp\n")
	{
		b, err := action.Timestamp.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Timestamp : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in Result from the byte slice
func (action *Result) write(b []byte) (int, error) {
	// fmt.Printf("Reading Result : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// AssetType (string)
	// fmt.Printf("Reading AssetType : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read AssetType : %d bytes remaining\n%+v\n", buf.Len(), action.AssetType)

	// AssetCode (AssetCode)
	// fmt.Printf("Reading AssetCode : %d bytes remaining\n", buf.Len())
	if err := action.AssetCode.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read AssetCode : %d bytes remaining\n%+v\n", buf.Len(), action.AssetCode)

	// Proposal (bool)
	// fmt.Printf("Reading Proposal : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.Proposal); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Proposal : %d bytes remaining\n%+v\n", buf.Len(), action.Proposal)

	// ProposedChanges ([]Amendment)
	// fmt.Printf("Reading ProposedChanges : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.ProposedChanges = make([]Amendment, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Amendment
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.ProposedChanges = append(action.ProposedChanges, newValue)
		}
	}

	// fmt.Printf("Read ProposedChanges : %d bytes remaining\n%+v\n", buf.Len(), action.ProposedChanges)

	// VoteTxID (TxId)
	// fmt.Printf("Reading VoteTxID : %d bytes remaining\n", buf.Len())
	if err := action.VoteTxID.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read VoteTxID : %d bytes remaining\n%+v\n", buf.Len(), action.VoteTxID)

	// VoteOptionsCount (uint8)
	// fmt.Printf("Reading VoteOptionsCount : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.VoteOptionsCount); err != nil {
		return 0, err
	}

	// fmt.Printf("Read VoteOptionsCount : %d bytes remaining\n%+v\n", buf.Len(), action.VoteOptionsCount)

	// OptionXTally (uint64)
	// fmt.Printf("Reading OptionXTally : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.OptionXTally); err != nil {
		return 0, err
	}

	// fmt.Printf("Read OptionXTally : %d bytes remaining\n%+v\n", buf.Len(), action.OptionXTally)

	// Result (string)
	// fmt.Printf("Reading Result : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.Result, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read Result : %d bytes remaining\n%+v\n", buf.Len(), action.Result)

	// Timestamp (Timestamp)
	// fmt.Printf("Reading Timestamp : %d bytes remaining\n", buf.Len())
	if err := action.Timestamp.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Timestamp : %d bytes remaining\n%+v\n", buf.Len(), action.Timestamp)

	// fmt.Printf("Read Result : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action Result) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action Result) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("AssetType:%#+v", action.AssetType))
	vals = append(vals, fmt.Sprintf("AssetCode:%#+v", action.AssetCode))
	vals = append(vals, fmt.Sprintf("Proposal:%#+v", action.Proposal))
	vals = append(vals, fmt.Sprintf("ProposedChanges:%#+v", action.ProposedChanges))
	vals = append(vals, fmt.Sprintf("VoteTxID:%#+v", action.VoteTxID))
	vals = append(vals, fmt.Sprintf("VoteOptionsCount:%v", action.VoteOptionsCount))
	vals = append(vals, fmt.Sprintf("OptionXTally:%v", action.OptionXTally))
	vals = append(vals, fmt.Sprintf("Result:%#+v", action.Result))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", action.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Message Message Action - the message action is a general purpose
// communication action. 'Twitter/SMS' for Issuers/Investors/Users. The
// message txn can also be used for passing partially signed txns on-chain,
// establishing private communication channels and EDI (receipting,
// invoices, PO, and private offers/bids). The messages are broken down by
// type for easy filtering in the a user's wallet. The Message Types are
// listed in the Message Types table.
type Message struct {
	Header         Header   `json:"header,omitempty"`          // Common header data for all actions
	AddressIndexes []uint16 `json:"address_indexes,omitempty"` // Associates the message to a particular output by the index.
	MessageType    string   `json:"message_type,omitempty"`    // Potential for up to 65,535 different message types
	MessagePayload []byte   `json:"message_payload,omitempty"` // Public or private (RSA public key, Diffie-Hellman). Issuers/Contracts can send the signifying amount of satoshis to themselves for public announcements or private 'notes' if encrypted. See Message Types for a full list of potential use cases.

}

// Type returns the type identifer for this message.
func (action Message) Type() string {
	return CodeMessage
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *Message) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *Message) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// AddressIndexes ([]uint16)
	// fmt.Printf("Serializing AddressIndexes\n")
	if err := WriteVariableSize(buf, uint64(len(action.AddressIndexes)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.AddressIndexes {
		if err := write(buf, value); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized AddressIndexes : buf len %d\n", buf.Len())

	// MessageType (string)
	// fmt.Printf("Serializing MessageType\n")
	if err := WriteFixedChar(buf, action.MessageType, 4); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized MessageType : buf len %d\n", buf.Len())

	// MessagePayload ([]byte)
	// fmt.Printf("Serializing MessagePayload\n")
	if err := WriteVarBin(buf, action.MessagePayload, 32); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized MessagePayload : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in Message from the byte slice
func (action *Message) write(b []byte) (int, error) {
	// fmt.Printf("Reading Message : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// AddressIndexes ([]uint16)
	// fmt.Printf("Reading AddressIndexes : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.AddressIndexes = make([]uint16, size, size)
		if err := read(buf, &action.AddressIndexes); err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read AddressIndexes : %d bytes remaining\n%+v\n", buf.Len(), action.AddressIndexes)

	// MessageType (string)
	// fmt.Printf("Reading MessageType : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.MessageType, err = ReadFixedChar(buf, 4)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read MessageType : %d bytes remaining\n%+v\n", buf.Len(), action.MessageType)

	// MessagePayload ([]byte)
	// fmt.Printf("Reading MessagePayload : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.MessagePayload, err = ReadVarBin(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read MessagePayload : %d bytes remaining\n%+v\n", buf.Len(), action.MessagePayload)

	// fmt.Printf("Read Message : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action Message) PayloadMessage() (PayloadMessage, error) {
	p := MessageTypeMapping(action.MessageType)
	if p == nil {
		return nil, fmt.Errorf("Undefined message type : %s", action.MessageType)
	}

	if _, err := p.Write(action.MessagePayload); err != nil {
		return nil, err
	}

	return p, nil
}

func (action Message) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("AddressIndexes:%v", action.AddressIndexes))
	vals = append(vals, fmt.Sprintf("MessageType:%#+v", action.MessageType))
	vals = append(vals, fmt.Sprintf("MessagePayload:%#x", action.MessagePayload))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Rejection Rejection Action - used to reject request actions that do not
// comply with the Contract. If money is to be returned to a User then it
// is used in lieu of the Settlement Action to properly account for token
// balances. All Issuer/User request Actions must be responded to by the
// Contract with an Action. The only exception to this rule is when there
// is not enough fees in the first Action for the Contract response action
// to remain revenue neutral. If not enough fees are attached to pay for
// the Contract response then the Contract will not respond.
type Rejection struct {
	Header                Header    `json:"header,omitempty"`                  // Common header data for all actions
	QtyReceivingAddresses uint8     `json:"qty_receiving_addresses,omitempty"` // 0-255 Message Receiving Addresses
	AddressIndexes        []uint16  `json:"address_indexes,omitempty"`         // Associates the message to a particular output by the index.
	RejectionType         uint8     `json:"rejection_type,omitempty"`          // Classifies the rejection by a type.
	MessagePayload        string    `json:"message_payload,omitempty"`         // Length 0-65,535 bytes. Message that explains the reasoning for a rejection, if needed.  Most rejection types will be captured by the Rejection Type Subfield.
	Timestamp             Timestamp `json:"timestamp,omitempty"`               // Timestamp in nanoseconds of when the smart contract created the action.
}

// Type returns the type identifer for this message.
func (action Rejection) Type() string {
	return CodeRejection
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *Rejection) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *Rejection) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// QtyReceivingAddresses (uint8)
	// fmt.Printf("Serializing QtyReceivingAddresses\n")
	if err := write(buf, action.QtyReceivingAddresses); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized QtyReceivingAddresses : buf len %d\n", buf.Len())

	// AddressIndexes ([]uint16)
	// fmt.Printf("Serializing AddressIndexes\n")
	if err := WriteVariableSize(buf, uint64(len(action.AddressIndexes)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.AddressIndexes {
		if err := write(buf, value); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized AddressIndexes : buf len %d\n", buf.Len())

	// RejectionType (uint8)
	// fmt.Printf("Serializing RejectionType\n")
	if err := write(buf, action.RejectionType); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized RejectionType : buf len %d\n", buf.Len())

	// MessagePayload (string)
	// fmt.Printf("Serializing MessagePayload\n")
	if err := WriteVarChar(buf, action.MessagePayload, 32); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized MessagePayload : buf len %d\n", buf.Len())

	// Timestamp (Timestamp)
	// fmt.Printf("Serializing Timestamp\n")
	{
		b, err := action.Timestamp.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Timestamp : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in Rejection from the byte slice
func (action *Rejection) write(b []byte) (int, error) {
	// fmt.Printf("Reading Rejection : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// QtyReceivingAddresses (uint8)
	// fmt.Printf("Reading QtyReceivingAddresses : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.QtyReceivingAddresses); err != nil {
		return 0, err
	}

	// fmt.Printf("Read QtyReceivingAddresses : %d bytes remaining\n%+v\n", buf.Len(), action.QtyReceivingAddresses)

	// AddressIndexes ([]uint16)
	// fmt.Printf("Reading AddressIndexes : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.AddressIndexes = make([]uint16, size, size)
		if err := read(buf, &action.AddressIndexes); err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read AddressIndexes : %d bytes remaining\n%+v\n", buf.Len(), action.AddressIndexes)

	// RejectionType (uint8)
	// fmt.Printf("Reading RejectionType : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.RejectionType); err != nil {
		return 0, err
	}

	// fmt.Printf("Read RejectionType : %d bytes remaining\n%+v\n", buf.Len(), action.RejectionType)

	// MessagePayload (string)
	// fmt.Printf("Reading MessagePayload : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.MessagePayload, err = ReadVarChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read MessagePayload : %d bytes remaining\n%+v\n", buf.Len(), action.MessagePayload)

	// Timestamp (Timestamp)
	// fmt.Printf("Reading Timestamp : %d bytes remaining\n", buf.Len())
	if err := action.Timestamp.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Timestamp : %d bytes remaining\n%+v\n", buf.Len(), action.Timestamp)

	// fmt.Printf("Read Rejection : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action Rejection) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action Rejection) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("QtyReceivingAddresses:%v", action.QtyReceivingAddresses))
	vals = append(vals, fmt.Sprintf("AddressIndexes:%v", action.AddressIndexes))
	vals = append(vals, fmt.Sprintf("RejectionType:%v", action.RejectionType))
	vals = append(vals, fmt.Sprintf("MessagePayload:%#+v", action.MessagePayload))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", action.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Establishment Establishment Action - Establishes an on-chain register.
type Establishment struct {
	Header  Header `json:"header,omitempty"`  // Common header data for all actions
	Message string `json:"message,omitempty"` //
}

// Type returns the type identifer for this message.
func (action Establishment) Type() string {
	return CodeEstablishment
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *Establishment) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *Establishment) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Message (string)
	// fmt.Printf("Serializing Message\n")
	if err := WriteVarChar(buf, action.Message, 32); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized Message : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in Establishment from the byte slice
func (action *Establishment) write(b []byte) (int, error) {
	// fmt.Printf("Reading Establishment : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// Message (string)
	// fmt.Printf("Reading Message : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.Message, err = ReadVarChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read Message : %d bytes remaining\n%+v\n", buf.Len(), action.Message)

	// fmt.Printf("Read Establishment : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action Establishment) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action Establishment) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("Message:%#+v", action.Message))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Addition Addition Action - Adds an entry to the Register.
type Addition struct {
	Header  Header `json:"header,omitempty"`  // Common header data for all actions
	Message string `json:"message,omitempty"` //
}

// Type returns the type identifer for this message.
func (action Addition) Type() string {
	return CodeAddition
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *Addition) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *Addition) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Message (string)
	// fmt.Printf("Serializing Message\n")
	if err := WriteVarChar(buf, action.Message, 32); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized Message : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in Addition from the byte slice
func (action *Addition) write(b []byte) (int, error) {
	// fmt.Printf("Reading Addition : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// Message (string)
	// fmt.Printf("Reading Message : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.Message, err = ReadVarChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read Message : %d bytes remaining\n%+v\n", buf.Len(), action.Message)

	// fmt.Printf("Read Addition : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action Addition) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action Addition) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("Message:%#+v", action.Message))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Alteration Alteration Action - A register entry/record can be altered.
type Alteration struct {
	Header  Header `json:"header,omitempty"`  // Common header data for all actions
	Message string `json:"message,omitempty"` //
}

// Type returns the type identifer for this message.
func (action Alteration) Type() string {
	return CodeAlteration
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *Alteration) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *Alteration) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Message (string)
	// fmt.Printf("Serializing Message\n")
	if err := WriteVarChar(buf, action.Message, 32); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized Message : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in Alteration from the byte slice
func (action *Alteration) write(b []byte) (int, error) {
	// fmt.Printf("Reading Alteration : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// Message (string)
	// fmt.Printf("Reading Message : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.Message, err = ReadVarChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read Message : %d bytes remaining\n%+v\n", buf.Len(), action.Message)

	// fmt.Printf("Read Alteration : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action Alteration) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action Alteration) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("Message:%#+v", action.Message))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Removal Removal Action - Removes an entry/record from the Register.
type Removal struct {
	Header  Header `json:"header,omitempty"`  // Common header data for all actions
	Message string `json:"message,omitempty"` //
}

// Type returns the type identifer for this message.
func (action Removal) Type() string {
	return CodeRemoval
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *Removal) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *Removal) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Message (string)
	// fmt.Printf("Serializing Message\n")
	if err := WriteVarChar(buf, action.Message, 32); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized Message : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in Removal from the byte slice
func (action *Removal) write(b []byte) (int, error) {
	// fmt.Printf("Reading Removal : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// Message (string)
	// fmt.Printf("Reading Message : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.Message, err = ReadVarChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read Message : %d bytes remaining\n%+v\n", buf.Len(), action.Message)

	// fmt.Printf("Read Removal : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action Removal) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action Removal) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("Message:%#+v", action.Message))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Transfer A Token Owner(s) Sends, Exchanges or Swaps a token(s) or
// Bitcoin for a token(s) or Bitcoin. Can be as simple as sending a single
// token to a receiver. Or can be as complex as many senders sending many
// different assets - controlled by many different smart contracts - to a
// number of receivers. This action also supports atomic swaps (tokens for
// tokens). Since many parties and contracts can be involved in a transfer
// and the corresponding settlement action, the partially signed T1 and T2
// actions will need to be passed around on-chain with an M1 action, or
// off-chain.
type Transfer struct {
	Header              Header          `json:"header,omitempty"`                // Common header data for all actions
	Assets              []AssetTransfer `json:"assets,omitempty"`                // The Assets involved in the Transfer Action.
	OfferExpiry         Timestamp       `json:"offer_expiry,omitempty"`          // This prevents any party from holding on to the partially signed message as a form of an option.  Eg. the exchange at this price is valid for 30 mins.
	ExchangeFeeCurrency string          `json:"exchange_fee_currency,omitempty"` // BSV, USD, AUD, EUR, etc.
	ExchangeFeeVar      float32         `json:"exchange_fee_var,omitempty"`      // Percent of the value of the transaction
	ExchangeFeeFixed    float32         `json:"exchange_fee_fixed,omitempty"`    // Fixed fee
	ExchangeFeeAddress  PublicKeyHash   `json:"exchange_fee_address,omitempty"`  // Identifies the public address that the exchange fee should be paid to.
}

// Type returns the type identifer for this message.
func (action Transfer) Type() string {
	return CodeTransfer
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *Transfer) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *Transfer) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Assets ([]AssetTransfer)
	// fmt.Printf("Serializing Assets\n")
	if err := WriteVariableSize(buf, uint64(len(action.Assets)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.Assets {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Assets : buf len %d\n", buf.Len())

	// OfferExpiry (Timestamp)
	// fmt.Printf("Serializing OfferExpiry\n")
	{
		b, err := action.OfferExpiry.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized OfferExpiry : buf len %d\n", buf.Len())

	// ExchangeFeeCurrency (string)
	// fmt.Printf("Serializing ExchangeFeeCurrency\n")
	if err := WriteFixedChar(buf, action.ExchangeFeeCurrency, 3); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ExchangeFeeCurrency : buf len %d\n", buf.Len())

	// ExchangeFeeVar (float32)
	// fmt.Printf("Serializing ExchangeFeeVar\n")
	if err := write(buf, action.ExchangeFeeVar); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ExchangeFeeVar : buf len %d\n", buf.Len())

	// ExchangeFeeFixed (float32)
	// fmt.Printf("Serializing ExchangeFeeFixed\n")
	if err := write(buf, action.ExchangeFeeFixed); err != nil {
		return nil, err
	}
	// fmt.Printf("Serialized ExchangeFeeFixed : buf len %d\n", buf.Len())

	// ExchangeFeeAddress (PublicKeyHash)
	// fmt.Printf("Serializing ExchangeFeeAddress\n")
	{
		b, err := action.ExchangeFeeAddress.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized ExchangeFeeAddress : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in Transfer from the byte slice
func (action *Transfer) write(b []byte) (int, error) {
	// fmt.Printf("Reading Transfer : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// Assets ([]AssetTransfer)
	// fmt.Printf("Reading Assets : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.Assets = make([]AssetTransfer, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue AssetTransfer
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.Assets = append(action.Assets, newValue)
		}
	}

	// fmt.Printf("Read Assets : %d bytes remaining\n%+v\n", buf.Len(), action.Assets)

	// OfferExpiry (Timestamp)
	// fmt.Printf("Reading OfferExpiry : %d bytes remaining\n", buf.Len())
	if err := action.OfferExpiry.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read OfferExpiry : %d bytes remaining\n%+v\n", buf.Len(), action.OfferExpiry)

	// ExchangeFeeCurrency (string)
	// fmt.Printf("Reading ExchangeFeeCurrency : %d bytes remaining\n", buf.Len())
	{
		var err error
		action.ExchangeFeeCurrency, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// fmt.Printf("Read ExchangeFeeCurrency : %d bytes remaining\n%+v\n", buf.Len(), action.ExchangeFeeCurrency)

	// ExchangeFeeVar (float32)
	// fmt.Printf("Reading ExchangeFeeVar : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ExchangeFeeVar); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ExchangeFeeVar : %d bytes remaining\n%+v\n", buf.Len(), action.ExchangeFeeVar)

	// ExchangeFeeFixed (float32)
	// fmt.Printf("Reading ExchangeFeeFixed : %d bytes remaining\n", buf.Len())
	if err := read(buf, &action.ExchangeFeeFixed); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ExchangeFeeFixed : %d bytes remaining\n%+v\n", buf.Len(), action.ExchangeFeeFixed)

	// ExchangeFeeAddress (PublicKeyHash)
	// fmt.Printf("Reading ExchangeFeeAddress : %d bytes remaining\n", buf.Len())
	if err := action.ExchangeFeeAddress.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read ExchangeFeeAddress : %d bytes remaining\n%+v\n", buf.Len(), action.ExchangeFeeAddress)

	// fmt.Printf("Read Transfer : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action Transfer) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action Transfer) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("Assets:%#+v", action.Assets))
	vals = append(vals, fmt.Sprintf("OfferExpiry:%#+v", action.OfferExpiry))
	vals = append(vals, fmt.Sprintf("ExchangeFeeCurrency:%#+v", action.ExchangeFeeCurrency))
	vals = append(vals, fmt.Sprintf("ExchangeFeeVar:%v", action.ExchangeFeeVar))
	vals = append(vals, fmt.Sprintf("ExchangeFeeFixed:%v", action.ExchangeFeeFixed))
	vals = append(vals, fmt.Sprintf("ExchangeFeeAddress:%#+v", action.ExchangeFeeAddress))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Settlement Settlement Action - Settles the transfer request of bitcoins
// and tokens from transfer (T1) actions.
type Settlement struct {
	Header    Header            `json:"header,omitempty"`    // Common header data for all actions
	Assets    []AssetSettlement `json:"assets,omitempty"`    // The Assets settled by the transfer action.
	Timestamp Timestamp         `json:"timestamp,omitempty"` // Timestamp in nanoseconds of when the smart contract created the action.
}

// Type returns the type identifer for this message.
func (action Settlement) Type() string {
	return CodeSettlement
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *Settlement) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *Settlement) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Assets ([]AssetSettlement)
	// fmt.Printf("Serializing Assets\n")
	if err := WriteVariableSize(buf, uint64(len(action.Assets)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range action.Assets {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Assets : buf len %d\n", buf.Len())

	// Timestamp (Timestamp)
	// fmt.Printf("Serializing Timestamp\n")
	{
		b, err := action.Timestamp.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
	// fmt.Printf("Serialized Timestamp : buf len %d\n", buf.Len())

	return buf.Bytes(), nil
}

// write populates the fields in Settlement from the byte slice
func (action *Settlement) write(b []byte) (int, error) {
	// fmt.Printf("Reading Settlement : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

	// Header (Header)
	// fmt.Printf("Reading Header : %d bytes remaining\n", buf.Len())
	if err := action.Header.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Header : %d bytes remaining\n%+v\n", buf.Len(), action.Header)

	// Assets ([]AssetSettlement)
	// fmt.Printf("Reading Assets : %d bytes remaining\n", buf.Len())
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		action.Assets = make([]AssetSettlement, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue AssetSettlement
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.Assets = append(action.Assets, newValue)
		}
	}

	// fmt.Printf("Read Assets : %d bytes remaining\n%+v\n", buf.Len(), action.Assets)

	// Timestamp (Timestamp)
	// fmt.Printf("Reading Timestamp : %d bytes remaining\n", buf.Len())
	if err := action.Timestamp.Write(buf); err != nil {
		return 0, err
	}

	// fmt.Printf("Read Timestamp : %d bytes remaining\n%+v\n", buf.Len(), action.Timestamp)

	// fmt.Printf("Read Settlement : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (action Settlement) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (action Settlement) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", action.Header))
	vals = append(vals, fmt.Sprintf("Assets:%#+v", action.Assets))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", action.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}