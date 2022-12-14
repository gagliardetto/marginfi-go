// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marginfi

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// MarginDepositCollateral is the `marginDepositCollateral` instruction.
type MarginDepositCollateral struct {
	Amount *uint64

	// [0] = [WRITE] marginfiAccount
	//
	// [1] = [WRITE] marginfiGroup
	//
	// [2] = [SIGNER] signer
	//
	// [3] = [WRITE] fundingAccount
	//
	// [4] = [WRITE] tokenVault
	//
	// [5] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewMarginDepositCollateralInstructionBuilder creates a new `MarginDepositCollateral` instruction builder.
func NewMarginDepositCollateralInstructionBuilder() *MarginDepositCollateral {
	nd := &MarginDepositCollateral{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
func (inst *MarginDepositCollateral) SetAmount(amount uint64) *MarginDepositCollateral {
	inst.Amount = &amount
	return inst
}

// SetMarginfiAccount sets the "marginfiAccount" account.
func (inst *MarginDepositCollateral) SetMarginfiAccount(marginfiAccount ag_solanago.PublicKey) *MarginDepositCollateral {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(marginfiAccount).WRITE()
	return inst
}

// GetMarginfiAccount gets the "marginfiAccount" account.
func (inst *MarginDepositCollateral) GetMarginfiAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMarginfiGroupAccount sets the "marginfiGroup" account.
func (inst *MarginDepositCollateral) SetMarginfiGroupAccount(marginfiGroup ag_solanago.PublicKey) *MarginDepositCollateral {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(marginfiGroup).WRITE()
	return inst
}

// GetMarginfiGroupAccount gets the "marginfiGroup" account.
func (inst *MarginDepositCollateral) GetMarginfiGroupAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSignerAccount sets the "signer" account.
func (inst *MarginDepositCollateral) SetSignerAccount(signer ag_solanago.PublicKey) *MarginDepositCollateral {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(signer).SIGNER()
	return inst
}

// GetSignerAccount gets the "signer" account.
func (inst *MarginDepositCollateral) GetSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetFundingAccount sets the "fundingAccount" account.
func (inst *MarginDepositCollateral) SetFundingAccount(fundingAccount ag_solanago.PublicKey) *MarginDepositCollateral {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(fundingAccount).WRITE()
	return inst
}

// GetFundingAccount gets the "fundingAccount" account.
func (inst *MarginDepositCollateral) GetFundingAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenVaultAccount sets the "tokenVault" account.
func (inst *MarginDepositCollateral) SetTokenVaultAccount(tokenVault ag_solanago.PublicKey) *MarginDepositCollateral {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenVault).WRITE()
	return inst
}

// GetTokenVaultAccount gets the "tokenVault" account.
func (inst *MarginDepositCollateral) GetTokenVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *MarginDepositCollateral) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *MarginDepositCollateral {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *MarginDepositCollateral) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst MarginDepositCollateral) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_MarginDepositCollateral,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MarginDepositCollateral) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MarginDepositCollateral) Validate() error {
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
			return errors.New("accounts.FundingAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TokenVault is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *MarginDepositCollateral) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MarginDepositCollateral")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     marginfi", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("marginfiGroup", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("       signer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      funding", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("   tokenVault", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta(" tokenProgram", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj MarginDepositCollateral) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *MarginDepositCollateral) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewMarginDepositCollateralInstruction declares a new MarginDepositCollateral instruction with the provided parameters and accounts.
func NewMarginDepositCollateralInstruction(
	// Parameters:
	amount uint64,
	// Accounts:
	marginfiAccount ag_solanago.PublicKey,
	marginfiGroup ag_solanago.PublicKey,
	signer ag_solanago.PublicKey,
	fundingAccount ag_solanago.PublicKey,
	tokenVault ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *MarginDepositCollateral {
	return NewMarginDepositCollateralInstructionBuilder().
		SetAmount(amount).
		SetMarginfiAccount(marginfiAccount).
		SetMarginfiGroupAccount(marginfiGroup).
		SetSignerAccount(signer).
		SetFundingAccount(fundingAccount).
		SetTokenVaultAccount(tokenVault).
		SetTokenProgramAccount(tokenProgram)
}
