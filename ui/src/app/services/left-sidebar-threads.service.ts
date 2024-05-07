import { EventEmitter, Injectable } from '@angular/core'

@Injectable({
  providedIn: 'root'
})
export class LeftSidebarThreadsService {
  refreshThreadsEmitter: EventEmitter<void> = new EventEmitter<void>();
  constructor() { }

  refreshThreads(): void {
    this.refreshThreadsEmitter.emit()
  }
}
