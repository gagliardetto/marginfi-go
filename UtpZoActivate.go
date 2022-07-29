// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marginfi

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UtpZoActivate is the `utpZoActivate` instruction.
type UtpZoActivate struct {
	AuthoritySeed *ag_solanago.PublicKey
	AuthorityBump *uint8
	ZoMarginNonce *uint8

	// [0] = [WRITE] marginfiAccount
	//
	// [1] = [] marginfiGroup
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] utpAuthority
	//
	// [4] = [] zoProgram
	//
	// [5] = [] zoState
	//
	// [6] = [WRITE] zoMargin
	//
	// [7] = [WRITE] zoControl
	//
	// [8] = [] rent
	//
	// [9] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUtpZoActivateInstructionBuilder creates a new `UtpZoActivate` instruction builder.
func NewUtpZoActivateInstructionBuilder() *UtpZoActivate {
	nd := &UtpZoActivate{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

// SetAuthoritySeed sets the "authoritySeed" parameter.
func (inst *UtpZoActivate) SetAuthoritySeed(authoritySeed ag_solanago.PublicKey) *UtpZoActivate {
	inst.AuthoritySeed = &authoritySeed
	return inst
}

// SetAuthorityBump sets the "authorityBump" parameter.
func (inst *UtpZoActivate) SetAuthorityBump(authorityBump uint8) *UtpZoActivate {
	inst.AuthorityBump = &authorityBump
	return inst
}

// SetZoMarginNonce sets the "zoMarginNonce" parameter.
func (inst *UtpZoActivate) SetZoMarginNonce(zoMarginNonce uint8) *UtpZoActivate {
	inst.ZoMarginNonce = &zoMarginNonce
	return inst
}

// SetMarginfiAccount sets the "marginfiAccount" account.
func (inst *UtpZoActivate) SetMarginfiAccount(marginfiAccount ag_solanago.PublicKey) *UtpZoActivate {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(marginfiAccount).WRITE()
	return inst
}

// GetMarginfiAccount gets the "marginfiAccount" account.
func (inst *UtpZoActivate) GetMarginfiAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMarginfiGroupAccount sets the "marginfiGroup" account.
func (inst *UtpZoActivate) SetMarginfiGroupAccount(marginfiGroup ag_solanago.PublicKey) *UtpZoActivate {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(marginfiGroup)
	return inst
}

// GetMarginfiGroupAccount gets the "marginfiGroup" account.
func (inst *UtpZoActivate) GetMarginfiGroupAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *UtpZoActivate) SetAuthorityAccount(authority ag_solanago.PublicKey) *UtpZoActivate {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *UtpZoActivate) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetUtpAuthorityAccount sets the "utpAuthority" account.
func (inst *UtpZoActivate) SetUtpAuthorityAccount(utpAuthority ag_solanago.PublicKey) *UtpZoActivate {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(utpAuthority)
	return inst
}

// GetUtpAuthorityAccount gets the "utpAuthority" account.
func (inst *UtpZoActivate) GetUtpAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetZoProgramAccount sets the "zoProgram" account.
func (inst *UtpZoActivate) SetZoProgramAccount(zoProgram ag_solanago.PublicKey) *UtpZoActivate {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(zoProgram)
	return inst
}

// GetZoProgramAccount gets the "zoProgram" account.
func (inst *UtpZoActivate) GetZoProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetZoStateAccount sets the "zoState" account.
func (inst *UtpZoActivate) SetZoStateAccount(zoState ag_solanago.PublicKey) *UtpZoActivate {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(zoState)
	return inst
}

// GetZoStateAccount gets the "zoState" account.
func (inst *UtpZoActivate) GetZoStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetZoMarginAccount sets the "zoMargin" account.
func (inst *UtpZoActivate) SetZoMarginAccount(zoMargin ag_solanago.PublicKey) *UtpZoActivate {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(zoMargin).WRITE()
	return inst
}

// GetZoMarginAccount gets the "zoMargin" account.
func (inst *UtpZoActivate) GetZoMarginAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetZoControlAccount sets the "zoControl" account.
func (inst *UtpZoActivate) SetZoControlAccount(zoControl ag_solanago.PublicKey) *UtpZoActivate {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(zoControl).WRITE()
	return inst
}

// GetZoControlAccount gets the "zoControl" account.
func (inst *UtpZoActivate) GetZoControlAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetRentAccount sets the "rent" account.
func (inst *UtpZoActivate) SetRentAccount(rent ag_solanago.PublicKey) *UtpZoActivate {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *UtpZoActivate) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *UtpZoActivate) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UtpZoActivate {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *UtpZoActivate) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

func (inst UtpZoActivate) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UtpZoActivate,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UtpZoActivate) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UtpZoActivate) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.AuthoritySeed == nil {
			return errors.New("AuthoritySeed parameter is not set")
		}
		if inst.AuthorityBump == nil {
			return errors.New("AuthorityBump parameter is not set")
		}
		if inst.ZoMarginNonce == nil {
			return errors.New("ZoMarginNonce parameter is not set")
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
			return errors.New("accounts.UtpAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.ZoProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ZoState is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.ZoMargin is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.ZoControl is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *UtpZoActivate) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UtpZoActivate")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("AuthoritySeed", *inst.AuthoritySeed))
						paramsBranch.Child(ag_format.Param("AuthorityBump", *inst.AuthorityBump))
						paramsBranch.Child(ag_format.Param("ZoMarginNonce", *inst.ZoMarginNonce))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     marginfi", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("marginfiGroup", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta(" utpAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("    zoProgram", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("      zoState", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("     zoMargin", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("    zoControl", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("         rent", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(9)))
					})
				})
		})
}

func (obj UtpZoActivate) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `AuthoritySeed` param:
	err = encoder.Encode(obj.AuthoritySeed)
	if err != nil {
		return err
	}
	// Serialize `AuthorityBump` param:
	err = encoder.Encode(obj.AuthorityBump)
	if err != nil {
		return err
	}
	// Serialize `ZoMarginNonce` param:
	err = encoder.Encode(obj.ZoMarginNonce)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UtpZoActivate) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `AuthoritySeed`:
	err = decoder.Decode(&obj.AuthoritySeed)
	if err != nil {
		return err
	}
	// Deserialize `AuthorityBump`:
	err = decoder.Decode(&obj.AuthorityBump)
	if err != nil {
		return err
	}
	// Deserialize `ZoMarginNonce`:
	err = decoder.Decode(&obj.ZoMarginNonce)
	if err != nil {
		return err
	}
	return nil
}

// NewUtpZoActivateInstruction declares a new UtpZoActivate instruction with the provided parameters and accounts.
func NewUtpZoActivateInstruction(
	// Parameters:
	authoritySeed ag_solanago.PublicKey,
	authorityBump uint8,
	zoMarginNonce uint8,
	// Accounts:
	marginfiAccount ag_solanago.PublicKey,
	marginfiGroup ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	utpAuthority ag_solanago.PublicKey,
	zoProgram ag_solanago.PublicKey,
	zoState ag_solanago.PublicKey,
	zoMargin ag_solanago.PublicKey,
	zoControl ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *UtpZoActivate {
	return NewUtpZoActivateInstructionBuilder().
		SetAuthoritySeed(authoritySeed).
		SetAuthorityBump(authorityBump).
		SetZoMarginNonce(zoMarginNonce).
		SetMarginfiAccount(marginfiAccount).
		SetMarginfiGroupAccount(marginfiGroup).
		SetAuthorityAccount(authority).
		SetUtpAuthorityAccount(utpAuthority).
		SetZoProgramAccount(zoProgram).
		SetZoStateAccount(zoState).
		SetZoMarginAccount(zoMargin).
		SetZoControlAccount(zoControl).
		SetRentAccount(rent).
		SetSystemProgramAccount(systemProgram)
}
