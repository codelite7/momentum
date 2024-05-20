import { Component, computed, inject, input, Input, InputSignal, signal, Signal } from '@angular/core'
import { BreadcrumbModule } from 'primeng/breadcrumb'
import { MenuItem } from 'primeng/api/menuitem'
import { ThreadFragment, ThreadQuery } from '../../../../../graphql/generated'
import { DropdownChangeEvent, DropdownModule } from 'primeng/dropdown'
import { Router } from '@angular/router'
import { EllipsisService } from '../../../services/ellipsis.service'

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
  ellipsisService: EllipsisService = inject(EllipsisService)
  thread: InputSignal<ThreadFragment | undefined> = input<ThreadFragment | undefined>(undefined, { alias: 'thread' })
  threadName: Signal<string> = computed(() => {
    return this.ellipsisService.ellipse(this.thread()?.name)
  })
  parent: Signal<any | undefined> = computed(() => {
    return this.thread()?.parent
  })
  parentName: Signal<any> = computed(() => {
    return this.ellipsisService.ellipse(this.parent()?.name)
  })
  firstChild: Signal<any | undefined> = computed((): any | undefined => {
    let children: any[] = this.thread()?.children?.edges ?? []
    if (children.length > 0) {
      return children[0].node
    }
    return undefined
  })
  firstChildName: Signal<string> = computed(() => {
    return this.ellipsisService.ellipse(this.firstChild()?.name)
  })
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
        label: this.ellipsisService.ellipse(child.name),
        value: child.id
      }
    })
  })
  home: MenuItem = { icon: 'pi pi-home', routerLink: '/' }

  async navigateToChild(event: DropdownChangeEvent) {
    await this.router.navigate([`/app/thread/${event.value}`])
  }
}
