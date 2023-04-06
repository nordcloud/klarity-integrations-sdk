# do not import all endpoints into this module because that uses a lot of memory and stack frames
# if you need the ability to import all endpoints from this module, import them with
# from klarity_integrations.apis.path_to_api import path_to_api

import enum


class PathValues(str, enum.Enum):
    V1_ESTATE_RECORDS = "/v1/estateRecords"
    V1_ENRICHMENTS_ESTATE_RECORDS = "/v1/enrichments/estateRecords"
