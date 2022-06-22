<script>
    import { fly } from 'svelte/transition'
    export let imageUrl
    export let imageThumbnailUrl
    export let caption
    export let identifier
    export let coordinates
    export let date
</script>

<div class="epic-item">
    <div class="epic-item-left" transition:fly={{ y: 20, duration: 300 }}>
        <div class="epic-item-label">
            Identifier: {identifier} <br />
            Date: {date} <br />
            Longitude: {coordinates.longitude} <br />
            Latitude: {coordinates.latitude}
        </div>
        <div class="epic-image">
            <a href={imageUrl} target="_blank"
                ><img src={imageThumbnailUrl} alt={caption} /></a
            >
            <div class="epic-image-caption">{caption}</div>
        </div>
    </div>
</div>

<style>
    .epic-image,
    .epic-image img {
        width: 150px;
        margin-left: 12px;
    }
    .epic-item {
        margin-bottom: 15px;
        display: flex;
        align-items: center;
        justify-content: space-between;
        animation-duration: 0.3s;
    }
    .epic-image-caption {
        font-size: small;
    }
    .epic-item-left {
        display: flex;
        align-items: center;
    }
    .epic-item-label {
        margin-left: 12px;
    }
</style>
