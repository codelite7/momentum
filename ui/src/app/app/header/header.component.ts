import { Component } from '@angular/core';
import { NgOptimizedImage } from '@angular/common'
import { AvatarModule } from 'primeng/avatar'
import { InputTextModule } from 'primeng/inputtext'
import { FormsModule } from '@angular/forms'
import { IconFieldModule } from 'primeng/iconfield'
import { InputIconModule } from 'primeng/inputicon'

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [
    NgOptimizedImage,
    AvatarModule,
    InputTextModule,
    FormsModule,
    IconFieldModule,
    InputIconModule
  ],
  templateUrl: './header.component.html',
  styleUrl: './header.component.css'
})
export class HeaderComponent {
  query: string = ''
}
