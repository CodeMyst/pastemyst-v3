import { apiBase } from "$lib/api/api";
import type { User } from "$lib/api/user";
import moment from "moment";
import type { Page } from "$lib/api/page";
import type { PageLoad } from "./$types";
import type { Paste } from "$lib/api/paste";
import { error } from "@sveltejs/kit";

export const load: PageLoad = async ({ params, fetch }) => {
    const userRes = await fetch(`${apiBase}/user/${params.user}`, {
        method: "get"
    });

    const meRes = await fetch(`${apiBase}/auth/self`, {
        method: "get",
        credentials: "include"
    });

    const userPastesRes = await fetch(`${apiBase}/user/${params.user}/pastes?page_size=5`, {
        method: "get",
        credentials: "include"
    });

    let user: User;
    let relativeJoined: string;
    let isCurrentUser = false;
    let pastes: Page<Paste>;
    if (userRes.ok) {
        user = await userRes.json();
        relativeJoined = moment(user.createdAt).fromNow();

        if (meRes.ok) {
            const loggedInUser: User = await meRes.json();

            isCurrentUser = loggedInUser.id === user.id;
        }

        if (userPastesRes.ok) {
            pastes = await userPastesRes.json();
        } else {
            throw error(userPastesRes.status);
        }

        return {
            user: user,
            isCurrentUser: isCurrentUser,
            relativeJoined: relativeJoined,
            pastes: pastes
        };
    } else {
        throw error(userRes.status);
    }
};
