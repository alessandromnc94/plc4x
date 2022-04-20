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

#include "szl_data_tree_item.h"

// Code generated by code-generation. DO NOT EDIT.


// Parse function.
plc4c_return_code plc4c_s7_read_write_szl_data_tree_item_parse(plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_szl_data_tree_item** _message) {
  uint16_t startPos = plc4c_spi_read_get_pos(readBuffer);
  plc4c_return_code _res = OK;

  // Allocate enough memory to contain this data structure.
  (*_message) = malloc(sizeof(plc4c_s7_read_write_szl_data_tree_item));
  if(*_message == NULL) {
    return NO_MEMORY;
  }

  // Simple Field (itemIndex)
  uint16_t itemIndex = 0;
  _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &itemIndex);
  if(_res != OK) {
    return _res;
  }
  (*_message)->item_index = itemIndex;

  // Array field (mlfb)
  plc4c_list* mlfb = NULL;
  plc4c_utils_list_create(&mlfb);
  if(mlfb == NULL) {
    return NO_MEMORY;
  }
  {
    // Count array
    uint16_t itemCount = (uint16_t) 20;
    for(int curItem = 0; curItem < itemCount; curItem++) {
      
      char* _value = malloc(sizeof(char));
      _res = plc4c_spi_read_char(readBuffer, (char*) _value);
      if(_res != OK) {
        return _res;
      }
      plc4c_utils_list_insert_head_value(mlfb, _value);
    }
  }
  (*_message)->mlfb = mlfb;

  // Simple Field (moduleTypeId)
  uint16_t moduleTypeId = 0;
  _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &moduleTypeId);
  if(_res != OK) {
    return _res;
  }
  (*_message)->module_type_id = moduleTypeId;

  // Simple Field (ausbg)
  uint16_t ausbg = 0;
  _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &ausbg);
  if(_res != OK) {
    return _res;
  }
  (*_message)->ausbg = ausbg;

  // Simple Field (ausbe)
  uint16_t ausbe = 0;
  _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &ausbe);
  if(_res != OK) {
    return _res;
  }
  (*_message)->ausbe = ausbe;

  return OK;
}

plc4c_return_code plc4c_s7_read_write_szl_data_tree_item_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_szl_data_tree_item* _message) {
  plc4c_return_code _res = OK;

  // Simple Field (itemIndex)
  _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->item_index);
  if(_res != OK) {
    return _res;
  }

  // Array field (mlfb)
  {
    uint8_t itemCount = plc4c_utils_list_size(_message->mlfb);
    for(int curItem = 0; curItem < itemCount; curItem++) {

      char* _value = (char*) plc4c_utils_list_get_value(_message->mlfb, curItem);
      plc4c_spi_write_char(writeBuffer, *_value);
    }
  }

  // Simple Field (moduleTypeId)
  _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->module_type_id);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (ausbg)
  _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->ausbg);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (ausbe)
  _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->ausbe);
  if(_res != OK) {
    return _res;
  }

  return OK;
}

uint16_t plc4c_s7_read_write_szl_data_tree_item_length_in_bytes(plc4c_s7_read_write_szl_data_tree_item* _message) {
  return plc4c_s7_read_write_szl_data_tree_item_length_in_bits(_message) / 8;
}

uint16_t plc4c_s7_read_write_szl_data_tree_item_length_in_bits(plc4c_s7_read_write_szl_data_tree_item* _message) {
  uint16_t lengthInBits = 0;

  // Simple field (itemIndex)
  lengthInBits += 16;

  // Array field
  lengthInBits += 8 * plc4c_utils_list_size(_message->mlfb);

  // Simple field (moduleTypeId)
  lengthInBits += 16;

  // Simple field (ausbg)
  lengthInBits += 16;

  // Simple field (ausbe)
  lengthInBits += 16;

  return lengthInBits;
}

