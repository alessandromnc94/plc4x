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
)

// Code generated by code-generation. DO NOT EDIT.

// FirmataCommandSetPinMode is the data-structure of this message
type FirmataCommandSetPinMode struct {
	*FirmataCommand
	Pin  uint8
	Mode PinMode

	// Arguments.
	Response bool
}

// IFirmataCommandSetPinMode is the corresponding interface of FirmataCommandSetPinMode
type IFirmataCommandSetPinMode interface {
	IFirmataCommand
	// GetPin returns Pin (property field)
	GetPin() uint8
	// GetMode returns Mode (property field)
	GetMode() PinMode
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

func (m *FirmataCommandSetPinMode) GetCommandCode() uint8 {
	return 0x4
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *FirmataCommandSetPinMode) InitializeParent(parent *FirmataCommand) {}

func (m *FirmataCommandSetPinMode) GetParent() *FirmataCommand {
	return m.FirmataCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *FirmataCommandSetPinMode) GetPin() uint8 {
	return m.Pin
}

func (m *FirmataCommandSetPinMode) GetMode() PinMode {
	return m.Mode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewFirmataCommandSetPinMode factory function for FirmataCommandSetPinMode
func NewFirmataCommandSetPinMode(pin uint8, mode PinMode, response bool) *FirmataCommandSetPinMode {
	_result := &FirmataCommandSetPinMode{
		Pin:            pin,
		Mode:           mode,
		FirmataCommand: NewFirmataCommand(response),
	}
	_result.Child = _result
	return _result
}

func CastFirmataCommandSetPinMode(structType interface{}) *FirmataCommandSetPinMode {
	if casted, ok := structType.(FirmataCommandSetPinMode); ok {
		return &casted
	}
	if casted, ok := structType.(*FirmataCommandSetPinMode); ok {
		return casted
	}
	if casted, ok := structType.(FirmataCommand); ok {
		return CastFirmataCommandSetPinMode(casted.Child)
	}
	if casted, ok := structType.(*FirmataCommand); ok {
		return CastFirmataCommandSetPinMode(casted.Child)
	}
	return nil
}

func (m *FirmataCommandSetPinMode) GetTypeName() string {
	return "FirmataCommandSetPinMode"
}

func (m *FirmataCommandSetPinMode) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *FirmataCommandSetPinMode) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (pin)
	lengthInBits += 8

	// Simple field (mode)
	lengthInBits += 8

	return lengthInBits
}

func (m *FirmataCommandSetPinMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func FirmataCommandSetPinModeParse(readBuffer utils.ReadBuffer, response bool) (*FirmataCommandSetPinMode, error) {
	if pullErr := readBuffer.PullContext("FirmataCommandSetPinMode"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (pin)
	_pin, _pinErr := readBuffer.ReadUint8("pin", 8)
	if _pinErr != nil {
		return nil, errors.Wrap(_pinErr, "Error parsing 'pin' field")
	}
	pin := _pin

	// Simple Field (mode)
	if pullErr := readBuffer.PullContext("mode"); pullErr != nil {
		return nil, pullErr
	}
	_mode, _modeErr := PinModeParse(readBuffer)
	if _modeErr != nil {
		return nil, errors.Wrap(_modeErr, "Error parsing 'mode' field")
	}
	mode := _mode
	if closeErr := readBuffer.CloseContext("mode"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("FirmataCommandSetPinMode"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &FirmataCommandSetPinMode{
		Pin:            pin,
		Mode:           mode,
		FirmataCommand: &FirmataCommand{},
	}
	_child.FirmataCommand.Child = _child
	return _child, nil
}

func (m *FirmataCommandSetPinMode) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataCommandSetPinMode"); pushErr != nil {
			return pushErr
		}

		// Simple Field (pin)
		pin := uint8(m.Pin)
		_pinErr := writeBuffer.WriteUint8("pin", 8, (pin))
		if _pinErr != nil {
			return errors.Wrap(_pinErr, "Error serializing 'pin' field")
		}

		// Simple Field (mode)
		if pushErr := writeBuffer.PushContext("mode"); pushErr != nil {
			return pushErr
		}
		_modeErr := m.Mode.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("mode"); popErr != nil {
			return popErr
		}
		if _modeErr != nil {
			return errors.Wrap(_modeErr, "Error serializing 'mode' field")
		}

		if popErr := writeBuffer.PopContext("FirmataCommandSetPinMode"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *FirmataCommandSetPinMode) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
