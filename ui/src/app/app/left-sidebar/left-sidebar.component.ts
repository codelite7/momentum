import { Component, inject } from '@angular/core'
import { ButtonModule } from 'primeng/button'
import { LeftSidebarThreadsComponent } from './left-sidebar-threads/left-sidebar-threads.component'
import { ThreadService } from '../../services/thread.service'
import { ToastService } from '../../services/toast.service'
import { LeftSidebarThreadsService } from '../../services/left-sidebar-threads.service'
import { Router } from '@angular/router'

@Component({
  selector: 'app-left-sidebar',
  standalone: true,
  imports: [
    ButtonModule,
    LeftSidebarThreadsComponent
  ],
  templateUrl: './left-sidebar.component.html',
  styleUrl: './left-sidebar.component.css'
})

export class LeftSidebarComponent {
  threadService: ThreadService = inject(ThreadService)
  toastService: ToastService = inject(ToastService)
  leftSidebarThreadsService: LeftSidebarThreadsService = inject(LeftSidebarThreadsService)
  router: Router = inject(Router)

  async createNewThread() {
    try {
      let thread = await this.threadService.createThread('New Thread')
      await this.router.navigate([`/app/thread/${thread.id}`])
    } catch (e) {
      this.toastService.error(`${e}`)
    }
  }
}
