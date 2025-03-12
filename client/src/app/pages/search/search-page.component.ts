import { Component } from '@angular/core';
import { GameHandler } from '../../services/game-fetch.service';
import { Game } from '../../interfaces/game';
import { Search } from '../../components/search/search.component';
import { GameComponent } from '../../components/game/game.component';

@Component({
    selector: 'search-page',
    templateUrl: './search-page.component.html',
    styleUrl: './search-page.component.scss',
    imports: [Search, GameComponent],
})
export class SearchPage {
    games: Game[] = [];

    handleSearch(games: Game[]) {
        this.games = games;
    }
}
