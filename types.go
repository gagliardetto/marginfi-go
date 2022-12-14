// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marginfi

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type UtpMangoPlacePerpOrderArgs struct {
	Side             MangoSide
	Price            int64
	MaxBaseQuantity  int64
	MaxQuoteQuantity int64
	ClientOrderId    uint64
	OrderType        MangoOrderType
	ReduceOnly       bool
	ExpiryTimestamp  *uint64 `bin:"optional"`
	Limit            uint8
	ExpiryType       MangoExpiryType
}

func (obj UtpMangoPlacePerpOrderArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Side` param:
	err = encoder.Encode(obj.Side)
	if err != nil {
		return err
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `MaxBaseQuantity` param:
	err = encoder.Encode(obj.MaxBaseQuantity)
	if err != nil {
		return err
	}
	// Serialize `MaxQuoteQuantity` param:
	err = encoder.Encode(obj.MaxQuoteQuantity)
	if err != nil {
		return err
	}
	// Serialize `ClientOrderId` param:
	err = encoder.Encode(obj.ClientOrderId)
	if err != nil {
		return err
	}
	// Serialize `OrderType` param:
	err = encoder.Encode(obj.OrderType)
	if err != nil {
		return err
	}
	// Serialize `ReduceOnly` param:
	err = encoder.Encode(obj.ReduceOnly)
	if err != nil {
		return err
	}
	// Serialize `ExpiryTimestamp` param (optional):
	{
		if obj.ExpiryTimestamp == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.ExpiryTimestamp)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Limit` param:
	err = encoder.Encode(obj.Limit)
	if err != nil {
		return err
	}
	// Serialize `ExpiryType` param:
	err = encoder.Encode(obj.ExpiryType)
	if err != nil {
		return err
	}
	return nil
}

func (obj *UtpMangoPlacePerpOrderArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Side`:
	err = decoder.Decode(&obj.Side)
	if err != nil {
		return err
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `MaxBaseQuantity`:
	err = decoder.Decode(&obj.MaxBaseQuantity)
	if err != nil {
		return err
	}
	// Deserialize `MaxQuoteQuantity`:
	err = decoder.Decode(&obj.MaxQuoteQuantity)
	if err != nil {
		return err
	}
	// Deserialize `ClientOrderId`:
	err = decoder.Decode(&obj.ClientOrderId)
	if err != nil {
		return err
	}
	// Deserialize `OrderType`:
	err = decoder.Decode(&obj.OrderType)
	if err != nil {
		return err
	}
	// Deserialize `ReduceOnly`:
	err = decoder.Decode(&obj.ReduceOnly)
	if err != nil {
		return err
	}
	// Deserialize `ExpiryTimestamp` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.ExpiryTimestamp)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Limit`:
	err = decoder.Decode(&obj.Limit)
	if err != nil {
		return err
	}
	// Deserialize `ExpiryType`:
	err = decoder.Decode(&obj.ExpiryType)
	if err != nil {
		return err
	}
	return nil
}

type UtpZoPlacePerpOrderIxArgs struct {
	IsLong           bool
	LimitPrice       uint64
	MaxBaseQuantity  uint64
	MaxQuoteQuantity uint64
	OrderType        OrderType
	Limit            uint16
	ClientId         uint64
}

func (obj UtpZoPlacePerpOrderIxArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `IsLong` param:
	err = encoder.Encode(obj.IsLong)
	if err != nil {
		return err
	}
	// Serialize `LimitPrice` param:
	err = encoder.Encode(obj.LimitPrice)
	if err != nil {
		return err
	}
	// Serialize `MaxBaseQuantity` param:
	err = encoder.Encode(obj.MaxBaseQuantity)
	if err != nil {
		return err
	}
	// Serialize `MaxQuoteQuantity` param:
	err = encoder.Encode(obj.MaxQuoteQuantity)
	if err != nil {
		return err
	}
	// Serialize `OrderType` param:
	err = encoder.Encode(obj.OrderType)
	if err != nil {
		return err
	}
	// Serialize `Limit` param:
	err = encoder.Encode(obj.Limit)
	if err != nil {
		return err
	}
	// Serialize `ClientId` param:
	err = encoder.Encode(obj.ClientId)
	if err != nil {
		return err
	}
	return nil
}

func (obj *UtpZoPlacePerpOrderIxArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `IsLong`:
	err = decoder.Decode(&obj.IsLong)
	if err != nil {
		return err
	}
	// Deserialize `LimitPrice`:
	err = decoder.Decode(&obj.LimitPrice)
	if err != nil {
		return err
	}
	// Deserialize `MaxBaseQuantity`:
	err = decoder.Decode(&obj.MaxBaseQuantity)
	if err != nil {
		return err
	}
	// Deserialize `MaxQuoteQuantity`:
	err = decoder.Decode(&obj.MaxQuoteQuantity)
	if err != nil {
		return err
	}
	// Deserialize `OrderType`:
	err = decoder.Decode(&obj.OrderType)
	if err != nil {
		return err
	}
	// Deserialize `Limit`:
	err = decoder.Decode(&obj.Limit)
	if err != nil {
		return err
	}
	// Deserialize `ClientId`:
	err = decoder.Decode(&obj.ClientId)
	if err != nil {
		return err
	}
	return nil
}

type UtpZoCancelPerpOrderIxArgs struct {
	OrderId  *ag_binary.Uint128 `bin:"optional"`
	IsLong   *bool              `bin:"optional"`
	ClientId *uint64            `bin:"optional"`
}

func (obj UtpZoCancelPerpOrderIxArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `OrderId` param (optional):
	{
		if obj.OrderId == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.OrderId)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `IsLong` param (optional):
	{
		if obj.IsLong == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.IsLong)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `ClientId` param (optional):
	{
		if obj.ClientId == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.ClientId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *UtpZoCancelPerpOrderIxArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `OrderId` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.OrderId)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `IsLong` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.IsLong)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `ClientId` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.ClientId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type MDecimal struct {
	Flags uint32
	Hi    uint32
	Lo    uint32
	Mid   uint32
}

func (obj MDecimal) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Flags` param:
	err = encoder.Encode(obj.Flags)
	if err != nil {
		return err
	}
	// Serialize `Hi` param:
	err = encoder.Encode(obj.Hi)
	if err != nil {
		return err
	}
	// Serialize `Lo` param:
	err = encoder.Encode(obj.Lo)
	if err != nil {
		return err
	}
	// Serialize `Mid` param:
	err = encoder.Encode(obj.Mid)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MDecimal) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Flags`:
	err = decoder.Decode(&obj.Flags)
	if err != nil {
		return err
	}
	// Deserialize `Hi`:
	err = decoder.Decode(&obj.Hi)
	if err != nil {
		return err
	}
	// Deserialize `Lo`:
	err = decoder.Decode(&obj.Lo)
	if err != nil {
		return err
	}
	// Deserialize `Mid`:
	err = decoder.Decode(&obj.Mid)
	if err != nil {
		return err
	}
	return nil
}

type UTPAccountConfig struct {
	Address        ag_solanago.PublicKey
	AuthoritySeed  ag_solanago.PublicKey
	AuthorityBump  uint8
	UtpAddressBook [4]ag_solanago.PublicKey
	ReservedSpace  [32]uint32
}

func (obj UTPAccountConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Address` param:
	err = encoder.Encode(obj.Address)
	if err != nil {
		return err
	}
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
	// Serialize `UtpAddressBook` param:
	err = encoder.Encode(obj.UtpAddressBook)
	if err != nil {
		return err
	}
	// Serialize `ReservedSpace` param:
	err = encoder.Encode(obj.ReservedSpace)
	if err != nil {
		return err
	}
	return nil
}

func (obj *UTPAccountConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Address`:
	err = decoder.Decode(&obj.Address)
	if err != nil {
		return err
	}
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
	// Deserialize `UtpAddressBook`:
	err = decoder.Decode(&obj.UtpAddressBook)
	if err != nil {
		return err
	}
	// Deserialize `ReservedSpace`:
	err = decoder.Decode(&obj.ReservedSpace)
	if err != nil {
		return err
	}
	return nil
}

type UTPConfig struct {
	UtpProgramId                   ag_solanago.PublicKey
	MarginRequirementDepositBuffer MDecimal
}

func (obj UTPConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `UtpProgramId` param:
	err = encoder.Encode(obj.UtpProgramId)
	if err != nil {
		return err
	}
	// Serialize `MarginRequirementDepositBuffer` param:
	err = encoder.Encode(obj.MarginRequirementDepositBuffer)
	if err != nil {
		return err
	}
	return nil
}

func (obj *UTPConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `UtpProgramId`:
	err = decoder.Decode(&obj.UtpProgramId)
	if err != nil {
		return err
	}
	// Deserialize `MarginRequirementDepositBuffer`:
	err = decoder.Decode(&obj.MarginRequirementDepositBuffer)
	if err != nil {
		return err
	}
	return nil
}

type GroupConfig struct {
	Admin  *ag_solanago.PublicKey `bin:"optional"`
	Bank   *BankConfig            `bin:"optional"`
	Paused *bool                  `bin:"optional"`
}

func (obj GroupConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Admin` param (optional):
	{
		if obj.Admin == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Admin)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Bank` param (optional):
	{
		if obj.Bank == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Bank)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Paused` param (optional):
	{
		if obj.Paused == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Paused)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *GroupConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Admin` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Admin)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Bank` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Bank)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Paused` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Paused)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type BankConfig struct {
	ScalingFactorC      *uint64 `bin:"optional"`
	FixedFee            *uint64 `bin:"optional"`
	InterestFee         *uint64 `bin:"optional"`
	InitMarginRatio     *uint64 `bin:"optional"`
	MaintMarginRatio    *uint64 `bin:"optional"`
	AccountDepositLimit *uint64 `bin:"optional"`
	LpDepositLimit      *uint64 `bin:"optional"`
}

func (obj BankConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ScalingFactorC` param (optional):
	{
		if obj.ScalingFactorC == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.ScalingFactorC)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `FixedFee` param (optional):
	{
		if obj.FixedFee == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.FixedFee)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `InterestFee` param (optional):
	{
		if obj.InterestFee == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.InterestFee)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `InitMarginRatio` param (optional):
	{
		if obj.InitMarginRatio == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.InitMarginRatio)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `MaintMarginRatio` param (optional):
	{
		if obj.MaintMarginRatio == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MaintMarginRatio)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `AccountDepositLimit` param (optional):
	{
		if obj.AccountDepositLimit == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.AccountDepositLimit)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `LpDepositLimit` param (optional):
	{
		if obj.LpDepositLimit == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.LpDepositLimit)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *BankConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ScalingFactorC` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.ScalingFactorC)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `FixedFee` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.FixedFee)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `InterestFee` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.InterestFee)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `InitMarginRatio` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.InitMarginRatio)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `MaintMarginRatio` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MaintMarginRatio)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `AccountDepositLimit` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.AccountDepositLimit)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `LpDepositLimit` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.LpDepositLimit)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type Bank struct {
	ScalingFactorC                     MDecimal
	FixedFee                           MDecimal
	InterestFee                        MDecimal
	DepositAccumulator                 MDecimal
	BorrowAccumulator                  MDecimal
	LastUpdate                         int64
	NativeDepositBalance               MDecimal
	NativeBorrowBalance                MDecimal
	Mint                               ag_solanago.PublicKey
	Vault                              ag_solanago.PublicKey
	VaultAuthorityPdaBump              uint8
	InsuranceVault                     ag_solanago.PublicKey
	InsuranceVaultAuthorityPdaBump     uint8
	InsuranceVaultOutstandingTransfers MDecimal
	FeeVault                           ag_solanago.PublicKey
	FeeVaultAuthorityPdaBump           uint8
	FeeVaultOutstandingTransfers       MDecimal
	InitMarginRatio                    MDecimal
	MaintMarginRatio                   MDecimal
	AccountDepositLimit                MDecimal
	LpDepositLimit                     MDecimal
	ReservedSpace                      [31]ag_binary.Uint128
}

func (obj Bank) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ScalingFactorC` param:
	err = encoder.Encode(obj.ScalingFactorC)
	if err != nil {
		return err
	}
	// Serialize `FixedFee` param:
	err = encoder.Encode(obj.FixedFee)
	if err != nil {
		return err
	}
	// Serialize `InterestFee` param:
	err = encoder.Encode(obj.InterestFee)
	if err != nil {
		return err
	}
	// Serialize `DepositAccumulator` param:
	err = encoder.Encode(obj.DepositAccumulator)
	if err != nil {
		return err
	}
	// Serialize `BorrowAccumulator` param:
	err = encoder.Encode(obj.BorrowAccumulator)
	if err != nil {
		return err
	}
	// Serialize `LastUpdate` param:
	err = encoder.Encode(obj.LastUpdate)
	if err != nil {
		return err
	}
	// Serialize `NativeDepositBalance` param:
	err = encoder.Encode(obj.NativeDepositBalance)
	if err != nil {
		return err
	}
	// Serialize `NativeBorrowBalance` param:
	err = encoder.Encode(obj.NativeBorrowBalance)
	if err != nil {
		return err
	}
	// Serialize `Mint` param:
	err = encoder.Encode(obj.Mint)
	if err != nil {
		return err
	}
	// Serialize `Vault` param:
	err = encoder.Encode(obj.Vault)
	if err != nil {
		return err
	}
	// Serialize `VaultAuthorityPdaBump` param:
	err = encoder.Encode(obj.VaultAuthorityPdaBump)
	if err != nil {
		return err
	}
	// Serialize `InsuranceVault` param:
	err = encoder.Encode(obj.InsuranceVault)
	if err != nil {
		return err
	}
	// Serialize `InsuranceVaultAuthorityPdaBump` param:
	err = encoder.Encode(obj.InsuranceVaultAuthorityPdaBump)
	if err != nil {
		return err
	}
	// Serialize `InsuranceVaultOutstandingTransfers` param:
	err = encoder.Encode(obj.InsuranceVaultOutstandingTransfers)
	if err != nil {
		return err
	}
	// Serialize `FeeVault` param:
	err = encoder.Encode(obj.FeeVault)
	if err != nil {
		return err
	}
	// Serialize `FeeVaultAuthorityPdaBump` param:
	err = encoder.Encode(obj.FeeVaultAuthorityPdaBump)
	if err != nil {
		return err
	}
	// Serialize `FeeVaultOutstandingTransfers` param:
	err = encoder.Encode(obj.FeeVaultOutstandingTransfers)
	if err != nil {
		return err
	}
	// Serialize `InitMarginRatio` param:
	err = encoder.Encode(obj.InitMarginRatio)
	if err != nil {
		return err
	}
	// Serialize `MaintMarginRatio` param:
	err = encoder.Encode(obj.MaintMarginRatio)
	if err != nil {
		return err
	}
	// Serialize `AccountDepositLimit` param:
	err = encoder.Encode(obj.AccountDepositLimit)
	if err != nil {
		return err
	}
	// Serialize `LpDepositLimit` param:
	err = encoder.Encode(obj.LpDepositLimit)
	if err != nil {
		return err
	}
	// Serialize `ReservedSpace` param:
	err = encoder.Encode(obj.ReservedSpace)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Bank) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ScalingFactorC`:
	err = decoder.Decode(&obj.ScalingFactorC)
	if err != nil {
		return err
	}
	// Deserialize `FixedFee`:
	err = decoder.Decode(&obj.FixedFee)
	if err != nil {
		return err
	}
	// Deserialize `InterestFee`:
	err = decoder.Decode(&obj.InterestFee)
	if err != nil {
		return err
	}
	// Deserialize `DepositAccumulator`:
	err = decoder.Decode(&obj.DepositAccumulator)
	if err != nil {
		return err
	}
	// Deserialize `BorrowAccumulator`:
	err = decoder.Decode(&obj.BorrowAccumulator)
	if err != nil {
		return err
	}
	// Deserialize `LastUpdate`:
	err = decoder.Decode(&obj.LastUpdate)
	if err != nil {
		return err
	}
	// Deserialize `NativeDepositBalance`:
	err = decoder.Decode(&obj.NativeDepositBalance)
	if err != nil {
		return err
	}
	// Deserialize `NativeBorrowBalance`:
	err = decoder.Decode(&obj.NativeBorrowBalance)
	if err != nil {
		return err
	}
	// Deserialize `Mint`:
	err = decoder.Decode(&obj.Mint)
	if err != nil {
		return err
	}
	// Deserialize `Vault`:
	err = decoder.Decode(&obj.Vault)
	if err != nil {
		return err
	}
	// Deserialize `VaultAuthorityPdaBump`:
	err = decoder.Decode(&obj.VaultAuthorityPdaBump)
	if err != nil {
		return err
	}
	// Deserialize `InsuranceVault`:
	err = decoder.Decode(&obj.InsuranceVault)
	if err != nil {
		return err
	}
	// Deserialize `InsuranceVaultAuthorityPdaBump`:
	err = decoder.Decode(&obj.InsuranceVaultAuthorityPdaBump)
	if err != nil {
		return err
	}
	// Deserialize `InsuranceVaultOutstandingTransfers`:
	err = decoder.Decode(&obj.InsuranceVaultOutstandingTransfers)
	if err != nil {
		return err
	}
	// Deserialize `FeeVault`:
	err = decoder.Decode(&obj.FeeVault)
	if err != nil {
		return err
	}
	// Deserialize `FeeVaultAuthorityPdaBump`:
	err = decoder.Decode(&obj.FeeVaultAuthorityPdaBump)
	if err != nil {
		return err
	}
	// Deserialize `FeeVaultOutstandingTransfers`:
	err = decoder.Decode(&obj.FeeVaultOutstandingTransfers)
	if err != nil {
		return err
	}
	// Deserialize `InitMarginRatio`:
	err = decoder.Decode(&obj.InitMarginRatio)
	if err != nil {
		return err
	}
	// Deserialize `MaintMarginRatio`:
	err = decoder.Decode(&obj.MaintMarginRatio)
	if err != nil {
		return err
	}
	// Deserialize `AccountDepositLimit`:
	err = decoder.Decode(&obj.AccountDepositLimit)
	if err != nil {
		return err
	}
	// Deserialize `LpDepositLimit`:
	err = decoder.Decode(&obj.LpDepositLimit)
	if err != nil {
		return err
	}
	// Deserialize `ReservedSpace`:
	err = decoder.Decode(&obj.ReservedSpace)
	if err != nil {
		return err
	}
	return nil
}

type MangoOrderType ag_binary.BorshEnum

const (
	MangoOrderTypeLimit MangoOrderType = iota
	MangoOrderTypeImmediateOrCancel
	MangoOrderTypePostOnly
	MangoOrderTypeMarket
	MangoOrderTypePostOnlySlide
)

func (value MangoOrderType) String() string {
	switch value {
	case MangoOrderTypeLimit:
		return "Limit"
	case MangoOrderTypeImmediateOrCancel:
		return "ImmediateOrCancel"
	case MangoOrderTypePostOnly:
		return "PostOnly"
	case MangoOrderTypeMarket:
		return "Market"
	case MangoOrderTypePostOnlySlide:
		return "PostOnlySlide"
	default:
		return ""
	}
}

type MangoSide ag_binary.BorshEnum

const (
	MangoSideBid MangoSide = iota
	MangoSideAsk
)

func (value MangoSide) String() string {
	switch value {
	case MangoSideBid:
		return "Bid"
	case MangoSideAsk:
		return "Ask"
	default:
		return ""
	}
}

type MangoExpiryType ag_binary.BorshEnum

const (
	MangoExpiryTypeAbsolute MangoExpiryType = iota
	MangoExpiryTypeRelative
)

func (value MangoExpiryType) String() string {
	switch value {
	case MangoExpiryTypeAbsolute:
		return "Absolute"
	case MangoExpiryTypeRelative:
		return "Relative"
	default:
		return ""
	}
}

type MarginRequirement ag_binary.BorshEnum

const (
	MarginRequirementInit MarginRequirement = iota
	MarginRequirementMaint
)

func (value MarginRequirement) String() string {
	switch value {
	case MarginRequirementInit:
		return "Init"
	case MarginRequirementMaint:
		return "Maint"
	default:
		return ""
	}
}

type EquityType ag_binary.BorshEnum

const (
	EquityTypeInitReqAdjusted EquityType = iota
	EquityTypeTotal
)

func (value EquityType) String() string {
	switch value {
	case EquityTypeInitReqAdjusted:
		return "InitReqAdjusted"
	case EquityTypeTotal:
		return "Total"
	default:
		return ""
	}
}

type BankVaultType ag_binary.BorshEnum

const (
	BankVaultTypeLiquidityVault BankVaultType = iota
	BankVaultTypeInsuranceVault
	BankVaultTypeProtocolFeeVault
)

func (value BankVaultType) String() string {
	switch value {
	case BankVaultTypeLiquidityVault:
		return "LiquidityVault"
	case BankVaultTypeInsuranceVault:
		return "InsuranceVault"
	case BankVaultTypeProtocolFeeVault:
		return "ProtocolFeeVault"
	default:
		return ""
	}
}

type InternalTransferType ag_binary.BorshEnum

const (
	InternalTransferTypeInsuranceFee InternalTransferType = iota
	InternalTransferTypeProtocolFee
)

func (value InternalTransferType) String() string {
	switch value {
	case InternalTransferTypeInsuranceFee:
		return "InsuranceFee"
	case InternalTransferTypeProtocolFee:
		return "ProtocolFee"
	default:
		return ""
	}
}

type LendingSide ag_binary.BorshEnum

const (
	LendingSideBorrow LendingSide = iota
	LendingSideDeposit
)

func (value LendingSide) String() string {
	switch value {
	case LendingSideBorrow:
		return "Borrow"
	case LendingSideDeposit:
		return "Deposit"
	default:
		return ""
	}
}

type OrderType ag_binary.BorshEnum

const (
	OrderTypeLimit OrderType = iota
	OrderTypeImmediateOrCancel
	OrderTypePostOnly
	OrderTypeReduceOnlyIoc
	OrderTypeReduceOnlyLimit
	OrderTypeFillOrKill
)

func (value OrderType) String() string {
	switch value {
	case OrderTypeLimit:
		return "Limit"
	case OrderTypeImmediateOrCancel:
		return "ImmediateOrCancel"
	case OrderTypePostOnly:
		return "PostOnly"
	case OrderTypeReduceOnlyIoc:
		return "ReduceOnlyIoc"
	case OrderTypeReduceOnlyLimit:
		return "ReduceOnlyLimit"
	case OrderTypeFillOrKill:
		return "FillOrKill"
	default:
		return ""
	}
}
