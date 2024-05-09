import { Component, ElementRef, inject, Input, ViewChild } from '@angular/core'
import { ThreadService } from '../../services/thread.service'
import { ToastService } from '../../services/toast.service'
import { BaseComponent } from '../base/base.component'
import { FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms'
import { InputTextareaModule } from 'primeng/inputtextarea'
import { BreadcrumbModule } from 'primeng/breadcrumb'
import { MenuItem } from 'primeng/api/menuitem'
import { ButtonModule } from 'primeng/button'
import { MessageService } from '../../services/message.service'
import { CdkTrapFocus } from '@angular/cdk/a11y'
import { OrderByDirection } from '../../services/graphql.service'
import { UserMessageComponent } from './user-message/user-message.component'
import { MessageComponent } from './message/message.component'
import { BreadcrumbComponent } from './breadcrumb/breadcrumb.component'

@Component({
  selector: 'app-thread',
  standalone: true,
  imports: [
    BaseComponent,
    FormsModule,
    InputTextareaModule,
    BreadcrumbModule,
    ReactiveFormsModule,
    ButtonModule,
    CdkTrapFocus,
    UserMessageComponent,
    MessageComponent,
    BreadcrumbComponent
  ],
  templateUrl: './thread.component.html',
  styleUrl: './thread.component.css'
})
export class ThreadComponent {
  @Input() id: string = ''
  @ViewChild('promptInput') promptInput!:ElementRef
  threadService: ThreadService = inject(ThreadService)
  messageService: MessageService = inject(MessageService)
  toastService: ToastService = inject(ToastService)
  thread: any = undefined
  messages: any = undefined
  home: MenuItem = { icon: 'pi pi-home', routerLink: '/' }

  loading: boolean = false
  promptForm: FormGroup = new FormGroup<any>({
    prompt: new FormControl({value: '', disabled: this.loading}, [Validators.required]),
  })

  async ngOnInit() {
    try {
      this.thread = await this.threadService.getThread(this.id)
    } catch (e) {
      this.toastService.error(`${e}`)
    }
    try {
      this.messages = await this.messageService.getThreadMessages(this.id, 50, [{field: "CREATED_AT", direction: OrderByDirection.Ascending}])
      console.log(this.messages)
    } catch (e) {
      this.toastService.error(`${e}`)
    }
  }

  async sendMessage(){
    try {
      this.toggleTextareaDisable()
      this.toggleLoading()
      let message = await this.messageService.createMessage(this.thread.node.id, this.promptForm.value.prompt)
      this.messages.edges.push({node: message})
      let control = this.getPromptControl()
      control.setValue('')
      control.clearValidators()
      // I'm not sure why but this only works if you use setTimeout
      setTimeout(() => this.promptInput.nativeElement.focus(), 0)
    } catch (e) {
      this.toastService.error(`${e}`)
    } finally {
      this.toggleTextareaDisable()
      this.toggleLoading()
    }
  }
  keyPressed(event: KeyboardEvent) {
    if (event.key === 'Enter') {
      event.preventDefault()
      if (event.shiftKey) {
        let value = this.promptForm.value.prompt
        value += "\n"
        this.promptForm.controls['prompt'].setValue(value)
      } else {
        this.sendMessage()
      }
    }
  }

  getPromptControl() {
    return this.promptForm.controls['prompt']
  }

  toggleTextareaDisable() {
    let control = this.getPromptControl()
    if (control.disabled) {
      control.enable()
    } else {
      control.disable()
    }
  }

  toggleLoading() {
    this.loading = !this.loading
  }
}

