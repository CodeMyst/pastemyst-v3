import { error } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import fs from "fs";
import { parse } from "yaml";

export const load: PageServerLoad = async ({ params }) => {
    let path = null;

    // first try to open /docs/page.md, if that doesn't exist, try and open /docs/page/index.md
    // this allows us to have the following docs tree structure with nice URLs:
    // - docs
    //   - v1
    //     - index.md
    //   - v2
    //     - index.md
    //     - paste.md
    // ...
    //
    // here the following URLs would be valid:
    // - docs/v1
    // - docs/v1/index
    // - docs/v2
    // - docs/v2/index
    // - docs/v2/paste
    if (fs.existsSync(`../docs/${params.path}.md`)) {
        path = `../docs/${params.path}.md`;
    } else if (fs.existsSync(`../docs/${params.path}/index.md`)) {
        path = `../docs/${params.path}/index.md`;
    } else {
        error(404);
    }

    const pageContent = fs.readFileSync(path, "utf8");

    const yamlFirstDivider = pageContent.indexOf("---");
    const yamlLastDivider = pageContent.indexOf("---", yamlFirstDivider + 1);
    const yamlMetaContent = pageContent.slice(yamlFirstDivider + 3, yamlLastDivider);
    const yamlMeta = parse(yamlMetaContent);

    const markdownContent = pageContent.slice(yamlLastDivider + 3);

    return {
        content: markdownContent,
        title: yamlMeta.title
    };
};
