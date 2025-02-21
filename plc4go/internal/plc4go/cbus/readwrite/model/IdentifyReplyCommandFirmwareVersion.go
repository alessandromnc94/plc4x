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

// IdentifyReplyCommandFirmwareVersion is the data-structure of this message
type IdentifyReplyCommandFirmwareVersion struct {
	*IdentifyReplyCommand
	FirmwareVersion string
}

// IIdentifyReplyCommandFirmwareVersion is the corresponding interface of IdentifyReplyCommandFirmwareVersion
type IIdentifyReplyCommandFirmwareVersion interface {
	IIdentifyReplyCommand
	// GetFirmwareVersion returns FirmwareVersion (property field)
	GetFirmwareVersion() string
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

func (m *IdentifyReplyCommandFirmwareVersion) GetAttribute() Attribute {
	return Attribute_FirmwareVersion
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *IdentifyReplyCommandFirmwareVersion) InitializeParent(parent *IdentifyReplyCommand) {}

func (m *IdentifyReplyCommandFirmwareVersion) GetParent() *IdentifyReplyCommand {
	return m.IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *IdentifyReplyCommandFirmwareVersion) GetFirmwareVersion() string {
	return m.FirmwareVersion
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandFirmwareVersion factory function for IdentifyReplyCommandFirmwareVersion
func NewIdentifyReplyCommandFirmwareVersion(firmwareVersion string) *IdentifyReplyCommandFirmwareVersion {
	_result := &IdentifyReplyCommandFirmwareVersion{
		FirmwareVersion:      firmwareVersion,
		IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	_result.Child = _result
	return _result
}

func CastIdentifyReplyCommandFirmwareVersion(structType interface{}) *IdentifyReplyCommandFirmwareVersion {
	if casted, ok := structType.(IdentifyReplyCommandFirmwareVersion); ok {
		return &casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandFirmwareVersion); ok {
		return casted
	}
	if casted, ok := structType.(IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandFirmwareVersion(casted.Child)
	}
	if casted, ok := structType.(*IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandFirmwareVersion(casted.Child)
	}
	return nil
}

func (m *IdentifyReplyCommandFirmwareVersion) GetTypeName() string {
	return "IdentifyReplyCommandFirmwareVersion"
}

func (m *IdentifyReplyCommandFirmwareVersion) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *IdentifyReplyCommandFirmwareVersion) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (firmwareVersion)
	lengthInBits += 64

	return lengthInBits
}

func (m *IdentifyReplyCommandFirmwareVersion) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandFirmwareVersionParse(readBuffer utils.ReadBuffer, attribute Attribute) (*IdentifyReplyCommandFirmwareVersion, error) {
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandFirmwareVersion"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (firmwareVersion)
	_firmwareVersion, _firmwareVersionErr := readBuffer.ReadString("firmwareVersion", uint32(64))
	if _firmwareVersionErr != nil {
		return nil, errors.Wrap(_firmwareVersionErr, "Error parsing 'firmwareVersion' field")
	}
	firmwareVersion := _firmwareVersion

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandFirmwareVersion"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &IdentifyReplyCommandFirmwareVersion{
		FirmwareVersion:      firmwareVersion,
		IdentifyReplyCommand: &IdentifyReplyCommand{},
	}
	_child.IdentifyReplyCommand.Child = _child
	return _child, nil
}

func (m *IdentifyReplyCommandFirmwareVersion) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandFirmwareVersion"); pushErr != nil {
			return pushErr
		}

		// Simple Field (firmwareVersion)
		firmwareVersion := string(m.FirmwareVersion)
		_firmwareVersionErr := writeBuffer.WriteString("firmwareVersion", uint32(64), "UTF-8", (firmwareVersion))
		if _firmwareVersionErr != nil {
			return errors.Wrap(_firmwareVersionErr, "Error serializing 'firmwareVersion' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandFirmwareVersion"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *IdentifyReplyCommandFirmwareVersion) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
