import { Component, inject, Input } from '@angular/core'
import { ThreadService } from '../../services/thread.service'
import { ToastService } from '../../services/toast.service'
import { BaseComponent } from '../base/base.component'

@Component({
  selector: 'app-thread',
  standalone: true,
  imports: [
    BaseComponent
  ],
  templateUrl: './thread.component.html',
  styleUrl: './thread.component.css'
})
export class ThreadComponent {
  @Input()  set id(id: string) {
    try {
      this.thread = this.threadService.getThread(id)
      console.log('thread: ' + this.thread)
    } catch (e) {
      this.toastService.error(`${e}`)
    }
  }
  threadService: ThreadService = inject(ThreadService)
  toastService: ToastService = inject(ToastService)
  thread: any = undefined
  protected readonly JSON = JSON
}
