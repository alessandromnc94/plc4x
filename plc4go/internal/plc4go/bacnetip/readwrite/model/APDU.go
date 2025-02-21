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

// APDU is the data-structure of this message
type APDU struct {

	// Arguments.
	ApduLength uint16
	Child      IAPDUChild
}

// IAPDU is the corresponding interface of APDU
type IAPDU interface {
	// GetApduType returns ApduType (discriminator field)
	GetApduType() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IAPDUParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IAPDU, serializeChildFunction func() error) error
	GetTypeName() string
}

type IAPDUChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *APDU)
	GetParent() *APDU

	GetTypeName() string
	IAPDU
}

// NewAPDU factory function for APDU
func NewAPDU(apduLength uint16) *APDU {
	return &APDU{ApduLength: apduLength}
}

func CastAPDU(structType interface{}) *APDU {
	if casted, ok := structType.(APDU); ok {
		return &casted
	}
	if casted, ok := structType.(*APDU); ok {
		return casted
	}
	if casted, ok := structType.(IAPDUChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *APDU) GetTypeName() string {
	return "APDU"
}

func (m *APDU) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *APDU) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *APDU) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (apduType)
	lengthInBits += 4

	return lengthInBits
}

func (m *APDU) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func APDUParse(readBuffer utils.ReadBuffer, apduLength uint16) (*APDU, error) {
	if pullErr := readBuffer.PullContext("APDU"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Discriminator Field (apduType) (Used as input to a switch field)
	apduType, _apduTypeErr := readBuffer.ReadUint8("apduType", 4)
	if _apduTypeErr != nil {
		return nil, errors.Wrap(_apduTypeErr, "Error parsing 'apduType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type APDUChild interface {
		InitializeParent(*APDU)
		GetParent() *APDU
	}
	var _child APDUChild
	var typeSwitchError error
	switch {
	case apduType == 0x0: // APDUConfirmedRequest
		_child, typeSwitchError = APDUConfirmedRequestParse(readBuffer, apduLength)
	case apduType == 0x1: // APDUUnconfirmedRequest
		_child, typeSwitchError = APDUUnconfirmedRequestParse(readBuffer, apduLength)
	case apduType == 0x2: // APDUSimpleAck
		_child, typeSwitchError = APDUSimpleAckParse(readBuffer, apduLength)
	case apduType == 0x3: // APDUComplexAck
		_child, typeSwitchError = APDUComplexAckParse(readBuffer, apduLength)
	case apduType == 0x4: // APDUSegmentAck
		_child, typeSwitchError = APDUSegmentAckParse(readBuffer, apduLength)
	case apduType == 0x5: // APDUError
		_child, typeSwitchError = APDUErrorParse(readBuffer, apduLength)
	case apduType == 0x6: // APDUReject
		_child, typeSwitchError = APDURejectParse(readBuffer, apduLength)
	case apduType == 0x7: // APDUAbort
		_child, typeSwitchError = APDUAbortParse(readBuffer, apduLength)
	case true: // APDUUnknown
		_child, typeSwitchError = APDUUnknownParse(readBuffer, apduLength)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("APDU"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent())
	return _child.GetParent(), nil
}

func (m *APDU) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *APDU) SerializeParent(writeBuffer utils.WriteBuffer, child IAPDU, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("APDU"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (apduType) (Used as input to a switch field)
	apduType := uint8(child.GetApduType())
	_apduTypeErr := writeBuffer.WriteUint8("apduType", 4, (apduType))

	if _apduTypeErr != nil {
		return errors.Wrap(_apduTypeErr, "Error serializing 'apduType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("APDU"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *APDU) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
