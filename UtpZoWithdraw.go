// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marginfi

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UtpZoWithdraw is the `utpZoWithdraw` instruction.
type UtpZoWithdraw struct {
	Amount *uint64

	// [0] = [WRITE] marginfiAccount
	//
	// [1] = [WRITE] marginfiGroup
	//
	// [2] = [SIGNER] signer
	//
	// [3] = [WRITE] marginCollateralVault
	//
	// [4] = [] utpAuthority
	//
	// [5] = [WRITE] zoMargin
	//
	// [6] = [] zoProgram
	//
	// [7] = [WRITE] zoState
	//
	// [8] = [WRITE] zoStateSigner
	//
	// [9] = [WRITE] zoCache
	//
	// [10] = [WRITE] zoControl
	//
	// [11] = [WRITE] zoVault
	//
	// [12] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUtpZoWithdrawInstructionBuilder creates a new `UtpZoWithdraw` instruction builder.
func NewUtpZoWithdrawInstructionBuilder() *UtpZoWithdraw {
	nd := &UtpZoWithdraw{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 13),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
func (inst *UtpZoWithdraw) SetAmount(amount uint64) *UtpZoWithdraw {
	inst.Amount = &amount
	return inst
}

// SetMarginfiAccount sets the "marginfiAccount" account.
func (inst *UtpZoWithdraw) SetMarginfiAccount(marginfiAccount ag_solanago.PublicKey) *UtpZoWithdraw {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(marginfiAccount).WRITE()
	return inst
}

// GetMarginfiAccount gets the "marginfiAccount" account.
func (inst *UtpZoWithdraw) GetMarginfiAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMarginfiGroupAccount sets the "marginfiGroup" account.
func (inst *UtpZoWithdraw) SetMarginfiGroupAccount(marginfiGroup ag_solanago.PublicKey) *UtpZoWithdraw {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(marginfiGroup).WRITE()
	return inst
}

// GetMarginfiGroupAccount gets the "marginfiGroup" account.
func (inst *UtpZoWithdraw) GetMarginfiGroupAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSignerAccount sets the "signer" account.
func (inst *UtpZoWithdraw) SetSignerAccount(signer ag_solanago.PublicKey) *UtpZoWithdraw {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(signer).SIGNER()
	return inst
}

// GetSignerAccount gets the "signer" account.
func (inst *UtpZoWithdraw) GetSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMarginCollateralVaultAccount sets the "marginCollateralVault" account.
func (inst *UtpZoWithdraw) SetMarginCollateralVaultAccount(marginCollateralVault ag_solanago.PublicKey) *UtpZoWithdraw {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(marginCollateralVault).WRITE()
	return inst
}

// GetMarginCollateralVaultAccount gets the "marginCollateralVault" account.
func (inst *UtpZoWithdraw) GetMarginCollateralVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUtpAuthorityAccount sets the "utpAuthority" account.
func (inst *UtpZoWithdraw) SetUtpAuthorityAccount(utpAuthority ag_solanago.PublicKey) *UtpZoWithdraw {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(utpAuthority)
	return inst
}

// GetUtpAuthorityAccount gets the "utpAuthority" account.
func (inst *UtpZoWithdraw) GetUtpAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetZoMarginAccount sets the "zoMargin" account.
func (inst *UtpZoWithdraw) SetZoMarginAccount(zoMargin ag_solanago.PublicKey) *UtpZoWithdraw {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(zoMargin).WRITE()
	return inst
}

// GetZoMarginAccount gets the "zoMargin" account.
func (inst *UtpZoWithdraw) GetZoMarginAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetZoProgramAccount sets the "zoProgram" account.
func (inst *UtpZoWithdraw) SetZoProgramAccount(zoProgram ag_solanago.PublicKey) *UtpZoWithdraw {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(zoProgram)
	return inst
}

// GetZoProgramAccount gets the "zoProgram" account.
func (inst *UtpZoWithdraw) GetZoProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetZoStateAccount sets the "zoState" account.
func (inst *UtpZoWithdraw) SetZoStateAccount(zoState ag_solanago.PublicKey) *UtpZoWithdraw {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(zoState).WRITE()
	return inst
}

// GetZoStateAccount gets the "zoState" account.
func (inst *UtpZoWithdraw) GetZoStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetZoStateSignerAccount sets the "zoStateSigner" account.
func (inst *UtpZoWithdraw) SetZoStateSignerAccount(zoStateSigner ag_solanago.PublicKey) *UtpZoWithdraw {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(zoStateSigner).WRITE()
	return inst
}

// GetZoStateSignerAccount gets the "zoStateSigner" account.
func (inst *UtpZoWithdraw) GetZoStateSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetZoCacheAccount sets the "zoCache" account.
func (inst *UtpZoWithdraw) SetZoCacheAccount(zoCache ag_solanago.PublicKey) *UtpZoWithdraw {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(zoCache).WRITE()
	return inst
}

// GetZoCacheAccount gets the "zoCache" account.
func (inst *UtpZoWithdraw) GetZoCacheAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetZoControlAccount sets the "zoControl" account.
func (inst *UtpZoWithdraw) SetZoControlAccount(zoControl ag_solanago.PublicKey) *UtpZoWithdraw {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(zoControl).WRITE()
	return inst
}

// GetZoControlAccount gets the "zoControl" account.
func (inst *UtpZoWithdraw) GetZoControlAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetZoVaultAccount sets the "zoVault" account.
func (inst *UtpZoWithdraw) SetZoVaultAccount(zoVault ag_solanago.PublicKey) *UtpZoWithdraw {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(zoVault).WRITE()
	return inst
}

// GetZoVaultAccount gets the "zoVault" account.
func (inst *UtpZoWithdraw) GetZoVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *UtpZoWithdraw) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *UtpZoWithdraw {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *UtpZoWithdraw) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

func (inst UtpZoWithdraw) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UtpZoWithdraw,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UtpZoWithdraw) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UtpZoWithdraw) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
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
			return errors.New("accounts.Signer is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.MarginCollateralVault is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.UtpAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ZoMargin is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.ZoProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.ZoState is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.ZoStateSigner is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.ZoCache is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.ZoControl is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.ZoVault is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *UtpZoWithdraw) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UtpZoWithdraw")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=13]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             marginfi", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        marginfiGroup", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("               signer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("marginCollateralVault", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("         utpAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("             zoMargin", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("            zoProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("              zoState", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("        zoStateSigner", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("              zoCache", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("            zoControl", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("              zoVault", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("         tokenProgram", inst.AccountMetaSlice.Get(12)))
					})
				})
		})
}

func (obj UtpZoWithdraw) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UtpZoWithdraw) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewUtpZoWithdrawInstruction declares a new UtpZoWithdraw instruction with the provided parameters and accounts.
func NewUtpZoWithdrawInstruction(
	// Parameters:
	amount uint64,
	// Accounts:
	marginfiAccount ag_solanago.PublicKey,
	marginfiGroup ag_solanago.PublicKey,
	signer ag_solanago.PublicKey,
	marginCollateralVault ag_solanago.PublicKey,
	utpAuthority ag_solanago.PublicKey,
	zoMargin ag_solanago.PublicKey,
	zoProgram ag_solanago.PublicKey,
	zoState ag_solanago.PublicKey,
	zoStateSigner ag_solanago.PublicKey,
	zoCache ag_solanago.PublicKey,
	zoControl ag_solanago.PublicKey,
	zoVault ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *UtpZoWithdraw {
	return NewUtpZoWithdrawInstructionBuilder().
		SetAmount(amount).
		SetMarginfiAccount(marginfiAccount).
		SetMarginfiGroupAccount(marginfiGroup).
		SetSignerAccount(signer).
		SetMarginCollateralVaultAccount(marginCollateralVault).
		SetUtpAuthorityAccount(utpAuthority).
		SetZoMarginAccount(zoMargin).
		SetZoProgramAccount(zoProgram).
		SetZoStateAccount(zoState).
		SetZoStateSignerAccount(zoStateSigner).
		SetZoCacheAccount(zoCache).
		SetZoControlAccount(zoControl).
		SetZoVaultAccount(zoVault).
		SetTokenProgramAccount(tokenProgram)
}
