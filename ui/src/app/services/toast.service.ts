import { inject, Injectable } from '@angular/core'
import { MessageService } from 'primeng/api'

@Injectable({
  providedIn: 'root'
})
export class ToastService {
  constructor(private messageService: MessageService) { }

  generalError() {
    this.error("We've encountered an error, please try again")
  }
  error(summary: string, detail?: string): void {
    this.messageService.add({severity: 'error', summary: summary, detail: detail})
  }

  success(summary: string, detail?: string): void {
    this.messageService.add({severity: 'success', summary: summary, detail: detail})
  }
}
