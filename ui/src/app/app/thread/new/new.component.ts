import { Component, inject, signal, WritableSignal } from '@angular/core'
import { BaseComponent } from '../../base/base.component'
import { BreadcrumbComponent } from '../breadcrumb/breadcrumb.component'
import { ButtonModule } from 'primeng/button'
import { InputTextareaModule } from 'primeng/inputtextarea'
import { MessageComponent } from '../message/message.component'
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms'
import { MenuItem } from 'primeng/api/menuitem'
import { createThread } from '../../../../gql_queries/threads'
import { ThreadService } from '../../../services/thread.service'
import { MessageService } from '../../../services/message.service'
import { ToastService } from '../../../services/toast.service'
import { Thread, ThreadFragment } from '../../../../../graphql/generated'
import { Router } from '@angular/router'

@Component({
  selector: 'app-new',
  standalone: true,
  imports: [
    BaseComponent,
    BreadcrumbComponent,
    ButtonModule,
    InputTextareaModule,
    MessageComponent,
    ReactiveFormsModule
  ],
  templateUrl: './new.component.html',
  styleUrl: './new.component.css'
})
export class NewComponent {
  threadService: ThreadService = inject(ThreadService)
  messageService: MessageService = inject(MessageService)
  toastService: ToastService = inject(ToastService)
  router: Router = inject(Router)
  loading: WritableSignal<boolean> = signal(false)
  thread: WritableSignal<ThreadFragment | undefined> = signal(undefined)
  promptForm: FormGroup = new FormGroup<any>({
    prompt: new FormControl({value: '', disabled: false}, [Validators.required]),
  })

  async createThread() {
    try {
      this.loading.set(true)
      let messageContent = this.getPromptControl().value
      // if the thread hasn't already been created, create it
      let thread = this.thread()
      if (thread == undefined) {
        // create the thread
        let newThread: ThreadFragment = await this.threadService.createThread(messageContent)
        this.thread.set(newThread)
        thread = newThread
      }
      // create the message
      await this.messageService.createMessage(messageContent, thread!.id)
      // navigate to the thread
      await this.router.navigate([`/app/thread/${thread!.id}`])
    } catch(e) {
      this.toastService.error("We've encountered an error, please try again")
      console.error(e)
    }
    finally {
      this.loading.set(false)
    }
  }

  getPromptControl() {
    return this.promptForm.controls['prompt']
  }

  async keyPressed(event: KeyboardEvent) {
    if (event.key === 'Enter') {
      event.preventDefault()
      if (event.shiftKey) {
        let value = this.promptForm.value.prompt
        value += "\n"
        this.promptForm.controls['prompt'].setValue(value)
      } else {
        await this.createThread()
      }
    }
  }
}
