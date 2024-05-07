import { Component, Input } from '@angular/core'
import { ButtonModule } from 'primeng/button'

@Component({
  selector: 'app-thread-button',
  standalone: true,
  imports: [
    ButtonModule
  ],
  templateUrl: './thread-button.component.html',
  styleUrl: './thread-button.component.css'
})
export class ThreadButtonComponent {
  @Input() thread: any

  getLinkUrl(): string {
    let url = `app/thread/${this.thread.id}`
    console.log(this.thread, url)
    return url
  }
}
