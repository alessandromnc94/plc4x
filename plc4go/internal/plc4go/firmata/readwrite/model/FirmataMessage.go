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

// FirmataMessage is the data-structure of this message
type FirmataMessage struct {

	// Arguments.
	Response bool
	Child    IFirmataMessageChild
}

// IFirmataMessage is the corresponding interface of FirmataMessage
type IFirmataMessage interface {
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IFirmataMessageParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IFirmataMessage, serializeChildFunction func() error) error
	GetTypeName() string
}

type IFirmataMessageChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *FirmataMessage)
	GetParent() *FirmataMessage

	GetTypeName() string
	IFirmataMessage
}

// NewFirmataMessage factory function for FirmataMessage
func NewFirmataMessage(response bool) *FirmataMessage {
	return &FirmataMessage{Response: response}
}

func CastFirmataMessage(structType interface{}) *FirmataMessage {
	if casted, ok := structType.(FirmataMessage); ok {
		return &casted
	}
	if casted, ok := structType.(*FirmataMessage); ok {
		return casted
	}
	if casted, ok := structType.(IFirmataMessageChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *FirmataMessage) GetTypeName() string {
	return "FirmataMessage"
}

func (m *FirmataMessage) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *FirmataMessage) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *FirmataMessage) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (messageType)
	lengthInBits += 4

	return lengthInBits
}

func (m *FirmataMessage) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func FirmataMessageParse(readBuffer utils.ReadBuffer, response bool) (*FirmataMessage, error) {
	if pullErr := readBuffer.PullContext("FirmataMessage"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType, _messageTypeErr := readBuffer.ReadUint8("messageType", 4)
	if _messageTypeErr != nil {
		return nil, errors.Wrap(_messageTypeErr, "Error parsing 'messageType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type FirmataMessageChild interface {
		InitializeParent(*FirmataMessage)
		GetParent() *FirmataMessage
	}
	var _child FirmataMessageChild
	var typeSwitchError error
	switch {
	case messageType == 0xE: // FirmataMessageAnalogIO
		_child, typeSwitchError = FirmataMessageAnalogIOParse(readBuffer, response)
	case messageType == 0x9: // FirmataMessageDigitalIO
		_child, typeSwitchError = FirmataMessageDigitalIOParse(readBuffer, response)
	case messageType == 0xC: // FirmataMessageSubscribeAnalogPinValue
		_child, typeSwitchError = FirmataMessageSubscribeAnalogPinValueParse(readBuffer, response)
	case messageType == 0xD: // FirmataMessageSubscribeDigitalPinValue
		_child, typeSwitchError = FirmataMessageSubscribeDigitalPinValueParse(readBuffer, response)
	case messageType == 0xF: // FirmataMessageCommand
		_child, typeSwitchError = FirmataMessageCommandParse(readBuffer, response)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("FirmataMessage"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent())
	return _child.GetParent(), nil
}

func (m *FirmataMessage) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *FirmataMessage) SerializeParent(writeBuffer utils.WriteBuffer, child IFirmataMessage, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("FirmataMessage"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType := uint8(child.GetMessageType())
	_messageTypeErr := writeBuffer.WriteUint8("messageType", 4, (messageType))

	if _messageTypeErr != nil {
		return errors.Wrap(_messageTypeErr, "Error serializing 'messageType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("FirmataMessage"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *FirmataMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
