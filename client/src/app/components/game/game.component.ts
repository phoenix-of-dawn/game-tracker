import { Component, computed, input } from '@angular/core';
import { GameHandler } from '../../services/game-fetch.service';
import { Game } from '../../interfaces/game';
import { MatCardModule } from '@angular/material/card';

@Component({
    selector: 'game',
    templateUrl: './game.component.html',
    styleUrl: './game.component.scss',
    imports: [MatCardModule],
})
export class GameComponent {
    gameInput = input<Game>();

    game = computed(() => this.gameInput());
}
