import { Injectable } from '@angular/core'

@Injectable({
  providedIn: 'root'
})
export class EllipsisService {

  constructor() {
  }

  ellipse(value?: string, maxLength: number = 30): string {
    if (!value) {
      return ''
    }
    if (value.length > maxLength) {
      return value.slice(0, maxLength) + '...'
    }
    return value
  }
}
