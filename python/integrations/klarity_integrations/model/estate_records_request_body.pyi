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


class EstateRecordsRequestBody(
    schemas.DictSchema
):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """


    class MetaOapg:
        required = {
            "records",
        }
        
        class properties:
            
            
            class records(
                schemas.ListSchema
            ):
            
            
                class MetaOapg:
                    
                    
                    class items(
                        schemas.DictSchema
                    ):
                    
                    
                        class MetaOapg:
                            required = {
                                "id",
                                "type",
                            }
                            
                            class properties:
                            
                                @staticmethod
                                def id() -> typing.Type['Id']:
                                    return Id
                            
                                @staticmethod
                                def name() -> typing.Type['Name']:
                                    return Name
                            
                                @staticmethod
                                def type() -> typing.Type['Type']:
                                    return Type
                                metadata = schemas.DictSchema
                            
                                @staticmethod
                                def tags() -> typing.Type['Tags']:
                                    return Tags
                            
                                @staticmethod
                                def costs() -> typing.Type['Costs']:
                                    return Costs
                            
                                @staticmethod
                                def validThrough() -> typing.Type['ValidThrough']:
                                    return ValidThrough
                                __annotations__ = {
                                    "id": id,
                                    "name": name,
                                    "type": type,
                                    "metadata": metadata,
                                    "tags": tags,
                                    "costs": costs,
                                    "validThrough": validThrough,
                                }
                        
                        id: 'Id'
                        type: 'Type'
                        
                        @typing.overload
                        def __getitem__(self, name: typing_extensions.Literal["id"]) -> 'Id': ...
                        
                        @typing.overload
                        def __getitem__(self, name: typing_extensions.Literal["name"]) -> 'Name': ...
                        
                        @typing.overload
                        def __getitem__(self, name: typing_extensions.Literal["type"]) -> 'Type': ...
                        
                        @typing.overload
                        def __getitem__(self, name: typing_extensions.Literal["metadata"]) -> MetaOapg.properties.metadata: ...
                        
                        @typing.overload
                        def __getitem__(self, name: typing_extensions.Literal["tags"]) -> 'Tags': ...
                        
                        @typing.overload
                        def __getitem__(self, name: typing_extensions.Literal["costs"]) -> 'Costs': ...
                        
                        @typing.overload
                        def __getitem__(self, name: typing_extensions.Literal["validThrough"]) -> 'ValidThrough': ...
                        
                        @typing.overload
                        def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
                        
                        def __getitem__(self, name: typing.Union[typing_extensions.Literal["id", "name", "type", "metadata", "tags", "costs", "validThrough", ], str]):
                            # dict_instance[name] accessor
                            return super().__getitem__(name)
                        
                        
                        @typing.overload
                        def get_item_oapg(self, name: typing_extensions.Literal["id"]) -> 'Id': ...
                        
                        @typing.overload
                        def get_item_oapg(self, name: typing_extensions.Literal["name"]) -> typing.Union['Name', schemas.Unset]: ...
                        
                        @typing.overload
                        def get_item_oapg(self, name: typing_extensions.Literal["type"]) -> 'Type': ...
                        
                        @typing.overload
                        def get_item_oapg(self, name: typing_extensions.Literal["metadata"]) -> typing.Union[MetaOapg.properties.metadata, schemas.Unset]: ...
                        
                        @typing.overload
                        def get_item_oapg(self, name: typing_extensions.Literal["tags"]) -> typing.Union['Tags', schemas.Unset]: ...
                        
                        @typing.overload
                        def get_item_oapg(self, name: typing_extensions.Literal["costs"]) -> typing.Union['Costs', schemas.Unset]: ...
                        
                        @typing.overload
                        def get_item_oapg(self, name: typing_extensions.Literal["validThrough"]) -> typing.Union['ValidThrough', schemas.Unset]: ...
                        
                        @typing.overload
                        def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
                        
                        def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["id", "name", "type", "metadata", "tags", "costs", "validThrough", ], str]):
                            return super().get_item_oapg(name)
                        
                    
                        def __new__(
                            cls,
                            *_args: typing.Union[dict, frozendict.frozendict, ],
                            id: 'Id',
                            type: 'Type',
                            name: typing.Union['Name', schemas.Unset] = schemas.unset,
                            metadata: typing.Union[MetaOapg.properties.metadata, dict, frozendict.frozendict, schemas.Unset] = schemas.unset,
                            tags: typing.Union['Tags', schemas.Unset] = schemas.unset,
                            costs: typing.Union['Costs', schemas.Unset] = schemas.unset,
                            validThrough: typing.Union['ValidThrough', schemas.Unset] = schemas.unset,
                            _configuration: typing.Optional[schemas.Configuration] = None,
                            **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
                        ) -> 'items':
                            return super().__new__(
                                cls,
                                *_args,
                                id=id,
                                type=type,
                                name=name,
                                metadata=metadata,
                                tags=tags,
                                costs=costs,
                                validThrough=validThrough,
                                _configuration=_configuration,
                                **kwargs,
                            )
            
                def __new__(
                    cls,
                    _arg: typing.Union[typing.Tuple[typing.Union[MetaOapg.items, dict, frozendict.frozendict, ]], typing.List[typing.Union[MetaOapg.items, dict, frozendict.frozendict, ]]],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'records':
                    return super().__new__(
                        cls,
                        _arg,
                        _configuration=_configuration,
                    )
            
                def __getitem__(self, i: int) -> MetaOapg.items:
                    return super().__getitem__(i)
        
            @staticmethod
            def insertInPeriod() -> typing.Type['InsertInPeriodEnum']:
                return InsertInPeriodEnum
            __annotations__ = {
                "records": records,
                "insertInPeriod": insertInPeriod,
            }
    
    records: MetaOapg.properties.records
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["records"]) -> MetaOapg.properties.records: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["insertInPeriod"]) -> 'InsertInPeriodEnum': ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["records", "insertInPeriod", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["records"]) -> MetaOapg.properties.records: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["insertInPeriod"]) -> typing.Union['InsertInPeriodEnum', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["records", "insertInPeriod", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *_args: typing.Union[dict, frozendict.frozendict, ],
        records: typing.Union[MetaOapg.properties.records, list, tuple, ],
        insertInPeriod: typing.Union['InsertInPeriodEnum', schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'EstateRecordsRequestBody':
        return super().__new__(
            cls,
            *_args,
            records=records,
            insertInPeriod=insertInPeriod,
            _configuration=_configuration,
            **kwargs,
        )

from klarity_integrations.model.costs import Costs
from klarity_integrations.model.id import Id
from klarity_integrations.model.insert_in_period_enum import InsertInPeriodEnum
from klarity_integrations.model.name import Name
from klarity_integrations.model.tags import Tags
from klarity_integrations.model.type import Type
from klarity_integrations.model.valid_through import ValidThrough
