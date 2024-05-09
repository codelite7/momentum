import { Component } from '@angular/core';
import { IconFieldModule } from 'primeng/iconfield'
import { InputIconModule } from 'primeng/inputicon'
import { InputTextModule } from 'primeng/inputtext'
import { PaginatorModule } from 'primeng/paginator'

@Component({
  selector: 'app-thread-search',
  standalone: true,
  imports: [
    IconFieldModule,
    InputIconModule,
    InputTextModule,
    PaginatorModule
  ],
  templateUrl: './thread-search.component.html',
  styleUrl: './thread-search.component.css'
})
export class ThreadSearchComponent {
  query: string = ''
}
