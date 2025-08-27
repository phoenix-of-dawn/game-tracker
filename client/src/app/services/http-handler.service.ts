import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Game } from '../interfaces/game';
import { User } from '../interfaces/user';

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
        password: string,
        username: string
    ): Observable<unknown> {
        return this.http.post<unknown>(`${this.baseUrl}/register`, {
            email: email,
            username: username,
            password: password,
        });
    }

    postLoginRequest(email: string, password: string): Observable<User | null> {
        return this.http.post<User | null>(`${this.baseUrl}/login`, {
            email: email,
            password: password,
        });
    }

    getUserRequest(id: number): Observable<User | null> {
        return this.http.get<User | null>(`${this.baseUrl}/user?id=${id}`);
    }
}
