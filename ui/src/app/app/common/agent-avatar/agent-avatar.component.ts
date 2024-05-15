import { Component } from '@angular/core';
import { ButtonModule } from 'primeng/button'
import { MenuModule } from 'primeng/menu'
import { SharedModule } from 'primeng/api'
import { ToastModule } from 'primeng/toast'

@Component({
  selector: 'app-agent-avatar',
  standalone: true,
  imports: [
    ButtonModule,
    MenuModule,
    SharedModule,
    ToastModule
  ],
  templateUrl: './agent-avatar.component.html',
  styleUrl: './agent-avatar.component.css'
})
export class AgentAvatarComponent {

}
