import { Component, inject, Input } from '@angular/core'
import { ButtonModule } from 'primeng/button'
import { TooltipModule } from 'primeng/tooltip'
import { CardModule } from 'primeng/card'
import { Router } from '@angular/router'
import { BookmarkService } from '../../../services/bookmark.service'
import { ToastService } from '../../../services/toast.service'
import { BookmarkFragment } from '../../../../../graphql/generated'
import { EllipsisModule } from 'ngx-ellipsis'

@Component({
  selector: 'app-bookmark',
  standalone: true,
  imports: [
    ButtonModule,
    TooltipModule,
    CardModule,
    EllipsisModule
  ],
  templateUrl: './bookmark.component.html',
  styleUrl: './bookmark.component.css'
})
export class BookmarkComponent {
  @Input() bookmark: BookmarkFragment | undefined;
  router: Router = inject(Router)
  bookmarkService: BookmarkService = inject(BookmarkService)
  toastService: ToastService = inject(ToastService)

  ngOnInit() {
  }

  async navigate() {
    try {
      if (this.bookmark?.message) {
        await this.router.navigate([`app/thread/${this.bookmark?.message.thread.id}`], {queryParams: {messageId: this.bookmark?.message.id}})
      } else if (this.bookmark?.thread) {
        await this.router.navigate([`app/thread/${this.bookmark.thread.id}`])
      } else if (this.bookmark?.response) {
        await this.router.navigate([`app/thread/${this.bookmark.response.message.thread.id}`], {fragment: this.bookmark.response.id})
      }
    } catch(e) {
      this.toastService.generalError()
      console.error(e)
    }
  }

  async removeBookmark() {
    try {
      await this.bookmarkService.deleteBookmark(this.bookmark!.id)
      this.toastService.success('Removed bookmark')
    } catch (e) {
      this.toastService.error(`${e}`)
      console.error(e)
    }
  }
}
