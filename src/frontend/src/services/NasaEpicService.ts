import nasaApi from '../config/nasa-api'

interface CentroidCoordinates {
    lat: number
    lon: number
}
interface DscovrJ2000Position {
    x: number
    y: number
    z: number
}
interface LunarJ2000Position {
    x: number
    y: number
    z: number
}

interface SunJ2000Position {
    x: number
    y: number
    z: number
}
interface AttitudeQuaternions {
    q0: number
    q1: number
    q2: number
    q3: number
}

export interface EpicItem {
    identifier: string
    caption: string
    image: string
    version: string
    centroid_coordinates: CentroidCoordinates
    dscovr_j2000_position: DscovrJ2000Position
    lunar_j2000_position: LunarJ2000Position
    sun_j2000_position: SunJ2000Position
    attitude_quaternions: AttitudeQuaternions
    date: string
    coords: {
        centroid_coordinates: CentroidCoordinates
        dscovr_j2000_position: DscovrJ2000Position
        lunar_j2000_position: LunarJ2000Position
        sun_j2000_position: SunJ2000Position
        attitude_quaternions: AttitudeQuaternions
    }
}

export default async function fetchEpicData(
    currentFilter: string,
): Promise<Array<EpicItem>> {
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
