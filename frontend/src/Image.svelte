<script lang="ts">
    import { OpenInFileExplorer } from "../wailsjs/go/main/App.js";
    import ImageModal from "./ImageModal.svelte";
    export let image;
    export let showEditTags = false;

    async function onOpenFileInExplorer() {
        await OpenInFileExplorer(image.Path);
    }

    function closeModal() {
        showEditTags = false;
    }
</script>

<div class="image-container">
    <div class="image-buttons">
        <button on:click={onOpenFileInExplorer}>Open file</button>
        <button on:click={() => (showEditTags = true)}>Edit tags</button>
    </div>
    {#if showEditTags}
        <ImageModal close={closeModal} {image} />
    {/if}
    <img src={image.Path} alt={image.Name} />
</div>

<style>
    .image-buttons {
        position: absolute;
        top: 5px;
        right: 5px;
    }

    .image-container {
        position: relative;
        width: calc(50% - 4px);
        border: black 2px solid;
        height: 300px;
        margin: 2px;
        overflow: hidden;
        box-sizing: border-box;
        display: flex;
        justify-content: center; /* Align horizontally */
        align-items: center;
    }

    /* Large (L) resolution */
    @media (min-width: 1024px) {
        .image-container {
            width: calc(33% - 4px);
        }
    }

    /* Extra Large (XL) resolution */
    @media (min-width: 1280px) {
        .image-container {
            width: calc(25% - 4px);
        }
    }

    .image-container img {
        max-width: 100%;
        height: auto;
    }
</style>
