/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// IdentifyReplyCommandExtendedDiagnosticSummary is the data-structure of this message
type IdentifyReplyCommandExtendedDiagnosticSummary struct {
	*IdentifyReplyCommand
	LowApplication         ApplicationIdContainer
	HighApplication        ApplicationIdContainer
	Area                   byte
	Crc                    uint16
	SerialNumber           uint32
	NetworkVoltage         byte
	OutputUnit             bool
	EnableChecksumAlarm    bool
	NetworkVoltageMarginal bool
	NetworkVoltageLow      bool
	UnitInLearnMode        bool
	MicroPowerReset        bool
	InternalStackOverflow  bool
	CommsTxError           bool
	MicroReset             bool
	EEDataError            bool
	EEChecksumError        bool
	EEWriteError           bool
	InstallationMMIError   bool
}

// IIdentifyReplyCommandExtendedDiagnosticSummary is the corresponding interface of IdentifyReplyCommandExtendedDiagnosticSummary
type IIdentifyReplyCommandExtendedDiagnosticSummary interface {
	IIdentifyReplyCommand
	// GetLowApplication returns LowApplication (property field)
	GetLowApplication() ApplicationIdContainer
	// GetHighApplication returns HighApplication (property field)
	GetHighApplication() ApplicationIdContainer
	// GetArea returns Area (property field)
	GetArea() byte
	// GetCrc returns Crc (property field)
	GetCrc() uint16
	// GetSerialNumber returns SerialNumber (property field)
	GetSerialNumber() uint32
	// GetNetworkVoltage returns NetworkVoltage (property field)
	GetNetworkVoltage() byte
	// GetOutputUnit returns OutputUnit (property field)
	GetOutputUnit() bool
	// GetEnableChecksumAlarm returns EnableChecksumAlarm (property field)
	GetEnableChecksumAlarm() bool
	// GetNetworkVoltageMarginal returns NetworkVoltageMarginal (property field)
	GetNetworkVoltageMarginal() bool
	// GetNetworkVoltageLow returns NetworkVoltageLow (property field)
	GetNetworkVoltageLow() bool
	// GetUnitInLearnMode returns UnitInLearnMode (property field)
	GetUnitInLearnMode() bool
	// GetMicroPowerReset returns MicroPowerReset (property field)
	GetMicroPowerReset() bool
	// GetInternalStackOverflow returns InternalStackOverflow (property field)
	GetInternalStackOverflow() bool
	// GetCommsTxError returns CommsTxError (property field)
	GetCommsTxError() bool
	// GetMicroReset returns MicroReset (property field)
	GetMicroReset() bool
	// GetEEDataError returns EEDataError (property field)
	GetEEDataError() bool
	// GetEEChecksumError returns EEChecksumError (property field)
	GetEEChecksumError() bool
	// GetEEWriteError returns EEWriteError (property field)
	GetEEWriteError() bool
	// GetInstallationMMIError returns InstallationMMIError (property field)
	GetInstallationMMIError() bool
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetAttribute() Attribute {
	return Attribute_ExtendedDiagnosticSummary
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) InitializeParent(parent *IdentifyReplyCommand) {
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetParent() *IdentifyReplyCommand {
	return m.IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetLowApplication() ApplicationIdContainer {
	return m.LowApplication
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetHighApplication() ApplicationIdContainer {
	return m.HighApplication
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetArea() byte {
	return m.Area
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetCrc() uint16 {
	return m.Crc
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetSerialNumber() uint32 {
	return m.SerialNumber
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetNetworkVoltage() byte {
	return m.NetworkVoltage
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetOutputUnit() bool {
	return m.OutputUnit
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetEnableChecksumAlarm() bool {
	return m.EnableChecksumAlarm
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetNetworkVoltageMarginal() bool {
	return m.NetworkVoltageMarginal
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetNetworkVoltageLow() bool {
	return m.NetworkVoltageLow
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetUnitInLearnMode() bool {
	return m.UnitInLearnMode
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetMicroPowerReset() bool {
	return m.MicroPowerReset
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetInternalStackOverflow() bool {
	return m.InternalStackOverflow
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetCommsTxError() bool {
	return m.CommsTxError
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetMicroReset() bool {
	return m.MicroReset
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetEEDataError() bool {
	return m.EEDataError
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetEEChecksumError() bool {
	return m.EEChecksumError
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetEEWriteError() bool {
	return m.EEWriteError
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetInstallationMMIError() bool {
	return m.InstallationMMIError
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandExtendedDiagnosticSummary factory function for IdentifyReplyCommandExtendedDiagnosticSummary
func NewIdentifyReplyCommandExtendedDiagnosticSummary(lowApplication ApplicationIdContainer, highApplication ApplicationIdContainer, area byte, crc uint16, serialNumber uint32, networkVoltage byte, outputUnit bool, enableChecksumAlarm bool, networkVoltageMarginal bool, networkVoltageLow bool, unitInLearnMode bool, microPowerReset bool, internalStackOverflow bool, commsTxError bool, microReset bool, EEDataError bool, EEChecksumError bool, EEWriteError bool, installationMMIError bool) *IdentifyReplyCommandExtendedDiagnosticSummary {
	_result := &IdentifyReplyCommandExtendedDiagnosticSummary{
		LowApplication:         lowApplication,
		HighApplication:        highApplication,
		Area:                   area,
		Crc:                    crc,
		SerialNumber:           serialNumber,
		NetworkVoltage:         networkVoltage,
		OutputUnit:             outputUnit,
		EnableChecksumAlarm:    enableChecksumAlarm,
		NetworkVoltageMarginal: networkVoltageMarginal,
		NetworkVoltageLow:      networkVoltageLow,
		UnitInLearnMode:        unitInLearnMode,
		MicroPowerReset:        microPowerReset,
		InternalStackOverflow:  internalStackOverflow,
		CommsTxError:           commsTxError,
		MicroReset:             microReset,
		EEDataError:            EEDataError,
		EEChecksumError:        EEChecksumError,
		EEWriteError:           EEWriteError,
		InstallationMMIError:   installationMMIError,
		IdentifyReplyCommand:   NewIdentifyReplyCommand(),
	}
	_result.Child = _result
	return _result
}

func CastIdentifyReplyCommandExtendedDiagnosticSummary(structType interface{}) *IdentifyReplyCommandExtendedDiagnosticSummary {
	if casted, ok := structType.(IdentifyReplyCommandExtendedDiagnosticSummary); ok {
		return &casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandExtendedDiagnosticSummary); ok {
		return casted
	}
	if casted, ok := structType.(IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandExtendedDiagnosticSummary(casted.Child)
	}
	if casted, ok := structType.(*IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandExtendedDiagnosticSummary(casted.Child)
	}
	return nil
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetTypeName() string {
	return "IdentifyReplyCommandExtendedDiagnosticSummary"
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lowApplication)
	lengthInBits += 8

	// Simple field (highApplication)
	lengthInBits += 8

	// Simple field (area)
	lengthInBits += 8

	// Simple field (crc)
	lengthInBits += 16

	// Simple field (serialNumber)
	lengthInBits += 32

	// Simple field (networkVoltage)
	lengthInBits += 8

	// Simple field (outputUnit)
	lengthInBits += 1

	// Simple field (enableChecksumAlarm)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (networkVoltageMarginal)
	lengthInBits += 1

	// Simple field (networkVoltageLow)
	lengthInBits += 1

	// Simple field (unitInLearnMode)
	lengthInBits += 1

	// Simple field (microPowerReset)
	lengthInBits += 1

	// Simple field (internalStackOverflow)
	lengthInBits += 1

	// Simple field (commsTxError)
	lengthInBits += 1

	// Simple field (microReset)
	lengthInBits += 1

	// Simple field (EEDataError)
	lengthInBits += 1

	// Simple field (EEChecksumError)
	lengthInBits += 1

	// Simple field (EEWriteError)
	lengthInBits += 1

	// Simple field (installationMMIError)
	lengthInBits += 1

	return lengthInBits
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandExtendedDiagnosticSummaryParse(readBuffer utils.ReadBuffer, attribute Attribute) (*IdentifyReplyCommandExtendedDiagnosticSummary, error) {
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandExtendedDiagnosticSummary"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (lowApplication)
	if pullErr := readBuffer.PullContext("lowApplication"); pullErr != nil {
		return nil, pullErr
	}
	_lowApplication, _lowApplicationErr := ApplicationIdContainerParse(readBuffer)
	if _lowApplicationErr != nil {
		return nil, errors.Wrap(_lowApplicationErr, "Error parsing 'lowApplication' field")
	}
	lowApplication := _lowApplication
	if closeErr := readBuffer.CloseContext("lowApplication"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (highApplication)
	if pullErr := readBuffer.PullContext("highApplication"); pullErr != nil {
		return nil, pullErr
	}
	_highApplication, _highApplicationErr := ApplicationIdContainerParse(readBuffer)
	if _highApplicationErr != nil {
		return nil, errors.Wrap(_highApplicationErr, "Error parsing 'highApplication' field")
	}
	highApplication := _highApplication
	if closeErr := readBuffer.CloseContext("highApplication"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (area)
	_area, _areaErr := readBuffer.ReadByte("area")
	if _areaErr != nil {
		return nil, errors.Wrap(_areaErr, "Error parsing 'area' field")
	}
	area := _area

	// Simple Field (crc)
	_crc, _crcErr := readBuffer.ReadUint16("crc", 16)
	if _crcErr != nil {
		return nil, errors.Wrap(_crcErr, "Error parsing 'crc' field")
	}
	crc := _crc

	// Simple Field (serialNumber)
	_serialNumber, _serialNumberErr := readBuffer.ReadUint32("serialNumber", 32)
	if _serialNumberErr != nil {
		return nil, errors.Wrap(_serialNumberErr, "Error parsing 'serialNumber' field")
	}
	serialNumber := _serialNumber

	// Simple Field (networkVoltage)
	_networkVoltage, _networkVoltageErr := readBuffer.ReadByte("networkVoltage")
	if _networkVoltageErr != nil {
		return nil, errors.Wrap(_networkVoltageErr, "Error parsing 'networkVoltage' field")
	}
	networkVoltage := _networkVoltage

	// Simple Field (outputUnit)
	_outputUnit, _outputUnitErr := readBuffer.ReadBit("outputUnit")
	if _outputUnitErr != nil {
		return nil, errors.Wrap(_outputUnitErr, "Error parsing 'outputUnit' field")
	}
	outputUnit := _outputUnit

	// Simple Field (enableChecksumAlarm)
	_enableChecksumAlarm, _enableChecksumAlarmErr := readBuffer.ReadBit("enableChecksumAlarm")
	if _enableChecksumAlarmErr != nil {
		return nil, errors.Wrap(_enableChecksumAlarmErr, "Error parsing 'enableChecksumAlarm' field")
	}
	enableChecksumAlarm := _enableChecksumAlarm

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 1)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 1)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 1)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (networkVoltageMarginal)
	_networkVoltageMarginal, _networkVoltageMarginalErr := readBuffer.ReadBit("networkVoltageMarginal")
	if _networkVoltageMarginalErr != nil {
		return nil, errors.Wrap(_networkVoltageMarginalErr, "Error parsing 'networkVoltageMarginal' field")
	}
	networkVoltageMarginal := _networkVoltageMarginal

	// Simple Field (networkVoltageLow)
	_networkVoltageLow, _networkVoltageLowErr := readBuffer.ReadBit("networkVoltageLow")
	if _networkVoltageLowErr != nil {
		return nil, errors.Wrap(_networkVoltageLowErr, "Error parsing 'networkVoltageLow' field")
	}
	networkVoltageLow := _networkVoltageLow

	// Simple Field (unitInLearnMode)
	_unitInLearnMode, _unitInLearnModeErr := readBuffer.ReadBit("unitInLearnMode")
	if _unitInLearnModeErr != nil {
		return nil, errors.Wrap(_unitInLearnModeErr, "Error parsing 'unitInLearnMode' field")
	}
	unitInLearnMode := _unitInLearnMode

	// Simple Field (microPowerReset)
	_microPowerReset, _microPowerResetErr := readBuffer.ReadBit("microPowerReset")
	if _microPowerResetErr != nil {
		return nil, errors.Wrap(_microPowerResetErr, "Error parsing 'microPowerReset' field")
	}
	microPowerReset := _microPowerReset

	// Simple Field (internalStackOverflow)
	_internalStackOverflow, _internalStackOverflowErr := readBuffer.ReadBit("internalStackOverflow")
	if _internalStackOverflowErr != nil {
		return nil, errors.Wrap(_internalStackOverflowErr, "Error parsing 'internalStackOverflow' field")
	}
	internalStackOverflow := _internalStackOverflow

	// Simple Field (commsTxError)
	_commsTxError, _commsTxErrorErr := readBuffer.ReadBit("commsTxError")
	if _commsTxErrorErr != nil {
		return nil, errors.Wrap(_commsTxErrorErr, "Error parsing 'commsTxError' field")
	}
	commsTxError := _commsTxError

	// Simple Field (microReset)
	_microReset, _microResetErr := readBuffer.ReadBit("microReset")
	if _microResetErr != nil {
		return nil, errors.Wrap(_microResetErr, "Error parsing 'microReset' field")
	}
	microReset := _microReset

	// Simple Field (EEDataError)
	_EEDataError, _EEDataErrorErr := readBuffer.ReadBit("EEDataError")
	if _EEDataErrorErr != nil {
		return nil, errors.Wrap(_EEDataErrorErr, "Error parsing 'EEDataError' field")
	}
	EEDataError := _EEDataError

	// Simple Field (EEChecksumError)
	_EEChecksumError, _EEChecksumErrorErr := readBuffer.ReadBit("EEChecksumError")
	if _EEChecksumErrorErr != nil {
		return nil, errors.Wrap(_EEChecksumErrorErr, "Error parsing 'EEChecksumError' field")
	}
	EEChecksumError := _EEChecksumError

	// Simple Field (EEWriteError)
	_EEWriteError, _EEWriteErrorErr := readBuffer.ReadBit("EEWriteError")
	if _EEWriteErrorErr != nil {
		return nil, errors.Wrap(_EEWriteErrorErr, "Error parsing 'EEWriteError' field")
	}
	EEWriteError := _EEWriteError

	// Simple Field (installationMMIError)
	_installationMMIError, _installationMMIErrorErr := readBuffer.ReadBit("installationMMIError")
	if _installationMMIErrorErr != nil {
		return nil, errors.Wrap(_installationMMIErrorErr, "Error parsing 'installationMMIError' field")
	}
	installationMMIError := _installationMMIError

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandExtendedDiagnosticSummary"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &IdentifyReplyCommandExtendedDiagnosticSummary{
		LowApplication:         lowApplication,
		HighApplication:        highApplication,
		Area:                   area,
		Crc:                    crc,
		SerialNumber:           serialNumber,
		NetworkVoltage:         networkVoltage,
		OutputUnit:             outputUnit,
		EnableChecksumAlarm:    enableChecksumAlarm,
		NetworkVoltageMarginal: networkVoltageMarginal,
		NetworkVoltageLow:      networkVoltageLow,
		UnitInLearnMode:        unitInLearnMode,
		MicroPowerReset:        microPowerReset,
		InternalStackOverflow:  internalStackOverflow,
		CommsTxError:           commsTxError,
		MicroReset:             microReset,
		EEDataError:            EEDataError,
		EEChecksumError:        EEChecksumError,
		EEWriteError:           EEWriteError,
		InstallationMMIError:   installationMMIError,
		IdentifyReplyCommand:   &IdentifyReplyCommand{},
	}
	_child.IdentifyReplyCommand.Child = _child
	return _child, nil
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandExtendedDiagnosticSummary"); pushErr != nil {
			return pushErr
		}

		// Simple Field (lowApplication)
		if pushErr := writeBuffer.PushContext("lowApplication"); pushErr != nil {
			return pushErr
		}
		_lowApplicationErr := m.LowApplication.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("lowApplication"); popErr != nil {
			return popErr
		}
		if _lowApplicationErr != nil {
			return errors.Wrap(_lowApplicationErr, "Error serializing 'lowApplication' field")
		}

		// Simple Field (highApplication)
		if pushErr := writeBuffer.PushContext("highApplication"); pushErr != nil {
			return pushErr
		}
		_highApplicationErr := m.HighApplication.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("highApplication"); popErr != nil {
			return popErr
		}
		if _highApplicationErr != nil {
			return errors.Wrap(_highApplicationErr, "Error serializing 'highApplication' field")
		}

		// Simple Field (area)
		area := byte(m.Area)
		_areaErr := writeBuffer.WriteByte("area", (area))
		if _areaErr != nil {
			return errors.Wrap(_areaErr, "Error serializing 'area' field")
		}

		// Simple Field (crc)
		crc := uint16(m.Crc)
		_crcErr := writeBuffer.WriteUint16("crc", 16, (crc))
		if _crcErr != nil {
			return errors.Wrap(_crcErr, "Error serializing 'crc' field")
		}

		// Simple Field (serialNumber)
		serialNumber := uint32(m.SerialNumber)
		_serialNumberErr := writeBuffer.WriteUint32("serialNumber", 32, (serialNumber))
		if _serialNumberErr != nil {
			return errors.Wrap(_serialNumberErr, "Error serializing 'serialNumber' field")
		}

		// Simple Field (networkVoltage)
		networkVoltage := byte(m.NetworkVoltage)
		_networkVoltageErr := writeBuffer.WriteByte("networkVoltage", (networkVoltage))
		if _networkVoltageErr != nil {
			return errors.Wrap(_networkVoltageErr, "Error serializing 'networkVoltage' field")
		}

		// Simple Field (outputUnit)
		outputUnit := bool(m.OutputUnit)
		_outputUnitErr := writeBuffer.WriteBit("outputUnit", (outputUnit))
		if _outputUnitErr != nil {
			return errors.Wrap(_outputUnitErr, "Error serializing 'outputUnit' field")
		}

		// Simple Field (enableChecksumAlarm)
		enableChecksumAlarm := bool(m.EnableChecksumAlarm)
		_enableChecksumAlarmErr := writeBuffer.WriteBit("enableChecksumAlarm", (enableChecksumAlarm))
		if _enableChecksumAlarmErr != nil {
			return errors.Wrap(_enableChecksumAlarmErr, "Error serializing 'enableChecksumAlarm' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 1, uint8(0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 1, uint8(0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 1, uint8(0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (networkVoltageMarginal)
		networkVoltageMarginal := bool(m.NetworkVoltageMarginal)
		_networkVoltageMarginalErr := writeBuffer.WriteBit("networkVoltageMarginal", (networkVoltageMarginal))
		if _networkVoltageMarginalErr != nil {
			return errors.Wrap(_networkVoltageMarginalErr, "Error serializing 'networkVoltageMarginal' field")
		}

		// Simple Field (networkVoltageLow)
		networkVoltageLow := bool(m.NetworkVoltageLow)
		_networkVoltageLowErr := writeBuffer.WriteBit("networkVoltageLow", (networkVoltageLow))
		if _networkVoltageLowErr != nil {
			return errors.Wrap(_networkVoltageLowErr, "Error serializing 'networkVoltageLow' field")
		}

		// Simple Field (unitInLearnMode)
		unitInLearnMode := bool(m.UnitInLearnMode)
		_unitInLearnModeErr := writeBuffer.WriteBit("unitInLearnMode", (unitInLearnMode))
		if _unitInLearnModeErr != nil {
			return errors.Wrap(_unitInLearnModeErr, "Error serializing 'unitInLearnMode' field")
		}

		// Simple Field (microPowerReset)
		microPowerReset := bool(m.MicroPowerReset)
		_microPowerResetErr := writeBuffer.WriteBit("microPowerReset", (microPowerReset))
		if _microPowerResetErr != nil {
			return errors.Wrap(_microPowerResetErr, "Error serializing 'microPowerReset' field")
		}

		// Simple Field (internalStackOverflow)
		internalStackOverflow := bool(m.InternalStackOverflow)
		_internalStackOverflowErr := writeBuffer.WriteBit("internalStackOverflow", (internalStackOverflow))
		if _internalStackOverflowErr != nil {
			return errors.Wrap(_internalStackOverflowErr, "Error serializing 'internalStackOverflow' field")
		}

		// Simple Field (commsTxError)
		commsTxError := bool(m.CommsTxError)
		_commsTxErrorErr := writeBuffer.WriteBit("commsTxError", (commsTxError))
		if _commsTxErrorErr != nil {
			return errors.Wrap(_commsTxErrorErr, "Error serializing 'commsTxError' field")
		}

		// Simple Field (microReset)
		microReset := bool(m.MicroReset)
		_microResetErr := writeBuffer.WriteBit("microReset", (microReset))
		if _microResetErr != nil {
			return errors.Wrap(_microResetErr, "Error serializing 'microReset' field")
		}

		// Simple Field (EEDataError)
		EEDataError := bool(m.EEDataError)
		_EEDataErrorErr := writeBuffer.WriteBit("EEDataError", (EEDataError))
		if _EEDataErrorErr != nil {
			return errors.Wrap(_EEDataErrorErr, "Error serializing 'EEDataError' field")
		}

		// Simple Field (EEChecksumError)
		EEChecksumError := bool(m.EEChecksumError)
		_EEChecksumErrorErr := writeBuffer.WriteBit("EEChecksumError", (EEChecksumError))
		if _EEChecksumErrorErr != nil {
			return errors.Wrap(_EEChecksumErrorErr, "Error serializing 'EEChecksumError' field")
		}

		// Simple Field (EEWriteError)
		EEWriteError := bool(m.EEWriteError)
		_EEWriteErrorErr := writeBuffer.WriteBit("EEWriteError", (EEWriteError))
		if _EEWriteErrorErr != nil {
			return errors.Wrap(_EEWriteErrorErr, "Error serializing 'EEWriteError' field")
		}

		// Simple Field (installationMMIError)
		installationMMIError := bool(m.InstallationMMIError)
		_installationMMIErrorErr := writeBuffer.WriteBit("installationMMIError", (installationMMIError))
		if _installationMMIErrorErr != nil {
			return errors.Wrap(_installationMMIErrorErr, "Error serializing 'installationMMIError' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandExtendedDiagnosticSummary"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *IdentifyReplyCommandExtendedDiagnosticSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
