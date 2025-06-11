import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Game } from '../interfaces/game';
import { UserCreationResponse } from '../interfaces/user';

@Injectable({
    providedIn: 'root',
})
export class HttpHandler {
    constructor(private http: HttpClient) {}

    baseUrl = 'http://localhost:8080';
    jwtKey = null;

    getGamesRequest(name: string): Observable<Game[]> {
        return this.http.get<Game[]>(`${this.baseUrl}/search?name=${name}`);
    }

    postRegisterRequest(
        email: string,
        password: string
    ): Observable<UserCreationResponse> {
        return this.http.post<UserCreationResponse>(
            `${this.baseUrl}/register`,
            {
                email: email,
                password: password,
            }
        );
    }
}
