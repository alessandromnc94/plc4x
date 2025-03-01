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

// NPDUControl is the data-structure of this message
type NPDUControl struct {
	MessageTypeFieldPresent bool
	DestinationSpecified    bool
	SourceSpecified         bool
	ExpectingReply          bool
	NetworkPriority         NPDUNetworkPriority
}

// INPDUControl is the corresponding interface of NPDUControl
type INPDUControl interface {
	// GetMessageTypeFieldPresent returns MessageTypeFieldPresent (property field)
	GetMessageTypeFieldPresent() bool
	// GetDestinationSpecified returns DestinationSpecified (property field)
	GetDestinationSpecified() bool
	// GetSourceSpecified returns SourceSpecified (property field)
	GetSourceSpecified() bool
	// GetExpectingReply returns ExpectingReply (property field)
	GetExpectingReply() bool
	// GetNetworkPriority returns NetworkPriority (property field)
	GetNetworkPriority() NPDUNetworkPriority
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *NPDUControl) GetMessageTypeFieldPresent() bool {
	return m.MessageTypeFieldPresent
}

func (m *NPDUControl) GetDestinationSpecified() bool {
	return m.DestinationSpecified
}

func (m *NPDUControl) GetSourceSpecified() bool {
	return m.SourceSpecified
}

func (m *NPDUControl) GetExpectingReply() bool {
	return m.ExpectingReply
}

func (m *NPDUControl) GetNetworkPriority() NPDUNetworkPriority {
	return m.NetworkPriority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNPDUControl factory function for NPDUControl
func NewNPDUControl(messageTypeFieldPresent bool, destinationSpecified bool, sourceSpecified bool, expectingReply bool, networkPriority NPDUNetworkPriority) *NPDUControl {
	return &NPDUControl{MessageTypeFieldPresent: messageTypeFieldPresent, DestinationSpecified: destinationSpecified, SourceSpecified: sourceSpecified, ExpectingReply: expectingReply, NetworkPriority: networkPriority}
}

func CastNPDUControl(structType interface{}) *NPDUControl {
	if casted, ok := structType.(NPDUControl); ok {
		return &casted
	}
	if casted, ok := structType.(*NPDUControl); ok {
		return casted
	}
	return nil
}

func (m *NPDUControl) GetTypeName() string {
	return "NPDUControl"
}

func (m *NPDUControl) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *NPDUControl) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (messageTypeFieldPresent)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (destinationSpecified)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (sourceSpecified)
	lengthInBits += 1

	// Simple field (expectingReply)
	lengthInBits += 1

	// Simple field (networkPriority)
	lengthInBits += 2

	return lengthInBits
}

func (m *NPDUControl) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NPDUControlParse(readBuffer utils.ReadBuffer) (*NPDUControl, error) {
	if pullErr := readBuffer.PullContext("NPDUControl"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (messageTypeFieldPresent)
	_messageTypeFieldPresent, _messageTypeFieldPresentErr := readBuffer.ReadBit("messageTypeFieldPresent")
	if _messageTypeFieldPresentErr != nil {
		return nil, errors.Wrap(_messageTypeFieldPresentErr, "Error parsing 'messageTypeFieldPresent' field")
	}
	messageTypeFieldPresent := _messageTypeFieldPresent

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

	// Simple Field (destinationSpecified)
	_destinationSpecified, _destinationSpecifiedErr := readBuffer.ReadBit("destinationSpecified")
	if _destinationSpecifiedErr != nil {
		return nil, errors.Wrap(_destinationSpecifiedErr, "Error parsing 'destinationSpecified' field")
	}
	destinationSpecified := _destinationSpecified

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

	// Simple Field (sourceSpecified)
	_sourceSpecified, _sourceSpecifiedErr := readBuffer.ReadBit("sourceSpecified")
	if _sourceSpecifiedErr != nil {
		return nil, errors.Wrap(_sourceSpecifiedErr, "Error parsing 'sourceSpecified' field")
	}
	sourceSpecified := _sourceSpecified

	// Simple Field (expectingReply)
	_expectingReply, _expectingReplyErr := readBuffer.ReadBit("expectingReply")
	if _expectingReplyErr != nil {
		return nil, errors.Wrap(_expectingReplyErr, "Error parsing 'expectingReply' field")
	}
	expectingReply := _expectingReply

	// Simple Field (networkPriority)
	if pullErr := readBuffer.PullContext("networkPriority"); pullErr != nil {
		return nil, pullErr
	}
	_networkPriority, _networkPriorityErr := NPDUNetworkPriorityParse(readBuffer)
	if _networkPriorityErr != nil {
		return nil, errors.Wrap(_networkPriorityErr, "Error parsing 'networkPriority' field")
	}
	networkPriority := _networkPriority
	if closeErr := readBuffer.CloseContext("networkPriority"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("NPDUControl"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewNPDUControl(messageTypeFieldPresent, destinationSpecified, sourceSpecified, expectingReply, networkPriority), nil
}

func (m *NPDUControl) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("NPDUControl"); pushErr != nil {
		return pushErr
	}

	// Simple Field (messageTypeFieldPresent)
	messageTypeFieldPresent := bool(m.MessageTypeFieldPresent)
	_messageTypeFieldPresentErr := writeBuffer.WriteBit("messageTypeFieldPresent", (messageTypeFieldPresent))
	if _messageTypeFieldPresentErr != nil {
		return errors.Wrap(_messageTypeFieldPresentErr, "Error serializing 'messageTypeFieldPresent' field")
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteUint8("reserved", 1, uint8(0))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (destinationSpecified)
	destinationSpecified := bool(m.DestinationSpecified)
	_destinationSpecifiedErr := writeBuffer.WriteBit("destinationSpecified", (destinationSpecified))
	if _destinationSpecifiedErr != nil {
		return errors.Wrap(_destinationSpecifiedErr, "Error serializing 'destinationSpecified' field")
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteUint8("reserved", 1, uint8(0))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (sourceSpecified)
	sourceSpecified := bool(m.SourceSpecified)
	_sourceSpecifiedErr := writeBuffer.WriteBit("sourceSpecified", (sourceSpecified))
	if _sourceSpecifiedErr != nil {
		return errors.Wrap(_sourceSpecifiedErr, "Error serializing 'sourceSpecified' field")
	}

	// Simple Field (expectingReply)
	expectingReply := bool(m.ExpectingReply)
	_expectingReplyErr := writeBuffer.WriteBit("expectingReply", (expectingReply))
	if _expectingReplyErr != nil {
		return errors.Wrap(_expectingReplyErr, "Error serializing 'expectingReply' field")
	}

	// Simple Field (networkPriority)
	if pushErr := writeBuffer.PushContext("networkPriority"); pushErr != nil {
		return pushErr
	}
	_networkPriorityErr := m.NetworkPriority.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("networkPriority"); popErr != nil {
		return popErr
	}
	if _networkPriorityErr != nil {
		return errors.Wrap(_networkPriorityErr, "Error serializing 'networkPriority' field")
	}

	if popErr := writeBuffer.PopContext("NPDUControl"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *NPDUControl) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
