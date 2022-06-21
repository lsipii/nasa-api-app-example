import { niceTrim } from '@lsipii/transformation-helpers/Strings'
import { getSetting } from '../utils/Settings'

export default {
    apiEndpointHost: niceTrim(
        getSetting('NASA_APP_DATA_ENDPOINT_HOST', 'http://localhost:8000'),
        '/',
        { onlyFromRight: true },
    ),
    clientAuthToken: getSetting('NASA_APP_DATA_ENDPOINT_ACCESS_TOKEN'),
}
