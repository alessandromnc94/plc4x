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

// APDUUnconfirmedRequest is the data-structure of this message
type APDUUnconfirmedRequest struct {
	*APDU
	ServiceRequest *BACnetUnconfirmedServiceRequest

	// Arguments.
	ApduLength uint16
}

// IAPDUUnconfirmedRequest is the corresponding interface of APDUUnconfirmedRequest
type IAPDUUnconfirmedRequest interface {
	IAPDU
	// GetServiceRequest returns ServiceRequest (property field)
	GetServiceRequest() *BACnetUnconfirmedServiceRequest
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

func (m *APDUUnconfirmedRequest) GetApduType() uint8 {
	return 0x1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *APDUUnconfirmedRequest) InitializeParent(parent *APDU) {}

func (m *APDUUnconfirmedRequest) GetParent() *APDU {
	return m.APDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *APDUUnconfirmedRequest) GetServiceRequest() *BACnetUnconfirmedServiceRequest {
	return m.ServiceRequest
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAPDUUnconfirmedRequest factory function for APDUUnconfirmedRequest
func NewAPDUUnconfirmedRequest(serviceRequest *BACnetUnconfirmedServiceRequest, apduLength uint16) *APDUUnconfirmedRequest {
	_result := &APDUUnconfirmedRequest{
		ServiceRequest: serviceRequest,
		APDU:           NewAPDU(apduLength),
	}
	_result.Child = _result
	return _result
}

func CastAPDUUnconfirmedRequest(structType interface{}) *APDUUnconfirmedRequest {
	if casted, ok := structType.(APDUUnconfirmedRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*APDUUnconfirmedRequest); ok {
		return casted
	}
	if casted, ok := structType.(APDU); ok {
		return CastAPDUUnconfirmedRequest(casted.Child)
	}
	if casted, ok := structType.(*APDU); ok {
		return CastAPDUUnconfirmedRequest(casted.Child)
	}
	return nil
}

func (m *APDUUnconfirmedRequest) GetTypeName() string {
	return "APDUUnconfirmedRequest"
}

func (m *APDUUnconfirmedRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *APDUUnconfirmedRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (serviceRequest)
	lengthInBits += m.ServiceRequest.GetLengthInBits()

	return lengthInBits
}

func (m *APDUUnconfirmedRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func APDUUnconfirmedRequestParse(readBuffer utils.ReadBuffer, apduLength uint16) (*APDUUnconfirmedRequest, error) {
	if pullErr := readBuffer.PullContext("APDUUnconfirmedRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 4)
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

	// Simple Field (serviceRequest)
	if pullErr := readBuffer.PullContext("serviceRequest"); pullErr != nil {
		return nil, pullErr
	}
	_serviceRequest, _serviceRequestErr := BACnetUnconfirmedServiceRequestParse(readBuffer, uint16(uint16(apduLength)-uint16(uint16(1))))
	if _serviceRequestErr != nil {
		return nil, errors.Wrap(_serviceRequestErr, "Error parsing 'serviceRequest' field")
	}
	serviceRequest := CastBACnetUnconfirmedServiceRequest(_serviceRequest)
	if closeErr := readBuffer.CloseContext("serviceRequest"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("APDUUnconfirmedRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &APDUUnconfirmedRequest{
		ServiceRequest: CastBACnetUnconfirmedServiceRequest(serviceRequest),
		APDU:           &APDU{},
	}
	_child.APDU.Child = _child
	return _child, nil
}

func (m *APDUUnconfirmedRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUUnconfirmedRequest"); pushErr != nil {
			return pushErr
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 4, uint8(0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (serviceRequest)
		if pushErr := writeBuffer.PushContext("serviceRequest"); pushErr != nil {
			return pushErr
		}
		_serviceRequestErr := m.ServiceRequest.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("serviceRequest"); popErr != nil {
			return popErr
		}
		if _serviceRequestErr != nil {
			return errors.Wrap(_serviceRequestErr, "Error serializing 'serviceRequest' field")
		}

		if popErr := writeBuffer.PopContext("APDUUnconfirmedRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *APDUUnconfirmedRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
