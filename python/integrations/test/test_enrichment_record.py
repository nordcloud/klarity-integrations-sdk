"""
    Klarity Integrations

    REST API for managing Estate Records using Klarity Integrations. You can enrich your estate by creating new kinds of estate records or extending existing ones. Before making use of the API, you must first register your External Integration in Klarity, which provides you with the required authentication credentials. Then, you use those credentials to obtain a Token that allows you to make authorized calls to Klarity’s REST API for External Integration.  # noqa: E501

    The version of the OpenAPI document: 0.0.4
    Contact: products@nordcloud.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import klarity_integrations
from klarity_integrations.model.enrichment_record import EnrichmentRecord


class TestEnrichmentRecord(unittest.TestCase):
    """EnrichmentRecord unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testEnrichmentRecord(self):
        """Test EnrichmentRecord"""
        # FIXME: construct object with mandatory attributes with example values
        # model = EnrichmentRecord()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
