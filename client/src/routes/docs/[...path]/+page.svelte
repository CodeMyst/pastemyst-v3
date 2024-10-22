<script lang="ts">
    import Markdown from "$lib/Markdown.svelte";
    import type { PageServerData } from "./$types";

    export let data: PageServerData;
</script>

<svelte:head>
    <title>pastemyst | {data.title}</title>
</svelte:head>

<section>
    <Markdown content={data.content} />
</section>

<style lang="scss">
    section {
        margin-bottom: 0;
    }
</style>

