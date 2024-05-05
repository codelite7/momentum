import { Injectable } from '@angular/core';
import Session from 'supertokens-web-js/recipe/session'

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor() { }

  async sessionExists(): Promise<boolean> {
    return Session.doesSessionExist()
  }

  async signOut(): Promise<void> {
    return await Session.signOut()
  }
}
