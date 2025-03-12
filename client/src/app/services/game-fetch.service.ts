import { Injectable } from '@angular/core';
import { HttpHandler } from './http-handler.service';
import { Game } from '../interfaces/game';

@Injectable({
    providedIn: 'root',
})
export class GameHandler {
    games: Game[];

    constructor(private http: HttpHandler) {
        this.games = [];
    }

    public getGames(name: string | null | undefined) {
        if (name == undefined) {
            return;
        }

        console.log('Got here!');

        this.http.getGamesRequest(name).subscribe((res) => {
            this.games = res;
            console.log(this.games);
        });
    }
}
