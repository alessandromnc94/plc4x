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
#ifndef PLC4C_DRIVER_PLC4X_ENCODE_DECODE_H_
#define PLC4C_DRIVER_PLC4X_ENCODE_DECODE_H_
#ifdef __cplusplus
extern "C" {
#endif

#include <stdlib.h>

#include "plc4c/driver_plc4x.h"

plc4c_return_code plc4c_driver_plc4x_encode_address(
    char* address, void** item);

#ifdef __cplusplus
}
#endif
#endif  // PLC4C_DRIVER_PLC4X_ENCODE_DECODE_H_