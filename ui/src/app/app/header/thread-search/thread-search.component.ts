import { Component, inject } from '@angular/core'
import { IconFieldModule } from 'primeng/iconfield'
import { InputIconModule } from 'primeng/inputicon'
import { InputTextModule } from 'primeng/inputtext'
import { PaginatorModule } from 'primeng/paginator'
import { debounce } from 'lodash'
import { ThreadService } from '../../../services/thread.service'
import { SidebarModule } from 'primeng/sidebar'
import { ThreadFragment } from '../../../../../graphql/generated'
import { NgOptimizedImage } from '@angular/common'
import { DialogModule } from 'primeng/dialog'
import { ToastService } from '../../../services/toast.service'
import { ProgressSpinnerModule } from 'primeng/progressspinner'
import { CardModule } from 'primeng/card'
import { Router, RouterLink } from '@angular/router'

@Component({
  selector: 'app-thread-search',
  standalone: true,
  imports: [
    IconFieldModule,
    InputIconModule,
    InputTextModule,
    PaginatorModule,
    SidebarModule,
    NgOptimizedImage,
    DialogModule,
    ProgressSpinnerModule,
    CardModule,
    RouterLink
  ],
  templateUrl: './thread-search.component.html',
  styleUrl: './thread-search.component.css'
})
export class ThreadSearchComponent {
  threadService: ThreadService = inject(ThreadService)
  toastService: ToastService = inject(ToastService)
  router: Router = inject(Router)
  query: string = ''
  loading: boolean = false;
  searchResults?: ThreadFragment[] = undefined
  searchDebounce = debounce(async () => {
    this.loading = true
    // I'm setting a timeout because the search is so fast that it looks janky without it
    setTimeout(async () => {
      try {
        this.searchResults = await this.threadService.threads(
          {
            where: {
              or: [
                {
                  nameContainsFold: this.query
                },
                {
                  hasMessagesWith: [{
                    or: [
                      {
                        contentContainsFold: this.query
                      },
                      {
                        hasResponseWith: [
                          {
                            contentContainsFold: this.query
                          }
                        ]
                      }
                    ]
                  }]
                }
              ]
            }
          }
        )
      } catch (e) {
        this.toastService.error('We encountered an error, please try again later.')
        console.error(e)
      } finally {
        this.loading = false
      }
    }, 1000)

  }, 1000)
  searchResultsVisible: boolean = false

  search() {
    if (this.query) {
      // if there is a query then execute the debounce function
      this.searchDebounce()
    } else {
      this.reset()
    }
  }

  reset() {
    this.searchDebounce.cancel()
    this.searchResults = undefined
    this.loading = false
  }

  async goToSearchResult(thread: ThreadFragment) {
    let url = `/app/thread/${thread.id}`
    try {
      await this.router.navigate([url])
    } catch (e) {
      console.error(e)
    }
  }

  lostFocus() {
    setTimeout(() => {
      this.searchResultsVisible = false
    }, 100)
  }

  protected readonly setTimeout = setTimeout
}
