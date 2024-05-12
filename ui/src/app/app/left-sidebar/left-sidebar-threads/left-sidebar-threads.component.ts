import { Component, inject } from '@angular/core'
import { ThreadService } from '../../../services/thread.service'
import { ToastService } from '../../../services/toast.service'
import { NgForOf, NgIf } from '@angular/common'
import { ThreadButtonComponent } from './thread-button/thread-button.component'
import { ThreadFragment, ThreadsQuery } from '../../../../../graphql/generated'

@Component({
  selector: 'app-left-sidebar-threads',
  standalone: true,
  imports: [
    NgIf,
    NgForOf,
    ThreadButtonComponent
  ],
  templateUrl: './left-sidebar-threads.component.html',
  styleUrl: './left-sidebar-threads.component.css'
})
export class LeftSidebarThreadsComponent {
  threadService: ThreadService = inject(ThreadService)
  toastService: ToastService = inject(ToastService)
  threads: ThreadFragment[] = []
  async ngOnInit(): Promise<void> {
    await this.loadThreads()
  }

  async loadThreads() {
    try {
      this.threads = await this.threadService.threads({first: 50})
    } catch (e) {
      this.toastService.error(`${e}`)
    }
  }
}
