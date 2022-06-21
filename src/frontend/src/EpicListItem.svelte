<script>
    import { fly } from 'svelte/transition'
    export let identifier
    export let caption
    export let image_url
</script>

<div class="epic-item">
    <div class="epic-item-left" transition:fly={{ y: 20, duration: 300 }}>
        <div class="epic-item-label">{identifier}: {caption}</div>
        <div class="epic-image">
            <img src={image_url} alt={caption} width="300" />
        </div>
    </div>
</div>

<style>
    .epic-item {
        margin-bottom: 15px;
        display: flex;
        align-items: center;
        justify-content: space-between;
        animation-duration: 0.3s;
    }
    .epic-item-left {
        display: flex;
        align-items: center;
    }
    .epic-item-label {
        border: 1px solid white;
        margin-left: 12px;
    }
</style>
