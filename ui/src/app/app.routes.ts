import {
  Routes
} from '@angular/router'
import { canActivateAppGuard } from './can-activate-app.guard'
import { SigninComponent } from './auth/signin/signin.component'
import { SignoutComponent } from './auth/signout/signout.component'
import { NotFoundComponent } from './not-found/not-found.component'
import { canActivateSignOutGuard } from './guards/can-activate-sign-out.guard'
import { canActivateSignInGuard } from './guards/can-activate-sign-in.guard'
import { HomeComponent } from './app/home/home.component'

export const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    redirectTo: '/app'
  },
  {
    path: 'signin',
    redirectTo: '/auth/signin',
    pathMatch: 'full',
  },
  {
    path: 'signout',
    redirectTo: '/auth/signout',
    pathMatch: 'full',
  },
  {
    path: 'auth',
    children: [
      {path: '', pathMatch: 'full', redirectTo: 'signin'},
      {path: 'signin', component: SigninComponent, canActivate: [canActivateSignInGuard]},
      {path: 'signout', component: SignoutComponent, canActivate: [canActivateSignOutGuard]},
    ]
  },
  {
    path: 'app',
    canActivateChild: [canActivateAppGuard],
    children: [
      {path: '', pathMatch: 'full', redirectTo: 'home'},
      {path: 'home', component: HomeComponent},
    ]
  },
  {path: '**', redirectTo: 'notfound'},
  {path: 'notfound', component: NotFoundComponent}
];
