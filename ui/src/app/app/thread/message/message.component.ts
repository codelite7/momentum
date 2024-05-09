import { Component, inject, Input } from '@angular/core'
import { AvatarModule } from 'primeng/avatar'
import { DatePipe } from '@angular/common'
import { AuthService } from '../../../auth.service'
import { UserAvatarComponent } from '../../common/user-avatar/user-avatar.component'

@Component({
  selector: 'app-message',
  standalone: true,
  imports: [
    AvatarModule,
    DatePipe,
    UserAvatarComponent
  ],
  templateUrl: './message.component.html',
  styleUrl: './message.component.css'
})
export class MessageComponent {
  @Input() message!: any
  authService: AuthService = inject(AuthService)
  email: string = ''
  async ngOnInit(): Promise<void> {
    this.email = await this.authService.email()
  }

  get sentAt(): string {
    let createdAtDate = new Date(this.message.node.createdAt)
    if (this.wasSentToday()) {
      return createdAtDate.toLocaleTimeString()
    }
    return createdAtDate.toLocaleDateString()
  }

  wasSentToday(): boolean {
    let sentAt = new Date(this.message.node.createdAt)
    return sentAt.setHours(0,0,0,0) == new Date().setHours(0,0,0,0)
  }
}
