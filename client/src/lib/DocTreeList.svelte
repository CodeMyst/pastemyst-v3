<script lang="ts">
    import type { FileTreeNode } from "./FileTreeNode";

    export let nodes: FileTreeNode[];
    export let selectedPath: string;
</script>

{#each nodes as node}
    <li>
        <details open>
            <summary>
                <a href="/docs/{node.path}" class:active={selectedPath === node.path}>{node.name}</a>
            </summary>

            {#if node.children.length > 0}
                <ul>
                    <svelte:self nodes={node.children} selectedPath={selectedPath} />
                </ul>
            {/if}
        </details>
    </li>
{/each}

<style>
    li {
        list-style: none;
        width: 100%;
        line-height: 1.5rem;
    }

    ul {
        padding-left: 1rem;
    }

    a {
        width: 100%;
        display: flex;

        &.active {
            color: var(--color-secondary);
            font-weight: bold;
        }
    }
</style>
