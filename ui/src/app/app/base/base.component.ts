import { Component } from '@angular/core';
import { HeaderComponent } from '../header/header.component'
import { LeftSidebarComponent } from '../left-sidebar/left-sidebar.component'
import { RightSidebarComponent } from '../right-sidebar/right-sidebar.component'
import { BreadcrumbModule } from 'primeng/breadcrumb'
import { MenuItem } from 'primeng/api'
import { EllipsisModule } from 'ngx-ellipsis'

@Component({
  selector: 'app-base',
  standalone: true,
  imports: [
    HeaderComponent,
    LeftSidebarComponent,
    RightSidebarComponent,
    BreadcrumbModule,
    EllipsisModule
  ],
  templateUrl: './base.component.html',
  styleUrl: './base.component.css'
})
export class BaseComponent {
  numbers = Array(100).fill(1)
  items: MenuItem[] = []
  home: MenuItem = {icon: 'pi pi-home', routerLink: '/'}

}
