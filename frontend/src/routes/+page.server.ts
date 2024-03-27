import type { PageServerLoad } from "./$types"
import type { Movie } from "$lib/types"
import { SECRET_BASE_API_URL } from '$env/static/private'
export const load : PageServerLoad = async ({ fetch}): Promise<{ movies: Movie[] }>  => {

    const response = await fetch(`${SECRET_BASE_API_URL}/movies`, {
        method: "GET",
    },
    )
    if (!response.ok) {
      return { movies : [] }
       
    }
    const movies: Movie[] = await response.json()
    return { movies :  movies}
}

