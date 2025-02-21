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

// ProjectInstallationIdentifier is the data-structure of this message
type ProjectInstallationIdentifier struct {
	ProjectNumber      uint8
	InstallationNumber uint8
}

// IProjectInstallationIdentifier is the corresponding interface of ProjectInstallationIdentifier
type IProjectInstallationIdentifier interface {
	// GetProjectNumber returns ProjectNumber (property field)
	GetProjectNumber() uint8
	// GetInstallationNumber returns InstallationNumber (property field)
	GetInstallationNumber() uint8
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

func (m *ProjectInstallationIdentifier) GetProjectNumber() uint8 {
	return m.ProjectNumber
}

func (m *ProjectInstallationIdentifier) GetInstallationNumber() uint8 {
	return m.InstallationNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewProjectInstallationIdentifier factory function for ProjectInstallationIdentifier
func NewProjectInstallationIdentifier(projectNumber uint8, installationNumber uint8) *ProjectInstallationIdentifier {
	return &ProjectInstallationIdentifier{ProjectNumber: projectNumber, InstallationNumber: installationNumber}
}

func CastProjectInstallationIdentifier(structType interface{}) *ProjectInstallationIdentifier {
	if casted, ok := structType.(ProjectInstallationIdentifier); ok {
		return &casted
	}
	if casted, ok := structType.(*ProjectInstallationIdentifier); ok {
		return casted
	}
	return nil
}

func (m *ProjectInstallationIdentifier) GetTypeName() string {
	return "ProjectInstallationIdentifier"
}

func (m *ProjectInstallationIdentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ProjectInstallationIdentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (projectNumber)
	lengthInBits += 8

	// Simple field (installationNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *ProjectInstallationIdentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ProjectInstallationIdentifierParse(readBuffer utils.ReadBuffer) (*ProjectInstallationIdentifier, error) {
	if pullErr := readBuffer.PullContext("ProjectInstallationIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (projectNumber)
	_projectNumber, _projectNumberErr := readBuffer.ReadUint8("projectNumber", 8)
	if _projectNumberErr != nil {
		return nil, errors.Wrap(_projectNumberErr, "Error parsing 'projectNumber' field")
	}
	projectNumber := _projectNumber

	// Simple Field (installationNumber)
	_installationNumber, _installationNumberErr := readBuffer.ReadUint8("installationNumber", 8)
	if _installationNumberErr != nil {
		return nil, errors.Wrap(_installationNumberErr, "Error parsing 'installationNumber' field")
	}
	installationNumber := _installationNumber

	if closeErr := readBuffer.CloseContext("ProjectInstallationIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewProjectInstallationIdentifier(projectNumber, installationNumber), nil
}

func (m *ProjectInstallationIdentifier) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("ProjectInstallationIdentifier"); pushErr != nil {
		return pushErr
	}

	// Simple Field (projectNumber)
	projectNumber := uint8(m.ProjectNumber)
	_projectNumberErr := writeBuffer.WriteUint8("projectNumber", 8, (projectNumber))
	if _projectNumberErr != nil {
		return errors.Wrap(_projectNumberErr, "Error serializing 'projectNumber' field")
	}

	// Simple Field (installationNumber)
	installationNumber := uint8(m.InstallationNumber)
	_installationNumberErr := writeBuffer.WriteUint8("installationNumber", 8, (installationNumber))
	if _installationNumberErr != nil {
		return errors.Wrap(_installationNumberErr, "Error serializing 'installationNumber' field")
	}

	if popErr := writeBuffer.PopContext("ProjectInstallationIdentifier"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *ProjectInstallationIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
