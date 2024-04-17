import { redirect } from "@sveltejs/kit";
import type { RequestHandler } from '@sveltejs/kit';
import { SECRET_BASE_API_URL } from '$env/static/private'

export const POST: RequestHandler = async ({ request }) => {

    const body = await request.json();

    const id = body.id

    const res = await fetch(`${SECRET_BASE_API_URL}/movies/${id}`)

    if (res.ok) {
       return res 
    }

    return redirect(303, "/")
};
