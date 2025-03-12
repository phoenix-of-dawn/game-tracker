import { Component, output } from '@angular/core';
import {
    FormControl,
    FormGroup,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { GameHandler } from '../../services/game-fetch.service';
import { Game } from '../../interfaces/game';

@Component({
    selector: 'search',
    templateUrl: './search.component.html',
    styleUrl: './search.component.scss',
    imports: [ReactiveFormsModule],
})
export class Search {
    constructor(private gameHandler: GameHandler) {}

    games = output<Game[]>();

    nameForm = new FormGroup({
        name: new FormControl('', Validators.required),
    });

    onSubmit() {
        this.gameHandler.getGames(this.nameForm.value.name);
        this.games.emit(this.gameHandler.games);
        this.nameForm.reset();
    }
}
