import { Component, inject, Signal, signal, WritableSignal } from '@angular/core'
import { FormsModule } from '@angular/forms'
import { IconFieldModule } from 'primeng/iconfield'
import { InputIconModule } from 'primeng/inputicon'
import { InputTextModule } from 'primeng/inputtext'
import { Subscription } from 'rxjs'
import { BookmarkService } from '../../services/bookmark.service'
import { ToastService } from '../../services/toast.service'
import { BookmarkComponent } from './bookmark/bookmark.component'
import {
  BookmarkFragment,
  BookmarkOrderField,
  BookmarksQuery,
  OrderDirection,
  ThreadFragment
} from '../../../../graphql/generated'
import { debounce } from 'lodash'
import { ProgressSpinnerModule } from 'primeng/progressspinner'

@Component({
  selector: 'app-right-sidebar',
  standalone: true,
  imports: [
    FormsModule,
    IconFieldModule,
    InputIconModule,
    InputTextModule,
    BookmarkComponent,
    ProgressSpinnerModule
  ],
  templateUrl: './right-sidebar.component.html',
  styleUrl: './right-sidebar.component.css'
})
export class RightSidebarComponent {
  query: string = ''
  bookmarks: WritableSignal<BookmarkFragment[] | undefined> = signal(undefined)
  bookmarkService: BookmarkService = inject(BookmarkService)
  toastService: ToastService = inject(ToastService)

  bookmarkDeletedSubscription?: Subscription
  bookmarkCreatedSubscription?: Subscription
  searchResults: WritableSignal<BookmarkFragment[] | undefined> = signal(undefined)
  searchLoading: WritableSignal<boolean> = signal(false)
  searchDebounce = debounce(async () => {
    if (this.query) {
      this.searchLoading.set(true)
      // I'm setting a timeout because the search is so fast that it looks janky without it
      setTimeout(async () => {
        try {
          let searchResults = await this.bookmarkService.bookmarks(
            {
              where: {
                or: [
                  {
                    hasThreadWith: [
                      {
                        nameContainsFold: this.query
                      }
                    ]
                  },
                  {
                    hasMessageWith: [{
                      contentContainsFold: this.query
                    }]
                  },
                  {
                    hasResponseWith: [{
                      contentContainsFold: this.query
                    }]
                  }
                ]
              }
            }
          )
          this.searchResults.set(searchResults)
        } catch (e) {
          this.toastService.error('We encountered an error, please try again later.')
          console.error(e)
        } finally {
          this.searchLoading.set(false)
        }
      }, 1000)
    }
  }, 1000)

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
        this.bookmarks.update((bookmarks: BookmarkFragment[] | undefined): BookmarkFragment[] | undefined => {
          if (!bookmarks) {
            return bookmarks
          }
          return bookmarks.filter((bookmark: BookmarkFragment) => bookmark.id !== id)
        })
      }
    )
    // when a bookmark is created, add it to the bookmarks
    this.bookmarkCreatedSubscription = this.bookmarkService.bookmarkCreatedEmitter.subscribe(
      async (id: string) => {
        let fetched = await this.bookmarkService.bookmark(id)
        if (fetched) {
          this.bookmarks.update((bookmarks: BookmarkFragment[] | undefined): BookmarkFragment[] | undefined => {
            let bookmark = fetched
            if (!bookmarks) {
              return [fetched as BookmarkFragment]
            } else {
              bookmarks.push(fetched)
              return bookmarks
            }
          })
        }
      }
    )
  }

  async loadBookmarks() {
    this.bookmarks.set(await this.bookmarkService.bookmarks({
      first: 50,
      orderBy: [{ field: BookmarkOrderField.CreatedAt, direction: OrderDirection.Desc }]
    }))
  }

  search(query: string) {
    if (query) {
      // if there is a query then execute the debounce function
      this.searchDebounce()
    } else {
      this.reset()
    }
  }

  reset() {
    this.searchDebounce.cancel()
    this.searchLoading.set(false)
    this.searchResults.set(undefined)
  }
}
