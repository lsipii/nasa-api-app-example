import nasaApi from '../config/nasa-api'

export default async function fetchEpicData(currentFilter: string): Promise<{
    success: boolean
    message: string
    items?: Array<any>
}> {
    const apiEndpoint = `${nasaApi.apiEndpointHost}/epic-api`

    const response = await fetch(apiEndpoint, {
        method: 'POST',
        mode: 'cors',
        headers: {
            'Content-Type': 'application/json; charset=utf-8',
            'X-Auth-Token': nasaApi.clientAuthToken,
        },
        body: JSON.stringify({
            filters: currentFilter,
        }),
    })

    const resposeItems = await response.json()
    validateResponseData(resposeItems)
    return resposeItems
}

function validateResponseData(resposeItems: any) {
    if (!(resposeItems instanceof Array)) {
        throw new Error('Bad response received')
    }
}
