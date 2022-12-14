// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marginfi

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ConfigureMarginfiGroup is the `configureMarginfiGroup` instruction.
type ConfigureMarginfiGroup struct {
	ConfigArg *GroupConfig

	// [0] = [WRITE] marginfiGroup
	//
	// [1] = [SIGNER] admin
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewConfigureMarginfiGroupInstructionBuilder creates a new `ConfigureMarginfiGroup` instruction builder.
func NewConfigureMarginfiGroupInstructionBuilder() *ConfigureMarginfiGroup {
	nd := &ConfigureMarginfiGroup{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetConfigArg sets the "configArg" parameter.
func (inst *ConfigureMarginfiGroup) SetConfigArg(configArg GroupConfig) *ConfigureMarginfiGroup {
	inst.ConfigArg = &configArg
	return inst
}

// SetMarginfiGroupAccount sets the "marginfiGroup" account.
func (inst *ConfigureMarginfiGroup) SetMarginfiGroupAccount(marginfiGroup ag_solanago.PublicKey) *ConfigureMarginfiGroup {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(marginfiGroup).WRITE()
	return inst
}

// GetMarginfiGroupAccount gets the "marginfiGroup" account.
func (inst *ConfigureMarginfiGroup) GetMarginfiGroupAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *ConfigureMarginfiGroup) SetAdminAccount(admin ag_solanago.PublicKey) *ConfigureMarginfiGroup {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *ConfigureMarginfiGroup) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst ConfigureMarginfiGroup) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ConfigureMarginfiGroup,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ConfigureMarginfiGroup) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ConfigureMarginfiGroup) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ConfigArg == nil {
			return errors.New("ConfigArg parameter is not set")
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
	}
	return nil
}

func (inst *ConfigureMarginfiGroup) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ConfigureMarginfiGroup")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("ConfigArg", *inst.ConfigArg))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("marginfiGroup", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        admin", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj ConfigureMarginfiGroup) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ConfigArg` param:
	err = encoder.Encode(obj.ConfigArg)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ConfigureMarginfiGroup) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ConfigArg`:
	err = decoder.Decode(&obj.ConfigArg)
	if err != nil {
		return err
	}
	return nil
}

// NewConfigureMarginfiGroupInstruction declares a new ConfigureMarginfiGroup instruction with the provided parameters and accounts.
func NewConfigureMarginfiGroupInstruction(
	// Parameters:
	configArg GroupConfig,
	// Accounts:
	marginfiGroup ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *ConfigureMarginfiGroup {
	return NewConfigureMarginfiGroupInstructionBuilder().
		SetConfigArg(configArg).
		SetMarginfiGroupAccount(marginfiGroup).
		SetAdminAccount(admin)
}
