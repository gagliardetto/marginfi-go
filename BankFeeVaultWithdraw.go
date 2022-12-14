// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marginfi

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// BankFeeVaultWithdraw is the `bankFeeVaultWithdraw` instruction.
type BankFeeVaultWithdraw struct {
	Amount *uint64

	// [0] = [WRITE] marginfiGroup
	//
	// [1] = [SIGNER] admin
	//
	// [2] = [WRITE] bankFeeVault
	//
	// [3] = [] bankFeeVaultAuthority
	//
	// [4] = [WRITE] recipientTokenAccount
	//
	// [5] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewBankFeeVaultWithdrawInstructionBuilder creates a new `BankFeeVaultWithdraw` instruction builder.
func NewBankFeeVaultWithdrawInstructionBuilder() *BankFeeVaultWithdraw {
	nd := &BankFeeVaultWithdraw{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
func (inst *BankFeeVaultWithdraw) SetAmount(amount uint64) *BankFeeVaultWithdraw {
	inst.Amount = &amount
	return inst
}

// SetMarginfiGroupAccount sets the "marginfiGroup" account.
func (inst *BankFeeVaultWithdraw) SetMarginfiGroupAccount(marginfiGroup ag_solanago.PublicKey) *BankFeeVaultWithdraw {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(marginfiGroup).WRITE()
	return inst
}

// GetMarginfiGroupAccount gets the "marginfiGroup" account.
func (inst *BankFeeVaultWithdraw) GetMarginfiGroupAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *BankFeeVaultWithdraw) SetAdminAccount(admin ag_solanago.PublicKey) *BankFeeVaultWithdraw {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *BankFeeVaultWithdraw) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetBankFeeVaultAccount sets the "bankFeeVault" account.
func (inst *BankFeeVaultWithdraw) SetBankFeeVaultAccount(bankFeeVault ag_solanago.PublicKey) *BankFeeVaultWithdraw {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(bankFeeVault).WRITE()
	return inst
}

// GetBankFeeVaultAccount gets the "bankFeeVault" account.
func (inst *BankFeeVaultWithdraw) GetBankFeeVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetBankFeeVaultAuthorityAccount sets the "bankFeeVaultAuthority" account.
func (inst *BankFeeVaultWithdraw) SetBankFeeVaultAuthorityAccount(bankFeeVaultAuthority ag_solanago.PublicKey) *BankFeeVaultWithdraw {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(bankFeeVaultAuthority)
	return inst
}

// GetBankFeeVaultAuthorityAccount gets the "bankFeeVaultAuthority" account.
func (inst *BankFeeVaultWithdraw) GetBankFeeVaultAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetRecipientTokenAccount sets the "recipientTokenAccount" account.
func (inst *BankFeeVaultWithdraw) SetRecipientTokenAccount(recipientTokenAccount ag_solanago.PublicKey) *BankFeeVaultWithdraw {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(recipientTokenAccount).WRITE()
	return inst
}

// GetRecipientTokenAccount gets the "recipientTokenAccount" account.
func (inst *BankFeeVaultWithdraw) GetRecipientTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *BankFeeVaultWithdraw) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *BankFeeVaultWithdraw {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *BankFeeVaultWithdraw) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst BankFeeVaultWithdraw) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_BankFeeVaultWithdraw,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst BankFeeVaultWithdraw) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *BankFeeVaultWithdraw) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.MarginfiGroup is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.BankFeeVault is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.BankFeeVaultAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.RecipientTokenAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *BankFeeVaultWithdraw) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("BankFeeVaultWithdraw")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        marginfiGroup", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                admin", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         bankFeeVault", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("bankFeeVaultAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       recipientToken", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("         tokenProgram", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj BankFeeVaultWithdraw) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *BankFeeVaultWithdraw) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewBankFeeVaultWithdrawInstruction declares a new BankFeeVaultWithdraw instruction with the provided parameters and accounts.
func NewBankFeeVaultWithdrawInstruction(
	// Parameters:
	amount uint64,
	// Accounts:
	marginfiGroup ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	bankFeeVault ag_solanago.PublicKey,
	bankFeeVaultAuthority ag_solanago.PublicKey,
	recipientTokenAccount ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *BankFeeVaultWithdraw {
	return NewBankFeeVaultWithdrawInstructionBuilder().
		SetAmount(amount).
		SetMarginfiGroupAccount(marginfiGroup).
		SetAdminAccount(admin).
		SetBankFeeVaultAccount(bankFeeVault).
		SetBankFeeVaultAuthorityAccount(bankFeeVaultAuthority).
		SetRecipientTokenAccount(recipientTokenAccount).
		SetTokenProgramAccount(tokenProgram)
}
