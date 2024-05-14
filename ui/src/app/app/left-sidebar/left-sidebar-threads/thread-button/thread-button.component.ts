import { Component, Input } from '@angular/core'
import { ButtonModule } from 'primeng/button'
import { ActivatedRoute, Router, RouterLink } from '@angular/router'
import { ThreadFragment } from '../../../../../../graphql/generated'

@Component({
  selector: 'app-thread-button',
  standalone: true,
  imports: [
    ButtonModule,
    RouterLink
  ],
  templateUrl: './thread-button.component.html',
  styleUrl: './thread-button.component.css'
})
export class ThreadButtonComponent {
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

  getLinkUrl(): string {
    let url = `/app/thread/${this.thread?.id}`
    return url
  }
}
