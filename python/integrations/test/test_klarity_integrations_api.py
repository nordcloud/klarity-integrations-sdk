"""
    Klarity Integrations

    REST API for managing Estate Records using Klarity Integrations. You can enrich your estate by creating new kinds of estate records or extending existing ones. Before making use of the API, you must first register your External Integration in Klarity, which provides you with the required authentication credentials. Then, you use those credentials to obtain a Token that allows you to make authorized calls to Klarity’s REST API for External Integration.  # noqa: E501

    The version of the OpenAPI document: 0.0.4
    Contact: products@nordcloud.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import klarity_integrations
from klarity_integrations.api.klarity_integrations_api import KlarityIntegrationsApi  # noqa: E501


class TestKlarityIntegrationsApi(unittest.TestCase):
    """KlarityIntegrationsApi unit test stubs"""

    def setUp(self):
        self.api = KlarityIntegrationsApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_v1_enrichments_estate_records_post(self):
        """Test case for v1_enrichments_estate_records_post

        Enrich Klarity estate records  # noqa: E501
        """
        pass

    def test_v1_estate_records_delete(self):
        """Test case for v1_estate_records_delete

        Delete Klarity estate records  # noqa: E501
        """
        pass

    def test_v1_estate_records_post(self):
        """Test case for v1_estate_records_post

        Manage Klarity estate records  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
