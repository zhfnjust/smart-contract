package state

import (
	"github.com/tokenized/smart-contract/pkg/inspector"
	"github.com/tokenized/smart-contract/pkg/protocol"
)

// Contract represents a Smart Contract.
type Contract struct {
	ID        protocol.PublicKeyHash `json:"id,omitempty"`
	Revision  uint32                 `json:"revision,omitempty"`
	CreatedAt protocol.Timestamp     `json:"created_at,omitempty"`
	UpdatedAt protocol.Timestamp     `json:"updated_at,omitempty"`
	Timestamp protocol.Timestamp     `json:"timestamp,omitempty"`

	Issuer   protocol.PublicKeyHash `json:"issuer,omitempty"`
	Operator protocol.PublicKeyHash `json:"operator,omitempty"`

	ContractName               string             `json:"contract_name,omitempty"`
	ContractFileType           uint8              `json:"contract_file_type,omitempty"`
	ContractFile               []byte             `json:"contract_file,omitempty"`
	GoverningLaw               string             `json:"governing_law,omitempty"`
	Jurisdiction               string             `json:"jurisdiction,omitempty"`
	ContractExpiration         protocol.Timestamp `json:"contract_expiration,omitempty"`
	ContractURI                string             `json:"contract_uri,omitempty"`
	IssuerName                 string             `json:"issuer_name,omitempty"`
	IssuerType                 byte               `json:"issuer_type,omitempty"`
	IssuerLogoURL              string             `json:"issuer_logo_url,omitempty"`
	ContractOperatorID         string             `json:"contract_operator_id,omitempty"`
	ContractAuthFlags          [16]byte           `json:"contract_auth_flags,omitempty"`
	VotingSystems              []VotingSystem     `json:"voting_systems,omitempty"`
	RestrictedQtyAssets        uint64             `json:"restricted_qty_assets,omitempty"`
	ReferendumProposal         bool               `json:"referendum_proposal,omitempty"`
	InitiativeProposal         bool               `json:"initiative_proposal,omitempty"`
	Registries                 []Registry         `json:"registries,omitempty"`
	IssuerAddress              bool               `json:"issuer_address,omitempty"`
	UnitNumber                 string             `json:"unit_number,omitempty"`
	BuildingNumber             string             `json:"building_number,omitempty"`
	Street                     string             `json:"street,omitempty"`
	SuburbCity                 string             `json:"suburb_city,omitempty"`
	TerritoryStateProvinceCode string             `json:"territory_state_province_code,omitempty"`
	CountryCode                string             `json:"country_code,omitempty"`
	PostalZIPCode              string             `json:"postal_zip_code,omitempty"`
	EmailAddress               string             `json:"email_address,omitempty"`
	PhoneNumber                string             `json:"phone_number,omitempty"`
	KeyRoles                   []KeyRole          `json:"key_roles,omitempty"`
	NotableRoles               []NotableRole      `json:"notable_roles,omitempty"`

	Assets map[protocol.AssetCode]Asset `json:"assets,omitempty"`
	Votes  map[protocol.TxId]Vote       `json:"votes,omitempty"`
}

type Asset struct {
	ID        protocol.AssetCode `json:"id,omitempty"`
	Revision  uint32             `json:"revision,omitempty"`
	CreatedAt protocol.Timestamp `json:"created_at,omitempty"`
	UpdatedAt protocol.Timestamp `json:"updated_at,omitempty"`
	Timestamp protocol.Timestamp `json:"timestamp,omitempty"`

	AssetType                   string  `json:"asset_type,omitempty"`
	AssetAuthFlags              [8]byte `json:"asset_auth_flags,omitempty"`
	TransfersPermitted          bool    `json:"transfers_permitted,omitempty"`
	TradeRestrictions           string  `json:"trade_restrictions,omitempty"`
	EnforcementOrdersPermitted  bool    `json:"enforcement_orders_permitted,omitempty"`
	VoteMultiplier              uint8   `json:"vote_multiplier,omitempty"`
	ReferendumProposal          bool    `json:"referendum_proposal,omitempty"`
	InitiativeProposal          bool    `json:"initiative_proposal,omitempty"`
	AssetModificationGovernance bool    `json:"asset_modification_governance,omitempty"`
	TokenQty                    uint64  `json:"token_qty,omitempty"`
	ContractFeeCurrency         string  `json:"contract_fee_currency,omitempty"`
	ContractFeeVar              float32 `json:"contract_fee_var,omitempty"`
	ContractFeeFixed            float32 `json:"contract_fee_fixed,omitempty"`
	AssetPayload                []byte  `json:"asset_payload,omitempty"`

	Holdings map[protocol.PublicKeyHash]Holding `json:"holdings,omitempty"`
}

type Holding struct {
	Address         protocol.PublicKeyHash `json:"address,omit_empty"`
	Balance         uint64                 `json:"balance,omit_empty"`
	HoldingStatuses []HoldingStatus        `json:"order_status,omitempty"`
	CreatedAt       protocol.Timestamp     `json:"created_at,omit_empty"`
}

type HoldingStatus struct {
	Code    string             `json:"code,omit_empty"`
	Expires protocol.Timestamp `json:"expires,omitempty"`
	Balance uint64             `json:"balance,omit_empty"`
	TxId    protocol.TxId      `json:"tx_id,omit_empty"`
}

type Vote struct {
	Address              protocol.PublicKeyHash `json:"address,omit_empty"`
	AssetType            string                 `json:"asset_type,omit_empty"`
	AssetCode            protocol.AssetCode     `json:"asset_id,omit_empty"`
	VoteType             byte                   `json:"vote_type,omit_empty"`
	VoteOptions          []uint8                `json:"vote_options,omit_empty"`
	VoteMax              uint8                  `json:"vote_max,omit_empty"`
	VoteLogic            byte                   `json:"vote_logic,omit_empty"`
	ProposalDescription  string                 `json:"proposal_description,omit_empty"`
	ProposalDocumentHash string                 `json:"proposal_document_hash,omit_empty"`
	VoteCutOffTimestamp  protocol.Timestamp     `json:"vote_cut_off_timestamp,omit_empty"`
	RefTxID              protocol.TxId          `json:"ref_txn_id_hash,omit_empty"`
	Ballots              map[string]Ballot      `json:"ballots,omit_empty,omit_empty"`
	UTXO                 inspector.UTXO         `json:"utxo,omit_empty,omit_empty"`
	Result               *Result                `json:"result,omitempty,omit_empty"`
	UsedBy               string                 `json:"used_by,omit_empty,omit_empty"`
	CreatedAt            protocol.Timestamp     `json:"created_at,omit_empty,omit_empty"`
}

type Ballot struct {
	Address   protocol.PublicKeyHash `json:"address,omit_empty"`
	AssetType string                 `json:"asset_type,omit_empty"`
	AssetCode protocol.AssetCode     `json:"asset_id,omit_empty"`
	VoteTxnID protocol.TxId          `json:"vote_txn_id,omit_empty"`
	Vote      []uint8                `json:"vote,omit_empty"`
	CreatedAt protocol.Timestamp     `json:"created_at,omit_empty"`
}

type Result map[uint8]uint64

type VotingSystem struct {
	Name                        string  `json:"name,omit_empty"`
	System                      [8]byte `json:"system,omit_empty"`
	Method                      byte    `json:"method,omit_empty"`
	Logic                       byte    `json:"logic,omit_empty"`
	ThresholdPercentage         uint8   `json:"threshold_percent,omit_empty"`
	VoteMultiplierPermitted     byte    `json:"multiplier_permitted,omit_empty"`
	InitiativeThreshold         float32 `json:"threshold,omit_empty"`
	InitiativeThresholdCurrency string  `json:"threshold_currency,omit_empty"`
}

type Registry struct {
	Name      string                 `json:"name,omit_empty"`
	URL       string                 `json:"url,omit_empty"`
	PublicKey protocol.PublicKeyHash `json:"public_key,omit_empty"`
}

type KeyRole struct {
	Type byte   `json:"type,omit_empty"`
	Name string `json:"name,omit_empty"`
}

type NotableRole struct {
	Type byte   `json:"type,omit_empty"`
	Name string `json:"name,omit_empty"`
}
