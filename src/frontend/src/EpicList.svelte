<script>
    import { onMount } from 'svelte'

    import EpicItem from './EpicItem.svelte'
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
            epics = await fetchEpicData(currentFilter)
        } catch (error) {
            errorMessage = String(error.message)
        }
    }

    //
    // Render hooks
    //

    onMount(fetchData)

    //
    // Filter definitions
    //
    let currentFilter = 'all'

    $: filteredEpics =
        currentFilter === 'all'
            ? epics
            : currentFilter === 'active'
            ? epics.filter((epic) => epic.completed)
            : epics.filter((epic) => !epic.completed)

    function updateFilter(newFilter) {
        currentFilter = newFilter
    }
</script>

<div class="container">
    <img class="logo" src={imgUrl} alt="Nasa Logo" />

    <h2>NASA Epic App</h2>
    {#if errorMessage.length > 0}
        ERROR: {errorMessage}
    {:else if epics === null}
        Loading Epics...
    {:else}
        {#each filteredEpics as epic}
            <div class="epic-item">
                <EpicItem {...epic} />
            </div>
        {/each}

        <div class="inner-container">
            <div>
                <button
                    on:click={() => updateFilter('all')}
                    class:active={currentFilter === 'all'}>All</button
                >
                <button
                    on:click={() => updateFilter('completed')}
                    class:active={currentFilter === 'completed'}
                    >Completed</button
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
        margin-bottom: 13px;
    }
    button {
        font-size: 14px;
        background-color: white;
        appearance: none;
    }
    button:hover {
        background: lightseagreen;
    }
    button:focus {
        outline: none;
    }
    .active {
        background: lightseagreen;
    }
</style>
