import { Component } from '@angular/core';
import { HeaderComponent } from '../header/header.component'
import { LeftSidebarComponent } from '../left-sidebar/left-sidebar.component'
import { RightSidebarComponent } from '../right-sidebar/right-sidebar.component'
import { BreadcrumbModule } from 'primeng/breadcrumb'
import { MenuItem } from 'primeng/api'

@Component({
  selector: 'app-base',
  standalone: true,
  imports: [
    HeaderComponent,
    LeftSidebarComponent,
    RightSidebarComponent,
    BreadcrumbModule
  ],
  templateUrl: './base.component.html',
  styleUrl: './base.component.css'
})
export class BaseComponent {
  items: MenuItem[] = []
  home: MenuItem = {icon: 'pi pi-home', routerLink: '/'}

}
