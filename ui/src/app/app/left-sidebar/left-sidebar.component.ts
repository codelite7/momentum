import { Component } from '@angular/core';
import { ButtonModule } from 'primeng/button'
import { LeftSidebarThreadsComponent } from './left-sidebar-threads/left-sidebar-threads.component'

@Component({
  selector: 'app-left-sidebar',
  standalone: true,
  imports: [
    ButtonModule,
    LeftSidebarThreadsComponent
  ],
  templateUrl: './left-sidebar.component.html',
  styleUrl: './left-sidebar.component.css'
})
export class LeftSidebarComponent {

}
