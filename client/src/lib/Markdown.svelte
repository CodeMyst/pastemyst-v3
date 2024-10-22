<script lang="ts">
    import { marked } from "marked";
    import { gfmHeadingId } from "marked-gfm-heading-id";

    export let content: string;

    marked.use(gfmHeadingId());

    $: markdownHtml = marked.parse(content, { gfm: true }) as string;
</script>

<div class="markdown">
    <!-- eslint-disable-next-line svelte/no-at-html-tags -->
    {@html markdownHtml}
</div>
