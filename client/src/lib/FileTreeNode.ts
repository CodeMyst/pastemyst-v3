export interface FileTreeNode {
    name: string;
    children: FileTreeNode[];
    path: string;
};

export const buildFileTree = async (paths: string[]): Promise<FileTreeNode> => {
    const root: FileTreeNode = {
        name: "",
        children: [],
    };

    for (const path of paths) {
        const pathParts = path.split("/");
        let current = root;

        for (const part of pathParts) {
            const existing = current.children.find((child) => child.name === part);

            if (existing) {
                current = existing;
            } else {
                const newNode: FileTreeNode = {
                    name: part,
                    children: [],
                    path: path
                };

                current.children.push(newNode);
                current = newNode;
            }
        }
    }

    return root;
};

