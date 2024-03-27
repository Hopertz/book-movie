import { redirect } from "@sveltejs/kit";
import type { RequestHandler } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ request }) => {
    
    const body = await request.json();

    const res = await fetch(`http://localhost:8448/bookers`, {
         method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(body),
      });

    if (res.ok) {
       return res.json() 
    }

    return redirect(303, "/")
};
