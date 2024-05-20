import {
  Component,
  computed,
  effect,
  ElementRef,
  inject,
  Signal,
  signal,
  ViewChild,
  WritableSignal
} from '@angular/core'
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
import { UserMessageComponent } from './user-message/user-message.component'
import { MessageComponent } from './message/message.component'
import { BreadcrumbComponent } from './breadcrumb/breadcrumb.component'
import {
  MessageOrderField,
  OrderDirection,
  ThreadFragment, ThreadMessageFragment,
} from '../../../../graphql/generated'
import { ActivatedRoute, Router } from '@angular/router'
import { ProgressSpinnerModule } from 'primeng/progressspinner'
import { cloneDeep } from 'lodash'


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
    BreadcrumbComponent,
    ProgressSpinnerModule
  ],
  templateUrl: './thread.component.html',
  styleUrl: './thread.component.css'
})
export class ThreadComponent {
  @ViewChild('promptInput') promptInput!:ElementRef
  threadService: ThreadService = inject(ThreadService)
  messageService: MessageService = inject(MessageService)
  toastService: ToastService = inject(ToastService)
  thread: WritableSignal<ThreadFragment | undefined> = signal(undefined)
  messages: WritableSignal<ThreadMessageFragment[]> = signal([])
  loading: WritableSignal<boolean> = signal(true)
  responses = computed(() => {
    let responses: any[] = []
    this.messages().forEach(message => {
      if (message.response) {
        responses.push(message.response)
      }
    })
    return responses
  })
  numMessages: Signal<number> = computed(() => {
    return this.messages().length
  })
  lastMessage: Signal<ThreadMessageFragment | undefined> = computed(() => {
    let messages = this.messages()
    if (messages.length > 0) {
      return messages[messages.length - 1]
    }
    return undefined
  })
  lastResponse = computed(() => {
    let responses = this.responses()
    if (responses.length > 0) {
      return responses[responses.length - 1]
    }
    return undefined
  })
  lastResponseHasContent = computed(() => {
    return this.lastResponse()?.content ?? '' != ''
  })
  lastMessageHasResponse = computed(() => {
    return this.lastMessage()?.response?.content ?? '' != ''
  })
  messageInputDisabled = computed(() => {
    return this.loading() || (this.numMessages() > 0 && !this.lastMessageHasResponse())
  })
  promptForm: FormGroup = new FormGroup<any>({
    prompt: new FormControl({value: '', disabled: true}, [Validators.required]),
  })

  constructor(private route: ActivatedRoute) {
    effect(() => {
      let messageInputDisabled = this.messageInputDisabled()
      if (this.messageInputDisabled()) {
        this.disableMessageInput()
      } else {
        this.enableMessageInput()
      }
    })
    effect(async () => {
      while (this.numMessages() > 0 && !this.lastResponseHasContent()) {
        await new Promise(r => setTimeout(r, 1000));
        try {
          let message = await this.messageService.threadMessage({id: this.lastMessage()!.id, userId: ''})
          let responseContent = message?.response?.content
          if (responseContent) {
            let messages = cloneDeep(this.messages())
            messages[messages.length - 1].response!.content = responseContent
            this.messages.set(messages)
          }
        } catch (e) {
          console.error(e)
        }
      }
    }, {allowSignalWrites: true})
  }
  async ngOnInit() {
    this.loadingTrue()
    try {
      this.route.paramMap.subscribe(
        (async (params) => {
          let id = params.get('id')
          if (id) {
            try {
              this.thread.set(await this.threadService.thread(id))
            } catch (e) {
              this.toastService.error(`${e}`)
            }
            try {
              this.messages.set(await this.messageService.threadMessages({userId: '', threadId: id, first: 50, orderBy: [{field: MessageOrderField.CreatedAt, direction: OrderDirection.Asc}]}))
            } catch (e) {
              this.toastService.error(`${e}`)
            }
          }
        })
      )
    } finally {
      this.loadingFalse()
    }

  }

  async sendMessage(){
    try {
      this.disableMessageInput()
      this.loadingTrue()
      let newMessage = await this.messageService.createMessage(this.promptForm.value.prompt, this.thread()!.id)
      let newThreadMessage = await this.messageService.threadMessage({id:newMessage.id, userId: ''})
      if (newThreadMessage) {
        this.messages.set(await this.messageService.threadMessages({userId: '', threadId: this.thread()!.id, first: 50, orderBy: [{field: MessageOrderField.CreatedAt, direction: OrderDirection.Asc}]}))
        let control = this.getPromptControl()
        control.setValue('')
        control.clearValidators()
        // I'm not sure why but this only works if you use setTimeout
        setTimeout(() => this.promptInput.nativeElement.focus(), 0)
      } else {
        throw new Error('error sending message')
      }
    } catch (e) {
      // @ts-ignore
      if (e?.response?.errors[0].message.includes('wait for a response')) {
        this.toastService.error('You must wait for a response to your previous message before sending another')
      } else {
        this.toastService.error(`${e}`)
        console.error(e)
      }
    } finally {
      this.loadingFalse()
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

  disableMessageInput() {
    this.getPromptControl().disable()
  }

  enableMessageInput() {
    this.getPromptControl().enable()
  }

  loadingTrue() {
    this.loading.set(true)
  }

  loadingFalse() {
    this.loading.set(false)
  }
}

