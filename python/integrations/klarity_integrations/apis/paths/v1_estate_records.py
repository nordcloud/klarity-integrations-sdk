from klarity_integrations.paths.v1_estate_records.post import ApiForpost
from klarity_integrations.paths.v1_estate_records.delete import ApiFordelete


class V1EstateRecords(
    ApiForpost,
    ApiFordelete,
):
    pass
