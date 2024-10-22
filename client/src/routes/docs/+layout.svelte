<script lang="ts">
    import DocTreeList from "$lib/DocTreeList.svelte";
    import type { LayoutData } from "./$types";

    export let data: LayoutData;
</script>

<div class="flex sm-row gap-l">
    <div class="left flex col">
        <ul>
            <DocTreeList nodes={data.fileTree} selectedPath={data.path} />
        </ul>
    </div>

    <div class="right">
        <slot />
    </div>
</div>

<style lang="scss">
    .left {
        margin-top: 2rem;
        background-color: var(--color-bg1);
        border: 1px solid var(--color-bg2);
        border-radius: $border-radius;
        min-width: 10rem;

        ul {
            list-style: none;
            margin: 0;
            padding: 0;
        }
    }

    .right {
        width: 100%;
    }

    @media screen and (max-width: $break-med) {
        .right :global(section) {
            margin-top: 0;
        }
    }
</style>
