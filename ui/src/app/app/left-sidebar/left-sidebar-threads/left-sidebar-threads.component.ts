import { Component, inject } from '@angular/core'
import { ThreadService } from '../../../services/thread.service'
import { ToastService } from '../../../services/toast.service'
import { NgForOf, NgIf } from '@angular/common'
import { ThreadButtonComponent } from './thread-button/thread-button.component'
import { LeftSidebarThreadsService } from '../../../services/left-sidebar-threads.service'

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
  leftSidebarThreadsService: LeftSidebarThreadsService = inject(LeftSidebarThreadsService)
  threads: any = []
  async ngOnInit(): Promise<void> {
    // on load refresh the threads
    this.refreshThreads()
    // on refresh emit event, refresh threads
    this.leftSidebarThreadsService.refreshThreadsEmitter.subscribe(
      () => this.refreshThreads()
    )
  }

  async refreshThreads() {
    try {
      let response = await this.threadService.listThreads(
        50,
        `
        {
          id
          name
        }
      `
      )
      if (response.threads) {
        this.threads = response.threads
      }
    } catch (e) {
      this.toastService.error(`${e}`)
    }
  }

  protected readonly JSON = JSON
}
