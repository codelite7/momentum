import { Component, Input } from '@angular/core';
import { BreadcrumbModule } from 'primeng/breadcrumb'
import { MenuItem } from 'primeng/api/menuitem'

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
  @Input() thread: any
  home: MenuItem = { icon: 'pi pi-home', routerLink: '/' }
  get breadcrumbItems() {
    return this.getBreadcrumbItems()
  }
  getBreadcrumbItems(): MenuItem[] {
    let items: MenuItem[] = []
    if (this.thread) {
      if (this.thread.node.parent) {
        items.push(this.getBreadcrumbItem(this.thread.node.parent))
      }
      items.push(this.getBreadcrumbItem(this.thread))
      if (this.thread.node.child) {
        items.push(this.getBreadcrumbItem(this.thread.node.child))
      }
    }
    return items
  }

  getBreadcrumbItem(thread: any): MenuItem {
    let menuItem: MenuItem = {
      label: thread.node.name
    }
    // only give url for other threads
    if (thread.node.id != this.thread.node.id) {
      menuItem.url = this.getThreadUrl(thread)
    }

    return menuItem
  }

  getThreadUrl(thread: any): string {
    return `/app/thread/${thread.id}`
  }
}
