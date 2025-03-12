import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Game } from '../interfaces/game';

@Injectable({
    providedIn: 'root',
})
export class HttpHandler {
    constructor(private http: HttpClient) {}

    baseUrl = 'http://localhost:8080';

    getGamesRequest(name: string): Observable<Game[]> {
        return this.http.get<Game[]>(`${this.baseUrl}/search?name=${name}`);
    }
}
