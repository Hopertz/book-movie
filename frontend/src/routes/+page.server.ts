import type { PageServerLoad } from "./$types"
import type { Movie } from "$lib/types"
export const load : PageServerLoad = async ({ fetch}): Promise<{ movies: Movie[] }>  => {

    const response = await fetch(`http://localhost:8448/movies`, {
        method: "GET",
    },
    )
    if (!response.ok) {
      return { movies : [] }
       
    }
    const movies: Movie[] = await response.json()
    return { movies :  movies}
}

