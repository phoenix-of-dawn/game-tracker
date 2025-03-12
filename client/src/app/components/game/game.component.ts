import { Component, computed, input } from '@angular/core';
import { GameHandler } from '../../services/game-fetch.service';
import { Game } from '../../interfaces/game';

@Component({
    selector: 'game',
    templateUrl: './game.component.html',
    styleUrl: './game.component.scss',
})
export class GameComponent {
    gameInput = input<Game>();

    game = computed(() => this.gameInput());
}
