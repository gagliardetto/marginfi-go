// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marginfi

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UtpMangoUseCancelPerpOrder is the `utpMangoUseCancelPerpOrder` instruction.
type UtpMangoUseCancelPerpOrder struct {
	OrderId     *ag_binary.Int128
	InvalidIdOk *bool

	// [0] = [] marginfiAccount
	//
	// [1] = [] marginfiGroup
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] mangoAuthority
	//
	// [4] = [WRITE] mangoAccount
	//
	// [5] = [] mangoProgram
	//
	// [6] = [] mangoGroup
	//
	// [7] = [WRITE] mangoPerpMarket
	//
	// [8] = [WRITE] mangoBids
	//
	// [9] = [WRITE] mangoAsks
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUtpMangoUseCancelPerpOrderInstructionBuilder creates a new `UtpMangoUseCancelPerpOrder` instruction builder.
func NewUtpMangoUseCancelPerpOrderInstructionBuilder() *UtpMangoUseCancelPerpOrder {
	nd := &UtpMangoUseCancelPerpOrder{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

// SetOrderId sets the "orderId" parameter.
func (inst *UtpMangoUseCancelPerpOrder) SetOrderId(orderId ag_binary.Int128) *UtpMangoUseCancelPerpOrder {
	inst.OrderId = &orderId
	return inst
}

// SetInvalidIdOk sets the "invalidIdOk" parameter.
func (inst *UtpMangoUseCancelPerpOrder) SetInvalidIdOk(invalidIdOk bool) *UtpMangoUseCancelPerpOrder {
	inst.InvalidIdOk = &invalidIdOk
	return inst
}

// SetMarginfiAccount sets the "marginfiAccount" account.
func (inst *UtpMangoUseCancelPerpOrder) SetMarginfiAccount(marginfiAccount ag_solanago.PublicKey) *UtpMangoUseCancelPerpOrder {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(marginfiAccount)
	return inst
}

// GetMarginfiAccount gets the "marginfiAccount" account.
func (inst *UtpMangoUseCancelPerpOrder) GetMarginfiAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMarginfiGroupAccount sets the "marginfiGroup" account.
func (inst *UtpMangoUseCancelPerpOrder) SetMarginfiGroupAccount(marginfiGroup ag_solanago.PublicKey) *UtpMangoUseCancelPerpOrder {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(marginfiGroup)
	return inst
}

// GetMarginfiGroupAccount gets the "marginfiGroup" account.
func (inst *UtpMangoUseCancelPerpOrder) GetMarginfiGroupAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *UtpMangoUseCancelPerpOrder) SetAuthorityAccount(authority ag_solanago.PublicKey) *UtpMangoUseCancelPerpOrder {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *UtpMangoUseCancelPerpOrder) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMangoAuthorityAccount sets the "mangoAuthority" account.
func (inst *UtpMangoUseCancelPerpOrder) SetMangoAuthorityAccount(mangoAuthority ag_solanago.PublicKey) *UtpMangoUseCancelPerpOrder {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mangoAuthority)
	return inst
}

// GetMangoAuthorityAccount gets the "mangoAuthority" account.
func (inst *UtpMangoUseCancelPerpOrder) GetMangoAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMangoAccount sets the "mangoAccount" account.
func (inst *UtpMangoUseCancelPerpOrder) SetMangoAccount(mangoAccount ag_solanago.PublicKey) *UtpMangoUseCancelPerpOrder {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(mangoAccount).WRITE()
	return inst
}

// GetMangoAccount gets the "mangoAccount" account.
func (inst *UtpMangoUseCancelPerpOrder) GetMangoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMangoProgramAccount sets the "mangoProgram" account.
func (inst *UtpMangoUseCancelPerpOrder) SetMangoProgramAccount(mangoProgram ag_solanago.PublicKey) *UtpMangoUseCancelPerpOrder {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(mangoProgram)
	return inst
}

// GetMangoProgramAccount gets the "mangoProgram" account.
func (inst *UtpMangoUseCancelPerpOrder) GetMangoProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMangoGroupAccount sets the "mangoGroup" account.
func (inst *UtpMangoUseCancelPerpOrder) SetMangoGroupAccount(mangoGroup ag_solanago.PublicKey) *UtpMangoUseCancelPerpOrder {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(mangoGroup)
	return inst
}

// GetMangoGroupAccount gets the "mangoGroup" account.
func (inst *UtpMangoUseCancelPerpOrder) GetMangoGroupAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetMangoPerpMarketAccount sets the "mangoPerpMarket" account.
func (inst *UtpMangoUseCancelPerpOrder) SetMangoPerpMarketAccount(mangoPerpMarket ag_solanago.PublicKey) *UtpMangoUseCancelPerpOrder {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(mangoPerpMarket).WRITE()
	return inst
}

// GetMangoPerpMarketAccount gets the "mangoPerpMarket" account.
func (inst *UtpMangoUseCancelPerpOrder) GetMangoPerpMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetMangoBidsAccount sets the "mangoBids" account.
func (inst *UtpMangoUseCancelPerpOrder) SetMangoBidsAccount(mangoBids ag_solanago.PublicKey) *UtpMangoUseCancelPerpOrder {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(mangoBids).WRITE()
	return inst
}

// GetMangoBidsAccount gets the "mangoBids" account.
func (inst *UtpMangoUseCancelPerpOrder) GetMangoBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetMangoAsksAccount sets the "mangoAsks" account.
func (inst *UtpMangoUseCancelPerpOrder) SetMangoAsksAccount(mangoAsks ag_solanago.PublicKey) *UtpMangoUseCancelPerpOrder {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(mangoAsks).WRITE()
	return inst
}

// GetMangoAsksAccount gets the "mangoAsks" account.
func (inst *UtpMangoUseCancelPerpOrder) GetMangoAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

func (inst UtpMangoUseCancelPerpOrder) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UtpMangoUseCancelPerpOrder,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UtpMangoUseCancelPerpOrder) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UtpMangoUseCancelPerpOrder) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.OrderId == nil {
			return errors.New("OrderId parameter is not set")
		}
		if inst.InvalidIdOk == nil {
			return errors.New("InvalidIdOk parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.MarginfiAccount is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MarginfiGroup is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.MangoAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.MangoAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.MangoProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.MangoGroup is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.MangoPerpMarket is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.MangoBids is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.MangoAsks is not set")
		}
	}
	return nil
}

func (inst *UtpMangoUseCancelPerpOrder) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UtpMangoUseCancelPerpOrder")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("    OrderId", *inst.OrderId))
						paramsBranch.Child(ag_format.Param("InvalidIdOk", *inst.InvalidIdOk))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       marginfi", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  marginfiGroup", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      authority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta(" mangoAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("          mango", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("   mangoProgram", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("     mangoGroup", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("mangoPerpMarket", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("      mangoBids", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("      mangoAsks", inst.AccountMetaSlice.Get(9)))
					})
				})
		})
}

func (obj UtpMangoUseCancelPerpOrder) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `OrderId` param:
	err = encoder.Encode(obj.OrderId)
	if err != nil {
		return err
	}
	// Serialize `InvalidIdOk` param:
	err = encoder.Encode(obj.InvalidIdOk)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UtpMangoUseCancelPerpOrder) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `OrderId`:
	err = decoder.Decode(&obj.OrderId)
	if err != nil {
		return err
	}
	// Deserialize `InvalidIdOk`:
	err = decoder.Decode(&obj.InvalidIdOk)
	if err != nil {
		return err
	}
	return nil
}

// NewUtpMangoUseCancelPerpOrderInstruction declares a new UtpMangoUseCancelPerpOrder instruction with the provided parameters and accounts.
func NewUtpMangoUseCancelPerpOrderInstruction(
	// Parameters:
	orderId ag_binary.Int128,
	invalidIdOk bool,
	// Accounts:
	marginfiAccount ag_solanago.PublicKey,
	marginfiGroup ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	mangoAuthority ag_solanago.PublicKey,
	mangoAccount ag_solanago.PublicKey,
	mangoProgram ag_solanago.PublicKey,
	mangoGroup ag_solanago.PublicKey,
	mangoPerpMarket ag_solanago.PublicKey,
	mangoBids ag_solanago.PublicKey,
	mangoAsks ag_solanago.PublicKey) *UtpMangoUseCancelPerpOrder {
	return NewUtpMangoUseCancelPerpOrderInstructionBuilder().
		SetOrderId(orderId).
		SetInvalidIdOk(invalidIdOk).
		SetMarginfiAccount(marginfiAccount).
		SetMarginfiGroupAccount(marginfiGroup).
		SetAuthorityAccount(authority).
		SetMangoAuthorityAccount(mangoAuthority).
		SetMangoAccount(mangoAccount).
		SetMangoProgramAccount(mangoProgram).
		SetMangoGroupAccount(mangoGroup).
		SetMangoPerpMarketAccount(mangoPerpMarket).
		SetMangoBidsAccount(mangoBids).
		SetMangoAsksAccount(mangoAsks)
}
