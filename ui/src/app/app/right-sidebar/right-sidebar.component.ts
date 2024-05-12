import { Component, inject } from '@angular/core'
import { FormsModule } from '@angular/forms'
import { IconFieldModule } from 'primeng/iconfield'
import { InputIconModule } from 'primeng/inputicon'
import { InputTextModule } from 'primeng/inputtext'
import { Subscription } from 'rxjs'
import { BookmarkService } from '../../services/bookmark.service'
import { ToastService } from '../../services/toast.service'
import { BookmarkComponent } from './bookmark/bookmark.component'
import { BookmarkFragment, BookmarkOrderField, BookmarksQuery, OrderDirection } from '../../../../graphql/generated'

@Component({
  selector: 'app-right-sidebar',
  standalone: true,
  imports: [
    FormsModule,
    IconFieldModule,
    InputIconModule,
    InputTextModule,
    BookmarkComponent
  ],
  templateUrl: './right-sidebar.component.html',
  styleUrl: './right-sidebar.component.css'
})
export class RightSidebarComponent {
  query: string = ''
  bookmarks: BookmarkFragment[] = []
  bookmarkService: BookmarkService = inject(BookmarkService)
  toastService: ToastService = inject(ToastService)

  bookmarkDeletedSubscription?: Subscription
  bookmarkCreatedSubscription?: Subscription

  async ngOnInit() {
    try {
      await this.loadBookmarks()
    } catch (e) {
      this.toastService.error(`${e}`)
      console.error(e)
    }
    // when a bookmark is deleted, remove it from the bookmarks if it's present
    this.bookmarkDeletedSubscription = this.bookmarkService.bookmarkDeletedEmitter.subscribe(
      (id) => {
        if (this.bookmarks) {
          this.bookmarks = this.bookmarks.filter((bookmark: any) => bookmark.id !==id )
        }
      }
    )
    // when a bookmark is created, add it to the bookmarks
    this.bookmarkCreatedSubscription = this.bookmarkService.bookmarkCreatedEmitter.subscribe(
      async (id: string) => {
        let fetched = await this.bookmarkService.bookmark(id)
        this.bookmarks.push(fetched as BookmarkFragment)
      }
    )
  }

  async loadBookmarks() {
    this.bookmarks = await this.bookmarkService.bookmarks({first: 50, orderBy: [{field: BookmarkOrderField.CreatedAt, direction: OrderDirection.Desc}]})
    console.log('bookmarks', this.bookmarks)
  }
}
