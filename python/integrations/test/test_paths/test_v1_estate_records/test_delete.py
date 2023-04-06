# coding: utf-8

"""


    Generated by: https://openapi-generator.tech
"""

import unittest
from unittest.mock import patch

import urllib3

import klarity_integrations
from klarity_integrations.paths.v1_estate_records import delete  # noqa: E501
from klarity_integrations import configuration, schemas, api_client

from .. import ApiTestMixin


class TestV1EstateRecords(ApiTestMixin, unittest.TestCase):
    """
    V1EstateRecords unit test stubs
        Delete Klarity estate records  # noqa: E501
    """
    _configuration = configuration.Configuration()

    def setUp(self):
        used_api_client = api_client.ApiClient(configuration=self._configuration)
        self.api = delete.ApiFordelete(api_client=used_api_client)  # noqa: E501

    def tearDown(self):
        pass

    response_status = 202




if __name__ == '__main__':
    unittest.main()
