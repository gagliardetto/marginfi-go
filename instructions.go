// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marginfi

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "Marginfi"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_InitMarginfiGroup = ag_binary.TypeID([8]byte{71, 42, 56, 194, 9, 180, 174, 163})

	Instruction_ConfigureMarginfiGroup = ag_binary.TypeID([8]byte{13, 26, 103, 157, 5, 25, 169, 123})

	Instruction_BankFeeVaultWithdraw = ag_binary.TypeID([8]byte{118, 24, 133, 25, 27, 115, 172, 176})

	Instruction_InitMarginfiAccount = ag_binary.TypeID([8]byte{225, 212, 68, 233, 62, 96, 242, 103})

	Instruction_BankInsuranceVaultWithdraw = ag_binary.TypeID([8]byte{34, 152, 185, 160, 88, 153, 26, 9})

	Instruction_MarginDepositCollateral = ag_binary.TypeID([8]byte{144, 208, 127, 5, 52, 211, 251, 74})

	Instruction_MarginWithdrawCollateral = ag_binary.TypeID([8]byte{76, 13, 209, 139, 220, 179, 255, 155})

	Instruction_Liquidate = ag_binary.TypeID([8]byte{223, 179, 226, 125, 48, 46, 39, 74})

	Instruction_DeactivateUtp = ag_binary.TypeID([8]byte{94, 212, 1, 215, 47, 46, 72, 47})

	Instruction_HandleBankruptcy = ag_binary.TypeID([8]byte{108, 115, 137, 210, 212, 178, 213, 29})

	Instruction_UpdateInterestAccumulator = ag_binary.TypeID([8]byte{224, 254, 233, 116, 37, 246, 218, 2})

	Instruction_UtpMangoActivate = ag_binary.TypeID([8]byte{253, 151, 120, 199, 12, 49, 243, 66})

	Instruction_UtpMangoDeposit = ag_binary.TypeID([8]byte{88, 40, 72, 72, 119, 127, 39, 229})

	Instruction_UtpMangoWithdraw = ag_binary.TypeID([8]byte{12, 23, 176, 92, 100, 155, 156, 138})

	Instruction_UtpMangoUsePlacePerpOrder = ag_binary.TypeID([8]byte{148, 18, 221, 203, 214, 255, 118, 52})

	Instruction_UtpMangoUseCancelPerpOrder = ag_binary.TypeID([8]byte{85, 99, 189, 41, 221, 46, 143, 229})

	Instruction_UtpZoActivate = ag_binary.TypeID([8]byte{144, 144, 99, 249, 96, 225, 84, 236})

	Instruction_UtpZoDeposit = ag_binary.TypeID([8]byte{138, 26, 171, 14, 16, 199, 195, 124})

	Instruction_UtpZoWithdraw = ag_binary.TypeID([8]byte{253, 203, 112, 49, 195, 165, 163, 209})

	Instruction_UtpZoCreatePerpOpenOrders = ag_binary.TypeID([8]byte{171, 232, 88, 85, 71, 180, 29, 92})

	Instruction_UtpZoPlacePerpOrder = ag_binary.TypeID([8]byte{3, 8, 52, 20, 9, 169, 127, 186})

	Instruction_UtpZoCancelPerpOrder = ag_binary.TypeID([8]byte{242, 14, 34, 160, 206, 108, 108, 3})

	Instruction_UtpZoSettleFunds = ag_binary.TypeID([8]byte{15, 126, 149, 93, 252, 92, 75, 246})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_InitMarginfiGroup:
		return "InitMarginfiGroup"
	case Instruction_ConfigureMarginfiGroup:
		return "ConfigureMarginfiGroup"
	case Instruction_BankFeeVaultWithdraw:
		return "BankFeeVaultWithdraw"
	case Instruction_InitMarginfiAccount:
		return "InitMarginfiAccount"
	case Instruction_BankInsuranceVaultWithdraw:
		return "BankInsuranceVaultWithdraw"
	case Instruction_MarginDepositCollateral:
		return "MarginDepositCollateral"
	case Instruction_MarginWithdrawCollateral:
		return "MarginWithdrawCollateral"
	case Instruction_Liquidate:
		return "Liquidate"
	case Instruction_DeactivateUtp:
		return "DeactivateUtp"
	case Instruction_HandleBankruptcy:
		return "HandleBankruptcy"
	case Instruction_UpdateInterestAccumulator:
		return "UpdateInterestAccumulator"
	case Instruction_UtpMangoActivate:
		return "UtpMangoActivate"
	case Instruction_UtpMangoDeposit:
		return "UtpMangoDeposit"
	case Instruction_UtpMangoWithdraw:
		return "UtpMangoWithdraw"
	case Instruction_UtpMangoUsePlacePerpOrder:
		return "UtpMangoUsePlacePerpOrder"
	case Instruction_UtpMangoUseCancelPerpOrder:
		return "UtpMangoUseCancelPerpOrder"
	case Instruction_UtpZoActivate:
		return "UtpZoActivate"
	case Instruction_UtpZoDeposit:
		return "UtpZoDeposit"
	case Instruction_UtpZoWithdraw:
		return "UtpZoWithdraw"
	case Instruction_UtpZoCreatePerpOpenOrders:
		return "UtpZoCreatePerpOpenOrders"
	case Instruction_UtpZoPlacePerpOrder:
		return "UtpZoPlacePerpOrder"
	case Instruction_UtpZoCancelPerpOrder:
		return "UtpZoCancelPerpOrder"
	case Instruction_UtpZoSettleFunds:
		return "UtpZoSettleFunds"
	default:
		return ""
	}
}

type Instruction struct {
	ag_binary.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent ag_treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = ag_binary.NewVariantDefinition(
	ag_binary.AnchorTypeIDEncoding,
	[]ag_binary.VariantType{
		{
			"init_marginfi_group", (*InitMarginfiGroup)(nil),
		},
		{
			"configure_marginfi_group", (*ConfigureMarginfiGroup)(nil),
		},
		{
			"bank_fee_vault_withdraw", (*BankFeeVaultWithdraw)(nil),
		},
		{
			"init_marginfi_account", (*InitMarginfiAccount)(nil),
		},
		{
			"bank_insurance_vault_withdraw", (*BankInsuranceVaultWithdraw)(nil),
		},
		{
			"margin_deposit_collateral", (*MarginDepositCollateral)(nil),
		},
		{
			"margin_withdraw_collateral", (*MarginWithdrawCollateral)(nil),
		},
		{
			"liquidate", (*Liquidate)(nil),
		},
		{
			"deactivate_utp", (*DeactivateUtp)(nil),
		},
		{
			"handle_bankruptcy", (*HandleBankruptcy)(nil),
		},
		{
			"update_interest_accumulator", (*UpdateInterestAccumulator)(nil),
		},
		{
			"utp_mango_activate", (*UtpMangoActivate)(nil),
		},
		{
			"utp_mango_deposit", (*UtpMangoDeposit)(nil),
		},
		{
			"utp_mango_withdraw", (*UtpMangoWithdraw)(nil),
		},
		{
			"utp_mango_use_place_perp_order", (*UtpMangoUsePlacePerpOrder)(nil),
		},
		{
			"utp_mango_use_cancel_perp_order", (*UtpMangoUseCancelPerpOrder)(nil),
		},
		{
			"utp_zo_activate", (*UtpZoActivate)(nil),
		},
		{
			"utp_zo_deposit", (*UtpZoDeposit)(nil),
		},
		{
			"utp_zo_withdraw", (*UtpZoWithdraw)(nil),
		},
		{
			"utp_zo_create_perp_open_orders", (*UtpZoCreatePerpOpenOrders)(nil),
		},
		{
			"utp_zo_place_perp_order", (*UtpZoPlacePerpOrder)(nil),
		},
		{
			"utp_zo_cancel_perp_order", (*UtpZoCancelPerpOrder)(nil),
		},
		{
			"utp_zo_settle_funds", (*UtpZoSettleFunds)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() ag_solanago.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*ag_solanago.AccountMeta) {
	return inst.Impl.(ag_solanago.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ag_binary.NewBorshEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *ag_text.Encoder, option *ag_text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst *Instruction) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := ag_binary.NewBorshDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(ag_solanago.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}