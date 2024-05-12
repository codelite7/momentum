import { Component, Input } from '@angular/core';
import { BreadcrumbModule } from 'primeng/breadcrumb'
import { MenuItem } from 'primeng/api/menuitem'
import { ThreadFragment, ThreadQuery } from '../../../../../graphql/generated'

@Component({
  selector: 'app-breadcrumb',
  standalone: true,
  imports: [
    BreadcrumbModule
  ],
  templateUrl: './breadcrumb.component.html',
  styleUrl: './breadcrumb.component.css'
})
export class BreadcrumbComponent {
  @Input() thread: ThreadFragment | undefined = undefined
  home: MenuItem = { icon: 'pi pi-home', routerLink: '/' }
  get breadcrumbItems() {
    return this.getBreadcrumbItems()
  }
  getBreadcrumbItems(): MenuItem[] {
    let items: MenuItem[] = []
    if (this.thread) {
      if (this.thread.parent) {
        items.push(this.getBreadcrumbItem(this.thread.parent))
      }
      items.push(this.getBreadcrumbItem(this.thread))
      if (this.thread.child) {
        items.push(this.getBreadcrumbItem(this.thread.child))
      }
    }
    return items
  }

  getBreadcrumbItem(thread: ThreadFragment): MenuItem {
    let menuItem: MenuItem = {
      label: thread.name
    }
    // only give url for other threads
    if (thread.id != this.thread?.id) {
      menuItem.url = this.getThreadUrl(thread)
    }

    return menuItem
  }

  getThreadUrl(thread: any): string {
    return `/app/thread/${thread.id}`
  }
}
