import { Component } from '@angular/core';
import { NgOptimizedImage } from '@angular/common'
import { AvatarModule } from 'primeng/avatar'
import { InputTextModule } from 'primeng/inputtext'
import { FormsModule } from '@angular/forms'
import { IconFieldModule } from 'primeng/iconfield'
import { InputIconModule } from 'primeng/inputicon'
import { UserAvatarComponent } from '../common/user-avatar/user-avatar.component'
import { ThreadSearchComponent } from './thread-search/thread-search.component'
import { LogoComponent } from './logo/logo.component'

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [
    NgOptimizedImage,
    AvatarModule,
    InputTextModule,
    FormsModule,
    IconFieldModule,
    InputIconModule,
    UserAvatarComponent,
    ThreadSearchComponent,
    LogoComponent
  ],
  templateUrl: './header.component.html',
  styleUrl: './header.component.css'
})
export class HeaderComponent {
}
