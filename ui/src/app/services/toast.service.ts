import { inject, Injectable } from '@angular/core'
import { MessageService } from 'primeng/api'

@Injectable({
  providedIn: 'root'
})
export class ToastService {
  constructor(private messageService: MessageService) { }

  error(summary: string, detail?: string): void {
    this.messageService.add({severity: 'error', summary: summary, detail: detail})
  }

  success(summary: string, detail?: string): void {
    this.messageService.add({severity: 'success', summary: summary, detail: detail})
  }
}
