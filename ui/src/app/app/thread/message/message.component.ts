import { Component, ElementRef, inject, Input, ViewChild } from '@angular/core'
import { AvatarModule } from 'primeng/avatar'
import { DatePipe, TitleCasePipe, UpperCasePipe, ViewportScroller } from '@angular/common'
import { AuthService } from '../../../auth.service'
import { UserAvatarComponent } from '../../common/user-avatar/user-avatar.component'
import { CardModule } from 'primeng/card'
import { ButtonModule } from 'primeng/button'
import { TooltipModule } from 'primeng/tooltip'
import { CdkCopyToClipboard } from '@angular/cdk/clipboard'
import { BookmarkService } from '../../../services/bookmark.service'
import { ToastService } from '../../../services/toast.service'
import { BookmarkFragment, MessageFragment, ThreadMessageFragment } from '../../../../../graphql/generated'
import { Subscription } from 'rxjs'
import { ThreadService } from '../../../services/thread.service'
import { Router } from '@angular/router'
import { MarkdownComponent, MarkdownModule } from 'ngx-markdown'
import { AgentAvatarComponent } from '../../common/agent-avatar/agent-avatar.component'
import { TypewriterComponent } from '../../common/typewriter/typewriter.component'

@Component({
  selector: 'app-message',
  standalone: true,
  imports: [
    AvatarModule,
    DatePipe,
    UserAvatarComponent,
    CardModule,
    ButtonModule,
    TooltipModule,
    CdkCopyToClipboard,
    MarkdownModule,
    UpperCasePipe,
    TitleCasePipe,
    AgentAvatarComponent,
    TypewriterComponent
  ],
  templateUrl: './message.component.html',
  styleUrl: './message.component.css'
})
export class MessageComponent {
  @Input() message: ThreadMessageFragment | undefined
  @Input() response: boolean = false;
  @Input() threadId?: string = ''
  @Input() last: boolean = false;
  // @ts-ignore
  @ViewChild('stuff') stuff:ElementRef;
  bookmark: BookmarkFragment | undefined
  authService: AuthService = inject(AuthService)
  bookmarkService: BookmarkService = inject(BookmarkService)
  threadService: ThreadService = inject(ThreadService)
  toastService: ToastService = inject(ToastService)
  router: Router = inject(Router)
  email: string = ''
  hover: boolean = false
  bookmarkDeletedSubscription?: Subscription

  async ngOnInit(): Promise<void> {
    if (!this.message?.response?.content) {
      setTimeout(() => {
        this.message!.response!.content = "Some ai response"
      }, 5000)
    }
    this.email = await this.authService.email()
    if (this.message && this.message.bookmarks.edges && this.message.bookmarks.edges.length > 0) {
      this.bookmark = this.message.bookmarks.edges[0]?.node as BookmarkFragment
    }
    this.bookmarkDeletedSubscription = this.bookmarkService.bookmarkDeletedEmitter.subscribe(id => {
      if (this.bookmark?.id === id) {
        this.bookmark = undefined
      }
    })
  }

  get sentAt(): string {
    if (this.message) {
      let createdAtDate = new Date(this.message.createdAt)
      if (this.wasSentToday()) {
        return createdAtDate.toLocaleTimeString()
      }
      return createdAtDate.toLocaleDateString()
    }
    return ''
  }

  get sentByName(): string {
    if (!this.response) {
      return 'me'
    }
    return `${this.message?.response?.sentBy.provider}`
  }

  wasSentToday(): boolean {
    if (this.message) {
      let sentAt = new Date(this.message.createdAt)
      return sentAt.setHours(0,0,0,0) == new Date().setHours(0,0,0,0)
    }
    return false
  }

  async removeBookmark() {
    try {
      await this.bookmarkService.deleteBookmark(this.bookmark!.id)
      this.toastService.success('Removed bookmark')
      this.bookmark = undefined
    } catch (e) {
      this.toastService.error(`${e}`)
      console.error(e)
    }
  }

  async bookmarkMessage() {
    try {
      if (this.message) {
      let result = await this.bookmarkService.bookmarkMessage(this.message.id)
      this.toastService.success('Bookmarked message')
      this.bookmark = result
      }
    } catch (e) {
      this.toastService.error(`${e}`)
      console.error(e)
    }
  }

  async newThread() {
    try {
      let result = await this.threadService.createThread('New Thread', this.threadId)
      this.toastService.success('Created new thread')
      await this.router.navigate([`/app/thread/${result.id}`])
    } catch (e) {
      this.toastService.error(`${e}`)
      console.error(e)
    }
  }

  get typewriterWords(): string[] {
    if (this.message?.response?.content) {
      return [this.message?.response?.content]
    }
    return []
  }
}
