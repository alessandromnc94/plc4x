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
)

// Code generated by code-generation. DO NOT EDIT.

// IdentifyReplyCommandNetworkTerminalLevels is the data-structure of this message
type IdentifyReplyCommandNetworkTerminalLevels struct {
	*IdentifyReplyCommand
}

// IIdentifyReplyCommandNetworkTerminalLevels is the corresponding interface of IdentifyReplyCommandNetworkTerminalLevels
type IIdentifyReplyCommandNetworkTerminalLevels interface {
	IIdentifyReplyCommand
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

func (m *IdentifyReplyCommandNetworkTerminalLevels) GetAttribute() Attribute {
	return Attribute_NetworkTerminalLevels
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *IdentifyReplyCommandNetworkTerminalLevels) InitializeParent(parent *IdentifyReplyCommand) {}

func (m *IdentifyReplyCommandNetworkTerminalLevels) GetParent() *IdentifyReplyCommand {
	return m.IdentifyReplyCommand
}

// NewIdentifyReplyCommandNetworkTerminalLevels factory function for IdentifyReplyCommandNetworkTerminalLevels
func NewIdentifyReplyCommandNetworkTerminalLevels() *IdentifyReplyCommandNetworkTerminalLevels {
	_result := &IdentifyReplyCommandNetworkTerminalLevels{
		IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	_result.Child = _result
	return _result
}

func CastIdentifyReplyCommandNetworkTerminalLevels(structType interface{}) *IdentifyReplyCommandNetworkTerminalLevels {
	if casted, ok := structType.(IdentifyReplyCommandNetworkTerminalLevels); ok {
		return &casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandNetworkTerminalLevels); ok {
		return casted
	}
	if casted, ok := structType.(IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandNetworkTerminalLevels(casted.Child)
	}
	if casted, ok := structType.(*IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandNetworkTerminalLevels(casted.Child)
	}
	return nil
}

func (m *IdentifyReplyCommandNetworkTerminalLevels) GetTypeName() string {
	return "IdentifyReplyCommandNetworkTerminalLevels"
}

func (m *IdentifyReplyCommandNetworkTerminalLevels) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *IdentifyReplyCommandNetworkTerminalLevels) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *IdentifyReplyCommandNetworkTerminalLevels) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandNetworkTerminalLevelsParse(readBuffer utils.ReadBuffer, attribute Attribute) (*IdentifyReplyCommandNetworkTerminalLevels, error) {
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandNetworkTerminalLevels"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandNetworkTerminalLevels"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &IdentifyReplyCommandNetworkTerminalLevels{
		IdentifyReplyCommand: &IdentifyReplyCommand{},
	}
	_child.IdentifyReplyCommand.Child = _child
	return _child, nil
}

func (m *IdentifyReplyCommandNetworkTerminalLevels) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandNetworkTerminalLevels"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandNetworkTerminalLevels"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *IdentifyReplyCommandNetworkTerminalLevels) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
