import { Component, inject, signal, WritableSignal } from '@angular/core'
import { BaseComponent } from '../base/base.component'
import { MessageService } from '../../services/message.service'
import { ToastService } from '../../services/toast.service'
import { Router } from '@angular/router'
import { ThreadMessageFragment } from '../../../../graphql/generated'

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [
    BaseComponent
  ],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {
  messageService:MessageService = inject(MessageService)
  toastService: ToastService = inject(ToastService)
  router: Router = inject(Router)
  loading: WritableSignal<boolean> = signal(true)
  async ngOnInit() {
    try {
      let mostRecentMessage: ThreadMessageFragment | undefined = await this.messageService.mostRecentMessage({userId: ''})
      if (mostRecentMessage) {
        await this.router.navigate([`/app/thread/${mostRecentMessage.thread.id}`])
      }
    } catch (e) {
      this.toastService.error("We've encountered an error, please try again")
    } finally {
      this.loading.set(false)
    }
  }

}
