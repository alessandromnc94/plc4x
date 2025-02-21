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

// BACnetConstructedDataLifeSafetyZone is the data-structure of this message
type BACnetConstructedDataLifeSafetyZone struct {
	*BACnetConstructedData
	Zones []*BACnetContextTagObjectIdentifier

	// Arguments.
	TagNumber                  uint8
	PropertyIdentifierArgument BACnetContextTagPropertyIdentifier
}

// IBACnetConstructedDataLifeSafetyZone is the corresponding interface of BACnetConstructedDataLifeSafetyZone
type IBACnetConstructedDataLifeSafetyZone interface {
	IBACnetConstructedData
	// GetZones returns Zones (property field)
	GetZones() []*BACnetContextTagObjectIdentifier
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

func (m *BACnetConstructedDataLifeSafetyZone) GetObjectType() BACnetObjectType {
	return BACnetObjectType_LIFE_SAFETY_ZONE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLifeSafetyZone) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLifeSafetyZone) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLifeSafetyZone) GetZones() []*BACnetContextTagObjectIdentifier {
	return m.Zones
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLifeSafetyZone factory function for BACnetConstructedDataLifeSafetyZone
func NewBACnetConstructedDataLifeSafetyZone(zones []*BACnetContextTagObjectIdentifier, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8, propertyIdentifierArgument BACnetContextTagPropertyIdentifier) *BACnetConstructedDataLifeSafetyZone {
	_result := &BACnetConstructedDataLifeSafetyZone{
		Zones:                 zones,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber, propertyIdentifierArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLifeSafetyZone(structType interface{}) *BACnetConstructedDataLifeSafetyZone {
	if casted, ok := structType.(BACnetConstructedDataLifeSafetyZone); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLifeSafetyZone); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLifeSafetyZone(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLifeSafetyZone(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLifeSafetyZone) GetTypeName() string {
	return "BACnetConstructedDataLifeSafetyZone"
}

func (m *BACnetConstructedDataLifeSafetyZone) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLifeSafetyZone) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Zones) > 0 {
		for _, element := range m.Zones {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataLifeSafetyZone) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLifeSafetyZoneParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType, propertyIdentifierArgument *BACnetContextTagPropertyIdentifier) (*BACnetConstructedDataLifeSafetyZone, error) {
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLifeSafetyZone"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Array field (zones)
	if pullErr := readBuffer.PullContext("zones", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	zones := make([]*BACnetContextTagObjectIdentifier, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_BACNET_OBJECT_IDENTIFIER)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'zones' field")
			}
			zones = append(zones, CastBACnetContextTagObjectIdentifier(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("zones", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLifeSafetyZone"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLifeSafetyZone{
		Zones:                 zones,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLifeSafetyZone) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLifeSafetyZone"); pushErr != nil {
			return pushErr
		}

		// Array Field (zones)
		if m.Zones != nil {
			if pushErr := writeBuffer.PushContext("zones", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Zones {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'zones' field")
				}
			}
			if popErr := writeBuffer.PopContext("zones", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLifeSafetyZone"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLifeSafetyZone) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
