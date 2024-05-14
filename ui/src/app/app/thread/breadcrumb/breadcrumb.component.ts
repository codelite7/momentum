import { Component, computed, inject, input, Input, InputSignal, signal, Signal } from '@angular/core'
import { BreadcrumbModule } from 'primeng/breadcrumb'
import { MenuItem } from 'primeng/api/menuitem'
import { ThreadFragment, ThreadQuery } from '../../../../../graphql/generated'
import { DropdownChangeEvent, DropdownModule } from 'primeng/dropdown'
import { Router } from '@angular/router'

@Component({
  selector: 'app-breadcrumb',
  standalone: true,
  imports: [
    BreadcrumbModule,
    DropdownModule
  ],
  templateUrl: './breadcrumb.component.html',
  styleUrl: './breadcrumb.component.css'
})
export class BreadcrumbComponent {
  router: Router = inject(Router)
  thread: InputSignal<ThreadFragment | undefined> = input<ThreadFragment | undefined>(undefined, {alias: 'thread'})
  children = computed(() => {
    let thread = this.thread()
    let children: ThreadFragment[] = []
    if (thread?.children) {
      thread.children.edges?.forEach(edge => {
        if (edge?.node) {
          children.push(edge.node as ThreadFragment)
        }
      })
    }
    return children
  })
  childrenOptions = computed(() => {
    return this.children().map((child: ThreadFragment) => {
      return {
        label: child.name,
        value: child.id
      }
    })
  })
  home: MenuItem = { icon: 'pi pi-home', routerLink: '/' }

  async navigateToChild(event: DropdownChangeEvent) {
    await this.router.navigate([`/app/thread/${event.value}`])
  }
}
