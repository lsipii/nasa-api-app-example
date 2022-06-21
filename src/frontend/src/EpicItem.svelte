<script>
    import { fly } from 'svelte/transition'
    export let author
    export let completed
</script>

<div class="todo-item">
    <div class="todo-item-left" transition:fly={{ y: 20, duration: 300 }}>
        <div class="todo-item-label" class:completed>{author}</div>
    </div>
</div>

<style>
    .todo-item {
        margin-bottom: 15px;
        display: flex;
        align-items: center;
        justify-content: space-between;
        animation-duration: 0.3s;
    }
    .todo-item-left {
        display: flex;
        align-items: center;
    }
    .todo-item-label {
        border: 1px solid white;
        margin-left: 12px;
    }
    .completed {
        text-decoration: line-through;
        color: grey;
    }
</style>
