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

// SysexCommandSysexRealtime is the data-structure of this message
type SysexCommandSysexRealtime struct {
	*SysexCommand
}

// ISysexCommandSysexRealtime is the corresponding interface of SysexCommandSysexRealtime
type ISysexCommandSysexRealtime interface {
	ISysexCommand
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

func (m *SysexCommandSysexRealtime) GetCommandType() uint8 {
	return 0x7F
}

func (m *SysexCommandSysexRealtime) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *SysexCommandSysexRealtime) InitializeParent(parent *SysexCommand) {}

func (m *SysexCommandSysexRealtime) GetParent() *SysexCommand {
	return m.SysexCommand
}

// NewSysexCommandSysexRealtime factory function for SysexCommandSysexRealtime
func NewSysexCommandSysexRealtime() *SysexCommandSysexRealtime {
	_result := &SysexCommandSysexRealtime{
		SysexCommand: NewSysexCommand(),
	}
	_result.Child = _result
	return _result
}

func CastSysexCommandSysexRealtime(structType interface{}) *SysexCommandSysexRealtime {
	if casted, ok := structType.(SysexCommandSysexRealtime); ok {
		return &casted
	}
	if casted, ok := structType.(*SysexCommandSysexRealtime); ok {
		return casted
	}
	if casted, ok := structType.(SysexCommand); ok {
		return CastSysexCommandSysexRealtime(casted.Child)
	}
	if casted, ok := structType.(*SysexCommand); ok {
		return CastSysexCommandSysexRealtime(casted.Child)
	}
	return nil
}

func (m *SysexCommandSysexRealtime) GetTypeName() string {
	return "SysexCommandSysexRealtime"
}

func (m *SysexCommandSysexRealtime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *SysexCommandSysexRealtime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *SysexCommandSysexRealtime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandSysexRealtimeParse(readBuffer utils.ReadBuffer, response bool) (*SysexCommandSysexRealtime, error) {
	if pullErr := readBuffer.PullContext("SysexCommandSysexRealtime"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandSysexRealtime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SysexCommandSysexRealtime{
		SysexCommand: &SysexCommand{},
	}
	_child.SysexCommand.Child = _child
	return _child, nil
}

func (m *SysexCommandSysexRealtime) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandSysexRealtime"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("SysexCommandSysexRealtime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *SysexCommandSysexRealtime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
