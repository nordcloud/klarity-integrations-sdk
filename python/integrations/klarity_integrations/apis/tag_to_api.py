import typing_extensions

from klarity_integrations.apis.tags import TagValues
from klarity_integrations.apis.tags.klarity_integrations_api import KlarityIntegrationsApi

TagToApi = typing_extensions.TypedDict(
    'TagToApi',
    {
        TagValues.KLARITY_INTEGRATIONS: KlarityIntegrationsApi,
    }
)

tag_to_api = TagToApi(
    {
        TagValues.KLARITY_INTEGRATIONS: KlarityIntegrationsApi,
    }
)
