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

  async userId(): Promise<string> {
    return Session.getUserId()
  }

  async payload(): Promise<any> {
    return Session.getAccessTokenPayloadSecurely()
  }

  async email(): Promise<string> {
    let payload = await this.payload()
    return payload.email
  }
}
