import { getUserTags, type User } from "$lib/api/user";
import moment from "moment";
import type { Page } from "$lib/api/page";
import type { PageLoad } from "./$types";
import type { Paste } from "$lib/api/paste";
import { error } from "@sveltejs/kit";
import { PUBLIC_API_BASE } from "$env/static/public";

export const load: PageLoad = async ({ params, fetch }) => {
    const userRes = await fetch(`${PUBLIC_API_BASE}/users/${params.user}`, {
        method: "get"
    });

    const meRes = await fetch(`${PUBLIC_API_BASE}/auth/self`, {
        method: "get",
        credentials: "include"
    });

    const userPastesRes = await fetch(`${PUBLIC_API_BASE}/users/${params.user}/pastes?pageSize=5`, {
        method: "get",
        credentials: "include"
    });

    const userPinnedPastesRes = await fetch(
        `${PUBLIC_API_BASE}/users/${params.user}/pastes/pinned?pageSize=5`,
        {
            method: "get",
            credentials: "include"
        }
    );

    let user: User;
    let relativeJoined: string;
    let isCurrentUser = false;
    let pastes: Page<Paste>;
    let pinnedPastes: Page<Paste>;
    if (userRes.ok) {
        user = await userRes.json();
        relativeJoined = moment(user.createdAt).fromNow();

        if (meRes.ok) {
            const loggedInUser: User = await meRes.json();

            isCurrentUser = loggedInUser.id === user.id;
        }

        if (userPastesRes.ok && userPinnedPastesRes.ok) {
            pastes = await userPastesRes.json();
            pinnedPastes = await userPinnedPastesRes.json();
        } else {
            throw error(userPastesRes.status);
        }

        const tags = await getUserTags(fetch, user.username);

        return {
            user: user,
            isCurrentUser: isCurrentUser,
            relativeJoined: relativeJoined,
            pastes: pastes,
            pinnedPastes: pinnedPastes,
            tags: tags
        };
    } else {
        throw error(userRes.status);
    }
};
