#[[
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
]]

include_directories("include" "../../api/include" "../../spi/include"
    "../../generated-sources/plc4x/include")

# Add the generated sources
file(GLOB generatedSources "../../generated-sources/plc4x/src/*.c")

add_library(plc4c-driver-plc4x
        src/driver_plc4x.c
        src/driver_plc4x_encode_decode.c
        src/driver_plc4x_packets.c
        src/driver_plc4x_sm_connect.c
        src/driver_plc4x_sm_disconnect.c
        src/driver_plc4x_sm_read.c
        src/driver_plc4x_sm_write.c
        src/driver_plc4x_static.c
        ${generatedSources})

target_link_libraries(plc4c-driver-plc4x
        plc4c-spi)

if (BUILD_PHASE STREQUAL test-compile)
    add_executable(plc4c-driver-plc4x-test
            test/plc4x_test.c
            ../../spi/test/utils/test_utils.c
            ../../generated-sources/plc4x/test/ParserSerializerTestsuite.c
            )

    IF (NOT WIN32)
        target_link_libraries(plc4c-driver-plc4x-test
                plc4c-driver-plc4x
                unity
                m
                )
    ELSE()
        target_link_libraries(plc4c-driver-plc4x-test
                plc4c-driver-plc4x
                unity
                )
    ENDIF()
    add_test(NAME plc4c-driver-plc4x-test COMMAND plc4c-driver-plc4x-test)
endif ()