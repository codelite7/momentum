import { Component, inject, Input } from '@angular/core'
import { ButtonModule } from 'primeng/button'
import { ActivatedRoute, Router, RouterLink } from '@angular/router'
import { ThreadFragment } from '../../../../../../graphql/generated'
import { EllipsisModule } from 'ngx-ellipsis'
import { ToastService } from '../../../../services/toast.service'
import { EllipsisService } from '../../../../services/ellipsis.service'

@Component({
  selector: 'app-thread-button',
  standalone: true,
  imports: [
    ButtonModule,
    RouterLink,
    EllipsisModule
  ],
  templateUrl: './thread-button.component.html',
  styleUrl: './thread-button.component.css'
})
export class ThreadButtonComponent {
  router: Router = inject(Router)
  toastService: ToastService = inject(ToastService)
  ellipsisService: EllipsisService = inject(EllipsisService)
  @Input() thread: ThreadFragment | undefined
  selectedThreadId: string | null = null
  constructor(private route: ActivatedRoute) {}
  ngOnInit() {
    this.route.paramMap.subscribe(
      (params => {
        this.selectedThreadId = params.get('id')
      })
    )
  }

  get buttonText(): string {
    return this.ellipsisService.ellipse(this.thread?.name)
  }

  getLinkUrl(): string {
    let url = `/app/thread/${this.thread?.id}`
    return url
  }

  async navigate() {
    try {
      await this.router.navigate([this.getLinkUrl()])
    }catch(e) {
      this.toastService.generalError()
      console.error(e)
    }
  }
}
