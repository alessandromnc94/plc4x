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

// BACnetUnconfirmedServiceRequest is the data-structure of this message
type BACnetUnconfirmedServiceRequest struct {

	// Arguments.
	Len   uint16
	Child IBACnetUnconfirmedServiceRequestChild
}

// IBACnetUnconfirmedServiceRequest is the corresponding interface of BACnetUnconfirmedServiceRequest
type IBACnetUnconfirmedServiceRequest interface {
	// GetServiceChoice returns ServiceChoice (discriminator field)
	GetServiceChoice() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetUnconfirmedServiceRequestParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetUnconfirmedServiceRequest, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetUnconfirmedServiceRequestChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetUnconfirmedServiceRequest)
	GetParent() *BACnetUnconfirmedServiceRequest

	GetTypeName() string
	IBACnetUnconfirmedServiceRequest
}

// NewBACnetUnconfirmedServiceRequest factory function for BACnetUnconfirmedServiceRequest
func NewBACnetUnconfirmedServiceRequest(len uint16) *BACnetUnconfirmedServiceRequest {
	return &BACnetUnconfirmedServiceRequest{Len: len}
}

func CastBACnetUnconfirmedServiceRequest(structType interface{}) *BACnetUnconfirmedServiceRequest {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequest); ok {
		return casted
	}
	if casted, ok := structType.(IBACnetUnconfirmedServiceRequestChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *BACnetUnconfirmedServiceRequest) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequest"
}

func (m *BACnetUnconfirmedServiceRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetUnconfirmedServiceRequest) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetUnconfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice, _serviceChoiceErr := readBuffer.ReadUint8("serviceChoice", 8)
	if _serviceChoiceErr != nil {
		return nil, errors.Wrap(_serviceChoiceErr, "Error parsing 'serviceChoice' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetUnconfirmedServiceRequestChild interface {
		InitializeParent(*BACnetUnconfirmedServiceRequest)
		GetParent() *BACnetUnconfirmedServiceRequest
	}
	var _child BACnetUnconfirmedServiceRequestChild
	var typeSwitchError error
	switch {
	case serviceChoice == 0x00: // BACnetUnconfirmedServiceRequestIAm
		_child, typeSwitchError = BACnetUnconfirmedServiceRequestIAmParse(readBuffer, len)
	case serviceChoice == 0x01: // BACnetUnconfirmedServiceRequestIHave
		_child, typeSwitchError = BACnetUnconfirmedServiceRequestIHaveParse(readBuffer, len)
	case serviceChoice == 0x02: // BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification
		_child, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationParse(readBuffer, len)
	case serviceChoice == 0x03: // BACnetUnconfirmedServiceRequestUnconfirmedEventNotification
		_child, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedEventNotificationParse(readBuffer, len)
	case serviceChoice == 0x04: // BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer
		_child, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransferParse(readBuffer, len)
	case serviceChoice == 0x05: // BACnetUnconfirmedServiceRequestUnconfirmedTextMessage
		_child, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedTextMessageParse(readBuffer, len)
	case serviceChoice == 0x06: // BACnetUnconfirmedServiceRequestTimeSynchronization
		_child, typeSwitchError = BACnetUnconfirmedServiceRequestTimeSynchronizationParse(readBuffer, len)
	case serviceChoice == 0x07: // BACnetUnconfirmedServiceRequestWhoHas
		_child, typeSwitchError = BACnetUnconfirmedServiceRequestWhoHasParse(readBuffer, len)
	case serviceChoice == 0x08: // BACnetUnconfirmedServiceRequestWhoIs
		_child, typeSwitchError = BACnetUnconfirmedServiceRequestWhoIsParse(readBuffer, len)
	case serviceChoice == 0x09: // BACnetUnconfirmedServiceRequestUTCTimeSynchronization
		_child, typeSwitchError = BACnetUnconfirmedServiceRequestUTCTimeSynchronizationParse(readBuffer, len)
	case serviceChoice == 0x0A: // BACnetUnconfirmedServiceRequestWriteGroup
		_child, typeSwitchError = BACnetUnconfirmedServiceRequestWriteGroupParse(readBuffer, len)
	case serviceChoice == 0x0B: // BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
		_child, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleParse(readBuffer, len)
	case true: // BACnetUnconfirmedServiceRequestUnconfirmedUnknown
		_child, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedUnknownParse(readBuffer, len)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent())
	return _child.GetParent(), nil
}

func (m *BACnetUnconfirmedServiceRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetUnconfirmedServiceRequest) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetUnconfirmedServiceRequest, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequest"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice := uint8(child.GetServiceChoice())
	_serviceChoiceErr := writeBuffer.WriteUint8("serviceChoice", 8, (serviceChoice))

	if _serviceChoiceErr != nil {
		return errors.Wrap(_serviceChoiceErr, "Error serializing 'serviceChoice' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequest"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetUnconfirmedServiceRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
