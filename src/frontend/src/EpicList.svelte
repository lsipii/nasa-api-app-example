<script lang="ts">
    import { onMount } from 'svelte'
    import { DateInput } from 'date-picker-svelte'
    import {
        ifEmptyObject,
        ifObject,
        omitEmptyObjAttrs,
    } from '@lsipii/transformation-helpers/Objects'

    import EpicListItem from './EpicListItem.svelte'
    // @ts-ignore
    import imgUrl from './public/nasa-app-logo.png'
    import fetchEpicData from './services/NasaEpicService'

    //
    // Data
    //

    // View data state
    let epics = null
    let errorMessage = ''

    // Data updater
    const fetchData = async () => {
        epics = null
        errorMessage = ''

        try {
            const epicItems = await fetchEpicData(currentFilters)

            currentlyShown = 5
            epics = transformToEpics(epicItems)
        } catch (error) {
            epics = [] // allows for re-searching
            errorMessage = String(error.message)
        }
    }

    // Transform data to for EpicListItem specs
    const transformToEpics = (epicItems) => {
        return epicItems.map((epic) => {
            return {
                imageUrl: epic.image_url,
                imageThumbnailUrl: epic.image_thumbnail_url,
                caption: epic.caption,
                identifier: epic.identifier,
                coordinates: {
                    latitude: epic.coords.centroid_coordinates.lat,
                    longitude: epic.coords.centroid_coordinates.lon,
                },
                date: epic.date,
            }
        })
    }

    //
    // Render hooks
    //

    onMount(fetchData)

    //
    // Filter definitions
    //
    let pageSize = 5
    let currentlyShown = 0
    $: epicsVisible = epics ? epics.slice(0, currentlyShown) : []
    function showMore() {
        currentlyShown += pageSize
    }

    //
    // Filter definitions
    //
    $: hasActiveFilters = false
    const onFiltersChanged = (filters = null) => {
        if (ifObject(filters)) {
            for (const attr in filters) {
                if (typeof currentFilters[attr] !== 'undefined') {
                    currentFilters[attr] = filters[attr]
                }
            }
        }
        hasActiveFilters = !ifEmptyObject(omitEmptyObjAttrs(currentFilters))
    }
    const currentFilters = {
        enhanced: null,
        date: null,
    }
</script>

<div class="container">
    <img class="logo" src={imgUrl} alt="Nasa Logo" />

    <h2>NASA Epic API App</h2>
    {#if errorMessage.length > 0}
        ERROR: {errorMessage}
    {/if}

    <div class="inner-container">
        <div class="search-inputs">
            <div class="search-btn">
                <button
                    on:click={() => fetchData()}
                    disabled={!hasActiveFilters}>Search</button
                >
            </div>

            Date filter
            <div>
                <DateInput
                    bind:value={currentFilters.date}
                    format="yyyy-MM-dd"
                    placeholder="date"
                    on:select={() => onFiltersChanged()}
                />
            </div>

            <label for="enhanced">Enhanced</label>
            <div>
                <input
                    id="enhanced"
                    type="checkbox"
                    bind:checked={currentFilters.enhanced}
                    on:change={() => onFiltersChanged()}
                />
            </div>
        </div>
    </div>

    {#if epics === null}
        Loading Epics...
    {:else}
        {#each epicsVisible as epic}
            <div class="epic-item">
                <EpicListItem {...epic} />
            </div>
        {/each}

        <div class="inner-container">
            <div>
                <button
                    on:click={() => showMore()}
                    disabled={!epics || currentlyShown >= epics.length}
                    >Show more</button
                >
            </div>
        </div>
    {/if}
</div>

<style>
    .container {
        max-width: 800px;
        margin: 10px auto;
    }
    .logo {
        display: block;
        margin: 20px auto;
        width: 50%;
    }
    .inner-container {
        display: flex;
        align-items: center;
        justify-content: space-between;
        font-size: 16px;
        border-top: 1px solid lightgrey;
        padding-top: 15px;
        margin-bottom: 25px;
    }
    .inner-container .search-inputs {
        float: left;
        width: 300px;
    }
    .inner-container .search-btn {
        float: right;
        padding-top: 30px;
    }

    button {
        font-size: 14px;
        background-color: white;
    }
    button:hover {
        background: lightseagreen;
    }
</style>
