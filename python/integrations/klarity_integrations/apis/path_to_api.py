import typing_extensions

from klarity_integrations.paths import PathValues
from klarity_integrations.apis.paths.v1_estate_records import V1EstateRecords
from klarity_integrations.apis.paths.v1_enrichments_estate_records import V1EnrichmentsEstateRecords

PathToApi = typing_extensions.TypedDict(
    'PathToApi',
    {
        PathValues.V1_ESTATE_RECORDS: V1EstateRecords,
        PathValues.V1_ENRICHMENTS_ESTATE_RECORDS: V1EnrichmentsEstateRecords,
    }
)

path_to_api = PathToApi(
    {
        PathValues.V1_ESTATE_RECORDS: V1EstateRecords,
        PathValues.V1_ENRICHMENTS_ESTATE_RECORDS: V1EnrichmentsEstateRecords,
    }
)
