import { redirect } from "@sveltejs/kit";
import type { RequestHandler } from '@sveltejs/kit';
import { SECRET_BASE_API_URL } from '$env/static/private'

export const POST: RequestHandler = async ({ request }) => {

    const body = await request.json();

    const res = await fetch(`${SECRET_BASE_API_URL}/bookers`, {
         method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(body),
      });

    if (res.ok) {
       return res 
    }

    return redirect(303, "/")
};
