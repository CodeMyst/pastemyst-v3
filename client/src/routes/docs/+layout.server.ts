import { readdir } from "fs/promises";
import type { LayoutServerLoad } from "./$types";
import { buildFileTree } from "$lib/FileTreeNode";

export const load: LayoutServerLoad = async ({ params }) => {
    const files = (await readdir('../docs', { recursive: true })).filter((file) => !file.endsWith('index.md'));

    // trim the extension from files if they have it
    files.forEach((file, index) => {
        if (file.endsWith('.md')) {
            files[index] = file.slice(0, -3);
        }
    });

    return {
        fileTree: (await buildFileTree(files)).children,
        path: params.path!
    };
};
