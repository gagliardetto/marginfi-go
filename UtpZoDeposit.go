// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marginfi

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UtpZoDeposit is the `utpZoDeposit` instruction.
type UtpZoDeposit struct {
	Amount *uint64

	// [0] = [WRITE] marginfiAccount
	//
	// [1] = [WRITE] marginfiGroup
	//
	// [2] = [WRITE, SIGNER] signer
	//
	// [3] = [WRITE] marginCollateralVault
	//
	// [4] = [] bankAuthority
	//
	// [5] = [WRITE] tempCollateralAccount
	//
	// [6] = [] utpAuthority
	//
	// [7] = [] zoProgram
	//
	// [8] = [] zoState
	//
	// [9] = [] zoStateSigner
	//
	// [10] = [WRITE] zoCache
	//
	// [11] = [WRITE] zoMargin
	//
	// [12] = [WRITE] zoVault
	//
	// [13] = [] rent
	//
	// [14] = [] tokenProgram
	//
	// [15] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUtpZoDepositInstructionBuilder creates a new `UtpZoDeposit` instruction builder.
func NewUtpZoDepositInstructionBuilder() *UtpZoDeposit {
	nd := &UtpZoDeposit{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 16),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
func (inst *UtpZoDeposit) SetAmount(amount uint64) *UtpZoDeposit {
	inst.Amount = &amount
	return inst
}

// SetMarginfiAccount sets the "marginfiAccount" account.
func (inst *UtpZoDeposit) SetMarginfiAccount(marginfiAccount ag_solanago.PublicKey) *UtpZoDeposit {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(marginfiAccount).WRITE()
	return inst
}

// GetMarginfiAccount gets the "marginfiAccount" account.
func (inst *UtpZoDeposit) GetMarginfiAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMarginfiGroupAccount sets the "marginfiGroup" account.
func (inst *UtpZoDeposit) SetMarginfiGroupAccount(marginfiGroup ag_solanago.PublicKey) *UtpZoDeposit {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(marginfiGroup).WRITE()
	return inst
}

// GetMarginfiGroupAccount gets the "marginfiGroup" account.
func (inst *UtpZoDeposit) GetMarginfiGroupAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSignerAccount sets the "signer" account.
func (inst *UtpZoDeposit) SetSignerAccount(signer ag_solanago.PublicKey) *UtpZoDeposit {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(signer).WRITE().SIGNER()
	return inst
}

// GetSignerAccount gets the "signer" account.
func (inst *UtpZoDeposit) GetSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMarginCollateralVaultAccount sets the "marginCollateralVault" account.
func (inst *UtpZoDeposit) SetMarginCollateralVaultAccount(marginCollateralVault ag_solanago.PublicKey) *UtpZoDeposit {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(marginCollateralVault).WRITE()
	return inst
}

// GetMarginCollateralVaultAccount gets the "marginCollateralVault" account.
func (inst *UtpZoDeposit) GetMarginCollateralVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetBankAuthorityAccount sets the "bankAuthority" account.
func (inst *UtpZoDeposit) SetBankAuthorityAccount(bankAuthority ag_solanago.PublicKey) *UtpZoDeposit {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(bankAuthority)
	return inst
}

// GetBankAuthorityAccount gets the "bankAuthority" account.
func (inst *UtpZoDeposit) GetBankAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTempCollateralAccount sets the "tempCollateralAccount" account.
func (inst *UtpZoDeposit) SetTempCollateralAccount(tempCollateralAccount ag_solanago.PublicKey) *UtpZoDeposit {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tempCollateralAccount).WRITE()
	return inst
}

// GetTempCollateralAccount gets the "tempCollateralAccount" account.
func (inst *UtpZoDeposit) GetTempCollateralAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetUtpAuthorityAccount sets the "utpAuthority" account.
func (inst *UtpZoDeposit) SetUtpAuthorityAccount(utpAuthority ag_solanago.PublicKey) *UtpZoDeposit {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(utpAuthority)
	return inst
}

// GetUtpAuthorityAccount gets the "utpAuthority" account.
func (inst *UtpZoDeposit) GetUtpAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetZoProgramAccount sets the "zoProgram" account.
func (inst *UtpZoDeposit) SetZoProgramAccount(zoProgram ag_solanago.PublicKey) *UtpZoDeposit {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(zoProgram)
	return inst
}

// GetZoProgramAccount gets the "zoProgram" account.
func (inst *UtpZoDeposit) GetZoProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetZoStateAccount sets the "zoState" account.
func (inst *UtpZoDeposit) SetZoStateAccount(zoState ag_solanago.PublicKey) *UtpZoDeposit {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(zoState)
	return inst
}

// GetZoStateAccount gets the "zoState" account.
func (inst *UtpZoDeposit) GetZoStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetZoStateSignerAccount sets the "zoStateSigner" account.
func (inst *UtpZoDeposit) SetZoStateSignerAccount(zoStateSigner ag_solanago.PublicKey) *UtpZoDeposit {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(zoStateSigner)
	return inst
}

// GetZoStateSignerAccount gets the "zoStateSigner" account.
func (inst *UtpZoDeposit) GetZoStateSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetZoCacheAccount sets the "zoCache" account.
func (inst *UtpZoDeposit) SetZoCacheAccount(zoCache ag_solanago.PublicKey) *UtpZoDeposit {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(zoCache).WRITE()
	return inst
}

// GetZoCacheAccount gets the "zoCache" account.
func (inst *UtpZoDeposit) GetZoCacheAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetZoMarginAccount sets the "zoMargin" account.
func (inst *UtpZoDeposit) SetZoMarginAccount(zoMargin ag_solanago.PublicKey) *UtpZoDeposit {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(zoMargin).WRITE()
	return inst
}

// GetZoMarginAccount gets the "zoMargin" account.
func (inst *UtpZoDeposit) GetZoMarginAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetZoVaultAccount sets the "zoVault" account.
func (inst *UtpZoDeposit) SetZoVaultAccount(zoVault ag_solanago.PublicKey) *UtpZoDeposit {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(zoVault).WRITE()
	return inst
}

// GetZoVaultAccount gets the "zoVault" account.
func (inst *UtpZoDeposit) GetZoVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetRentAccount sets the "rent" account.
func (inst *UtpZoDeposit) SetRentAccount(rent ag_solanago.PublicKey) *UtpZoDeposit {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *UtpZoDeposit) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *UtpZoDeposit) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *UtpZoDeposit {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *UtpZoDeposit) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *UtpZoDeposit) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UtpZoDeposit {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *UtpZoDeposit) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

func (inst UtpZoDeposit) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UtpZoDeposit,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UtpZoDeposit) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UtpZoDeposit) Validate() error {
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
			return errors.New("accounts.BankAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TempCollateralAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.UtpAuthority is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.ZoProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.ZoState is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.ZoStateSigner is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.ZoCache is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.ZoMargin is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.ZoVault is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *UtpZoDeposit) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UtpZoDeposit")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=16]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             marginfi", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        marginfiGroup", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("               signer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("marginCollateralVault", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        bankAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("       tempCollateral", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("         utpAuthority", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("            zoProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("              zoState", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("        zoStateSigner", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("              zoCache", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("             zoMargin", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("              zoVault", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("                 rent", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("         tokenProgram", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("        systemProgram", inst.AccountMetaSlice.Get(15)))
					})
				})
		})
}

func (obj UtpZoDeposit) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UtpZoDeposit) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewUtpZoDepositInstruction declares a new UtpZoDeposit instruction with the provided parameters and accounts.
func NewUtpZoDepositInstruction(
	// Parameters:
	amount uint64,
	// Accounts:
	marginfiAccount ag_solanago.PublicKey,
	marginfiGroup ag_solanago.PublicKey,
	signer ag_solanago.PublicKey,
	marginCollateralVault ag_solanago.PublicKey,
	bankAuthority ag_solanago.PublicKey,
	tempCollateralAccount ag_solanago.PublicKey,
	utpAuthority ag_solanago.PublicKey,
	zoProgram ag_solanago.PublicKey,
	zoState ag_solanago.PublicKey,
	zoStateSigner ag_solanago.PublicKey,
	zoCache ag_solanago.PublicKey,
	zoMargin ag_solanago.PublicKey,
	zoVault ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *UtpZoDeposit {
	return NewUtpZoDepositInstructionBuilder().
		SetAmount(amount).
		SetMarginfiAccount(marginfiAccount).
		SetMarginfiGroupAccount(marginfiGroup).
		SetSignerAccount(signer).
		SetMarginCollateralVaultAccount(marginCollateralVault).
		SetBankAuthorityAccount(bankAuthority).
		SetTempCollateralAccount(tempCollateralAccount).
		SetUtpAuthorityAccount(utpAuthority).
		SetZoProgramAccount(zoProgram).
		SetZoStateAccount(zoState).
		SetZoStateSignerAccount(zoStateSigner).
		SetZoCacheAccount(zoCache).
		SetZoMarginAccount(zoMargin).
		SetZoVaultAccount(zoVault).
		SetRentAccount(rent).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram)
}