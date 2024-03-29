# coding: utf-8

"""
    Klarity Integrations

    REST API for managing Estate Records using Klarity Integrations. You can enrich your estate by creating new kinds of estate records or extending existing ones. Before making use of the API, you must first register your External Integration in Klarity, which provides you with the required authentication credentials. Then, you use those credentials to obtain a Token that allows you to make authorized calls to Klarity’s REST API for External Integration.  # noqa: E501

    The version of the OpenAPI document: 0.0.5
    Contact: products@nordcloud.com
    Generated by: https://openapi-generator.tech
"""

from datetime import date, datetime  # noqa: F401
import decimal  # noqa: F401
import functools  # noqa: F401
import io  # noqa: F401
import re  # noqa: F401
import typing  # noqa: F401
import typing_extensions  # noqa: F401
import uuid  # noqa: F401

import frozendict  # noqa: F401

from klarity_integrations import schemas  # noqa: F401


class ValidThrough(
    schemas.StrSchema
):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.

    Determines in which period the resource will be closed.
If specified, it has to be in `YYYY-MM` format and be at least current period - it can not be past period.
If you need to insert Estate Records in previous period, use insertInPeriod field - when set it will treat previous period as current.
If not specified, the record will exist indefinitely.
E.g. a resource with `validThrough` set to `2022-05` will be active till May 2022 and start being inactive in June 2022.

    """


    class MetaOapg:
        regex=[{
            'pattern': r'^2\d{3}-(0[1-9]|1[0-2])$',  # noqa: E501
        }]
