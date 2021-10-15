"""
    Klarity Integrations

    REST API for managing Estate Records using Klarity Integrations. You can enrich your estate by creating new kinds of estate records or extending existing ones. Before making use of the API, you must first register your External Integration in Klarity, which provides you with the required authentication credentials. Then, you use those credentials to obtain a Token that allows you to make authorized calls to Klarity’s REST API for External Integration.  # noqa: E501

    The version of the OpenAPI document: 0.0.3
    Contact: products@nordcloud.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import klarity_integrations
from klarity_integrations.model.estate_records_delete_body_records import EstateRecordsDeleteBodyRecords
from klarity_integrations.model.period_enum import PeriodEnum
globals()['EstateRecordsDeleteBodyRecords'] = EstateRecordsDeleteBodyRecords
globals()['PeriodEnum'] = PeriodEnum
from klarity_integrations.model.estate_records_delete_body import EstateRecordsDeleteBody


class TestEstateRecordsDeleteBody(unittest.TestCase):
    """EstateRecordsDeleteBody unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testEstateRecordsDeleteBody(self):
        """Test EstateRecordsDeleteBody"""
        # FIXME: construct object with mandatory attributes with example values
        # model = EstateRecordsDeleteBody()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
