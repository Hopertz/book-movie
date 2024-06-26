export interface BookerType {
    name: string;
    phone: string;
    seat: {
        row: number;
        seat: number;
    }[];
    movie: string;
    amount: number;
    seats: string[][];
}


export interface Movie {
    id : string
    name: string;
    amount: number;
}
