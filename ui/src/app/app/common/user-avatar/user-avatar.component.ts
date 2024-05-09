import { Component, Input } from '@angular/core';
import { NgOptimizedImage } from '@angular/common'
import { ToastModule } from 'primeng/toast'
import { MenuModule } from 'primeng/menu'
import { ButtonModule } from 'primeng/button'
import { MenuItem } from 'primeng/api'

@Component({
  selector: 'app-user-avatar',
  standalone: true,
  imports: [
    NgOptimizedImage,
    ToastModule,
    MenuModule,
    ButtonModule
  ],
  templateUrl: './user-avatar.component.html',
  styleUrl: './user-avatar.component.css'
})
export class UserAvatarComponent {
  @Input() showMenu: boolean = false;
  menuItems: MenuItem[] = [{
    label: 'Sign out',
    icon: 'pi pi-power-off',
    url: '/signout'
  }]
}
