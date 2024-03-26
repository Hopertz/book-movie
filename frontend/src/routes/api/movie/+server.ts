import { redirect } from "@sveltejs/kit";
import type { RequestHandler } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ request }) => {
    const body = await request.json();
    console.log(typeof body)
   
    const res = await fetch(`http://localhost:8448/bookers`, {
        method: "POST",
        headers: {
				'content-type': 'application/json'
			},
            body: body

    },)

    if (res.ok) {
       return res.json() 
    }

    return redirect(303, "/")
};
