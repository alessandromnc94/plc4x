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

#include <stdio.h>
#include <plc4c/spi/evaluation_helper.h>
#include <plc4c/driver_s7_static.h>

#include "s7_var_request_parameter_item.h"

// Code generated by code-generation. DO NOT EDIT.

// Array of discriminator values that match the enum type constants.
// (The order is identical to the enum constants, so we can use the
// enum constant to directly access a given type's discriminator values)
const plc4c_s7_read_write_s7_var_request_parameter_item_discriminator plc4c_s7_read_write_s7_var_request_parameter_item_discriminators[] = {
  {/* plc4c_s7_read_write_s7_var_request_parameter_item_address */
   .itemType = 0x12 }

};

// Function returning the discriminator values for a given type constant.
plc4c_s7_read_write_s7_var_request_parameter_item_discriminator plc4c_s7_read_write_s7_var_request_parameter_item_get_discriminator(plc4c_s7_read_write_s7_var_request_parameter_item_type type) {
  return plc4c_s7_read_write_s7_var_request_parameter_item_discriminators[type];
}

// Create an empty NULL-struct
static const plc4c_s7_read_write_s7_var_request_parameter_item plc4c_s7_read_write_s7_var_request_parameter_item_null_const;

plc4c_s7_read_write_s7_var_request_parameter_item plc4c_s7_read_write_s7_var_request_parameter_item_null() {
  return plc4c_s7_read_write_s7_var_request_parameter_item_null_const;
}


// Parse function.
plc4c_return_code plc4c_s7_read_write_s7_var_request_parameter_item_parse(plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_s7_var_request_parameter_item** _message) {
  uint16_t startPos = plc4c_spi_read_get_pos(readBuffer);
  plc4c_return_code _res = OK;

  // Allocate enough memory to contain this data structure.
  (*_message) = malloc(sizeof(plc4c_s7_read_write_s7_var_request_parameter_item));
  if(*_message == NULL) {
    return NO_MEMORY;
  }
        // Discriminator Field (itemType)

  // Discriminator Field (itemType) (Used as input to a switch field)
  uint8_t itemType = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &itemType);
  if(_res != OK) {
    return _res;
  }

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
if( itemType == 0x12 ) { /* S7VarRequestParameterItemAddress */
    (*_message)->_type = plc4c_s7_read_write_s7_var_request_parameter_item_type_plc4c_s7_read_write_s7_var_request_parameter_item_address;

  // Implicit Field (itemLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  uint8_t itemLength = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &itemLength);
  if(_res != OK) {
    return _res;
  }


  // Simple Field (address)
  plc4c_s7_read_write_s7_address* address;
  _res = plc4c_s7_read_write_s7_address_parse(readBuffer, (void*) &address);
  if(_res != OK) {
    return _res;
  }
  (*_message)->s7_var_request_parameter_item_address_address = address;
  }

  return OK;
}

plc4c_return_code plc4c_s7_read_write_s7_var_request_parameter_item_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_s7_var_request_parameter_item* _message) {
  plc4c_return_code _res = OK;

  // Discriminator Field (itemType)
  plc4c_spi_write_unsigned_byte(writeBuffer, 8, plc4c_s7_read_write_s7_var_request_parameter_item_get_discriminator(_message->_type).itemType);

  // Switch Field (Depending on the current type, serialize the subtype elements)
  switch(_message->_type) {
    case plc4c_s7_read_write_s7_var_request_parameter_item_type_plc4c_s7_read_write_s7_var_request_parameter_item_address: {

  // Implicit Field (itemLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
  _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, plc4c_s7_read_write_s7_address_length_in_bytes(_message->s7_var_request_parameter_item_address_address));
  if(_res != OK) {
    return _res;
  }

  // Simple Field (address)
  _res = plc4c_s7_read_write_s7_address_serialize(writeBuffer, _message->s7_var_request_parameter_item_address_address);
  if(_res != OK) {
    return _res;
  }

      break;
    }
  }

  return OK;
}

uint16_t plc4c_s7_read_write_s7_var_request_parameter_item_length_in_bytes(plc4c_s7_read_write_s7_var_request_parameter_item* _message) {
  return plc4c_s7_read_write_s7_var_request_parameter_item_length_in_bits(_message) / 8;
}

uint16_t plc4c_s7_read_write_s7_var_request_parameter_item_length_in_bits(plc4c_s7_read_write_s7_var_request_parameter_item* _message) {
  uint16_t lengthInBits = 0;

  // Discriminator Field (itemType)
  lengthInBits += 8;

  // Depending of the current type, add the length of sub-type elements ...
  switch(_message->_type) {
    case plc4c_s7_read_write_s7_var_request_parameter_item_type_plc4c_s7_read_write_s7_var_request_parameter_item_address: {

  // Implicit Field (itemLength)
  lengthInBits += 8;


  // Simple field (address)
  lengthInBits += plc4c_s7_read_write_s7_address_length_in_bits(_message->s7_var_request_parameter_item_address_address);

      break;
    }
  }

  return lengthInBits;
}

