// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marginfi

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UtpMangoDeposit is the `utpMangoDeposit` instruction.
type UtpMangoDeposit struct {
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
	// [6] = [] mangoAuthority
	//
	// [7] = [WRITE] mangoAccount
	//
	// [8] = [] mangoProgram
	//
	// [9] = [] mangoGroup
	//
	// [10] = [WRITE] mangoCache
	//
	// [11] = [WRITE] mangoRootBank
	//
	// [12] = [WRITE] mangoNodeBank
	//
	// [13] = [WRITE] mangoVault
	//
	// [14] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUtpMangoDepositInstructionBuilder creates a new `UtpMangoDeposit` instruction builder.
func NewUtpMangoDepositInstructionBuilder() *UtpMangoDeposit {
	nd := &UtpMangoDeposit{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 15),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
func (inst *UtpMangoDeposit) SetAmount(amount uint64) *UtpMangoDeposit {
	inst.Amount = &amount
	return inst
}

// SetMarginfiAccount sets the "marginfiAccount" account.
func (inst *UtpMangoDeposit) SetMarginfiAccount(marginfiAccount ag_solanago.PublicKey) *UtpMangoDeposit {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(marginfiAccount).WRITE()
	return inst
}

// GetMarginfiAccount gets the "marginfiAccount" account.
func (inst *UtpMangoDeposit) GetMarginfiAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMarginfiGroupAccount sets the "marginfiGroup" account.
func (inst *UtpMangoDeposit) SetMarginfiGroupAccount(marginfiGroup ag_solanago.PublicKey) *UtpMangoDeposit {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(marginfiGroup).WRITE()
	return inst
}

// GetMarginfiGroupAccount gets the "marginfiGroup" account.
func (inst *UtpMangoDeposit) GetMarginfiGroupAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSignerAccount sets the "signer" account.
func (inst *UtpMangoDeposit) SetSignerAccount(signer ag_solanago.PublicKey) *UtpMangoDeposit {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(signer).WRITE().SIGNER()
	return inst
}

// GetSignerAccount gets the "signer" account.
func (inst *UtpMangoDeposit) GetSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMarginCollateralVaultAccount sets the "marginCollateralVault" account.
func (inst *UtpMangoDeposit) SetMarginCollateralVaultAccount(marginCollateralVault ag_solanago.PublicKey) *UtpMangoDeposit {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(marginCollateralVault).WRITE()
	return inst
}

// GetMarginCollateralVaultAccount gets the "marginCollateralVault" account.
func (inst *UtpMangoDeposit) GetMarginCollateralVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetBankAuthorityAccount sets the "bankAuthority" account.
func (inst *UtpMangoDeposit) SetBankAuthorityAccount(bankAuthority ag_solanago.PublicKey) *UtpMangoDeposit {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(bankAuthority)
	return inst
}

// GetBankAuthorityAccount gets the "bankAuthority" account.
func (inst *UtpMangoDeposit) GetBankAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTempCollateralAccount sets the "tempCollateralAccount" account.
func (inst *UtpMangoDeposit) SetTempCollateralAccount(tempCollateralAccount ag_solanago.PublicKey) *UtpMangoDeposit {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tempCollateralAccount).WRITE()
	return inst
}

// GetTempCollateralAccount gets the "tempCollateralAccount" account.
func (inst *UtpMangoDeposit) GetTempCollateralAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMangoAuthorityAccount sets the "mangoAuthority" account.
func (inst *UtpMangoDeposit) SetMangoAuthorityAccount(mangoAuthority ag_solanago.PublicKey) *UtpMangoDeposit {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(mangoAuthority)
	return inst
}

// GetMangoAuthorityAccount gets the "mangoAuthority" account.
func (inst *UtpMangoDeposit) GetMangoAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetMangoAccount sets the "mangoAccount" account.
func (inst *UtpMangoDeposit) SetMangoAccount(mangoAccount ag_solanago.PublicKey) *UtpMangoDeposit {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(mangoAccount).WRITE()
	return inst
}

// GetMangoAccount gets the "mangoAccount" account.
func (inst *UtpMangoDeposit) GetMangoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetMangoProgramAccount sets the "mangoProgram" account.
func (inst *UtpMangoDeposit) SetMangoProgramAccount(mangoProgram ag_solanago.PublicKey) *UtpMangoDeposit {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(mangoProgram)
	return inst
}

// GetMangoProgramAccount gets the "mangoProgram" account.
func (inst *UtpMangoDeposit) GetMangoProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetMangoGroupAccount sets the "mangoGroup" account.
func (inst *UtpMangoDeposit) SetMangoGroupAccount(mangoGroup ag_solanago.PublicKey) *UtpMangoDeposit {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(mangoGroup)
	return inst
}

// GetMangoGroupAccount gets the "mangoGroup" account.
func (inst *UtpMangoDeposit) GetMangoGroupAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetMangoCacheAccount sets the "mangoCache" account.
func (inst *UtpMangoDeposit) SetMangoCacheAccount(mangoCache ag_solanago.PublicKey) *UtpMangoDeposit {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(mangoCache).WRITE()
	return inst
}

// GetMangoCacheAccount gets the "mangoCache" account.
func (inst *UtpMangoDeposit) GetMangoCacheAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetMangoRootBankAccount sets the "mangoRootBank" account.
func (inst *UtpMangoDeposit) SetMangoRootBankAccount(mangoRootBank ag_solanago.PublicKey) *UtpMangoDeposit {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(mangoRootBank).WRITE()
	return inst
}

// GetMangoRootBankAccount gets the "mangoRootBank" account.
func (inst *UtpMangoDeposit) GetMangoRootBankAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetMangoNodeBankAccount sets the "mangoNodeBank" account.
func (inst *UtpMangoDeposit) SetMangoNodeBankAccount(mangoNodeBank ag_solanago.PublicKey) *UtpMangoDeposit {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(mangoNodeBank).WRITE()
	return inst
}

// GetMangoNodeBankAccount gets the "mangoNodeBank" account.
func (inst *UtpMangoDeposit) GetMangoNodeBankAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetMangoVaultAccount sets the "mangoVault" account.
func (inst *UtpMangoDeposit) SetMangoVaultAccount(mangoVault ag_solanago.PublicKey) *UtpMangoDeposit {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(mangoVault).WRITE()
	return inst
}

// GetMangoVaultAccount gets the "mangoVault" account.
func (inst *UtpMangoDeposit) GetMangoVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *UtpMangoDeposit) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *UtpMangoDeposit {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *UtpMangoDeposit) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

func (inst UtpMangoDeposit) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UtpMangoDeposit,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UtpMangoDeposit) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UtpMangoDeposit) Validate() error {
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
			return errors.New("accounts.MangoAuthority is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.MangoAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.MangoProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.MangoGroup is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.MangoCache is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.MangoRootBank is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.MangoNodeBank is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.MangoVault is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *UtpMangoDeposit) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UtpMangoDeposit")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=15]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             marginfi", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        marginfiGroup", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("               signer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("marginCollateralVault", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        bankAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("       tempCollateral", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("       mangoAuthority", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                mango", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("         mangoProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("           mangoGroup", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("           mangoCache", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("        mangoRootBank", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("        mangoNodeBank", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("           mangoVault", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("         tokenProgram", inst.AccountMetaSlice.Get(14)))
					})
				})
		})
}

func (obj UtpMangoDeposit) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UtpMangoDeposit) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewUtpMangoDepositInstruction declares a new UtpMangoDeposit instruction with the provided parameters and accounts.
func NewUtpMangoDepositInstruction(
	// Parameters:
	amount uint64,
	// Accounts:
	marginfiAccount ag_solanago.PublicKey,
	marginfiGroup ag_solanago.PublicKey,
	signer ag_solanago.PublicKey,
	marginCollateralVault ag_solanago.PublicKey,
	bankAuthority ag_solanago.PublicKey,
	tempCollateralAccount ag_solanago.PublicKey,
	mangoAuthority ag_solanago.PublicKey,
	mangoAccount ag_solanago.PublicKey,
	mangoProgram ag_solanago.PublicKey,
	mangoGroup ag_solanago.PublicKey,
	mangoCache ag_solanago.PublicKey,
	mangoRootBank ag_solanago.PublicKey,
	mangoNodeBank ag_solanago.PublicKey,
	mangoVault ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *UtpMangoDeposit {
	return NewUtpMangoDepositInstructionBuilder().
		SetAmount(amount).
		SetMarginfiAccount(marginfiAccount).
		SetMarginfiGroupAccount(marginfiGroup).
		SetSignerAccount(signer).
		SetMarginCollateralVaultAccount(marginCollateralVault).
		SetBankAuthorityAccount(bankAuthority).
		SetTempCollateralAccount(tempCollateralAccount).
		SetMangoAuthorityAccount(mangoAuthority).
		SetMangoAccount(mangoAccount).
		SetMangoProgramAccount(mangoProgram).
		SetMangoGroupAccount(mangoGroup).
		SetMangoCacheAccount(mangoCache).
		SetMangoRootBankAccount(mangoRootBank).
		SetMangoNodeBankAccount(mangoNodeBank).
		SetMangoVaultAccount(mangoVault).
		SetTokenProgramAccount(tokenProgram)
}