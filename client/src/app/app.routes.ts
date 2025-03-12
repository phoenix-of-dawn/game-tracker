import { Routes } from '@angular/router';
import { AppComponent } from './app.component';
import { SearchPage } from './pages/search/search-page.component';

export const routes: Routes = [
    {
        path: '',
        component: AppComponent,
    },
    {
        path: 'search',
        component: SearchPage,
    },
];
