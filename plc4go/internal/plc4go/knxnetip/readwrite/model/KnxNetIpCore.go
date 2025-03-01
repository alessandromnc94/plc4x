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

// KnxNetIpCore is the data-structure of this message
type KnxNetIpCore struct {
	*ServiceId
	Version uint8
}

// IKnxNetIpCore is the corresponding interface of KnxNetIpCore
type IKnxNetIpCore interface {
	IServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
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

func (m *KnxNetIpCore) GetServiceType() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *KnxNetIpCore) InitializeParent(parent *ServiceId) {}

func (m *KnxNetIpCore) GetParent() *ServiceId {
	return m.ServiceId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *KnxNetIpCore) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxNetIpCore factory function for KnxNetIpCore
func NewKnxNetIpCore(version uint8) *KnxNetIpCore {
	_result := &KnxNetIpCore{
		Version:   version,
		ServiceId: NewServiceId(),
	}
	_result.Child = _result
	return _result
}

func CastKnxNetIpCore(structType interface{}) *KnxNetIpCore {
	if casted, ok := structType.(KnxNetIpCore); ok {
		return &casted
	}
	if casted, ok := structType.(*KnxNetIpCore); ok {
		return casted
	}
	if casted, ok := structType.(ServiceId); ok {
		return CastKnxNetIpCore(casted.Child)
	}
	if casted, ok := structType.(*ServiceId); ok {
		return CastKnxNetIpCore(casted.Child)
	}
	return nil
}

func (m *KnxNetIpCore) GetTypeName() string {
	return "KnxNetIpCore"
}

func (m *KnxNetIpCore) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *KnxNetIpCore) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *KnxNetIpCore) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxNetIpCoreParse(readBuffer utils.ReadBuffer) (*KnxNetIpCore, error) {
	if pullErr := readBuffer.PullContext("KnxNetIpCore"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (version)
	_version, _versionErr := readBuffer.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field")
	}
	version := _version

	if closeErr := readBuffer.CloseContext("KnxNetIpCore"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &KnxNetIpCore{
		Version:   version,
		ServiceId: &ServiceId{},
	}
	_child.ServiceId.Child = _child
	return _child, nil
}

func (m *KnxNetIpCore) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetIpCore"); pushErr != nil {
			return pushErr
		}

		// Simple Field (version)
		version := uint8(m.Version)
		_versionErr := writeBuffer.WriteUint8("version", 8, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetIpCore"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *KnxNetIpCore) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
