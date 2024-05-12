import { Component, inject, Input } from '@angular/core'
import { ButtonModule } from 'primeng/button'
import { TooltipModule } from 'primeng/tooltip'
import { CardModule } from 'primeng/card'
import { Router } from '@angular/router'
import { BookmarkService } from '../../../services/bookmark.service'
import { ToastService } from '../../../services/toast.service'
import { BookmarkFragment } from '../../../../../graphql/generated'

@Component({
  selector: 'app-bookmark',
  standalone: true,
  imports: [
    ButtonModule,
    TooltipModule,
    CardModule
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

  navigate() {
    // if (this.bookmark?.message) {
    //   this.router.navigate([`app/thread/${this.bookmark?.message.thread.edges[0].node.id}/message/${this.bookmark.node.message.id}`])
    // } else if (this.bookmark.node.thread) {
    //   this.router.navigate([`app/thread/${this.bookmark.node.thread.id}`])
    // }
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
