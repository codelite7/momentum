import { Component, inject } from '@angular/core'
import { ThreadService } from '../../../services/thread.service'
import { ToastService } from '../../../services/toast.service'
import { NgForOf, NgIf } from '@angular/common'
import { ThreadButtonComponent } from './thread-button/thread-button.component'
import {
  BookmarkFragment,
  MessageOrderField,
  OrderDirection,
  ThreadFragment,
  ThreadsQuery
} from '../../../../../graphql/generated'
import { Subscription } from 'rxjs'
import { ActivatedRoute } from '@angular/router'

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
  threads: ThreadFragment[] = Array(30).fill({name: 'Some thread name that is longer than it should be and should be truncated'} as ThreadFragment)
  threadCreatedSubscription: Subscription

  constructor(private route: ActivatedRoute) {
    this.threadCreatedSubscription = this.threadService.threadCreatedEmitter.subscribe(
      async (id) => {
        let thread = await this.threadService.thread(id)
        if (thread) {
          this.threads.unshift(thread)
        }
      }
    )
  }

  async ngOnInit(): Promise<void> {
    await this.loadThreads()
  }

  async loadThreads() {
    try {
      // this.threads = await this.threadService.threads({first: 50})
    } catch (e) {
      this.toastService.error(`${e}`)
    }
  }
}
