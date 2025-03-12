export interface Game {
    id: number;
    name: string;
    aggregated_rating: number;
    cover: Cover;
}

export interface Cover {
    id: number;
    url: string;
}
