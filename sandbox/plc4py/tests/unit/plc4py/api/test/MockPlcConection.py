#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#
import asyncio
from dataclasses import dataclass
from typing import Awaitable

from plc4py.api.PlcConnection import PlcConnection
from plc4py.api.messages.PlcRequest import ReadRequestBuilder, PlcReadRequest
from plc4py.api.messages.PlcResponse import PlcReadResponse, PlcResponse
from plc4py.api.value.PlcValue import PlcResponseCode
from tests.unit.plc4py.api.test.MockReadRequestBuilder import MockReadRequestBuilder


@dataclass
class MockPlcConnection(PlcConnection):
    _is_connected: bool = False

    def connect(self):
        """
        Connect the Mock PLC connection
        :return:
        """
        self._is_connected = True

    def is_connected(self) -> bool:
        """
        Return the current status of the Mock PLC Connection
        :return bool: True is connected
        """
        return self._is_connected

    def close(self) -> None:
        """
        Closes the connection to the remote PLC.
        :return:
        """
        self._is_connected = False

    def read_request_builder(self) -> ReadRequestBuilder:
        """
        :return: read request builder.
        """
        return MockReadRequestBuilder()

    def _default_failed_request(
        self, code: PlcResponseCode
    ) -> Awaitable[PlcReadResponse]:
        """
        Returns a default PlcResponse, mainly used in case of a failed request
        :param code: The response code to return
        :return: The PlcResponse
        """
        loop = asyncio.get_running_loop()
        fut = loop.create_future()
        fut.set_result(PlcResponse(code))
        return fut

    def execute(self, request: PlcReadRequest) -> Awaitable[PlcReadResponse]:
        """
        Executes a PlcRequest as long as it's already connected
        :param PlcRequest: Plc Request to execute
        :return: The response from the Plc/Device
        """
        if not self.is_connected():
            return self._default_failed_request(PlcResponseCode.NOT_CONNECTED)

        return self._default_failed_request(PlcResponseCode.NOT_CONNECTED)
