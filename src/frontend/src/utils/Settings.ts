const secretSettings = {
    NASA_APP_DATA_ENDPOINT_HOST: 'http://localhost:8000',
    NASA_APP_DATA_ENDPOINT_ACCESS_TOKEN: '05f717e5',
}

export function getSetting(settingName: string, defaultValue?: string): string {
    if (typeof secretSettings[settingName] !== 'undefined') {
        return secretSettings[settingName]
    }
    return defaultValue
}
